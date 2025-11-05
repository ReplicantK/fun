import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "export",
  distDir: "dist",
  basePath: "/qr",
  images: {
    unoptimized: true
  }
};

export default nextConfig;
