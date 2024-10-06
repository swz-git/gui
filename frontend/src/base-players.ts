import type { DraggablePlayer } from ".";
import { HumanInfo, PsyonixBotInfo } from "../bindings/gui";
import controller from "./assets/controller.svg";
import rlbot_mono from "./assets/rlbot_mono.png";

export const BASE_PLAYERS: DraggablePlayer[] = [
  {
    displayName: "Human",
    icon: controller,
    id: Math.random(),
    player: new HumanInfo(),
  },
  {
    displayName: "Psyonix Beginner",
    icon: "",
    id: Math.random(),
    player: new PsyonixBotInfo({
      skill: 0,
    }),
  },
  {
    displayName: "Psyonix Rookie",
    icon: "",
    id: Math.random(),
    player: new PsyonixBotInfo({
      skill: 1,
    }),
  },
  {
    displayName: "Psyonix Pro",
    icon: "",
    id: Math.random(),
    player: new PsyonixBotInfo({
      skill: 2,
    }),
  },
  {
    displayName: "Psyonix Allstar",
    icon: "",
    id: Math.random(),
    player: new PsyonixBotInfo({
      skill: 3,
    }),
  },
];
