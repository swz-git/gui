<script lang="ts">
    import { flip } from "svelte/animate";
    import {
        dndzone,
        TRIGGERS,
        SHADOW_ITEM_MARKER_PROPERTY_NAME,
    } from "svelte-dnd-action";
    import defaultIcon from "../assets/rlbot_mono.png";
    import type { DraggablePlayer } from "../index";
    export let items: DraggablePlayer[] = [];
    const flipDurationMs = 100;
    let shouldIgnoreDndEvents = false;
    function handleDndConsider(e: any) {
        const { trigger, id } = e.detail.info;
        if (trigger === TRIGGERS.DRAG_STARTED) {
            const idx = items.findIndex((item) => item.id === id);
            const newId = `${id}_copy_${Math.round(Math.random() * 100000)}`;
            // the line below was added in order to be compatible with version svelte-dnd-action 0.7.4 and above
            e.detail.items = e.detail.items.filter(
                (item: any) => !item[SHADOW_ITEM_MARKER_PROPERTY_NAME],
            );
            e.detail.items.splice(idx, 0, { ...items[idx], id: newId });
            items = e.detail.items;
            shouldIgnoreDndEvents = true;
        } else if (!shouldIgnoreDndEvents) {
            // with this uncommented, this accepts bots dragged in from
            // the team lists, we don't want that
            // items = e.detail.items;
        } else {
            items = [...items];
        }
    }
    function handleDndFinalize(e: any) {
        if (!shouldIgnoreDndEvents) {
            items = e.detail.items;
        } else {
            items = [...items];
            shouldIgnoreDndEvents = false;
        }
    }
</script>

<div
    class="bots"
    use:dndzone={{
        items,
        flipDurationMs,
        centreDraggedOnCursor: true,
        dropTargetStyle: {},
        dropTargetClasses: ["dropTarget"],
    }}
    on:consider={handleDndConsider}
    on:finalize={handleDndFinalize}
>
    {#each items as bot (bot.id)}
        <div class="bot" animate:flip={{ duration: flipDurationMs }}>
            <img src={bot.icon || defaultIcon} alt="icon" />
            <p>{bot.displayName}</p>
        </div>
    {/each}
</div>

<style>
    .bots {
        display: flex;
        gap: 0.5rem;
        flex-wrap: wrap;
    }
    .bot {
        display: flex;
        align-items: center;
        background-color: var(--background-alt);
        height: 2.25rem;
        padding: 0.2rem;
        padding-right: 0.6rem;
        gap: 0.5rem;
        border-radius: 0.2rem;
    }
    .bot img {
        height: 2rem;
        width: auto;
    }
</style>
