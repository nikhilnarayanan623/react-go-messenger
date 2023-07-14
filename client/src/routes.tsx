import { createBrowserRouter } from "react-router-dom";
import UserLogin from "./components/auth/UserLogin";
import App from "./App";
import ErrorElement from "./components/common/ErrorElement";

const AppRouter = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorElement />,
    children: [],
  },
  {
    path: "/login",
    element: <UserLogin />,
  }
]);
export default AppRouter;
