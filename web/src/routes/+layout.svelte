<script>
    let { children } = $props();

    import "../styles/pico.min.css";
    import "../styles/pico.colors.min.css";
    import "../styles/general.css";
</script>

<div
    class="vazirmatn min-h-screen bg-[url(/images/bg1.jpg)] bg-cover bg-center sm:pt-14 sm:pb-7 sm:px-1"
>
    <div class="container-fluid sm:min-h-full min-h-screen max-w-[960px] mx-auto rounded-[2px] py-[19px]">
        <header></header>

        <main class="flex-grow">
            {@render children()}
        </main>

        <footer></footer>
    </div>
</div>
