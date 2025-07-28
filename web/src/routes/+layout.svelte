<script>
    import "../app.css";
    import { ModeWatcher } from "mode-watcher";
    import Header from "$lib/components/header.svelte";
    import Footer from "$lib/components/footer.svelte";

    let { children } = $props();
</script>

<div class="min-h-screen w-[512px] mx-auto sm:px-0 px-2">
    <header><Header /></header>
    <div class="sm:mt-10 mt-5">
        <ModeWatcher />
        {@render children()}
    </div>
    <!-- <footer><Footer /></footer> -->
</div>
