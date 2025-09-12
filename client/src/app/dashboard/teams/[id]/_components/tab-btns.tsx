"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { useRouter } from "next/navigation";

const btnArray = ["tasks", "members", "overview"];

export default function TabBtns({ section = "members" }: { section: string }) {
  const router = useRouter();

  const handleAddQuery = (text: string) => {
    router.replace(`?section=${text}`);
    return;
  };
  return (
    <Card className="bg-transparent w-fit shadow-none py-2.5">
      <CardContent>
        <div className="flex space-x-5">
          {btnArray.map((text) => {
            const isActiveTab = section === text;
            return (
              <Button
                onClick={() => handleAddQuery(text)}
                variant={isActiveTab ? "default" : "secondary"}
                key={text}
                className="capitalize"
              >
                {text}
              </Button>
            );
          })}
        </div>
      </CardContent>
    </Card>
  );
}
