<script lang="ts">
import ExtraFlags from "../ExtraFlags.svelte";
import { replyBoxStore } from "../../../stores";
import { onMount } from "svelte";
import Icon from "@iconify/svelte";

export let content
export let threadID
export let isOpen


function parseDateStringToDate(datetimeString: string): string {
  const date = new Date(datetimeString);
  const options: Intl.DateTimeFormatOptions = {
    year: "2-digit",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: false
  };

  return (
    date.toLocaleDateString("en-US", options).replace(",", "")
  );
}

const openReply = (threadID: number, postType: string, targetPost: string) => {
  $replyBoxStore.threadID = threadID;
  $replyBoxStore.open = true;

  if (postType === "post") {
    $replyBoxStore.content = $replyBoxStore.content + ">>" + targetPost + "\n";
  }

};


onMount(async function () {

  if (window.innerWidth <= 768) {
    isMobile = true;
  }

});

let isMobile = false;


</script>

<footer class="p-4 flex justify-start items-center space-x-4" style="margin-top: auto;">
  <div class="flex-auto flex justify-between items-center">
    <div class="grid">
      <div class="flex justify-center">
        <small style={'color: #' + content.UserHash}>{content.Name}</small>
      </div>
      <div class="flex space-x-1 justify-center items-center">
        <i class={'flag ' + content.Country}></i>
        <ExtraFlags Flags={content.ExtraFlags} />
      </div>
    </div>
    <small on:click={() => openReply(threadID, "post" ,content.ID) }
       style="color: blue; text-decoration: underline; cursor: pointer;">#{content.ID}</small>
    <small>{content.Topic ? content.PostCount + ' Posts' : ''}</small>
    <small>{parseDateStringToDate(content.UnixTime)}</small>
    <small style={'color: #' + content.UserHash}>
      {"#" + content.UserHash}
    </small>

  </div>

  {#if content.Topic && isOpen === false}
    <a href={`thread/${content.ID}`}>
      {#if isMobile}
        <Icon icon="ooui:link-external-ltr" />
        {:else}
      <button class="btn variant-filled" type="button">
        Open Thread
      </button>
        {/if}
    </a>
  {/if}
</footer>