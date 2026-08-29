import { defineConfig } from '@playwright/test';

export default defineConfig({
	webServer: [
		{ command: 'E2E_API_PORT=18082 node tests/e2e/mock-content-api.mjs', port: 18082 },
		{ command: 'npm run build && npm run preview -- --port 4174', port: 4174, env: { PUBLIC_API_BASE_URL: 'http://localhost:18082' } },
		{ command: 'E2E_API_PORT=18081 node tests/e2e/mock-api.mjs', port: 18081 },
		{ command: 'npm run preview -- --port 4173', port: 4173, env: { PUBLIC_API_BASE_URL: 'http://localhost:18081' } }
	],
	testMatch: '**/*.e2e.{ts,js}'
});
