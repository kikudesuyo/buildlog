export interface DiaryEntry {
	id: number;
	title: string;
	content: string;
	createdAt: string;
	updatedAt: string;
}

export interface TechArticle {
	id: number;
	title: string;
	excerpt: string;
	category: string;
	readTime: string;
	views?: string;
	isNewsletter?: boolean;
	createdAt: string;
	updatedAt: string;
}

export type FeaturedTechArticle = Omit<TechArticle, 'id' | 'views' | 'isNewsletter'>;

export interface AppProject {
	id: string;
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
	finalQuote: string;
}
