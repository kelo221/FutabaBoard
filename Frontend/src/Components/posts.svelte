<script>
    import {Convert} from "../dataFormats/post"
    import sampleData from "/src/sampleData/samplePost.json"
    const postData = Convert.toPostData(JSON.stringify(sampleData));

    console.log(postData);

    import { Card } from "flowbite-svelte";
    import { Gallery } from 'flowbite-svelte';
</script>



<Gallery class="gap-2 grid">
    <Card img="Sample/cow.jpg" padding="none" class="bg-slate-200" style="opacity: 0.9">
        <p class="mb-3 text-xs text-center text-gray-700 dark:text-gray-400 leading-tight">
            1664934427285917.jpg  (69 KB, 1125x1014)
        </p>
        <h5 class="m-2 text font-bold tracking-tight text-gray-900 dark:text-white">
            Anonymous 06/22/23(Thu)20:11:52 No.183608416
        </h5>
        <p class="m-2 font-normal text-gray-700 dark:text-gray-400 leading-tight" style="min-height: 100px">
            I HECKING LOVE DIXIE
        </p>
    </Card>
</Gallery>
