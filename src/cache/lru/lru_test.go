package lru


import "testing"

func TestLru(t *testing.T) {
	lruCache := NewLruCache(3)
	lruCache.Add("a", 1)
	lruCache.Add("b", 2)
	lruCache.Add("c", 3)
	value := lruCache.Get("b")
	if value != 2 {
		t.Error("uncorrect value ", value, " by key b")
	}
	value = lruCache.Get("e")
	if value != nil {
		t.Error("uncorrect value ", value, " by key e")
	}
	lruCache.Add("d", 4)
	lruCache.Add("e", 5)
	value = lruCache.Get("a")
	if value != nil {
		t.Error("uncorrect value ", value, " by key a")
	}
	value = lruCache.Get("c")
	if value != nil {
		t.Error("uncorrect value ", value, " by key c")
	}
	value = lruCache.Get("b")
	if value != 2 {
		t.Error("uncorrect value ", value, " by key b")
	}
}
