<script>
    import Post from "../../Components/posts.svelte";
</script>

{#each {length: 3} as _, _}
    <Post/>
{/each}

