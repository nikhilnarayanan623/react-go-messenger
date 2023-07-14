import axios from "axios";
import { UserRegisterInfo } from "../../types/User";
import END_POINTS from "../../constants/endpoints";

const auth = () => {
  const registerUser = async (userInfo: UserRegisterInfo) => {
    const response = await axios.post(END_POINTS.SIGNUP, userInfo);
    return response.data;
  };
  return {
    registerUser,
  };
};
export default auth;
