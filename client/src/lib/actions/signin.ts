import { getErrorMesage, SERVER_URl } from "@/lib/utils";
import axios from "axios";

type ApiResponse = APIRes & {
  email_verified: boolean;
};

export async function signIn(formData: SigninFormData) {
  try {
    const { data } = await axios.post<ApiResponse>(
      `${SERVER_URl}/auth/signin`,
      formData,
      {
        withCredentials: true,
      }
    );

    return { ...data };
  } catch (error) {
    console.error("Error signing up: ", error);
    return {
      success: false,
      message: getErrorMesage(error),
      email_verified: false,
    };
  }
}
