<script lang='ts'>
	// The ordering of these imports is critical to your app working properly
	import '@skeletonlabs/skeleton/themes/theme-crimson.css';
	// If you have source.organizeImports set to true in VSCode, then it will auto change this ordering
	import '@skeletonlabs/skeleton/styles/skeleton.css';

	import { computePosition, autoUpdate, offset, shift, flip, arrow } from '@floating-ui/dom';

	import { storePopup } from "@skeletonlabs/skeleton";
	storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });


	// Most of your app wide CSS should be put in this file
	import '../app.postcss';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	import Icon from '@iconify/svelte';

	import '@skeletonlabs/skeleton/themes/theme-skeleton.css';
	import '@skeletonlabs/skeleton/themes/theme-crimson.css';
	import '@skeletonlabs/skeleton/themes/theme-hamlindigo.css';
	import '@skeletonlabs/skeleton/themes/theme-gold-nouveau.css';
	import '@skeletonlabs/skeleton/themes/theme-modern.css';
	import '@skeletonlabs/skeleton/themes/theme-rocket.css';
	import '@skeletonlabs/skeleton/themes/theme-sahara.css';
	import '@skeletonlabs/skeleton/themes/theme-seafoam.css';
	import '@skeletonlabs/skeleton/themes/theme-seasonal.css';
	import '@skeletonlabs/skeleton/themes/theme-vintage.css';

	import { onMount } from "svelte";
	import ReplyBox from "./Components/ReplyBox.svelte";
	import { currentThreadStore, postPreview, userSettings } from "../stores";

	let mousePos: { x: number; y: number }

	onMount(async function () {
		document.getElementById("themeHooker").setAttribute('data-theme', $userSettings.Theme);
		window.addEventListener('mousemove', (event) => {
			mousePos = { x: event.clientX, y: event.clientY };
		});
	});


	import Post from "./Components/Post.svelte";

</script>

<style>
	.floating-element {
		position: absolute;
		top: calc( var(--y) * 1px + -280px);
		left: calc( var(--x) * 1px + -10px );
		z-index: 1000;
		user-select: none;
		width: 300px;

		height: 200px;
	}

	.hidden{
		display: none;
	}

</style>


<AppShell>
	<svelte:fragment slot="pageHeader">
		<AppBar>
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">Hesat</strong>
			</svelte:fragment>

			<svelte:fragment slot="trail">
				{#if ($currentThreadStore !== undefined)}
				<small class="text-xl ">
					{$currentThreadStore.PostCount} Posts | Page {$currentThreadStore.Page} |
					</small>
					{/if}


				<a href={`/rules`}>
				<Icon icon="ooui:article-ltr" />
				</a>
				<a href={`/forms`}>
				<Icon icon="ooui:message" />
				</a>
				<a href={`/settings`}>
				<Icon icon="ooui:settings" />
				</a>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>


	<slot />
	<ReplyBox/>

	{#if (mousePos)}
	<div class="card floating-element" style="border: 4px outset
	rgba(var(--color-secondary-900) / 1);
	--x:{mousePos.x}; --y:{mousePos.y}" class:hidden="{$postPreview.open === false}">
	{#if ($postPreview.open && $currentThreadStore)}
		<Post content={$postPreview.postData} isOpen="true" threadID={$currentThreadStore.ID}/>
	{/if}
		{#if (!$currentThreadStore)}
			FAULTY
			{/if}
	</div>
	{/if}



</AppShell>