<script>
    import Select from "./NiceSelect.svelte";
    import maps from "../arena-names.js";
    export let map = maps.DFHStadium;
    const modes = [
        "Soccer",
        "Hoops",
        "Dropshot",
        "Hockey",
        "Rumble",
        "Heatseeker",
        "Gridiron",
    ];
    export let mode = "Soccer";

    export let onStart = () => {};
    export let onStop = () => {};
</script>

<div class="matchSettings">
    <h1>Match Settings</h1>
    <div class="settings">
        <Select options={maps} bind:value={map} placeholder="Select map" />
        <Select
            options={Object.fromEntries(modes.map((x) => [x, x]))}
            bind:value={mode}
            placeholder="Select mode"
        />
        <button on:click={alert.bind(null, "TODO: not implemented yet")}
            >Mutators</button
        >
        <button on:click={alert.bind(null, "TODO: not implemented yet")}
            >Extra</button
        >
    </div>
    <div class="controls">
        <button class="start" on:click={onStart()}>Start</button>
        <button class="stop" on:click={onStop()}>Stop</button>
    </div>
</div>

<style>
    h1 {
        margin-bottom: 0.6rem;
    }
    .settings {
        display: flex;
    }

    button.start {
        background-color: darkgreen;
    }
    button.stop {
        background-color: darkred;
    }
</style>
