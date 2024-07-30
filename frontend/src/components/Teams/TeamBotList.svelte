<script lang="ts">
    import { dndzone } from "svelte-dnd-action";
    import { flip } from "svelte/animate";
    import closeIcon from "../../assets/close.svg";
    import type { DraggableBotInfo } from "../..";
    const flipDurationMs = 100;
    function handleSort(e) {
        items = e.detail.items;
    }
    export let items: DraggableBotInfo[] = [];
    function remove(id) {
        items = items.filter((x) => x.id !== id);
    }
</script>

<div class="teamBotList">
    <p
        class="placeholder"
        style={items.length == 0
            ? "margin-top:.6rem;"
            : "opacity: 0;z-index: -999"}
    >
        Drop bots here...
    </p>
    <div
        class="bots"
        use:dndzone={{
            items,
            flipDurationMs,
            dropTargetStyle: {},
            dropTargetClasses: ["dropTarget"],
        }}
        on:consider={handleSort}
        on:finalize={handleSort}
    >
        {#each items as bot (bot.id)}
            <div class="bot" animate:flip={{ duration: flipDurationMs }}>
                <img src={bot?.config?.settings.logoFile} alt="icon" />
                <p>{bot?.config?.settings.name}</p>
                <div style="flex: 1;"></div>
                <button class="close" on:click={remove(bot.id)}>
                    <img src={closeIcon} alt="X" />
                </button>
            </div>
        {/each}
    </div>
</div>

<style>
    * {
        user-select: none;
        -webkit-user-select: none;
    }
    .teamBotList {
        padding: 0.6rem;
        overflow: auto;
        height: 100%;
        min-height: 4rem;
        display: flex;
        flex-direction: column;
        position: relative;
    }
    .placeholder {
        position: absolute;
        transition: 100ms;
    }
    .bots {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        min-height: 100%;
    }
    .bot {
        display: flex;
        align-items: center;
        background-color: var(--background-alt);
        height: 2rem;
        padding: 0.2rem;
        gap: 0.5rem;
        border-radius: 0.2rem;
    }
    .bot img {
        height: 1.8rem;
        width: auto;
    }
    .close {
        background-color: transparent;
        height: 100%;
        padding: 0;
    }
    .close img {
        height: 100%;
        width: auto;
        filter: invert();
    }
</style>
