package service

import (
	"testing"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
)

func TestSortTechFeedItems(t *testing.T) {
	items := []entity.TechFeedItem{
		{ID: 1, CreatedAt: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC)},
		{ID: 2, CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	sortTechFeedItems(items, "asc")
	if items[0].ID != 2 || items[1].ID != 1 {
		t.Fatalf("ascending order = [%d, %d], want [2, 1]", items[0].ID, items[1].ID)
	}

	sortTechFeedItems(items, "desc")
	if items[0].ID != 1 || items[1].ID != 2 {
		t.Fatalf("descending order = [%d, %d], want [1, 2]", items[0].ID, items[1].ID)
	}
}
