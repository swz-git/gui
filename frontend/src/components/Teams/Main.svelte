<script>
    import TeamBotList from "./TeamBotList.svelte";

    // biome-ignore lint/style/useConst: its bound
    export let blue_bots = [];
    // biome-ignore lint/style/useConst: its bound
    export let orange_bots = [];
</script>

<div class="teams">
    <div class="team box">
        <header class="blue">
            <h3>Blue team</h3>
            <div style="flex: 1;"></div>
            <h3 class="dimmed">{blue_bots?.length || 0} bots</h3>
        </header>
        <TeamBotList bind:items={blue_bots} />
    </div>
    <div class="team box">
        <header class="orange">
            <h3>Orange team</h3>
            <div style="flex: 1;"></div>
            <h3 class="dimmed">{orange_bots?.length || 0} bots</h3>
        </header>
        <TeamBotList bind:items={orange_bots} />
    </div>
</div>

<style>
    .box {
        border-radius: 0.4rem;
        background-color: var(--background);
        padding: 0px 0.6rem;
    }
    .teams {
        display: flex;
        gap: 1rem;
    }
    .teams > .team {
        flex: 1;
        padding: 0px 0;
        /* Nice transparent blur, not supported yet though. Therefor commented */
        /* background-color: rgba(0, 0, 0, 0.5);
        -webkit-backdrop-filter: blur(10px);
        backdrop-filter: blur(10px); */
        display: flex;
        flex-direction: column;
    }
    .teams > .team > header {
        border-radius: 0.4rem 0.4rem 0px 0px;
        padding: 0px 0.6rem;
        text-transform: uppercase;
        display: flex;
    }
    header.blue {
        background-color: #0054a6;
    }
    header.orange {
        background-color: #f26522;
    }
    .dimmed {
        color: var(--foreground-alt);
    }
</style>
