<script lang="ts">
  import PostContent from "./PostComponents/PostContent.svelte";
  import PostFooter from "./PostComponents/PostFooter.svelte";

  export let content;
  export let threadID;
  export let isOpen;



</script>

<div class="card" id={content.ID} style="display: flex; flex-direction: column;">

  {#if content.Topic}
  <div class="container flex flex-col items-center justify-center m-2">
      <h1 style="font-size: x-large">{content.Topic}</h1>
  </div>
  {/if}


  <PostContent content={content} threadID={threadID} isOpen={isOpen}/>
  <PostFooter content={content} threadID={threadID} isOpen={isOpen}/>

  <hr class="opacity-50" />

</div>
