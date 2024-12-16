import { defineConfig } from "vite";
import deno from "@deno/vite-plugin";
import viteReact from "@vitejs/plugin-react";
import { TanStackRouterVite } from "@tanstack/router-plugin/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    deno(),
    TanStackRouterVite(),
    viteReact(),
  ],
  build: {
    outDir: "../server/pb_public",
  },
});
