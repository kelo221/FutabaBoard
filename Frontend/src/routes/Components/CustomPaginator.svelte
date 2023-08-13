<script lang="ts">
  import { onMount } from "svelte";
  import { createEventDispatcher } from 'svelte';

  // Create an event dispatcher
  const dispatch = createEventDispatcher();

  export let currentPage;
  let totalPages = 0; // Total number of pages in your content

  onMount(async function () {
    try {
      const endpoint = `http://${window.location.hostname}:8000/api/pageCount`;
      const response = await fetch(endpoint);
      if (!response.ok) {
        throw new Error("Failed to fetch data");
      }
      const tempData = await response.json();
      totalPages = tempData
      console.log(totalPages);
    } catch (error) {
      console.error(error);
    }
  });

  function gotoPage(page) {
    if (page !== currentPage) {
      dispatch("countUpdated", page); // Emit an event with updated count
      currentPage = page;
    }
  }

</script>

<style>


    button:hover {
        background-color: #ada5a5;
    }

    button.selected {
        background-color: rgba(var(--color-secondary-900) / 1);
        color: rgba(var(--color-secondary-900) / 1);
        font-weight: bold;
    }
</style>

<div class="card p-1">
  <div class="btn-group variant-filled">
    <ul>
      {#each Array.from({ length: totalPages + 1 }) as _, index}
        <button
          class:selected={index === currentPage}
          on:click={() => gotoPage(index)}>
          {index}
        </button>
      {/each}
    </ul>
  </div>
</div>