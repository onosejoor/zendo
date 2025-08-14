import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Verified } from "lucide-react";
import { redirect } from "next/navigation";

export default function SuccessComp() {
  
  setTimeout(() => {
    redirect("/dashboard");
  }, 1000);


  return (
    <Card className="max-w-md mx-auto my-5">
      <CardContent>
        <CardHeader className="grid gap-3">
          <CardTitle>
            <Verified className="size-5 text-blue-500" />
          </CardTitle>
          <CardTitle>Signin Successfully</CardTitle>
          <CardDescription>
            Oauth Signing Successfull, redirecting now
          </CardDescription>
        </CardHeader>
      </CardContent>
    </Card>
  );
}
