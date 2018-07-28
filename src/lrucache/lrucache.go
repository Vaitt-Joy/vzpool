package lrucache

import (
	"sync"
	"container/list"
	"time"
	"fmt"
)

type LruCache struct {
	mu sync.Mutex

	// list & table of *entry objects
	list  *list.List
	table map[string]*list.Element

	size uint64

	maxSize uint64
}

type Value interface {
	Size() int
}

type Item struct {
	key   string
	value Value
}

type entry struct {
	key           string
	value         Value
	size          int
	time_accessed time.Time
}

func NewLruCache(maxSize uint64) *LruCache {
	return &LruCache{
		list:    list.New(),
		table:   make(map[string]*list.Element),
		maxSize: maxSize,
	}
}

func (this *LruCache) addNew(key string, value Value) {
	/* 添加新的 value*/
	newElement := &entry{key, value, value.Size(), time.Now()}
	element := this.list.PushFront(newElement)
	this.table[key] = element
	this.size += uint64(newElement.size)
	this.checkCapacity()
}

func (this *LruCache) Get(key string) (v Value, ok bool) {
	/*获取缓存的值*/
	this.mu.Lock()
	defer this.mu.Unlock()

	element := this.table[key]
	if element == nil {
		return nil, false
	}
	this.moveToFirst(element)
	return element.Value.(*entry).value, true
}

func (this *LruCache) Set(key string, value Value) {
	/* 更新值*/
	this.mu.Lock()
	defer this.mu.Unlock()

	if element := this.table[key]; element != nil {
		this.updateInplace(element, value)
	} else {
		this.addNew(key, value)
	}
}

func (this *LruCache) SetIfAbsent(key string, value Value) {
	/*不存在就添加一个 缓存*/
	this.mu.Lock()
	defer this.mu.Unlock()

	if element := this.table[key]; element != nil {
		this.moveToFirst(element)
	} else {
		this.addNew(key, value)
	}
}

func (this *LruCache) Delete(key string) bool {
	/*删除*/
	this.mu.Lock()
	defer this.mu.Unlock()

	element := this.table[key]
	if element == nil {
		return false
	}
	this.list.Remove(element)
	delete(this.table, key)
	this.size -= uint64(element.Value.(*entry).size)
	return true
}

func (this *LruCache) Clear() {
	/*重置 lru 空间*/
	this.mu.Lock()
	defer this.mu.Unlock()

	this.list.Init()
	this.table = make(map[string]*list.Element)
	this.size = 0
}

func (this *LruCache) SetMaxSize(maxSize uint64) {
	/*设置缓存结构的缓存最大值 空间*/
	this.mu.Lock()
	defer this.mu.Unlock()

	this.size = maxSize
	this.checkCapacity()
}

func (this *LruCache) Count() (length, size, maxSize uint64, oldest time.Time) {
	/*统计使用情况*/
	this.mu.Lock()
	defer this.mu.Unlock()

	if lastElem := this.list.Back(); lastElem != nil {
		oldest = lastElem.Value.(*entry).time_accessed
	}
	return uint64(this.list.Len()), this.size, this.maxSize, oldest
}

func (this *LruCache) CountJson() string {
	/*统计使用情况*/
	this.mu.Lock()
	defer this.mu.Unlock()

	if this == nil {
		return "{}"
	}
	l, s, c, t := this.Count()

	return fmt.Sprintf("{\"len\": %v ,\"size\": %v , \"maxSize\": %v , \"OldestAccess\": \"%v\"}", l, s, c, t)
}

func (this *LruCache) Keys() []string {
	this.mu.Lock()
	defer this.mu.Unlock()

	keys := make([]string, 0, this.list.Len())
	for e := this.list.Front(); e != nil; e = e.Next() {
		keys = append(keys, e.Value.(*entry).key)
	}
	return keys
}

func (this *LruCache) Items() []Item {
	this.mu.Lock()
	defer this.mu.Unlock()

	items := make([]Item, 0, this.list.Len())
	for e := this.list.Front(); e != nil; e = e.Next() {
		items = append(items, Item{key: e.Value.(*entry).key, value: e.Value.(*entry).value})
	}
	return items
}

func (this *LruCache) updateInplace(element *list.Element, value Value) {
	/* 更新一个key 栈顶*/
	valueSize := value.Size()
	sizeDiff := valueSize - element.Value.(*entry).size
	element.Value.(*entry).value = value
	element.Value.(*entry).size = valueSize
	this.size += uint64(sizeDiff)
	this.moveToFirst(element)

}

func (this *LruCache) moveToFirst(element *list.Element) {
	/*更新栈顶*/
	this.list.MoveToFront(element)
	element.Value.(*entry).time_accessed = time.Now()
}

func (this *LruCache) checkCapacity() {
	/*更新越界的缓存*/
	for this.size > this.maxSize {
		delElement := this.list.Back()
		delValue := delElement.Value.(*entry)
		this.list.Remove(delElement)
		delete(this.table, delValue.key)
		this.size -= uint64(delValue.size)
	}
}
