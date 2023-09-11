<script lang="ts">
  import Post from "./Components/Post.svelte";
  import CustomPaginator from "./Components/CustomPaginator.svelte";
  import { currentPageStore, replyBoxStore } from "../stores";
  import { PUBLIC_BACKEND_ADDRESS } from "$env/static/public";

  /** @type {import("./$types").PageData} */
  export let data;

  async function handleCountUpdated(event) {
    let newPage = event.detail;
    try {
      const endpoint = `${PUBLIC_BACKEND_ADDRESS}/api/page/${newPage}`;
      const response = await fetch(endpoint);
      if (!response.ok) {
        throw new Error("Failed to fetch data");
      }
      const tempData = await response.json();
      data.threadPreviews.length = 0;
      data.threadPreviews.push(...tempData);
      $currentPageStore = newPage;
    } catch (error) {
      console.error(error);
    }
  }

  const handleNewThread = () => {
    $replyBoxStore.open = true;
  };

  $replyBoxStore.newThread = true;

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
  <div class="container h-full mx-auto m-2 flex justify-center items-center">
    <button class="btn variant-filled" on:click={() => handleNewThread()} type="button">
      New Thread
    </button>
  </div>
  <div class="container">
    <div class="flex-1 space-y-3">

      {#if data.threadPreviews}
        {#each data.threadPreviews as thread, index}
          <Post content={thread} threadID={thread.ID} isOpen={false} />
          <div class="w-full grid grid-cols-1 md:grid-cols-2 gap-2">
            {#each thread.Posts as post}
              <Post content={post} threadID={post.ParentThread} isOpen={false} />
            {/each}
          </div>
          {#if index < data.threadPreviews.length - 1}
            <hr class="full-width-hr" />
            <hr class="full-width-hr" />
          {/if}
        {/each}

      {:else}
        <div>No threads.</div>
      {/if}

    </div>
  </div>
  <div class="container h-full mx-auto mt-4 flex justify-center items-center">
    <CustomPaginator currentPage={$currentPageStore} on:countUpdated={handleCountUpdated} totalPages={data.pageCount} />
  </div>
</main>
