package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	if cache.data == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "key1",
			val: []byte("val1"),
		},

		{
			key: "key2",
			val: []byte("val2"),
		},
		{
			key: "key3",
			val: []byte("val3"),
		},
	}
	cache := NewCache(time.Millisecond)
	for _, cas := range cases {
		cache.Add(cas.key, cas.val)
		actual, ok := cache.Get(cas.key)
		if !ok {
			t.Errorf("%s key not found", cas.key)
			continue
		}
		if string(actual) != string(cas.val) {
			t.Errorf("%s does'nt match %s", actual, cas.val)
		}
	}
}

func TestReapedLoop(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	key := "key1"
	cache.Add(key, []byte("val1"))
	time.Sleep(interval + time.Millisecond)
	if _, ok := cache.Get(key); ok {
		t.Errorf("%s key should be reaped", key)
	}
}

func TestFailReapedLoop(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	key := "key1"
	cache.Add(key, []byte("val1"))
	time.Sleep(time.Millisecond)
	if _, ok := cache.Get(key); !ok {
		t.Errorf("%s key should have not beet reaped", key)
	}
}
