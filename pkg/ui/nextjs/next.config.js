/** @type {import('next').NextConfig} */
const nextConfig = {
    distDir: "build",
    output: 'export',
    experimental: {
        cpus: 4,
        workerThreads: false,
        appDir: false,
    }
};

module.exports = nextConfig;
