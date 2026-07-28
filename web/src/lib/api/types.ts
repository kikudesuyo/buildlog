export interface DiaryEntry {
	id: number;
	title: string;
	content: string;
	status?: 'draft' | 'published';
	createdAt: string;
	updatedAt: string;
	likesCount: number;
	hasLiked?: boolean;
}

export interface TechArticle {
	id: number;
	title: string;
	content: string;
	category: string;
	views?: string;
	status?: 'draft' | 'published';
	createdAt: string;
	updatedAt: string;
	likesCount: number;
	hasLiked?: boolean;
}

export type FeaturedTechArticle = TechArticle;

export interface AppProject {
	id: string;
	slug?: string;
	name: string;
	category: string;
	tags: string[];
	description: string;
	icon: string;
	iconUrl?: string;
	demoUrl?: string;
	codeUrl?: string;
}

export interface ProfileData {
	name: string;
	subtitle: string;
	title: string;
	avatarUrl: string;
	quote: string;
	bio: string[];
	highlights: {
		title: string;
		period: string;
		description: string;
	}[];
	award?: string;
	expertise: string[];
	contactEmail: string;
	githubUrl?: string;
	xUrl?: string;
	finalQuote: string;
}

export interface TrashEntry {
	id: number;
	type: string;
	title: string;
	content: string;
	category?: string;
	createdAt: string;
	deletedAt: string;
}
