import axios, { isAxiosError, type AxiosRequestConfig } from "axios";
import { getCookie } from "@/lib/actions/cookie";

export const axiosInstance = axios.create({
  baseURL: "/api",
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

const refreshToken = async (useBackup = false) => {
  const baseURL = useBackup ? "/backup-api" : "/api"; // Use Next.js rewrite routes
  try {
    await axios.get(`${baseURL}/auth/refresh-token`, {
      withCredentials: true,
    });
    const accessToken = await getCookie("zendo_access_token");
    if (!accessToken) throw new Error("No access token received");
    return accessToken;
  } catch (error) {
    throw error;
  }
};

axiosInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config as AxiosRequestConfig & {
      _retry?: boolean;
      _retriedWithBackup?: boolean;
    };

    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
          .then(() => {
            return axiosInstance(originalRequest);
          })
          .catch((err) => Promise.reject(err));
      }

      originalRequest._retry = true;
      isRefreshing = true;

      try {
        const accessToken = await refreshToken(
          !!originalRequest._retriedWithBackup
        );

        processQueue(null, accessToken);
        return axiosInstance(originalRequest);
      } catch (refreshError) {
        if (isAxiosError(refreshError)) {
          console.error("Refresh token error:", refreshError.response?.data);
        } else {
          console.error("Unexpected refresh error:", refreshError);
        }
        processQueue(refreshError, null);
        if (typeof window !== "undefined") {
          window.location.href = "/auth/signin";
        }
        return Promise.reject(refreshError);
      } finally {
        isRefreshing = false;
      }
    }

    if (
      (error.response?.status === 503 || error.code === "ERR_NETWORK") &&
      !originalRequest._retriedWithBackup
    ) {
      console.warn(
        "Primary server unavailable. Retrying with backup server..."
      );
      originalRequest._retriedWithBackup = true;
      originalRequest.baseURL = "/backup-api";
      return axiosInstance(originalRequest);
    }

    return Promise.reject(error);
  }
);
