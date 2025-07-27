import { axiosInstance } from "@/api/api";
import { isAxiosError } from "axios";
import { clsx, type ClassValue } from "clsx";
import dayjs from "dayjs";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

type Fields = SignUpFormData | SigninFormData;

export function validateFields(fields: Fields) {
  const keys = Object.keys(fields);

  const isEmptyFields = keys.some(
    (key) =>
      typeof fields[key as keyof Fields] === "string" &&
      !fields[key as keyof Fields].trim()
  );

  return { isEmptyFields };
}

export const SERVER_URl = "/api";

export const fetcher = async (url: string) =>
  axiosInstance.get(url).then((res) => res.data);

export const checkExpired = (date: ITask["dueDate"]) => {
  const dueDate = dayjs(date);
  return dueDate.isBefore(dayjs());
};

export function getErrorMesage(error: unknown) {
  if (isAxiosError(error)) {
    return error.response?.data.message || error.response?.data;
  }
  return error instanceof Error ? error.message : "Internal Error";
}

export function containsOnly(letter: string, str?: string) {
  return str && str.length > 0 && [...str].every((c) => c === letter);
}
