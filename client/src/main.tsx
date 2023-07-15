import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { RouterProvider } from "react-router-dom";
import AppRouter from "./routes.tsx";
import "react-toastify/dist/ReactToastify.css";
import { ToastContainer } from "react-toastify";
import { GoogleOAuthProvider } from "@react-oauth/google";
import ENV from "./constants/const.ts";
import { ThemeProvider } from "@material-tailwind/react";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <GoogleOAuthProvider clientId={ENV.VITE_GOOGLE_CLIENT_ID}>
      <ThemeProvider>
      <RouterProvider router={AppRouter} />        
      <ToastContainer />
      </ThemeProvider>
    </GoogleOAuthProvider>
  </React.StrictMode>
);
