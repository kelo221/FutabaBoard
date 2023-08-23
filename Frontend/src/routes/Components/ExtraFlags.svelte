<script lang="ts">

  export let Flags;


  const flagSplitter = (inputFlag: string): string[] => {
    const parts = inputFlag.split('/');
    const resultArray: string[] = [];

    let currentPart = parts[0];
    for (let i = 1; i < parts.length; i++) {
      currentPart += '/' + parts[i];
      resultArray.push(currentPart);
    }

    return resultArray;
  };

</script>

{#if Flags !== ""}
  {#each flagSplitter(Flags) as flag}
  <img alt={flag} src={`http://127.0.0.1:8000/Flags/Regional/${flag}.png`}/>
    {/each}
{/if}
