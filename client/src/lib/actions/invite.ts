import { axiosInstance } from "@/api/api";
import { mutateMember } from "./members";
import { toast } from "sonner";
import { getErrorMesage } from "../utils";

export async function removeInvite(id: string, teamId: ITeam["_id"]) {
  try {
    const { data } = await axiosInstance.delete<APIRes>(
      `/teams/${teamId}/invites/${id}`
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

export const handleRemoveInvite = async (id: string, teamId: string) => {
  const { success, message } = await removeInvite(id, teamId);
  const options = success ? "success" : "error";

  toast[options](message);
};
