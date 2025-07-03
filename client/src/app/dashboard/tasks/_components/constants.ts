import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.extend(relativeTime);

export const formatDate = (dateString: Date) => {
  return dayjs(new Date(dateString)).fromNow();
};

//    export const formatDate = (dateString: string) => {
//     return dayjs(new Date(dateString)).fromNow();
//   };

export const getStatusColor = (status: ITask["status"]) => {
  switch (status) {
    case "completed":
      return "bg-green-500";
    case "in-progress":
      return "bg-blue-500";
    case "pending":
      return "bg-yellow-500";
    default:
      return "bg-gray-500";
  }
};
