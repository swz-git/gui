<script lang="ts">
    // @ts-ignore
    import {
        App,
        BotInfo,
        HumanInfo,
        PsyonixBotInfo,
        type StartMatchOptions,
    } from "../../bindings/gui/index.js";
    /** @import * from '../../bindings/gui' */
    import toast from "svelte-french-toast";
    // @ts-ignore
    import arenaImages from "../arena-images.ts";
    import rlbotMono from "../assets/rlbot_mono.png";
    import closeIcon from "../assets/close.svg";
    import reloadIcon from "../assets/reload.svg";
    import BotList from "../components/BotList.svelte";
    // @ts-ignore
    import Teams from "../components/Teams/Main.svelte";
    // @ts-ignore
    import MatchSettings from "../components/MatchSettings/Main.svelte";
    import { type DraggablePlayer, draggablePlayerToPlayerJs } from "../index";

    const backgroundImage =
        arenaImages[Math.floor(Math.random() * arenaImages.length)];
    // const backgroundImage = arenaImages.find((x) =>
    //     x.includes("Mannfield_Stormy"),
    // );

    let paths = JSON.parse(
        window.localStorage.getItem("BOT_SEARCH_PATHS") || "[]",
    );

    const BASE_PLAYERS: DraggablePlayer[] = [
        {
            displayName: "Human",
            icon: "None",
            id: Math.random(),
            player: new HumanInfo(),
        },
        {
            displayName: "Psyonix Beginner",
            icon: "None",
            id: Math.random(),
            player: new PsyonixBotInfo({
                skill: 0,
            }),
        },
        {
            displayName: "Psyonix Rookie",
            icon: "None",
            id: Math.random(),
            player: new PsyonixBotInfo({
                skill: 1,
            }),
        },
        {
            displayName: "Psyonix Pro",
            icon: "None",
            id: Math.random(),
            player: new PsyonixBotInfo({
                skill: 2,
            }),
        },
        {
            displayName: "Psyonix Allstar",
            icon: "None",
            id: Math.random(),
            player: new PsyonixBotInfo({
                skill: 3,
            }),
        },
    ];

    let players: DraggablePlayer[] = [...BASE_PLAYERS];

    let loadingPlayers = false;
    let latestBotUpdateTime = null;
    async function updateBots() {
        loadingPlayers = true;
        let internalUpdateTime = new Date();
        latestBotUpdateTime = internalUpdateTime;
        const result = await App.GetBots(paths);
        if (latestBotUpdateTime !== internalUpdateTime) {
            return; // if newer "search" already started, dont write old data
        }
        players = result.map((x: BotInfo) => {
            // @ts-ignore
            const n: typeof DraggablePlayer = {
                displayName: x.config.settings.name,
                icon: x.config.settings.logoFile,
                player: new BotInfo(x),
                id: Math.random(),
            };
            return n;
        });
        players = [...BASE_PLAYERS, ...players];
        loadingPlayers = false;
        console.log("Loaded bots:", result);
    }
    // this closure will get called if paths updates
    $: {
        paths;
        window.localStorage.setItem("BOT_SEARCH_PATHS", JSON.stringify(paths));
        updateBots();
    }

    let bluePlayers: DraggablePlayer[] = [];
    let orangePlayers: DraggablePlayer[] = [];

    let map: any;
    let mode: any;
    let mutatorSettings: any;

    async function onMatchStart() {
        let options: StartMatchOptions = {
            map,
            gameMode: mode,
            bluePlayers: bluePlayers.map((x: DraggablePlayer) => {
                // @ts-ignore
                return draggablePlayerToPlayerJs(x);
            }),
            orangePlayers: orangePlayers.map((x: DraggablePlayer) => {
                // @ts-ignore
                return draggablePlayerToPlayerJs(x);
            }),
            mutatorSettings,
        };

        toast("Starting match...", {
            position: "bottom-right",
        });

        let response = await App.StartMatch(options);

        if (response.success) {
            toast.success("Sent start match command", {
                position: "bottom-right",
                duration: 5000,
            });
        } else {
            toast.error(`Match start failed\n${response.message}`, {
                position: "bottom-right",
                duration: 5000,
            });
        }
    }

    async function onMatchStop() {
        toast("Stopping match...", {
            position: "bottom-right",
        });
        let response = await App.StopMatch(false);

        if (response.success) {
            toast.success("Sent stop match command", {
                position: "bottom-right",
                duration: 5000,
            });
        } else {
            toast.error(`Match stop failed\n${response.message}`, {
                position: "bottom-right",
                duration: 5000,
            });
        }
    }
</script>

<div class="page" style={`background-image: url("${backgroundImage}")`}>
    <div class="avaliableBots box">
        <header>
            <h1>Bots</h1>
            <div class="dropdown">
                <button>Add/Remove</button>
                <div class="dropmenu">
                    {#each paths as path, i}
                        <div class="path">
                            <pre>{path}</pre>
                            <button
                                class="close"
                                on:click={() => {
                                    paths.splice(i, 1);
                                    // makes reactivity work
                                    paths = paths;
                                }}
                            >
                                <img src={closeIcon} alt="X" />
                            </button>
                        </div>
                    {/each}
                    <button
                        on:click={async () => {
                            let result = await App.PickFolder();
                            console.log("PickFolder returned:", result);
                            if (result != "") {
                                paths = [...paths, result];
                            }
                        }}>Add folder</button
                    >
                </div>
            </div>
            {#if loadingPlayers}
                <h3>Searching...</h3>
            {:else}
                <button class="reloadButton" on:click={updateBots}
                    ><img src={reloadIcon} alt="reload" /></button
                >
            {/if}
            <div style="flex:1"></div>
            <input type="text" class="botSearch" placeholder="Search..." />
        </header>
        <BotList items={players} />
    </div>

    <div><Teams bind:bluePlayers bind:orangePlayers /></div>

    <div class="box">
        <MatchSettings
            onStart={onMatchStart}
            onStop={onMatchStop}
            bind:map
            bind:mode
            bind:mutators={mutatorSettings}
        />
    </div>
</div>

<style>
    .page {
        padding: 1rem;
        height: 100%;
        display: flex;
        flex-direction: column;
        overflow: auto;
        background-size: cover;
        background-repeat: no-repeat;
        background-position: center;
        background-attachment: fixed;
    }
    .page * {
        user-select: none;
        -webkit-user-select: none;
    }
    .box {
        border-radius: 0.4rem;
        background-color: var(--background);
        padding: 0.6rem;
    }
    .page > div:not(:first-child) {
        margin-top: 1rem;
    }
    .avaliableBots {
        padding-bottom: 0.6rem;
    }
    .avaliableBots header {
        display: flex;
        align-items: center;
        gap: 1rem;
        margin-bottom: 0.6rem;
    }
    .reloadButton {
        padding: 0px;
    }
    .reloadButton img {
        filter: invert();
    }
    .path {
        display: flex;
        align-items: center;
        gap: 1rem;
        justify-content: space-between;
    }
    .path pre {
        font-size: 1rem;
        margin: 0px;
    }
    .path button {
        padding: 0px;
    }
    .path button img {
        filter: invert();
    }
</style>
