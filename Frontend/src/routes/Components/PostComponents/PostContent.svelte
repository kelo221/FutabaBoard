<script lang="ts">

  import { currentThreadStore, postPreview } from "../../../stores";
  import { onMount } from "svelte";

  let enlargedImage : boolean
  export let content

  export let threadID

  export let isOpen : boolean

  const handleImageClick = imageHash => {
    enlargedImage = enlargedImage === imageHash ? null : imageHash;
  };

  onMount(async () => {
    //content.Text = processReplyMatches();
    attachMouseOver();
  });


  const getPostPreview= async (postID: number) => {
    try {
      const endpoint = `http://${window.location.hostname}:8000/api/post/${postID}`;
      const response = await fetch(endpoint);
      if (!response.ok) {
        throw new Error("Failed to fetch data");
      }
      return await response.json()
    } catch (error) {
      console.error(error);
    }
  }

  // TODO, work on replies outside of current thread
  function attachMouseOver() {
    const links = document.querySelectorAll("a");
    links.forEach(link => {

      link.addEventListener("mouseover", async (event) => {
        if (event.target.getAttribute("href") !== null) {
          const href = Number(event.target.getAttribute("href").replace(/\D/g, ""));
          if (Number.isInteger(href)) {
            if (isOpen) {
              let foundObjects = $currentThreadStore.Posts.filter(obj => obj.ID === href);
              $postPreview.open = true;
              if (foundObjects[0]) {
                $postPreview.postData = foundObjects[0];
              } else {
                $postPreview.postData = $currentThreadStore;
              }
            } else {
              const fetch = await getPostPreview(href);
              if (fetch) {
                console.log(threadID);
                $postPreview.postData = fetch;
                $postPreview.open = true;
              }
            }
          }
        }
      });

      link.addEventListener("mouseleave", () => {
        $postPreview.open = false;
      });

    });
  }


</script>

<div class="flex" style="text-align: left; flex-grow: 1;">
  {#if !enlargedImage}
    <div class="container m-4" style="text-align: left; flex-grow: 1;">
      {@html content.Text}
    </div>
  {/if}

  {#if content.PostImage.ImageHash }
    <div class="container flex flex-col items-center justify-center m-4">
      <img
        src={`http://localhost:8000/ThreadContent/${threadID}/${content.PostImage.ImageHash}${
                                enlargedImage === content.PostImage.ImageHash ? content.PostImage.ImageInfo.split(" ").pop() : '_small.png'}`}
        alt="Post Image"
        on:click={() => handleImageClick(content.PostImage.ImageHash)}
      />
      {#if !enlargedImage}
        <small>{content.PostImage.ImageInfo}</small>
      {/if}
    </div>
  {/if}
</div>