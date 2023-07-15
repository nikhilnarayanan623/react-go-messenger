import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { RouterProvider } from "react-router-dom";
import AppRouter from "./routes.tsx";
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from "react-toastify";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={AppRouter} />
    <ToastContainer />
  </React.StrictMode>
);
