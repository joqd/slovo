<script>
    import { Button } from "$lib/components/ui/button/index.js";
    import SunIcon from "@lucide/svelte/icons/sun";
    import MoonIcon from "@lucide/svelte/icons/moon";
    import { Menu } from "lucide-svelte";

    import { toggleMode } from "mode-watcher";
</script>

<div class="h-12 flex justify-between items-center">
    <div class="text-lg font-bold">SLOVO</div>

    <div>
        <Button onclick={toggleMode} variant="ghost" size="icon">
            <SunIcon
                class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 !transition-all dark:-rotate-90 dark:scale-0"
            />
            <MoonIcon
                class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 !transition-all dark:rotate-0 dark:scale-100"
            />
            <span class="sr-only">Toggle theme</span>
        </Button>
        <Button variant="ghost" size="icon">
            <Menu class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100" />
        </Button>
    </div>
</div>
