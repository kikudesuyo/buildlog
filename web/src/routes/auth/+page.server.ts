import { fail, redirect } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';

const ADMIN_SESSION_COOKIE = 'buildlog_admin_session';

function getSafeRedirectTarget(value: string | null): string {
	return value?.startsWith('/admin') ? value : '/admin';
}

export const load: PageServerLoad = async ({ cookies, fetch, url }) => {
	const session = cookies.get(ADMIN_SESSION_COOKIE);
	if (!session) return;

	const response = await fetch('/api/v1/auth/session', {
		headers: { cookie: `${ADMIN_SESSION_COOKIE}=${session}` }
	});
	if (response.ok) throw redirect(303, getSafeRedirectTarget(url.searchParams.get('redirect')));
	cookies.delete(ADMIN_SESSION_COOKIE, { path: '/' });
};

export const actions: Actions = {
	default: async ({ cookies, fetch, request, url }) => {
		const formData = await request.formData();
		const password = formData.get('password');

		if (typeof password !== 'string') return fail(401, { error: 'パスワードが正しくありません。' });

		const response = await fetch('/api/v1/auth/login', {
			method: 'POST',
			headers: { 'content-type': 'application/json' },
			body: JSON.stringify({ password })
		});
		if (!response.ok) return fail(response.status === 503 ? 500 : 401, { error: 'パスワードが正しくありません。' });
		const result = (await response.json()) as { data?: { session?: string } };
		if (!result.data?.session) return fail(500, { error: '認証設定が不足しています。' });

		cookies.set(ADMIN_SESSION_COOKIE, result.data.session, {
			path: '/',
			httpOnly: true,
			sameSite: 'lax',
			secure: url.protocol === 'https:',
			maxAge: 60 * 60 * 24 * 30
		});

		throw redirect(303, getSafeRedirectTarget(url.searchParams.get('redirect')));
	}
};
