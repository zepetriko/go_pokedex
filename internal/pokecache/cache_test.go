package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(50 * time.Millisecond)

	key := "test"
	val := []byte("hello world")

	cache.Add(key, val)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatal("expected key to exist in cache")
	}

	if string(got) != string(val) {
		t.Fatalf("expected %s, got %s", val, got)
	}

	time.Sleep(60 * time.Millisecond)

	_, ok = cache.Get(key)
	if ok {
		t.Fatal("expected key to be reaped from cache")
	}
}