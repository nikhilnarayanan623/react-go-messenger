import axios, { AxiosInstance, AxiosResponse, AxiosError } from "axios";
import { API_BASE_URL } from "../../constants/endpoints";
import ApiError from "../../utils/ApiError";

const axiosInstance: AxiosInstance = axios.create({
  baseURL: API_BASE_URL,
});

axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error: AxiosError) => {
    if (error.response) {
      const { data, status } = error.response;
      console.log(data)
      switch (status) {
        case 400:
          throw new ApiError("Bad request", data);
        case 401:
          throw new ApiError("Unauthorized", data);
        case 404:
          throw new ApiError("Not Found", data);
        case 409:
          throw new ApiError("Conflict", data);
        default:
          throw new ApiError(`Request failed with status ${status}`, data);
      }
    } else if (error.request) {
      throw new ApiError(`No response received`, error.request);
    } else {
      console.log("Error:", error.message);
    }

    return Promise.reject(error);
  }
);

export default axiosInstance