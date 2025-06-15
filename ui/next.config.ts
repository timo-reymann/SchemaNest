import {NextConfig} from "next";

const nextConfig: NextConfig = {
    distDir: "build",
    output: 'export',
    experimental: {
        cpus: 4,
        workerThreads: false,
    },
    eslint: {
        ignoreDuringBuilds: true,
    },
};

export default nextConfig;
