import type { MetadataRoute } from "next";

export default function manifest(): MetadataRoute.Manifest {
  return {
    name: "Zendo",
    short_name: "Zendo",
    description:
      "Zendo is a modern task management application to organize work, track progress, and increase productivity.",
    start_url: "/",
    display: "standalone",
    background_color: "#ffffff",
    theme_color: "#0f172a",
    // icons: [
    //   {
    //     src: "/icons/icon-192x192.png",
    //     sizes: "192x192",
    //     type: "image/png",
    //   },
    //   {
    //     src: "/icons/icon-512x512.png",
    //     sizes: "512x512",
    //     type: "image/png",
    //   },
    //   {
    //     src: "/icons/icon-512x512.png",
    //     sizes: "512x512",
    //     type: "image/png",
    //     purpose: "maskable",
    //   },
    // ],
    lang: "en",
    categories: ["productivity", "organization", "tasks"],
  };
}
