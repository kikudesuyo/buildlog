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

export interface AnalyticsArticleItem {
	id: number;
	type: string;
	title: string;
	views: number;
	likes: number;
}

export interface MonthlyActivityItem {
	month: string;
	count: number;
}

export interface AnalyticsData {
	totalViews: number;
	totalLikes: number;
	totalPosts: number;
	diaryCount: number;
	techCount: number;
	topViewsArticles: AnalyticsArticleItem[];
	topLikesArticles: AnalyticsArticleItem[];
	monthlyActivities: MonthlyActivityItem[];
}

export interface HistoryItem {
	id: number;
	type: 'diary' | 'tech';
	title: string;
	createdAt: string;
}

export interface CommentEntry {
	id: number;
	postId: number;
	parentId: number | null;
	content: string;
	createdAt: string;
	updatedAt: string;
	replies: CommentEntry[];
}
