import END_POINTS from "../../constants/endpoints";
import axiosInstance from "../middleWares/interceptors";
const UserChat = () => {
  const getRecentlyChattedFriends = async () => {
    const response = await axiosInstance.get(END_POINTS.GET_MY_CHATS);
    return response.data;
  };

  const getRecentChats = async (chatId: number) => {
    const response = await axiosInstance.get(
      `${END_POINTS.GET_RECENT_MESSAGES}/${chatId}/messages`
    );
    return response.data;
  };

  return {
    getRecentlyChattedFriends,
    getRecentChats,
  };
};

export default UserChat;
