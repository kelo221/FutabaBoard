// eslint-disable-next-line svelte/no-unused-vars
/** @type {import("./$types").PageLoad} */
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore

import { PUBLIC_BACKEND_ADDRESS } from "$env/static/public";

// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
// @ts-ignore
// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export const load = async ({ fetch }) => {
	const fetchPage = async () => {
		try {
			const response = await fetch(`${PUBLIC_BACKEND_ADDRESS}/api/page/0`);
			if (!response.ok) {
				throw new Error('Network response was not ok');
			}

			return await response.json();
		} catch (error) {
			throw error;
		}
	};

	const fetchPageCount = async () => {
		try {
			const response = await fetch('http://127.0.0.1:8000/api/pageCount');
			if (!response.ok) {
				throw new Error('Network response was not ok');
			}

			const tempData = await response.json();
			return tempData;
		} catch (error) {
			throw error;
		}
	};

	return {
		threadPreviews: fetchPage(),
		pageCount: fetchPageCount()
	};
};
