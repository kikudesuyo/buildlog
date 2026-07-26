import type { PageLoad } from './$types';
import type { ProfileData } from '$lib/api/types';

const profileData: ProfileData = {
	name: 'kikudesuyo',
	subtitle: 'Hriomu Kikuchi — Product Engineer',
	title: 'Creative Director',
	avatarUrl: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=300&q=80',
	quote: '「余白とは、単なる空白ではなく、思考が呼吸するための空間である。」',
	bio: [
		'15年間にわたり、デザインとテクノロジーの交差点で「静寂」を追求してきました。情報の過剰な現代において、真に価値のある体験とは、引き算によってのみ生まれると信じています。',
		'現在は Essence Editorial のリードデザイナーとして、執筆者と読者が深く繋がれるデジタル空間の構築に注力しています。物理的なノートのような手触り感と、デジタルの効率性を融合させた、新しい編集体験を提案しています。'
	],
	highlights: [
		{
			title: 'Essence Platform',
			period: '2021 — Present',
			description: '次世代の執筆環境をデザイン。月間100万人のアクティブユーザーを持つプラットフォームへと成長。'
		},
		{
			title: 'Mono Design Studio',
			period: '2018 — 2021',
			description: '創設者。12の国際的なデザイン賞を受賞し、ミニマリズムの先駆者として認知される。'
		}
	],
	award: 'Global Design Excellence 2023',
	expertise: [
		'インタラクション設計',
		'タイポグラフィ',
		'UXライティング',
		'ブランド戦略',
		'ミニマリストUI',
		'コンテンツ戦略'
	],
	contactEmail: 'contact@essence.editorial',
	finalQuote: '「物語は、余白の中で最も力強く響く。あなたの想いを聞かせてください。」'
};

export const load: PageLoad = async () => ({
	profileData
});
