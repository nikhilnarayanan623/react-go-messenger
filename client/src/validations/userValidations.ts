import { object, string, number, ref } from "yup";
export const userLoginValidationSchema = object().shape({
  email: string().email("Invalid email").required("Email is required"),
  password: string().required("Password is required"),
});

export const userRegisterValidationSchema = object().shape({
  first_name: string().trim().required("First Name is required"),
  last_name: string().trim().required("Last Name is required"),
  user_name:string().trim().required("User Name is required"),
  email: string().email("Invalid email").trim().required("Email is required"),
  age: number()
    .typeError("Age must be a number")
    .required("Age is required")
    .min(13, "Minimum age must be 13")
    .max(100, "Maximum age is 100"),
  password: string().required("Password is required"),
  confirm_password: string()
    .oneOf([ref("password")], "Passwords must match")
    .required("Confirm Password is required"),
});
