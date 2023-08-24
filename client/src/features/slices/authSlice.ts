import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { RootState } from "../store";

const access_token = localStorage.getItem("access_token");
const refresh_token = localStorage.getItem("refresh_token");

const initialState = {
  data: {
    access_token,
    refresh_token,
  },
  isLoggedIn: access_token !== null ? true : false,
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setToken(
      state,
      action: PayloadAction<{ access_token: string; refresh_token: string }>
    ) {
      localStorage.setItem(
        "access_token",
        JSON.stringify({
          access_token: action.payload.access_token,
        })
      );
      localStorage.setItem(
        "refresh_token",
        JSON.stringify({
          refresh_token: action.payload.refresh_token,
        })
      );
      state.data = {
        access_token: action.payload.access_token,
        refresh_token: action.payload.refresh_token,
      };
      state.isLoggedIn = true;
    },
    clearToken(state) {
      state.data = {
        access_token: "",
        refresh_token: "",
      };
      localStorage.removeItem("access_token");
      localStorage.removeItem("refresh_token");
      state.isLoggedIn = false;
    },
  },
});

export const { setToken, clearToken } = authSlice.actions;

export const selectAuth = (state: RootState) => state.auth.data;

export const selectAccessToken = () => {
  const accessTokenString = localStorage.getItem("access_token")
  const token =accessTokenString&&JSON.parse(accessTokenString);
  return token.access_token
};

export const selectIsLoggedIn = () => {
  const accessToken = localStorage.getItem("access_token");
  return accessToken ? true : false;
};

export const authReducer = authSlice.reducer;
