import React from "react";
import Friends from "./Friends";
import { Link, Outlet } from "react-router-dom";
import { PiMessengerLogo } from "react-icons/pi";
import { List } from "@material-tailwind/react";
import { useParams } from "react-router-dom";

const Chats: React.FC = () => {
  const { userId } = useParams();
  return (
    <div className='pl-10 flex h-screen'>
      <div className=' w-3/12 '>
        <div className=' w-full '>
          <div className='p-5 border-b border-gray-300 pb-10'>Username</div>
        </div>
        <List className='p-0'>
          <Link to='message/8'>
            <Friends />
          </Link>
          <Friends />
          <Friends />
          <Friends />
          <Friends />
          <Friends />
          <Friends />
          <Friends />
        </List>
      </div>
      <div className='pl-10 w-8/12 flex justify-center items-center'>
        {userId ? (
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