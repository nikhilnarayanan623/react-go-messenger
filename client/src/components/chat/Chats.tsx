import React, { useEffect, useState } from "react";
import Friends from "./Friends";
import { Link, Outlet } from "react-router-dom";
import { PiMessengerLogo } from "react-icons/pi";
import { List } from "@material-tailwind/react";
import { useParams } from "react-router-dom";
import UserChat from "../../api/chat/chats";
import { RecentlyChattedFriends } from "../../types/Chat";
import { useDispatch, useSelector } from "react-redux";
import {
  setCurrentChat,
  clearCurrentChat,
} from "../../features/slices/chatSlice";
import { toast } from "react-toastify";
import { sortByProperty } from "../../utils/helpers";
import { selectAccessToken } from "../../features/slices/authSlice";
import { useWebSocketContext } from "../../features/contexts/socketContext";

const Chats: React.FC = () => {
  const { chatId } = useParams();
  const [chats, setChats] = useState<RecentlyChattedFriends[] | null>(null);
  const userChat = UserChat();
  const dispatch = useDispatch();
  const token = useSelector(selectAccessToken);
  const { setSocket } = useWebSocketContext();

  const fetchChats = async () => {
    try {
      const response = await userChat.getRecentlyChattedFriends();
      const sortedData: RecentlyChattedFriends[] = sortByProperty(
        response?.data,
        "last_message_at",
        true
      );
      setChats(sortedData);
    } catch (error: any) {
      toast.error(error?.data?.error[0] || "Something went wrong", {
        position: toast.POSITION.BOTTOM_RIGHT,
      });
    }
  };

  useEffect(() => {
    const Socket = new WebSocket("ws://localhost:8080/api/ws");
    if (token) {
      Socket.onopen = () => {

        Socket.send(
          JSON.stringify({
            token,
          })
        );
        setSocket(Socket);
      };
      Socket.onclose = (event) => {
        console.log("Socket Closed Connection: ", event);
      };

      Socket.onerror = (error) => {
        console.log("Socket Error: ", error);
      };
      return () => {
        Socket.close();
      };
    }
  }, [token]);

  useEffect(() => {
    fetchChats();
    return () => {
      dispatch(clearCurrentChat());
    };
  }, []);

  return (
    <div className='pl-10 flex h-screen'>
      <div className=' w-3/12 '>
        <div className=' w-full '>
          <div className='p-5 border-b border-gray-300 pb-10'>Username</div>
        </div>
        <List className='p-0'>
          {chats?.map((chat) => (
            <Link
              to={`chat/${chat.chat_id}/message/${chat?.user_id}`}
              key={chat?.chat_id}
              onClick={() => {
                dispatch(setCurrentChat({ currentChat: chat }));

              }}

            >
              <Friends {...chat} />
            </Link>
          ))}
        </List>
      </div>
      <div className='pl-10 w-8/12 flex justify-center items-center'>
        {chatId ? (
          <Outlet />
        ) : (
          <div className='flex flex-col items-center'>
            <PiMessengerLogo className=' h-14 w-14' />
            <h2 className='text-2xl pt-1 font-bold text-blue-gray-900'>
              Your messages
            </h2>
            <h5 className='text-blue-gray-900 text-lg '>
              Send messages with your friends
            </h5>
            <button className='mt-2 bg-blue-500 rounded-md p-1 text-sm text-white hover:bg-blue-600'>
              Send message
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default Chats;
