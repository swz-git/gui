<script lang="ts">
    import { GetBots } from "../../wailsjs/go/main/App.js";
    import arenaImages from "../arena-images.js";
    import rlbotMono from "../assets/rlbot_mono.png";
    import BotList from "../components/BotList.svelte";
    import Teams from "../components/Teams/Main.svelte";
    import MatchSettings from "../components/MatchSettings.svelte";
    import type { DraggableBotInfo } from "../index.js";

    let count = 0;
    // const backgroundImage =
    //     arenaImages[Math.floor(Math.random() * arenaImages.length)];
    const backgroundImage = arenaImages.find((x) =>
        x.includes("Mannfield_Stormy"),
    );

    let paths = [];
    let bots: DraggableBotInfo[] = [];
    let loadingBots = false;
    async function updateBots() {
        loadingBots = true;
        const result = await GetBots(paths);
        bots = result.map((x) => {
            const n: DraggableBotInfo = {
                ...x,
                id: Math.random(),
            };
            return n;
        });
        loadingBots = false;
        console.log("Loaded bots:", result);
    }
    $: {
        paths;
        updateBots();
    }

    console.log(bots);

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
                        <p>{path}</p>
                    {/each}
                </div>
            </div>
            <div style="flex:1"></div>
            <input type="text" class="botSearch" placeholder="Search..." />
        </header>
        {#if loadingBots}
            <h3>loading...</h3>
        {/if}
        <BotList items={bots} />
    </div>

    <div><Teams /></div>

    <div class="box"><MatchSettings /></div>
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
</style>
