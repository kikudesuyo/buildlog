package service

import (
	"testing"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
)

func TestSortTechFeed_List(t *testing.T) {
	item_List := []entity.TechFeedItem{
		{ID: 1, CreatedAt: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC)},
		{ID: 2, CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	sortTechFeed_List(item_List, "asc")
	if item_List[0].ID != 2 || item_List[1].ID != 1 {
		t.Fatalf("ascending order = [%d, %d], want [2, 1]", item_List[0].ID, item_List[1].ID)
	}

	sortTechFeed_List(item_List, "desc")
	if item_List[0].ID != 1 || item_List[1].ID != 2 {
		t.Fatalf("descending order = [%d, %d], want [1, 2]", item_List[0].ID, item_List[1].ID)
	}
}
