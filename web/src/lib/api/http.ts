import * as env from '$env/static/public';
import type { LoadEvent } from '@sveltejs/kit';

export type ApiFetch = LoadEvent['fetch'];

export type ApiListResponse<T> = {
	data_list: T[];
};

export type ApiObjectResponse<T> = {
	data: T;
};

const apiBaseUrl = (() => {
	const rawUrl = env.PUBLIC_API_BASE_URL || 'http://localhost:8081';
	return rawUrl.endsWith('/api/v1') ? rawUrl : `${rawUrl}/api/v1`;
})();

export async function get<T>(fetchFn: ApiFetch, path: string): Promise<T> {
	const response = await fetchFn(`${apiBaseUrl}${path}`);
	if (!response.ok) {
		throw new Error(`API request failed: ${response.status} ${response.statusText}`);
	}
	return response.json() as Promise<T>;
}

export async function sendRequest<T>(method: string, path: string, body?: unknown): Promise<T> {
	const response = await fetch(`${apiBaseUrl}${path}`, {
		method,
		headers: body ? { 'Content-Type': 'application/json' } : undefined,
		body: body ? JSON.stringify(body) : undefined
	});
	if (!response.ok) {
		throw new Error(`API request failed: ${response.status} ${response.statusText}`);
	}
	return response.json() as Promise<T>;
}
