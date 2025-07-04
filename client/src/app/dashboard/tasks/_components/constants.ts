import { checkExpired } from "@/lib/utils";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.extend(relativeTime);

export const formatDate = (dateString: Date) => {
  return dayjs(new Date(dateString)).fromNow();
};

export const getStatusColor = (
  status: ITask["status"],
  dueDate: ITask["dueDate"]
) => {
  if (checkExpired(dueDate) && status !== "completed") {
    return "bg-gray-500";
  }
  switch (status) {
    case "completed":
      return "bg-green-500";
    case "in-progress":
      return "bg-accent-blue";
    case "pending":
      return "bg-yellow-500";
    default:
      return "bg-gray-500";
  }
};
