package lrucache

import (
	"testing"
	"fmt"
)

type CacheValue struct {
	size int
}

func (cv *CacheValue) Size() int {
	return cv.size
}

func TestInitState(t *testing.T) {
	cache := getCache(5)
	l, sz, c, _ := cache.Count()
	if l != 0 {
		t.Errorf("length = %v, want 0", l)
	}

	if sz != 0 {
		t.Errorf("size = %v, want 0", sz)
	}
	if c != 5 {
		t.Errorf("maxSize = %v, want 5", c)
	}
	fmt.Println("run finish !!!!")
}

func TestInsertValue(t *testing.T) {
	cache := getCache(100)

	for i := 0; i < 120; i++ {
		cache.Set(fmt.Sprintf("key%v", i), &CacheValue{0})
	}

	for i := 0; i < 40; i++ {
		key := fmt.Sprintf("key%v", i)
		v, ok := cache.Get(key)
		if ok {
			fmt.Printf("key: %v ,value: %v \n", key, v)
		} else {
			fmt.Printf("key: %v is not found in cache! \n", key)
		}
	}
	cache.Get("key5")
	cache.Get("key19")
	cache.Get("key13")
	fmt.Println(cache.Keys())

	fmt.Println("run finish !!!!")
}

func getCache(capacity uint64) *LruCache {
	return NewLruCache(capacity)
}
