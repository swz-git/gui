import { defineConfig } from "@rsbuild/core";
import { pluginSvelte } from "@rsbuild/plugin-svelte";

export default defineConfig({
  plugins: [pluginSvelte()],
  dev: {
    // the go code fails to compile if no
    // files are present in the dist folder
    writeToDisk: true,
  },
});