import { env } from '$env/dynamic/private';
import { fail, redirect } from '@sveltejs/kit';
import { createAdminSession } from '$lib/server/admin-session';
import type { Actions, PageServerLoad } from './$types';

const ADMIN_SESSION_COOKIE = 'buildlog_admin_session';

function getSafeRedirectTarget(value: string | null): string {
	return value?.startsWith('/admin') ? value : '/admin';
}

export const load: PageServerLoad = ({ cookies, url }) => {
	if (cookies.get(ADMIN_SESSION_COOKIE)) {
		throw redirect(303, getSafeRedirectTarget(url.searchParams.get('redirect')));
	}
};

export const actions: Actions = {
	default: async ({ cookies, request, url }) => {
		const formData = await request.formData();
		const password = formData.get('password');

		if (!env.ADMIN_PASSWORD || !env.ADMIN_SESSION_SECRET) {
			return fail(500, { error: '認証設定が不足しています。' });
		}
		if (typeof password !== 'string' || password !== env.ADMIN_PASSWORD) {
			return fail(401, { error: 'パスワードが正しくありません。' });
		}

		cookies.set(ADMIN_SESSION_COOKIE, await createAdminSession(env.ADMIN_SESSION_SECRET), {
			path: '/',
			httpOnly: true,
			sameSite: 'lax',
			secure: url.protocol === 'https:',
			maxAge: 60 * 60 * 24 * 30
		});

		throw redirect(303, getSafeRedirectTarget(url.searchParams.get('redirect')));
	}
};
