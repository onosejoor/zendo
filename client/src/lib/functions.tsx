import { Badge } from "@/components/ui/badge";
import { checkExpired } from "./utils";
import dayjs from "dayjs";

export const getStatusBadge = (
  status: ITask["status"],
  dueDate: ITask["dueDate"]
) => {
  let statusMap: {
    label: string;
    className: string;
  };

  switch (status) {
    case "completed":
      statusMap = { label: "Completed", className: "text-green-500" };
      break;
    case "in-progress":
      statusMap = { label: "In Progress", className: "text-accent-blue" };
      break;
    case "pending":
      statusMap = { label: "Not Started", className: "text-orange-500" };
    default:
      statusMap = { label: "Not Started", className: "text-orange-500" };
      break;
  }

  if (checkExpired(dueDate) && status !== "completed") {
    statusMap = { label: "Expired", className: "text-gray-600" };
  }

  return (
    <Badge className={`${statusMap.className} bg-accent border-0 font-medium`}>
      {statusMap.label}
    </Badge>
  );
};

export function getTextNewLength({ id, value }: { id: string; value: string }) {
  if (value.length > 70 && (id === "title" || id === "subtask")) {
    return { value: value.slice(0, 70), isLong: true };
  }

  if (value.length > 300 && id === "description") {
    return { value: value.slice(0, 300), isLong: true };
  }
  return { value, isLong: false };
}

export const generateId = () => {
  const chars = "abcdefghijklmnopqrstuvwxyz1234567890";

  let id = "";

  while (id.length < 16) {
    const randomStr = chars[Math.floor(Math.random() * chars.length)];

    id += randomStr;
  }

  return id;
};

export function getLocalISOString() {
  const now = dayjs();
  return now.format("YYYY-MM-DDTHH:mm");
}
