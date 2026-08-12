import { defineConfig } from '@playwright/test';

export default defineConfig({
	webServer: [
		{ command: 'E2E_API_PORT=18083 node tests/e2e/mock-interactions-api.mjs', port: 18083 },
		{ command: 'npm run build && npm run preview -- --port 4176', port: 4176, env: { PUBLIC_API_BASE_URL: 'http://localhost:18083' } }
	],
	testMatch: '**/*.e2e.{ts,js}'
});
