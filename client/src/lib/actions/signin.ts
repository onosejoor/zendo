import { axiosInstance } from "@/api/api";
import { getErrorMesage } from "@/lib/utils";

type ApiResponse = APIRes & {
  email_verified: boolean;
};

export async function signIn(formData: SigninFormData) {
  try {
    const { data } = await axiosInstance.post<ApiResponse>(`/auth/signin`, formData);

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
