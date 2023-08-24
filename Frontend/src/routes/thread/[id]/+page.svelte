<script lang="ts">
  import Post from "../../Components/Post.svelte";
  import { currentThreadStore, replyBoxStore } from "../../../stores";
  export let data;

  $currentThreadStore = data.thread;

  // Once a thread is opened, set the replyBox to reply to the current thread.
  $replyBoxStore.newThread = false
</script>

<style>
    @import '../../../country-flags/flags.css';
</style>

  <main class="m-3">
    <div class="container">
      <div class="flex-1 space-y-3">
          <Post content={$currentThreadStore} threadID={$currentThreadStore.ID} isOpen={true}/>
          <div class="w-full grid grid-cols-1 md:grid-cols-1 gap-2">
            {#each $currentThreadStore.Posts as post}
                <Post content={post} threadID={post.ParentThread} isOpen={true}/>
            {/each}
          </div>
      </div>
    </div>
  </main>
