import { SERVER_URl } from "@/lib/utils";
import axios from "axios";

export async function signIn(formData: SigninFormData) {
  try {
    const { data } = await axios.post<APIRes>(
      `${SERVER_URl}/auth/signin`,
      formData,
      {
        withCredentials: true,
      }
    );

    return { ...data };
  } catch (error) {
    console.error("Error signing up: ", error);

    if (axios.isAxiosError(error)) {
      return { success: false, message: error.response?.data.message || error.response?.data };
    } else {
      return { success: false, message: "Internal error" };
    }
  }
}
