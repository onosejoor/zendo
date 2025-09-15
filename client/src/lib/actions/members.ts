import { axiosInstance } from "@/api/api";
import { getErrorMesage } from "../utils";
import { mutate } from "swr";
import { toast } from "sonner";

type CreateProps = {
  email: string;
  role: string;
};

export async function sendMemberInvite(payload: CreateProps, teamId: string) {
  try {
    const { data } = await axiosInstance.post<APIRes>(
      `/teams/${teamId}/members/invite`,
      payload
    );
    return data;
  } catch (error) {
    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export async function removeMember(
  memberId: IMember["_id"],
  teamId: ITeam["_id"]
) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(
      `/teams/${teamId}/members/${memberId}`
    );
    if (data.success) {
      mutateMember();
    }
    return data;
  } catch (error) {
    console.log("Error removing member: ", error);

    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export const handleRemoveMember = async (
  memberId: IProject["_id"],
  teamId: string
) => {
  const { success, message } = await removeMember(memberId, teamId);
  const options = success ? "success" : "error";

  toast[options](message);

  if (success) {
    mutateMember();
  }
};

export function mutateMember() {
  mutate((key) => typeof key === "string" && key.startsWith("/teams"));
}
