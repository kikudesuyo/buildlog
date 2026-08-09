import { redirect } from '@sveltejs/kit';

const ADMIN_SESSION_COOKIE = 'buildlog_admin_session';

export const handle = async ({ event, resolve }) => {
	if (event.url.pathname.startsWith('/admin')) {
		const session = event.cookies.get(ADMIN_SESSION_COOKIE);
		const response = session
			? await event.fetch('/api/v1/auth/session', { headers: { cookie: `${ADMIN_SESSION_COOKIE}=${session}` } })
			: null;
		if (!response?.ok) {
			const redirectTo = `${event.url.pathname}${event.url.search}`;
			throw redirect(303, `/auth?redirect=${encodeURIComponent(redirectTo)}`);
		}
	}

	return resolve(event);
};
