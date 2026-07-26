export const techCategories = ['Frontend', 'Backend', 'Database', 'Infrastructure'] as const;

export type TechCategory = (typeof techCategories)[number];

export const defaultTechCategory: TechCategory = 'Frontend';
