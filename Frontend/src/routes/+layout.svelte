<script lang='ts'>
	// The ordering of these imports is critical to your app working properly
	import '@skeletonlabs/skeleton/themes/theme-crimson.css';
	// If you have source.organizeImports set to true in VSCode, then it will auto change this ordering
	import '@skeletonlabs/skeleton/styles/skeleton.css';


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

	import {localStorageStore} from "@skeletonlabs/skeleton"; //Backup data
	import type { Writable } from "svelte/store";
	import type { UserPrefences } from "../app";
	import { onMount } from "svelte";								//Write data

	const userSettings : Writable<UserPrefences> = localStorageStore("userPrefs", {});

	//document.getElementById("themeHooker").setAttribute('data-theme', $userSettings.Theme);

	onMount(async function () {
		document.getElementById("themeHooker").setAttribute('data-theme', $userSettings.Theme);
	});


</script>

<AppShell>
	<svelte:fragment slot="header">
		<AppBar class="hidden lg:block">
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">Hesat</strong>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<a href={`rules`}>
				<Icon icon="ooui:article-ltr" />
				</a>
				<a href={`contact`}>
				<Icon icon="ooui:message" />
				</a>
				<a href={`settings`}>
				<Icon icon="ooui:settings" />
				</a>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<slot />
</AppShell>