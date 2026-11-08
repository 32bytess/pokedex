package pokecache

import (
	"testing"
)

func TestAddGet(t *testing.T) {

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	cache := NewCache()
	for i := range cases {
		cache.Add(cases[i].key, cases[i].val)
		if len(cache.data) != i+1 {
			t.Errorf("case %d: expected to add key %s, val: %s", i, cases[i].key, cases[i].val)
			continue
		}

		got, ok := cache.Get(cases[i].key)
		if !ok {
			t.Errorf("case %d: expected to find key %q", i, cases[i].key)
			continue
		}
		if string(got) != string(cases[i].val) {
			t.Errorf("case %d: expected %q, got %q", i, cases[i].val, got)
		}
	}

}

func TestGetMissing(t *testing.T) {
	cache := NewCache()

	got, ok := cache.Get("not-there")
	if ok {
		t.Errorf("expected ok to be false for a missing key, got true")
	}
	if got != nil {
		t.Errorf("expected nil for a missing key, got %q", got)
	}
}
