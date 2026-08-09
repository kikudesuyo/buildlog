import { redirect, type Handle } from '@sveltejs/kit';

const JWT_TOKEN_COOKIE = 'buildlog_jwt_token';

export const handle: Handle = async ({ event, resolve }) => {
	if (event.url.pathname.startsWith('/admin')) {
		const session = event.cookies.get(JWT_TOKEN_COOKIE);
		const response = session
			? await event.fetch('/api/v1/auth/session', { headers: { cookie: `${JWT_TOKEN_COOKIE}=${session}` } })
			: null;
		if (!response?.ok) {
			const redirectTo = `${event.url.pathname}${event.url.search}`;
			throw redirect(303, `/auth?redirect=${encodeURIComponent(redirectTo)}`);
		}
	}

	return resolve(event);
};
