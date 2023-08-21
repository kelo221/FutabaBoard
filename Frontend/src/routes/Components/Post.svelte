<script lang="ts">
  import PostContent from "./PostComponents/PostContent.svelte";
  import PostFooter from "./PostComponents/PostFooter.svelte";

  export let content;
  export let threadID;
  export let isOpen;



  /*const processReplyMatches = () => {
    const pattern = /href="#(\d+)"/;
    const matches = content.Text.match(pattern);

    if (matches) {
      const parser = new DOMParser();
      const doc = parser.parseFromString(content.Text, "text/html");

      const aTags = doc.querySelectorAll("a");
      aTags.forEach((aTag) => {
          aTag.setAttribute("on:mouseenter", `{()=>
            console.log("enter")}
          `
          );
      });

      content.Text = doc.body.innerHTML;
    } else {
      console.log("No number found.");
    }

    return content.Text;
  };*/


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
