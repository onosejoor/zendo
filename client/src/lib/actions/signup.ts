import { axiosInstance } from "@/api/api";
import { getErrorMesage } from "@/lib/utils";

export async function signup(formData: SignUpFormData) {
  try {
    const { data } = await axiosInstance.post<
      APIRes & {
        email_sent: boolean;
      }
    >(`/auth/signup`, formData);

    return data;
  } catch (error) {
    console.error("Error signing up: ", error);
    return {
      success: false,
      message: getErrorMesage(error),
      email_sent: false,
    };
  }
}
