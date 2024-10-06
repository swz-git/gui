import "./global.css";
import App from "./App.svelte";
import type { BotInfo } from "../bindings/gui";

const app = new App({
  target: document.body,
  // props: {
  //   name: "world",
  // },
});

export interface DraggableBotInfo extends BotInfo {
  id: number;
}

export default app;
