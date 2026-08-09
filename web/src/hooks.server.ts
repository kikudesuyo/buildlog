import { redirect } from '@sveltejs/kit';
import { env } from '$env/dynamic/private';
import { isValidAdminSession } from '$lib/server/admin-session';

const ADMIN_SESSION_COOKIE = 'buildlog_admin_session';

export const handle = async ({ event, resolve }) => {

	if (event.url.pathname.startsWith('/admin') && !(env.ADMIN_SESSION_SECRET && await isValidAdminSession(event.cookies.get(ADMIN_SESSION_COOKIE), env.ADMIN_SESSION_SECRET))) {
		const redirectTo = `${event.url.pathname}${event.url.search}`;
		throw redirect(303, `/auth?redirect=${encodeURIComponent(redirectTo)}`);
	}

	return resolve(event);
};
