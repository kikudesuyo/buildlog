import { createServer } from 'node:http';
const post = { id: 1, type: 'diary', title: '操作対象の日記', content: 'E2E本文', status: 'published', created_at: '2026-08-13T00:00:00Z', updated_at: '2026-08-13T00:00:00Z', likes_count: 1, comments_count: 1, has_liked: false };
const more = [2, 3, 4].map((id) => ({ ...post, id, title: `追加日記${id}`, likes_count: id, comments_count: 0 }));
const json = (res, body, status = 200) => { res.writeHead(status, { 'content-type': 'application/json' }); res.end(JSON.stringify(body)); };
createServer((req, res) => {
	const url = new URL(req.url ?? '/', 'http://localhost');
	if (url.pathname.endsWith('/diaries/1')) return json(res, { data: post });
	if (url.pathname.endsWith('/diaries')) return json(res, { data_list: Number(url.searchParams.get('offset') ?? 0) ? [more[2]] : [post, ...more.slice(0, 2)] });
	if (url.pathname.endsWith('/posts/1/comments')) return json(res, { data_list: [{ id: 1, post_id: 1, content: '既存コメント', created_at: post.created_at, updated_at: post.updated_at }] });
	if (url.pathname.endsWith('/like')) return json(res, { data: { likes_count: 2, has_liked: true } });
	if (url.pathname.endsWith('/profile')) return json(res, { data: { name: 'E2E', subtitle: '', title: '', quote: '', bio: [], highlights: [], expertise: [], contactEmail: 'test@example.com', finalQuote: '' } });
	return json(res, { data_list: [] });
}).listen(Number(process.env.E2E_API_PORT ?? 18083), '127.0.0.1');
