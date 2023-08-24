<script lang="ts">
  import { createEventDispatcher} from "svelte";
  const dispatch = createEventDispatcher();

  export let currentPage;
  export let totalPages;


  const gotoPage = page => {
    if (page !== currentPage) {
      dispatch("countUpdated", page);
      currentPage = page;
    }
  };

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