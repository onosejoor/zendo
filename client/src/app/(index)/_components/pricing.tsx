'use client'

import { Crown, Cuboid } from "lucide-react";
import TitleHeader from "./title-header";
import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { toast } from "sonner";

export const pricingPlans = [
  {
    name: "Free",
    price: "$0",
    period: "montly",
    description:
      "For individuals and small users who want to stay organized with core features.",
    features: [
      "Create up to 3 projects",
      "Maximum of 50 tasks per project",
      "Basic task reminders",
      "Create Up to 3 teams",
      "Invite up to 10 team members",
      "Access from any device",
      "Email support only",
    ],
    icon: Cuboid,
    buttonLabel: "Get Started",
  },
  {
    name: "Pro",
    price: "$1",
    period: "montly",
    description:
      "For teams and professionals who need advanced productivity and collaboration tools.",
    features: [
      "Unlimited projects and tasks",
      "Unlimited team members",
      "Advanced reminders and recurring tasks",
      "Team collaboration and roles",
      "weekly summary emails",
      "Calendar integrations (Google, Outlook)",
      "3rd-party app connections (Slack, Notion, etc.)",
      "Priority customer support",
    ],
    icon: Crown,
    buttonLabel: "Upgrade to Pro",
  },
];

export default function PricingSection() {
  return (
    <section className="bg-zinc-50/50 px-4 space-y-7.5 py-7.5">
      <TitleHeader
        title="Pricing"
        subtitle="Our product prices, and what each bill offers"
      />

      <div className="flex md:flex-row flex-col gap-7.5 justify-center container mx-auto">
        {pricingPlans.map((price, idx) => (
          <Card key={idx} className="bg-white max-w-sm w-full">
            <CardContent className="flex flex-col h-full space-y-5">
              <div className="space-y-5">
                <price.icon className="size-7.5 text-black mb-4" />

                <div className="flex justify-between items-center">
                  <h3 className="text-xl font-semibold mb-2">{price.price}</h3>
                  <Badge variant="default">{price.period}</Badge>
                </div>
              </div>

              <hr className="border-muted-foreground/50" />

              <p className="text-muted-foreground">{price.description}</p>

              <hr className="border-muted-foreground/50" />

              <ul className="space-y-2.5 list-inside list-disc">
                {price.features.map((feature, idx) => (
                  <li key={idx} className="text-foreground">
                    {feature}
                  </li>
                ))}
              </ul>

              {renderBtn(price.name as "Free" | "Pro", price.buttonLabel)}
            </CardContent>
          </Card>
        ))}
      </div>
    </section>
  );
}

function renderBtn(text: "Free" | "Pro", label: string) {
  switch (text) {
    case "Free":
      return (
        <Link href={"/auth/signin"} className="w-full mt-auto">
          <Button variant="default" size="lg" className="w-full mt-auto">
            {label}
          </Button>
        </Link>
      );

    case "Pro":
      return (
        <Button
          onClick={() => toast.error("Not available yet...")}
          variant="default"
          size="lg"
          className="w-full mt-auto"
        >
          {label}
        </Button>
      );
    default:
      break;
  }
}
