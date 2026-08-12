package service

import (
	"testing"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
)

func TestSortTechFeed_List(t *testing.T) {
	itemList := []entity.TechFeedItem{
		{ID: 1, CreatedAt: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC)},
		{ID: 2, CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	sortTechFeed_List(itemList, "asc")
	if itemList[0].ID != 2 || itemList[1].ID != 1 {
		t.Fatalf("ascending order = [%d, %d], want [2, 1]", itemList[0].ID, itemList[1].ID)
	}

	sortTechFeed_List(itemList, "desc")
	if itemList[0].ID != 1 || itemList[1].ID != 2 {
		t.Fatalf("descending order = [%d, %d], want [1, 2]", itemList[0].ID, itemList[1].ID)
	}
}
