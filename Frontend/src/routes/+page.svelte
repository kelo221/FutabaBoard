<script lang="ts">
	import { onMount } from "svelte";
	import type { Root } from "../dataFormats/ThreadPreview";
	import {Avatar} from "@skeletonlabs/skeleton";

	let threadPreviews: Root[] = []; // Initialize with an empty array

	onMount(async function () {
		try {
			const endpoint = `http://${window.location.hostname}:8000/api/page/0`;
			const response = await fetch(endpoint);

			if (!response.ok) {
				throw new Error("Failed to fetch data");
			}

			const data = await response.json();
			threadPreviews = data; // Assign the JSON response directly to the array
			console.log(threadPreviews);
		} catch (error) {
			console.error(error);
		}
	});


	function parseDateStringToDate(dateString: string): Date {
		const [datePart, timePart] = dateString.split("T");
		const [timeWithoutMilliseconds, timezoneOffset] = timePart.split("+");
		const timeWithoutMillis = timeWithoutMilliseconds.slice(0, -4);
		const combinedDateTime = `${datePart}T${timeWithoutMillis}+${timezoneOffset}`;
		return new Date(combinedDateTime);
	}
	function getArrayOfStrings(input: string): string[] {
		return input
				.split(',') // Split the input string by commas
				.map(word => word.trim()) // Remove whitespace from each element using trim()
				.filter(word => word !== ''); // Filter out empty strings if any (e.g., for cases like "US, Georgia, , Canada")
	}


</script>

<style>
	@import '../country-flags/flags.css';
</style>

<div class="container h-full mx-auto flex justify-center items-center">
	<div class="space-y-10 text-center flex flex-col items-center">
		{#each threadPreviews as thread}
		<div class="w-full text-token grid grid-cols-1 md:grid-cols-2 gap-4">
			<!-- Detailed -->
			<a class="card overflow-hidden">
				<div class="p-4 space-y-4">
					<h3 class="h3" data-toc-ignore>{thread.Topic}</h3>
					<article>
						<p>
							<!-- cspell:disable -->
							Lorem ipsum dolor sit amet consectetur adipisicing elit. Numquam aspernatur provident eveniet eligendi cumque consequatur tempore
							sint nisi sapiente. Iste beatae laboriosam iure molestias cum expedita architecto itaque quae rem.
							<!-- cspell:enable -->
						</p>
					</article>
				</div>
				<hr class="opacity-50" />
				<footer class="p-4 flex justify-start items-center space-x-4">
					<div class="flex-auto flex justify-between items-center">
						<div class="flex-auto flex justify-between items-center">
							<p>{thread.Name}</p>
							<p>No. {thread.ID}</p>
							<p>{parseDateStringToDate(thread.UnixTime).toLocaleString()}</p>
							<i class={'flag ' + getArrayOfStrings(thread.Flags)[0].toLocaleLowerCase()}></i>
						</div>
					</div>
				</footer>
			</a>
		</div>
		{/each}
	</div>
</div>