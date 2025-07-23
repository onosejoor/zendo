import { axiosInstance } from "@/api/api";
import { getErrorMesage } from "../utils";

export async function sendEmailToken(token: string) {
  try {
    const { data } = await axiosInstance.get<APIRes>(
      `/auth/verify_email?token=${token}`
    );

    return data;
  } catch (error) {
    console.log("SEND EMAIL ERROR: ", error);
    
    return { success: false, message: getErrorMesage(error) };
  }
}

export async function postToken() {
  try {
    const { data } = await axiosInstance.post<APIRes>(
      `/auth/verify_email`
    );

    return data;
  } catch (error) {
    console.log("SEND NEW TOKEN TO EMAIL ERROR: ", error);
    
    return { success: false, message: getErrorMesage(error) };
  }
}
