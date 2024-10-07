<script lang="ts">
    import Select from "../NiceSelect.svelte";
    import Modal from "../Modal.svelte";
    import maps from "../../arena-names.js";
    import { mutators as mutatorOptions } from "./rlmutators";
    import type { ExtraOptions } from "../../../bindings/gui";

    export let map = localStorage.getItem("MS_MAP") || maps.DFHStadium;
    $: {
        localStorage.setItem("MS_MAP", map);
    }

    const modes = [
        "Soccer",
        "Hoops",
        "Dropshot",
        "Hockey",
        "Rumble",
        "Heatseeker",
        "Gridiron",
    ];
    export let mode = localStorage.getItem("MS_MODE") || "Soccer";
    $: {
        localStorage.setItem("MS_MODE", mode);
    }

    const existingMatchBehaviors: { [n: string]: number } = {
        "Restart if different": 0,
        Restart: 1,
        "Continue and spawn": 2,
    };

    let showExtraOptions = false;
    export let extraOptions: ExtraOptions = JSON.parse(
        localStorage.getItem("MS_EXTRAOPTIONS") || "{}",
    );
    $: {
        localStorage.setItem("MS_EXTRAOPTIONS", JSON.stringify(extraOptions));
    }

    let showMutators = false;
    export let mutators = JSON.parse(
        localStorage.getItem("MS_MUTATORS") || "{}",
    );
    $: {
        localStorage.setItem("MS_MUTATORS", JSON.stringify(mutators));
    }

    function cleanCamelCase(toClean: string) {
        return toClean
            .replace(/[A-Z]/g, (letter) => ` ${letter.toLowerCase()}`)
            .replace(/^[a-z]/, (letter) => letter.toUpperCase());
    }

    export let onStart: () => any = () => {};
    export let onStop: () => any = () => {};
</script>

<div class="matchSettings">
    <h1>Match Settings</h1>
    <div class="content">
        <div class="settings">
            <Select options={maps} bind:value={map} placeholder="Select map" />
            <Select
                options={Object.fromEntries(modes.map((x) => [x, x]))}
                bind:value={mode}
                placeholder="Select mode"
            />
            <button
                on:click={() => {
                    showMutators = true;
                }}>Mutators</button
            >
            <button
                on:click={() => {
                    showExtraOptions = true;
                }}>Extra</button
            >
        </div>
        <div class="controls">
            <button class="start" on:click={onStart()}>Start</button>
            <button class="stop" on:click={onStop()}>Stop</button>
        </div>
    </div>
</div>

<Modal title="Rocket League Mutators" bind:visible={showMutators}>
    <div class="mutators">
        {#each Object.keys(mutatorOptions) as mutatorKey}
            <div class="mutator">
                <label
                    style={mutators[mutatorKey] == 0 ? "color:grey" : ""}
                    for={mutatorKey}>{cleanCamelCase(mutatorKey)}</label
                >

                <select
                    name={mutatorKey}
                    id={mutatorKey}
                    bind:value={mutators[mutatorKey]}
                >
                    {#each mutatorOptions[mutatorKey] as value, i}
                        <option value={i}>{value.replaceAll("_", " ")}</option>
                    {/each}
                </select>
            </div>
        {/each}
    </div>
    <div class="bottomButtons">
        <p>Settings are saved automatically</p>
        <button
            class="mutatorResetButton"
            on:click={() => {
                for (let key of Object.keys(mutators)) {
                    mutators[key] = 0;
                }
            }}>Reset</button
        >
    </div>
</Modal>

<Modal title="RLBot Extra Options" bind:visible={showExtraOptions}>
    <div class="extraoptions">
        <input
            type="checkbox"
            id="enableRendering"
            bind:checked={extraOptions.enableRendering}
        />
        <label for="enableRendering">
            Enable Rendering (bots can draw on screen)
        </label>
        <br />
        <input
            type="checkbox"
            id="enableStateSetting"
            bind:checked={extraOptions.enableStateSetting}
        />
        <label for="enableStateSetting">
            Enable State Setting (bots can teleport)
        </label>
        <br />
        <input
            type="checkbox"
            id="autoSaveReplay"
            bind:checked={extraOptions.autoSaveReplay}
        />
        <label for="autoSaveReplay"> Auto Save Replay </label>
        <br />
        <input
            type="checkbox"
            id="skipReplays"
            bind:checked={extraOptions.skipReplays}
        />
        <label for="skipReplays"> Skip Replays </label>
        <br />
        <input
            type="checkbox"
            id="instantStart"
            bind:checked={extraOptions.instantStart}
        />
        <label for="instantStart"> Instant Start </label>
        <br />
        <select
            name="cock"
            id="emb"
            bind:value={extraOptions.existingMatchBehavior}
        >
            {#each Object.keys(existingMatchBehaviors) as key}
                <option value={existingMatchBehaviors[key]}>{key}</option>
            {/each}
        </select>
        <label for="emb"></label>
    </div>
    <div class="bottomButtons">
        <p>Settings are saved automatically</p>
    </div>
</Modal>

<style>
    h1 {
        margin-bottom: 0.6rem;
    }
    .settings,
    .controls {
        display: flex;
        gap: 0.5rem;
    }
    .content {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    .mutators {
        display: grid;
        grid-template-columns: auto auto auto;
        gap: 1rem;
    }
    .mutator {
        display: flex;
        flex-direction: column;
    }
    .bottomButtons {
        display: flex;
        margin-top: 1rem;
        justify-content: space-between;
        align-items: end;
    }
    .mutatorResetButton {
        background-color: red;
    }

    button.start {
        background-color: #15680e;
    }
    button.stop {
        background-color: #cc1414;
    }
</style>
