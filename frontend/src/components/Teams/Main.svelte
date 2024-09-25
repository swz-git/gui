<script>
    import TeamBotList from "./TeamBotList.svelte";

    export let blueBots = [];
    export let orangeBots = [];
</script>

<div class="teams">
    <div class="team box">
        <header class="blue">
            <h3>Blue team</h3>
            <div style="flex: 1;"></div>
            <h3 class="dimmed">{blueBots?.length || 0} bots</h3>
        </header>
        <TeamBotList bind:items={blueBots} />
    </div>
    <div class="team box">
        <header class="orange">
            <h3>Orange team</h3>
            <div style="flex: 1;"></div>
            <h3 class="dimmed">{orangeBots?.length || 0} bots</h3>
        </header>
        <TeamBotList bind:items={orangeBots} />
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
        color: white;
    }
    header.blue {
        background-color: #0054a6;
    }
    header.orange {
        background-color: #f26522;
    }
    .dimmed {
        color: #ffffffcc;
    }
</style>
