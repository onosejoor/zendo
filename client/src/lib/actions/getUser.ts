import { SERVER_URl } from "@/lib/utils";
import axios from "axios";

export async function getUser() {
  try {
    const { data } = await axios.get<{
      success: boolean;
      user: IUser;
    }>(`${SERVER_URl}/auth/user`, {
      withCredentials: true,
    });

    return { ...data };
  } catch (error) {
    console.error("Error getting user ", error);

    if (axios.isAxiosError(error)) {
      return {
        success: false,
        message: error.response?.data.message || error.response?.data,
        user: null,
      };
    } else {
      return { success: false, message: "Internal error", user: null };
    }
  }
}
