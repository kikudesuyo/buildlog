import type { PageLoad } from './$types';
import type { AppProject } from '$lib/api/types';

const appProjects: AppProject[] = [
	{
		id: 'whichway',
		name: 'Whichway',
		category: 'Tool / Decision Support',
		tags: ['TypeScript', 'Go'],
		description: '公共交通機関の乗換案内アプリです。',
		icon: 'explore',
		iconUrl: '/whichway-icon.svg',
		demoUrl: 'https://whichway-6862oinh7-kikudeusyos-projects.vercel.app/',
		codeUrl: 'https://github.com/kikudesuyo/whichway'
	},
	{
		id: 'mahjong-scoreboard',
		name: '麻雀スコアボード管理',
		category: 'Tool / Utility',
		tags: ['TypeScript', 'Go'],
		description: '対局のスコア記録・計算・成績管理をスムーズに行える麻雀専用スコアボード。',
		icon: 'score',
		iconUrl: '/mahjong-icon.svg',
		demoUrl: 'https://mahjong-scoreboard-management.vercel.app/'
	},
		{
		id: 'pratan',
		name: 'Pratan',
		category: 'Education / Language',
		tags: ['TypeScript', 'React', 'Firebase'],
		description: '英単語の効率的な学習と記憶定着をサポートする語学学習アプリケーション。クイズ機能を搭載しています',
		icon: 'translate',
		iconUrl: '/pratan-icon.svg',
		demoUrl: 'https://pratan-714.web.app/',
		codeUrl: 'https://github.com/kikudesuyo/pratan'
	},
	{
		id: 'econom-eye',
		name: 'economEye',
		category: 'Finance / Tracking',
		tags: ['TypeScript', 'React', 'Firebase'],
		description: '商品の価格推移を視覚的に追跡・監視する商品価格追跡ツール。最初に作成したWebアプリケーションです。',
		icon: 'monitoring',
		iconUrl: '/economeye-icon.svg',
		demoUrl: 'https://economeye-d5146.web.app/',
		codeUrl: 'https://github.com/kikudesuyo/economEye'
	},
];

export const load: PageLoad = async () => ({
	appProjects
});
