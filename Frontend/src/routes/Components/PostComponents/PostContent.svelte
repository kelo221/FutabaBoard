<script lang="ts">

  import { currentThreadStore, postPreview } from "../../../stores";
  import { PUBLIC_BACKEND_ADDRESS } from "$env/static/public";

  let enlargedImage: boolean;
  export let content;

  export let threadID;

  export let isOpen: boolean;

  const handleImageClick = imageHash => {
    enlargedImage = enlargedImage === imageHash ? null : imageHash;
  };


  const getPostPreview = async (postID: number) => {
    try {
      const endpoint = `${PUBLIC_BACKEND_ADDRESS}/api/post/${postID}`;
      const response = await fetch(endpoint);
      if (!response.ok) {
        throw new Error("Failed to fetch data");
      }
      return await response.json();
    } catch (error) {
      console.error(error);
    }
  };

  const handleMouseIn = async (event) => {
    if (event.target.getAttribute("href") !== null) {
      const href = Number(event.target.getAttribute("href").replace(/\D/g, ""));
      if (Number.isInteger(href)) {
        if (isOpen) {
          console.log("local");
          let foundObjects = $currentThreadStore.Posts.filter(obj => obj.ID === href);
          $postPreview.open = true;
          if (foundObjects[0]) {
            $postPreview.postData = foundObjects[0];
          } else {
            $postPreview.postData = $currentThreadStore;
          }
        } else {
          console.log("fetch");
          console.log(href);
          const fetch = await getPostPreview(href);
          if (fetch) {

            if (content.ParentThread) {
              $postPreview.tempThreadID = content.ParentThread;
            } else {
              $postPreview.tempThreadID = content.ID;
            }


            $postPreview.postData = fetch;
            $postPreview.open = true;
          }
        }
      }
    }
  };

  const handleMouseOut = () => {
    $postPreview.open = false;
  };


</script>

<style>
    .quote {
        color: green;
    }

    .reply {
        color: dodgerblue;
        font-weight: bold;
    }
</style>

<div class="flex" style="text-align: left; flex-grow: 1;">
  {#if !enlargedImage}
    <div class="container m-4" style="text-align: left; flex-grow: 1;">
      {#each content.TextRaw.split(/\r?\n/) as text}
        {#if text.startsWith(">>") && text.slice(2).trim().match(/^\d+$/)}
          <a class="reply" href={"#"+text.slice(2).trim()} on:mouseenter={handleMouseIn}
             on:mouseleave={handleMouseOut}>{text}</a>
        {:else if text.startsWith(">")}
          <p class="quote">{text}</p>
        {:else}
          <p>{text}</p>
        {/if}
      {/each}
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