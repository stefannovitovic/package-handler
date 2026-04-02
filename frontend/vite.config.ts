import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

const backendUrl =
  process.env.VITE_BACKEND_URL ?? "http://host.docker.internal:8080";

export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      "/calculate-packs": {
        target: backendUrl,
        changeOrigin: true,
      },
    },
  },
});