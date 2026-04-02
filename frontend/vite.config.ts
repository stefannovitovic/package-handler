import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      "/calculate-packs": {
        target: "http://host.docker.internal:8080",
        changeOrigin: true,
      },
    },
  },
});
