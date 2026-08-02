package service

import "github.com/kikudesuyo/buildlog/api/entity"

func mapTagsToStrings(tags []entity.DBTableTag) []string {
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	return names
}
