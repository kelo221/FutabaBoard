<script lang="ts">
  import { onMount } from 'svelte';
  import Icon from '@iconify/svelte';
  import { replyBoxStore } from '../../stores'
  const closeReplyWindow = () =>{
    $replyBoxStore.open = false
    $replyBoxStore.content = ""
  }

  let threadTopic = ""

  onMount(() => {
      const dragBar = document.getElementById("dragBar");
      const replyBox = document.querySelector(".replyBox");

      let isDragging = false;
      let offsetX = 0;
      let offsetY = 0;

      const handleDragStart = (event: MouseEvent) => {
        isDragging = true;
        offsetX = event.clientX - replyBox.getBoundingClientRect().left;
        offsetY = event.clientY - replyBox.getBoundingClientRect().top;
      };

      const handleDrag = (event: MouseEvent) => {
        if (isDragging) {
          const x = event.clientX - offsetX;
          const y = event.clientY - offsetY;
          replyBox.style.transform = `translate(${x}px, ${y}px)`;
        }
      };

      const handleDragEnd = () => {
        isDragging = false;
      };

      dragBar.addEventListener("mousedown", handleDragStart);
      window.addEventListener("mousemove", handleDrag);
      window.addEventListener("mouseup", handleDragEnd);

      return () => {
        dragBar.removeEventListener("mousedown", handleDragStart);
        window.removeEventListener("mousemove", handleDrag);
        window.removeEventListener("mouseup", handleDragEnd);
      };
  });
  const sendData = () => {

  };

</script>

<style>
    .floating-element {
        position: absolute;
        left: 0vw;
        top: 0vh;
        z-index: 1000;
        user-select: none;
    }

    .overLength {
        color: red;
    }

    .hidden{
        display: none;
    }
</style>



<div class="card floating-element replyBox" style="border: 4px outset rgba(var(--color-secondary-900) / 1);" class:hidden="{$replyBoxStore.open === false}">
  <div id="dragBar" style="border: 10px outset rgba(var(--color-secondary-300) / 1); height: 10px;"></div>
  <header class="card-header flex">

    {#if $replyBoxStore.newThread}
      <h4 style="font-size: large">Creating a new thread</h4>
    {:else }
      <h4 style="font-size: large">Replying to thread #{$replyBoxStore.threadID}</h4>
      {/if}
    <div class="ml-auto" on:click={closeReplyWindow} style="cursor: pointer">
      <Icon class="icon" icon="ooui:close" />
    </div>
  </header>
  <section class="p-4">
    {#if $replyBoxStore.newThread}
    <label class="label">
      <textarea bind:value={threadTopic} class="textarea" rows="1" placeholder="Topic (Required)"></textarea>
    </label>
      {/if}
    <label class="label">
      <textarea bind:value={$replyBoxStore.content} class="textarea" rows="4"></textarea>
    </label>
    <small class:overLength="{$replyBoxStore.content.length > 3000}">
      {$replyBoxStore.content.length}/3000
    </small>
  </section>
  <footer class="card-footer">
    <input class="input" type="file" />
  </footer>
  <button type="button" class="btn variant-filled m-2" on:click={()=> sendData()}>Reply</button>
</div>
