import { createServer } from 'node:http';

const profile = { name: 'E2Eユーザー', subtitle: '固定プロフィール', title: 'Engineer', quote: '固定データです。', bio: ['E2E用の自己紹介です。'], highlights: [], award: '', expertise: ['Go'], contact_email: 'test@example.com', final_quote: 'Build with care.' };
const app = { id: 'app-1', name: 'E2Eアプリ', category: 'Tool', tags: ['Go'], description: '固定アプリデータです。', icon: 'apps', demo_url: 'https://example.com', code_url: 'https://github.com/example/app' };
const json = (res, body, status = 200) => { res.writeHead(status, { 'content-type': 'application/json' }); res.end(JSON.stringify(body)); };
createServer((req, res) => {
	const path = new URL(req.url ?? '/', 'http://localhost').pathname;
	if (path.endsWith('/profile')) return json(res, { data: profile });
	if (path.endsWith('/apps')) return json(res, { data_list: [app] });
	if (path.endsWith('/techs')) return json(res, { data_list: [] });
	if (path.endsWith('/posts/history') || path.endsWith('/diaries')) return json(res, { data_list: [] });
	return json(res, { data_list: [] });
}).listen(Number(process.env.E2E_API_PORT ?? 18082), '127.0.0.1');
