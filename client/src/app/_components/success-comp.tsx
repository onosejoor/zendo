import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { getAndDeleteCookie } from "@/lib/actions/cookie";
import { Verified } from "lucide-react";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

type Props = { redirectRoute: string; title: string; message: string };

export default function SuccessComp({ redirectRoute, title, message }: Props) {
  const router = useRouter();

useEffect(() => {
  let id: NodeJS.Timeout;

  getAndDeleteCookie("zendo_redirect_url")
    .then((url) => {
      id = setTimeout(
        () => router.push(url ? decodeURIComponent(url) : redirectRoute),
        1000
      );
    })
    .catch(() => {
      id = setTimeout(() => router.push(redirectRoute), 1000);
    });

  return () => clearTimeout(id);
}, [redirectRoute, router]);


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
