import { UserRegisterInfo,UserLoginInfo } from "../../types/User";
import END_POINTS from "../../constants/endpoints";
import axiosInstance from "../../middleWares/interceptors";
const UserAuth = () => {
    const signUp = async (userInfo: UserRegisterInfo) => {
      const response = await axiosInstance.post(END_POINTS.SIGN_UP, userInfo);
      return response.data;
    };
  
    const signIn = async (userInfo: UserLoginInfo) => {
      const response = await axiosInstance.post(END_POINTS.SIGN_IN, userInfo);
      return response.data;
    };
  
    return {
      signUp,
      signIn,
    };
  };
  
  export default UserAuth;
  