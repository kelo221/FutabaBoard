/** @type {import('./$types').PageLoad} */
/** @type {import('$sveltekit/kit').PageLoad} */
export async function load({ params, fetch }) {
	try {
		const endpoint = `http://127.0.0.1:8000/api/thread/${params.id}`;
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
