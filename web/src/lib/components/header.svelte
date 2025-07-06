<script>
    import { toggleMode } from "mode-watcher";
    import { Button } from "$lib/components/ui/button/index.js";
</script>

<header
    class="fixed top-0 left-0 right-0 z-50 w-full px-[24px] py-3 flex items-center justify-between border-b border-border backdrop-blur-2xl bg-[var(--background)/70] text-[var(--foreground)]"
    dir="ltr"
>
    <div class="text-lg font-semibold">SLOVO</div>

    <Button onclick={toggleMode} variant="outline" size="icon">
        <img
            src="/icons/sun.svg"
            alt="Sun"
            class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 !transition-all dark:-rotate-90 dark:scale-0"
        />

        <img
            src="/icons/moon.svg"
            alt="Moon"
            class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 !transition-all dark:rotate-0 dark:scale-100 dark:invert"
        />
        <span class="sr-only">Toggle theme</span>
    </Button>
</header>
