<script>
    import Header from "$lib/components/header.svelte";
    import Footer from "$lib/components/footer.svelte";
    import { ModeWatcher } from "mode-watcher";
    import "../app.css";

    let { children } = $props();
</script>

<div class="min-h-screen flex flex-col sohne" dir="rtl">
    <ModeWatcher />
    <Header />

    <div class="mt-[100px] flex-1 relative">
        <div class="max-w-[640px] mx-auto px-6 sm:px-0 min-h-full flex">
            <div class="w-full">{@render children()}</div>
        </div>
    </div>

    <Footer />
</div>
