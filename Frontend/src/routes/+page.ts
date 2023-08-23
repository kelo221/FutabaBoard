/** @type {import('./$types').PageLoad} */

const API_URL = 'http://127.0.0.1:8000/api/page/0';
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
export const load = async ({ fetch }) => {
	try {
		const response = await fetch(API_URL);
		if (!response.ok) {
			throw new Error('Network response was not ok');
		}

		const tempData = await response.json();
		return {
			threadPreviews: tempData
		};
	} catch (error) {
		console.error('An error occurred:', error);
		throw error;
	}
};
