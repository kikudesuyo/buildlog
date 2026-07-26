import type { PageLoad } from './$types';
import { load as parentLoad } from '../tech/+page';

export const load: PageLoad = async (event) => {
	return parentLoad(event as any);
};
