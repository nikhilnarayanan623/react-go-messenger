export interface UserLoginInfo {
    email:string;
    password:string;
}

export interface UserRegisterInfo {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
    confirmPassword: string;
    mobile: string;
    interests: string[];
  }