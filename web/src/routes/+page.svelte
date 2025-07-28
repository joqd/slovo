<script lang="ts">
    import { goto } from "$app/navigation";
    import { buttonVariants } from "$lib/components/ui/button";
    import { SquareArrowOutUpLeft } from "lucide-svelte";
    import { BinocularsIcon } from "lucide-svelte";

    const words = [
        {
            _id: "6835a2db5a859aff5197007a",
            bare: "весь",
            accented: "весь",
            level: "B1",
            type: "noun",
            translations: [
                { en: "all", fa: "همه" },
                {
                    en: "all gone (no more left)",
                    fa: "همه رفته‌اند (دیگر چیزی نمانده است)",
                },
                { en: "whole, entire", fa: "تمام، سراسر" },
            ],
        },
        {
            _id: "6835a2db5a859aff5197007a",
            bare: "весь",
            accented: "весь",
            level: "B1",
            type: "adverb",
            translations: [{ en: "all", fa: "همه" }],
        },
        {
            _id: "6835a2db5a859aff5197007a",
            bare: "весь",
            accented: "весь",
            level: "B1",
            type: "adverb",
            translations: [
                {
                    en: "all gone (no more left)",
                    fa: null,
                },
                { en: "whole, entire", fa: "تمام، سراسر" },
            ],
        },
    ];

    let visibleCount = 20;
    const increment = 10;

    function showMore() {
        visibleCount += increment;
    }

    function goToWord(id: string) {
        goto(`words/${id}`);
    }
</script>

<!-- intro -->
<div class="w-full space-y-2" dir="rtl">
    <div class="text-2xl mx-auto w-max">جستجو و کاوش کلمات</div>
    <div class="text-sm text-center">
        جستجو از بین بیش از ۹۰ هزار کلمه انجام خواهد شد، در صورت عدم وجود کلمه
        لطفا به <a href="/report" class={buttonVariants({ variant: "link" })}
            >این صفحه مراجعه کنید <SquareArrowOutUpLeft
                class="inline h-[11px] w-[11px]"
            /></a
        >
    </div>
</div>

<!-- search bar -->
<div class="w-full mt-10">
    <div class="relative">
        <BinocularsIcon
            class="absolute left-6 top-1/2 -translate-y-1/2 w-5 h-5 text-primary opacity-50 pointer-events-none"
        />

        <input
            type="text"
            dir="auto"
            placeholder="Search..."
            class="w-full rounded-md placeholder:text-primary placeholder:italic py-3 pl-16 pr-4 focus:outline-none transition focus:placeholder:opacity-50 placeholder:opacity-80 bg-secondary"
        />
    </div>
</div>

<!-- search results -->
<div class="mt-5 space-y-5 px-1">
    {#if words.length > 0}
        {#each words.slice(0, visibleCount) as word}
            <div>
                <div class="flex justify-between">
                    <button on:click={() => goToWord(word._id)}>
                        <p
                            class="hover:underline cursor-pointer text-lg text-blue-500"
                        >
                            {word.accented}
                        </p>
                    </button>
                    <div class="opacity-50 text-sm">
                        {word.type}
                    </div>
                </div>
                <div class="opacity-70 text-sm">
                    {#each word.translations as tr}
                        <div class="flex items-center gap-2">
                            <div>{tr.en}</div>
                            <div class="flex-grow border-b border-dotted"></div>
                            <div>{tr.fa ?? "-"}</div>
                        </div>
                    {/each}
                </div>
            </div>
        {/each}

        {#if visibleCount < words.length}
            <div class="text-center mt-4">
                <button
                    on:click={showMore}
                    class={buttonVariants({ variant: "outline" })}
                >
                    نمایش بیشتر
                </button>
            </div>
        {/if}
    {:else}
        <div class="text-center opacity-80">
            هیچ داده ای برای نمایش یافت نشد
        </div>
    {/if}
</div>
