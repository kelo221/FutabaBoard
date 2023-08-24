/** @type {import('./$types').PageLoad} */
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
export const load = async ({ fetch }) => {
	const fetchPage = async () => {
		try {
			const response = await fetch('http://127.0.0.1:8000/api/page/0');
			if (!response.ok) {
				throw new Error('Network response was not ok');
			}

			const tempData = await response.json();
			return tempData;
		} catch (error) {
			console.error('An error occurred:', error);
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
			console.error('An error occurred:', error);
			throw error;
		}
	};

	return {
		threadPreviews: fetchPage(),
		pageCount: fetchPageCount()
	};
};
