<script lang="ts">
  import { onMount } from "svelte";
  import type { Thread } from "../../../dataFormats/Thread";
  import Post from "../../Components/Post.svelte";

  /** @type {import('./$types').PageData} */
  export let data;

  let notFoundError = false

  let threads: Thread;

  onMount(async function () {
    try {
      const endpoint = `http://${window.location.hostname}:8000/api/thread/${data.id}`;
      const response = await fetch(endpoint);
      if (!response.ok) {
        throw new Error("Failed to fetch data");
      }
      threads = await response.json()
      console.log(threads);
    } catch (error) {
      console.error(error);
      notFoundError = true
    }
  })
</script>

<style>
    @import '../../../country-flags/flags.css';
</style>



{#if (notFoundError)}
  <div class="container h-full mx-auto mt-4 flex justify-center items-center">
    <div class="space-y-5 text-center flex flex-col items-center">
          <img src={`http://${window.location.hostname}:8000/404.jpg`} alt="Not Found" />
        <small>Nothing here man.</small>
  </div>
  </div>
{/if}

{#if (threads)}
  <main class="m-3">
    <div class="container">
      <div class="flex-1 space-y-3">
        <Post postType="thread" content={threads} threadID={threads.ID} isOpen={true}/>
          <div class="w-full grid grid-cols-1 md:grid-cols-1 gap-2">
            {#each threads.Posts as post, index}
              <Post postType="post" content={post} threadID={threads.ID}/>
            {/each}
          </div>
      </div>
    </div>
  </main>
{:else}
  <main class="m-3">
    <div class="container">
      <div class="flex-1 space-y-3">
    </div>
    </div>
  </main>
{/if}
