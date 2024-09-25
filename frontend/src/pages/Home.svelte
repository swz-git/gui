<script lang="ts">
    import {
        GetBots,
        PickFolder,
        StartMatch,
        StopMatch,
    } from "../../wailsjs/go/main/App.js";
    import type * as go from "../../wailsjs/go/models";
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
    import type { DraggableBotInfo } from "../index.js";

    let count = 0;
    // const backgroundImage =
    //     arenaImages[Math.floor(Math.random() * arenaImages.length)];
    const backgroundImage = arenaImages.find((x) =>
        x.includes("Mannfield_Stormy"),
    );

    let paths = JSON.parse(
        window.localStorage.getItem("BOT_SEARCH_PATHS") || "[]",
    );
    let bots: DraggableBotInfo[] = [];
    let loadingBots = false;
    let latestBotUpdateTime = null;
    async function updateBots() {
        loadingBots = true;
        let internalUpdateTime = new Date();
        latestBotUpdateTime = internalUpdateTime;
        const result = await GetBots(paths);
        if (latestBotUpdateTime !== internalUpdateTime) {
            return; // if newer "search" already started, dont write old data
        }
        bots = result.map((x) => {
            // @ts-ignore
            const n: DraggableBotInfo = {
                ...x,
                id: Math.random(),
            };
            return n;
        });
        loadingBots = false;
        console.log("Loaded bots:", result);
    }
    // this closure will get called if paths updates
    $: {
        paths;
        window.localStorage.setItem("BOT_SEARCH_PATHS", JSON.stringify(paths));
        updateBots();
    }

    let blueBots: any[] = [];
    let orangeBots: any[] = [];

    let map: any;
    let mode: any;
    let mutatorSettings: any;

    async function onMatchStart() {
        // @ts-ignore
        let options: go.main.StartMatchOptions = {
            map,
            gameMode: mode,
            blueBots,
            orangeBots,
            mutatorSettings,
        };
        console.log("starting with options", options);
        toast("Starting match...", {
            position: "bottom-right",
        });
        let response = await StartMatch(options);

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
        let response = await StopMatch(false);

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

    function handleClick() {
        count += 1;
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
                            let result = await PickFolder();
                            console.log("PickFolder returned:", result);
                            if (result != "") {
                                paths = [...paths, result];
                            }
                        }}>Add folder</button
                    >
                </div>
            </div>
            {#if loadingBots}
                <h3>Searching...</h3>
            {:else}
                <button class="reloadButton" on:click={updateBots}
                    ><img src={reloadIcon} alt="reload" /></button
                >
            {/if}
            <div style="flex:1"></div>
            <input type="text" class="botSearch" placeholder="Search..." />
        </header>
        <BotList items={bots} />
    </div>

    <div><Teams bind:blueBots bind:orangeBots /></div>

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
