import { env } from '$env/dynamic/public';
import type { RequestHandler } from './$types';

const forward: RequestHandler = async ({ request, params, url, fetch }) => {
	const apiBaseUrl = (env.PUBLIC_API_BASE_URL || 'http://localhost:8081').replace(/\/$/, '');
	const headers = new Headers();
	for (const name of ['content-type', 'cookie', 'accept']) {
		const value = request.headers.get(name);
		if (value) headers.set(name, value);
	}

	const response = await fetch(`${apiBaseUrl}/api/v1/${params.path}${url.search}`, {
		method: request.method,
		headers,
		body: request.method === 'GET' || request.method === 'HEAD' ? undefined : await request.arrayBuffer()
	});

	const responseHeaders = new Headers();
	for (const name of ['content-type', 'cache-control']) {
		const value = response.headers.get(name);
		if (value) responseHeaders.set(name, value);
	}
	return new Response(response.body, { status: response.status, headers: responseHeaders });
};

export const GET = forward;
export const POST = forward;
export const PUT = forward;
export const DELETE = forward;
