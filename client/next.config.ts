import type { NextConfig } from "next";

const nextConfig: NextConfig = {

  images: {
    remotePatterns: [new URL("https://res.cloudinary.com/**")]
  },
  /* config options here */
  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: `${process.env.NEXT_PUBLIC_SERVER_URL}/:path*`,
      },
      {
        source: "/backup-api/:path*",
        destination: `${process.env.NEXT_PUBLIC_BACKUP_SERVER_URL}/:path*`,
      },
    ];
  },
};

export default nextConfig;
