import { object, string,ref } from "yup";
export const userLoginValidationSchema = object().shape({
    email: string().email("Invalid email").required("Email is required"),
    password: string().required("Password is required"),
  });


