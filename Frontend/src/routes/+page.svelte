<script lang="ts">
    import {onMount} from "svelte";
    import type {Root} from "../dataFormats/ThreadPreview";

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
            console.log(threadPreviews)
        } catch (error) {
            console.error(error);
        }
    });


    function parseDateStringToDate(dateString: string): Date {
        const [datePart, timePart] = dateString.split("T");
        const [timeWithoutMilliseconds, timezoneOffset] = timePart.split("+");
        const timeWithoutMillis = timeWithoutMilliseconds.slice(0, -4);
        const combinedDateTime = `${datePart}T${timeWithoutMillis}+${timezoneOffset}`;
        return new Date(combinedDateTime);
    }

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

<main>
    <div class="container h-full mx-auto mt-4 flex justify-center items-center">
        <div class="space-y-5 text-center flex flex-col items-center">

            {#each threadPreviews as thread, index}
                <div class="w-full text-token grid grid-cols-1 md:grid-cols-1">
                    <div class="card">
                        <div class="p-4 space-y-4">
                            <h3 class="h3" data-toc-ignore>{thread.Topic}</h3>
                            <article class="flex items-center">
                                <img src={"http://localhost:8000/ThreadContent/Sample/cow_s.jpg"}>
                                <p style="text-align: left; padding-left: 10px">
                                    {thread.Text}
                                </p>
                            </article>
                        </div>
                        <hr class="opacity-50"/>
                        <footer class="p-4 flex justify-start items-center space-x-4">
                            <div class="flex-auto flex justify-between items-center">
                                <b style={'color: #' + thread.Hash}>{thread.Name}</b>

                                <div class="flex space-x-2">
                                    <i class={'flag ' + thread.Country}></i>
                                    <img src={`http://${window.location.hostname}:8000/Flags/Regional/${thread.ExtraFlags}.png`}/>
                                </div>

                                <p>No. {thread.ID}</p>
                                <p>{thread.PostCount} Posts</p>
                                <p>{parseDateStringToDate(thread.UnixTime).toLocaleString()}</p>
                                <p style={'color: #' + thread.Hash}>
                                    {"#" + thread.Hash.toUpperCase()}
                                </p>

                                <button type="button" class="btn variant-filled">
                                    Open Thread
                                </button>
                            </div>
                        </footer>
                    </div>
                </div>

                <!-- Posts under each thread -->
                <div class="w-full text-token grid grid-cols-1 md:grid-cols-2 gap-4">
                    {#each thread.Posts as post}
                        <div class="card">
                            <div class="p-4 space-y-4">
                                <article>
                                    <p style="text-align: left; padding-left: 10px">
                                        {post.Text}
                                    </p>
                                </article>
                            </div>
                            <hr class="opacity-50"/>
                            <footer class="p-4 flex justify-start items-center space-x-4">
                                <div class="flex-auto flex justify-between items-center">
                                    <b style={'color: #' + post.Hash}>{post.Name}</b>

                                    <div class="flex space-x-2">
                                        <i class={'flag ' + post.Country}></i>
                                        <img src={`http://${window.location.hostname}:8000/Flags/Regional/${post.ExtraFlags}.png`}/>
                                    </div>

                                    <p>No. {post.ID}</p>
                                    <p>{parseDateStringToDate(post.UnixTime).toLocaleString()}</p>
                                    <p style={'color: #' + post.Hash}>
                                        {"#" + post.Hash.toUpperCase()}
                                    </p>
                                </div>
                            </footer>
                        </div>
                    {/each}
                </div>
                <!-- Modified hr tag with full-width-hr class -->
                {#if index < threadPreviews.length - 1}
                    <hr class="full-width-hr"/>
                    <hr class="full-width-hr"/>
                {/if}
            {/each}
        </div>
    </div>
</main>