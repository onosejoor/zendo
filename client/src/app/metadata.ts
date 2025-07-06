import { Metadata } from "next";

export const layoutMetadata: Metadata = {
  title: "Zendo — Task Management Made Simple",
  description:
    "Stay productive and organized with Zendo. Manage your tasks and projects in one sleek dashboard.",
  keywords: [
    "task management",
    "project management",
    "Zendo",
    "Next.js app",
    "productivity tool",
  ],
  authors: [{ name: "Onos Ejoor", url: "https://onos-ejoor.vercel.app" }],
  creator: "Zendo",
  openGraph: {
    title: "Zendo — Task Management Made Simple",
    description:
      "An intuitive platform to manage tasks and projects for individuals and teams.",
    url: "https://zendo.vercel.app",
    siteName: "Zendo",
    type: "website",
    images: [
      {
        url: "/og.png",
        width: 1200,
        height: 630,
        alt: "Zendo Dashboard",
      },
    ],
  },
};
