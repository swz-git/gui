<script lang="ts">
    import close from "../assets/close.svg";

    export let title = "Modal";
    export let visible = true;
</script>

<div class={visible ? "" : "hidden"}>
    <div class="modalContainer">
        <div class="modal">
            <header>
                <h2>{title}</h2>
                <button
                    on:click={() => {
                        visible = false;
                    }}><img src={close} alt="close" /></button
                >
            </header>
            <div class="modalBody">
                <slot />
            </div>
        </div>
    </div>
</div>

<style>
    .modalContainer {
        display: flex;
        height: 100vh;
        width: 100vw;
        position: fixed;
        top: 0px;
        left: 0px;
        justify-content: center;
        align-items: center;
        z-index: 100;
        background-color: #00000080;
        transition: opacity 0.2s;
    }

    .hidden * {
        opacity: 0;
        z-index: -100 !important;
    }

    .modal {
        background-color: var(--background);
        padding: 0.2rem;
        border-radius: 0.6rem;
        min-width: 20vw;
        min-height: 20vh;
    }
    header {
        padding: 0.2rem;
        padding-left: 0.8rem;
        border-bottom: 1px solid var(--background-alt);
        display: flex;
        justify-content: space-between;
    }
    header button {
        padding: 0px;
        background-color: transparent;
        filter: invert();
    }
    .modalBody {
        padding: 1rem;
        max-width: 80vw;
        max-height: 80vh;
    }
</style>
