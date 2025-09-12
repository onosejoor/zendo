import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Verified } from "lucide-react";
import { redirect } from "next/navigation";

type Props = { redirectRoute: string; title: string; message: string };

export default function SuccessComp({ redirectRoute, title, message }: Props) {
  setTimeout(() => {
    redirect(redirectRoute);
  }, 1000);

  return (
    <Card className="max-w-md mx-auto my-5">
      <CardContent>
        <CardHeader className="grid gap-3">
          <CardTitle>
            <Verified className="size-5 text-blue-500" />
          </CardTitle>
          <CardTitle>{title}</CardTitle>
          <CardDescription>{message}</CardDescription>
        </CardHeader>
      </CardContent>
    </Card>
  );
}
