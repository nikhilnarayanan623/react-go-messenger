import { configureStore } from "@reduxjs/toolkit";
import { authReducer } from "./slices/authSlice";
import { chatReducer } from "./slices/chatSlice";

export const store = configureStore({
  reducer: {
    auth: authReducer,
    chat: chatReducer,
  },
});

export type State = typeof store;

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
