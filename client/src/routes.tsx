import { createBrowserRouter } from "react-router-dom";
import UserLogin from "./components/auth/UserLogin";
import UserRegister from "./components/auth/UserRegister";
import App from "./App";
import ErrorElement from "./components/common/ErrorElement";
import Chats from "./components/chat/Chats";
import SendMessages from "./components/chat/SendMessages";

const AppRouter = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorElement />,
  },
  {
    path: "/sign-in",
    element: <UserLogin />,
  },{
    path:"/sign-up",
    element:<UserRegister/>
  },
  {
    path:"/chats",
    element:<Chats/>,
    children:[
      {
        path:"chat/:chatId/message/:receiverId",
        element:<SendMessages/>
      }
    ]
  }
]);
export default AppRouter;
