import { defineConfig } from '@playwright/test';

export default defineConfig({
	webServer: [
<<<<<<< HEAD
		{ command: 'E2E_API_PORT=18082 node tests/e2e/mock-content-api.mjs', port: 18082 },
		{ command: 'npm run build && npm run preview -- --port 4174', port: 4174, env: { PUBLIC_API_BASE_URL: 'http://localhost:18082' } }
		{ command: 'E2E_API_PORT=18083 node tests/e2e/mock-interactions-api.mjs', port: 18083 },
		{ command: 'npm run preview -- --port 4176', port: 4176, env: { PUBLIC_API_BASE_URL: 'http://localhost:18083' } }
	],
	testMatch: '**/*.e2e.{ts,js}'
});
