import { createServer } from 'node:http';

const profile = {
	name: 'E2Eユーザー', subtitle: '固定モックプロフィール', title: 'Software Engineer',
	quote: 'E2E用プロフィールです。', bio: ['固定データの自己紹介です。'], highlights: [],
	expertise: ['Go', 'Svelte'], contactEmail: 'test@example.com', githubUrl: 'https://github.com/example', finalQuote: 'Build with care.'
};
const apps = [{ id: 'app-1', name: 'E2Eアプリ', category: 'Tool', tags: ['Svelte'], description: '固定アプリデータです。', icon: 'apps', demo_url: 'https://example.com/demo', code_url: 'https://github.com/example/app' }];
const diaries = [1, 2, 3, 4].map((id) => ({ id, type: 'diary', title: `日記${id}`, content: `E2E本文${id}`, status: 'published', created_at: '2026-08-13T00:00:00Z', updated_at: '2026-08-13T00:00:00Z', likes_count: id, comments_count: id === 1 ? 1 : 0, has_liked: false }));
const json = (res, body, status = 200) => { res.writeHead(status, { 'content-type': 'application/json' }); res.end(JSON.stringify(body)); };

createServer((req, res) => {
	const url = new URL(req.url ?? '/', 'http://localhost');
	const path = url.pathname;
	if (path.endsWith('/profile')) return json(res, { data: profile });
	if (path.endsWith('/apps')) return json(res, { data_list: apps });
	if (path.endsWith('/posts/history')) return json(res, { data_list: diaries });
	if (path.endsWith('/diaries')) {
		const offset = Number(url.searchParams.get('offset') ?? 0);
		return json(res, { data_list: diaries.slice(offset, offset + 3) });
	}
	if (path.endsWith('/techs')) return json(res, { data_list: [] });
	if (path.endsWith('/diaries/1')) return json(res, { data: diaries[0] });
	if (path.endsWith('/posts/1/comments')) return json(res, { data_list: [{ id: 1, post_id: 1, content: 'E2Eコメント', created_at: diaries[0].created_at, updated_at: diaries[0].updated_at }] });
	if (path.endsWith('/like')) return json(res, { data: { likes_count: 2, has_liked: true } });
	return json(res, { data_list: [] });
}).listen(Number(process.env.E2E_API_PORT ?? 18081), '127.0.0.1');
