/** @type {import("./$types").PageLoad} */
import { PUBLIC_BACKEND_ADDRESS } from "$env/static/public";

/** @type {import("$sveltekit/kit").PageLoad} */
export async function load({ params, fetch }) {
	try {
		const endpoint = `${PUBLIC_BACKEND_ADDRESS}/api/thread/${params.id}`;
		const response = await fetch(endpoint);
		if (!response.ok) {
			throw new Error('Not found');
		}
		const tempData = await response.json();
		return {
			thread: tempData
		};
	} catch (err) {
		throw new Error('Not found');
	}
}
