"use client";

import Img from "@/components/Img";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { useHomeData } from "@/hooks/use-user";
import { ArrowRight, CloudLightning } from "lucide-react";
import Link from "next/link";

export default function HeroSection() {
  const { data, isLoading } = useHomeData();

  const arrayToMapFrom = !isLoading ? data?.data.avatars || [] : [...Array(5)];

  return (
    <header className="px-4 py-20 min-h-[90vh] bg-gray-50 bg-[url(https://www.transparenttextures.com/patterns/ag-square.png)]">
      <div className="container mx-auto text-center space-y-6">
        <Badge variant={"default"} className="bg-blue-500">
          <CloudLightning /> Task management made simple
        </Badge>
        <h1 className="text-2xl sm:!text-5xl md:!text-6xl font-bold text-gray-900 leading-[1]">
          Simplify how you{" "}
          <span className="text-blue-600">{"plan, track"}</span> <br />
          and <span className="text-blue-600">complete</span> your work with{" "}
          <span className="text-blue-600">Ease</span>
        </h1>
        <p className="text-gray-600 md:text-lg max-w-2xl mx-auto">
          Organize your projects, track your tasks, and collaborate with your
          team. Everything you need to stay productive in one place.
        </p>

        <div className="items-center justify-center flex sm:flex-row flex-col flex-wrap gap-5">
          <Link href="/auth/signup">
            <Button size="lg" className="rounded-full px-8 py-3">
              Start Free Trial
              <ArrowRight className="ml-2 h-5 w-5" />
            </Button>
          </Link>

          <div className="flex justify-center items-center text-sm font-medium ">
            {arrayToMapFrom.map((avatar, idx) => (
              <div className="-ml-2.5" key={idx}>
                <Img
                  src={
                    avatar ||
                    "https://res.cloudinary.com/dog3ihaqs/image/upload/v1740561527/c5bgr7ybdhuqlkcn6vuc.jpg"
                  }
                  alt="Image"
                  className="size-10 object-cover rounded-full border-2 border-blue-500"
                />
              </div>
            ))}
            {data && (
              <p className="pl-2.5">
                +{data.data.total_users - data.data.avatars.length} more!
              </p>
            )}
          </div>
        </div>
      </div>
    </header>
  );
}
