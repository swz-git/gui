import "./global.css";
import App from "./App.svelte";
import type { main } from "../wailsjs/go/models.js";

const app = new App({
  target: document.body,
  // props: {
  //   name: "world",
  // },
});

export interface DraggableBotInfo extends main.BotInfo {
  id: number;
}

export default app;
