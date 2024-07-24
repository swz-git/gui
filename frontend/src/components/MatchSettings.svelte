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
        <button class="start">Start</button>
        <button class="stop" disabled>Stop</button>
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
