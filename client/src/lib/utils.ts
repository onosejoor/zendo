import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

type Fields = SignUpFormData | SigninFormData;

export function validateFields(fields: Fields) {
  const keys = Object.keys(fields);

  const isEmptyFields = keys.some((key) => typeof fields[key as keyof Fields] === "string" && !fields[key as keyof Fields].trim());

  return { isEmptyFields };
}

export const SERVER_URl = process.env.SERVER_URL