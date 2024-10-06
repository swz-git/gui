import "./global.css";
import App from "./App.svelte";
import {
  BotInfo,
  type PlayerJs,
  PsyonixBotInfo,
  HumanInfo,
} from "../bindings/gui";

const app = new App({
  target: document.body,
  // props: {
  //   name: "world",
  // },
});

export interface DraggablePlayer {
  id: number;
  displayName: string;
  icon: string;
  player: BotInfo | PsyonixBotInfo | HumanInfo;
}

export function draggablePlayerToPlayerJs(d: DraggablePlayer): PlayerJs {
  let sort = "";

  if (d.player instanceof BotInfo) {
    sort = "rlbot";
  }
  if (d.player instanceof PsyonixBotInfo) {
    sort = "psyonix";
  }
  if (d.player instanceof HumanInfo) {
    sort = "human";
  }

  return {
    sort: sort,
    player: d.player,
  };
}

export default app;
