package cache

import (
	"testing"
	"time"
)

func TestStoreExpiresEntries(t *testing.T) {
	store := New(10 * time.Millisecond)
	store.Set("key", "value")

	value, ok := store.Get("key")
	if !ok || value != "value" {
		t.Fatalf("Get() = (%v, %t), want (value, true)", value, ok)
	}

	time.Sleep(20 * time.Millisecond)
	if _, ok := store.Get("key"); ok {
		t.Fatal("Get() returned an expired entry")
	}
}

func TestStoreDeleteRemovesMultipleEntries(t *testing.T) {
	store := New(time.Minute)
	store.Set("first", 1)
	store.Set("second", 2)

	store.Delete("first", "second")
	if _, ok := store.Get("first"); ok {
		t.Fatal("first entry was not deleted")
	}
	if _, ok := store.Get("second"); ok {
		t.Fatal("second entry was not deleted")
	}
}
