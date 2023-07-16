import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "../store";
import { RecentlyChattedFriends } from "../../types/Chat";

interface InitialState {
    data: {
      currentChat: RecentlyChattedFriends|null
    };
  }
  const initialState: InitialState = {
    data: {
      currentChat:null
    }
  };
const chatSlice = createSlice({
  name: "chat",
  initialState,
  reducers: {
    setCurrentChat(
      state,
      action: PayloadAction<{currentChat:RecentlyChattedFriends }>
    ) {
      state.data = {
        currentChat:action.payload.currentChat
      };
    },
    clearCurrentChat(state) {
      state.data = {
        currentChat:null
      };
    },
  },
});

export const { setCurrentChat,clearCurrentChat} = chatSlice.actions;

export const selectCurrentChat = (state: RootState) => state.chat.data.currentChat


export const chatReducer = chatSlice.reducer