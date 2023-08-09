<script lang="ts">
    import {onMount} from "svelte";
    import type {Root} from "../dataFormats/ThreadPreview";
    import Post from "./Components/Post.svelte";

    let threadPreviews: Root[] = [];

    onMount(async function () {
        try {
            const endpoint = `http://${window.location.hostname}:8000/api/page/0`;
            const response = await fetch(endpoint);
            if (!response.ok) {
                throw new Error("Failed to fetch data");
            }
            const data = await response.json();
            threadPreviews = data;
        } catch (error) {
            console.error(error);
        }
    })

</script>

<style>
    @import '../country-flags/flags.css';

    .full-width-hr {
        width: 100%;
        margin: 1rem 0; /* Add any additional margin if needed */
        border: none;
        border-top: 3px solid #ddd; /* Replace #ddd with your desired hr color */
    }

</style>

<main class="m-3">
    <div class="container">
        <div class="flex-1 space-y-3" id="HELLO">
            {#each threadPreviews as thread, index}
                <Post postType="thread" content={thread} threadID={thread.ID}/>
                <div class="w-full grid grid-cols-1 md:grid-cols-2 gap-2">
                    {#each thread.Posts as post}
                        <Post postType="post" content={post} threadID={thread.ID}/>
                    {/each}
                </div>
                {#if index < threadPreviews.length - 1}
                    <hr class="full-width-hr"/>
                    <hr class="full-width-hr"/>
                {/if}
            {/each}
        </div>
    </div>
</main>