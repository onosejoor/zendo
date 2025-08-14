import { axiosInstance } from "@/api/api";
import { getErrorMesage } from "../utils";
import { mutate } from "swr";

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

    if (data.success) {
      mutate("/auth/user")
    }
    return data;
    
  } catch (error) {
    console.log("SEND NEW TOKEN TO EMAIL ERROR: ", error);
    
    return { success: false, message: getErrorMesage(error) };
  }
}

export async function sendOauthCode(code:string) {
  try {
    const { data } = await axiosInstance.post<APIRes>(
      `/auth/oauth/exchange?code=${code}`
    );

    if (data.success) {
      mutate("/auth/user")
    }
    return data;
    
  } catch (error) {
    console.log("SEND OAUTH CODE ERROR: ", error);
    
    return { success: false, message: getErrorMesage(error) };
  }
}