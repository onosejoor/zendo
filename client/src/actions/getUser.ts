"use server";

import { SERVER_URl } from "@/lib/utils";
import axios from "axios";

export async function getUser() {
  try {
    const { data } = await axios.get<UserRes>(`${SERVER_URl}/auth/user`, {
      withCredentials: true,
    });

    return { ...data };
  } catch (error) {
    console.error("Error signing up: ", error);

    if (axios.isAxiosError(error)) {
      return {
        success: false,
        message: error.response?.data.message,
        user: null,
      };
    } else {
      return { success: false, message: "Internal error", user: null };
    }
  }
}
