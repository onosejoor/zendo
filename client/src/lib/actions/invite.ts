import { axiosInstance } from "@/api/api";
import { mutateMember } from "./members";
import { toast } from "sonner";
import { getErrorMesage } from "../utils";

export async function removeInvite(email: string, teamId: ITeam["_id"]) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(
      `/teams/${teamId}/invites/${email}`
    );
    if (data.success) {
      mutateMember();
    }
    return data;
  } catch (error) {
    console.log("Error removing member invite: ", error);

    return {
      success: false,
      message: getErrorMesage(error),
    };
  }
}

export const handleRemoveInvite = async (email: string, teamId: string) => {
  const { success, message } = await removeInvite(email, teamId);
  const options = success ? "success" : "error";

  toast[options](message);
};
