import END_POINTS from "../../constants/endpoints";
import axiosInstance from "../middleWares/interceptors";
const UserChat = () => {
  const getRecentlyChattedFriends = async () => {
    const response = await axiosInstance.get(END_POINTS.GET_MY_CHATS);
    return response.data;
  };

  const getRecentMessages = async (chatId: string) => {
    const response = await axiosInstance.get(
      `${END_POINTS.GET_RECENT_MESSAGES}/${chatId}/messages`
    );
    return response.data;
  };

  const sendMessage = async (chatId: string, message: string) => {
    const response = await axiosInstance.post(
      `${END_POINTS.SEND_MESSAGE}/${chatId}/messages`,
      { content: message }
    );
    return response.data;
  };

  return {
    getRecentlyChattedFriends,
    getRecentMessages,
    sendMessage,
  };
};

export default UserChat;
