// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
/** @type {import('./$types').PageLoad} */
export async function load({ params }) {
	return {
		id: params.id
	};
}
