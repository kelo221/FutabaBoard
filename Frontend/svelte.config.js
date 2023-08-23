import adapter from '@sveltejs/adapter-auto';
import { vitePreprocess } from '@sveltejs/kit/vite';

const config = {
	preprocess: vitePreprocess(),
	vitePlugin: {
		inspector: true
	},
	kit: {
		adapter: adapter({
			hydrate: false
		})
	}
};

export default config;
