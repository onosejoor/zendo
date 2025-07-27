import { getErrorMesage, SERVER_URl } from "@/lib/utils";

import axios from "axios";

export async function signup(formData: SignUpFormData) {
  try {
    const { data } = await axios.post<
      APIRes & {
        email_sent: boolean;
      }
    >(`${SERVER_URl}/auth/signup`, formData, {
      withCredentials: true,
    });

    return data;
  } catch (error) {
    console.error("Error signing up: ", error);
    return { success: false, message: getErrorMesage(error), email_sent: false };
  }
}
