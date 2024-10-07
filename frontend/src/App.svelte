<script lang="ts">
    import Home from "./pages/Home.svelte";
    import logo from "./assets/rlbot_logo.svg";
    import { Toaster } from "svelte-french-toast";

    const activePage = "home";
</script>

<Toaster />

<main>
    <div class="navbar">
        <div>
            <img class="logo" src={logo} alt="logo" />
            <h1>RLBot</h1>
        </div>
        <div class="navbuttons">
            <button on:click={alert.bind(null, "TODO: not implemented yet")}
                >Events</button
            >
            <button on:click={alert.bind(null, "TODO: not implemented yet")}
                >State Setting Sandbox</button
            >
            <button on:click={alert.bind(null, "TODO: not implemented yet")}
                >Story Mode</button
            >
            <div class="spacer"></div>
            <div class="dropdown">
                <button>Menu</button>
                <div class="dropmenu right">
                    <button
                        on:click={alert.bind(null, "TODO: not implemented yet")}
                        >Repair botpack</button
                    >
                </div>
            </div>
        </div>
    </div>

    {#if activePage == "home"}
        <Home />
    {/if}
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
    }
    .navbar {
        display: flex;
        height: 3rem;
        justify-content: space-between;
        padding: 0.1rem;
        /* background: var(--background-alt);
        color: var(--foreground-alt); */
    }
    .navbar > div {
        display: flex;
        align-items: center;
    }
    h1 {
        margin: 0px;
        margin-bottom: 0.2rem;
    }
    .logo {
        height: 3rem;
        margin-right: 0.2rem;
        padding: 0.3rem;
    }
    .navbuttons > * {
        margin: 0px 0.25rem;
    }
    .navbar .dropmenu {
        padding: 0.2rem;
    }
    .navbar .dropmenu > * {
        margin: 0.2rem;
    }
</style>
