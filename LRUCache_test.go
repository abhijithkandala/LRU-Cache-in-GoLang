package main

import "testing"

func TestPutAndGet(t *testing.T) {
	cache := newLRUCache(3)

	cache.Put("A", 1)
	cache.Put("B", 2)

	value, ok := cache.Get("A")

	if !ok {
		t.Fatal("expected A to exist")
	}

	if value != 1 {
		t.Fatalf("expected A = 1, got %d", value)
	}
}

func TestGetMissingKey(t *testing.T) {
	cache := newLRUCache(3)

	value, ok := cache.Get("A")

	if ok {
		t.Fatal("expected A to be missing")
	}

	if value != 0 {
		t.Fatalf("expected missing value to be 0, got %d", value)
	}
}

func TestUpdateExistingKey(t *testing.T) {
	cache := newLRUCache(3)

	cache.Put("A", 1)
	cache.Put("A", 100)

	value, ok := cache.Get("A")

	if !ok {
		t.Fatal("expected A to exist")
	}

	if value != 100 {
		t.Fatalf("expected A = 100, got %d", value)
	}
}

func TestEviction(t *testing.T) {
	cache := newLRUCache(3)

	cache.Put("A", 1)
	cache.Put("B", 2)
	cache.Put("C", 3)

	// A becomes the most recently used.
	cache.Get("A")

	// B should now be the least recently used.
	cache.Put("D", 4)

	_, ok := cache.Get("B")

	if ok {
		t.Fatal("expected B to be evicted")
	}

	// These should still exist.
	for _, key := range []string{"A", "C", "D"} {
		_, ok := cache.Get(key)

		if !ok {
			t.Fatalf("expected %s to still exist", key)
		}
	}
}

func TestCapacityOne(t *testing.T) {
	cache := newLRUCache(1)

	cache.Put("A", 1)
	cache.Put("B", 2)

	_, ok := cache.Get("A")

	if ok {
		t.Fatal("expected A to be evicted")
	}

	value, ok := cache.Get("B")

	if !ok || value != 2 {
		t.Fatal("expected B = 2")
	}
}
