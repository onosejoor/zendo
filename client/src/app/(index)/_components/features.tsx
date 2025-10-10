import {
  Bell,
  CheckSquare,
  FolderKanban,
  Rocket,
  Shield,
  Users,
} from "lucide-react";
import TitleHeader from "./title-header";

const teamFeatures = [
  {
    title: "Task Management",
    text: "Create, organize, and track personal or team tasks with deadlines and status indicators to stay productive every day.",
    icon: CheckSquare,
  },
  {
    title: "Projects",
    text: "Group related tasks into projects for better organization. Monitor overall progress and manage deliverables in one view.",
    icon: FolderKanban,
  },
  {
    title: "Teams & Collaboration",
    text: "Collaborate with others by creating teams, inviting members, assigning roles, and working together on shared goals.",
    icon: Users,
  },
  {
    title: "Reminders & Notifications",
    text: "Get timely email for tasks and deadlines to ensure you never miss important updates.",
    icon: Bell,
  },
  //   {
  //     title: "Daily & Weekly Summaries",
  //     text: "Receive automatic summaries of your tasks and projects, helping you stay informed and plan ahead effortlessly.",
  //     icon: "Calendar",
  //   },
  {
    title: "Secure Authentication",
    text: "Your account is protected with token-based authentication, and secure session handling.",
    icon: Shield,
  },
  {
    title: "Fast & Reliable Performance",
    text: "Powered by optimized APIs and Redis caching, Zendo ensures quick responses and smooth navigation across the app.",
    icon: Rocket,
  },
];

export default function FeaturesSection() {
  return (
    <section className="bg-zinc-50/50 px-4 space-y-7.5 py-7.5">
      <TitleHeader
        title="Features We Offer"
        subtitle="What makes us your best task managemnt platform"
      />

      <div className="grid lg:grid-cols-3 md:grid-cols-2 grid-cols-1 container mx-auto">
        {teamFeatures.map((feature, idx) => (
          <div
            key={idx}
            className="p-6 group m-4 border border-gray-200 rounded-lg hover:shadow-lg transition-shadow"
          >
            <feature.icon className="h-10 w-10 group-hover:animate-pulse text-blue-600 mb-4" />
            <h3 className="text-xl font-semibold mb-2">{feature.title}</h3>
            <p className="text-muted-foreground">{feature.text}</p>
          </div>
        ))}
      </div>
    </section>
  );
}
