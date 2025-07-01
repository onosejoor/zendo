import axios, { isAxiosError } from "axios";

import { getCookie } from "@/lib/actions/cookie";
import { SERVER_URl } from "@/lib/utils";

export const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_SERVER_URL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

interface QueueItem {
  resolve: (value?: string | PromiseLike<string> | undefined | null) => void;
  reject: (reason?: any) => void;
}

let isRefreshing = false;
let failedQueue: QueueItem[] = [];

const processQueue = (
  error: any,
  token: string | undefined | null = undefined
) => {
  failedQueue.forEach((prom) => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });

  failedQueue = [];
};

axiosInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (
      error.response &&
      error.response.status === 401 &&
      !originalRequest?._retry
    ) {
      if (isRefreshing) {
        return new Promise(function (resolve, reject) {
          failedQueue.push({ resolve, reject });
        })
          .then(() => {
            return axiosInstance(originalRequest);
          })
          .catch((err) => {
            return Promise.reject(err);
          });
      }
      originalRequest._retry = true;
      isRefreshing = true;

      try {
        await axios.get(`${SERVER_URl}/auth/refresh-token`, {
          withCredentials: true,
        });

        const accessToken = await getCookie("zendo_access_token");
        processQueue(null, accessToken!);

        return axiosInstance(originalRequest);
      } catch (error) {
        if (isAxiosError(error)) {
          console.log(error.response?.data);
        }

        console.error("ERROR", error);
        processQueue(error, null);
        window.location.href = "/auth/signin";
        return Promise.reject(error);
      } finally {
        isRefreshing = false;
      }
    }

    return Promise.reject(error);
  }
);
