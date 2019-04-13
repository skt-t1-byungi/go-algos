package hashtable

import "github.com/skt-t1-byungi/go-algos/linkedlist"

type tableItem struct {
	key string
	val interface{}
}

type HashTable struct {
	bucket []*linkedlist.LinkedList
	size   int
}

func New() *HashTable {
	return &HashTable{make([]*linkedlist.LinkedList, 10), 0}
}

func (ht *HashTable) resize() {
	old := ht.bucket
	ht.bucket = make([]*linkedlist.LinkedList, len(old)*2)
	ht.size = 0
	each(old, func(val interface{}, key string) {
		ht.Set(key, val)
	})
}

func (ht *HashTable) Set(key string, val interface{}) {
	if ht.size == len(ht.bucket)+2 {
		ht.resize()
	}

	pos := ht.pos(key)
	list := ht.bucket[pos]
	if list == nil {
		list = linkedlist.New()
		ht.bucket[pos] = list
	}
	if _, idx := find(list, key); idx > -1 {
		list.RemoveAt(idx)
	}

	ht.size++
	list.Append(&tableItem{key, val})
}

func (ht *HashTable) pos(key string) int {
	hash := 0
	for _, char := range key {
		hash = hash<<5 + int(char)
	}
	return hash % len(ht.bucket)
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	pos := ht.pos(key)
	list := ht.bucket[pos]
	if list == nil {
		return nil, false
	}
	found, idx := find(list, key)
	return found, idx > -1
}

func (ht *HashTable) Delete(key string) bool {
	pos := ht.pos(key)
	list := ht.bucket[pos]
	if list == nil {
		return false
	}
	_, idx := find(list, key)
	if idx == -1 {
		return false
	}
	if ht.size == len(ht.bucket)-2 {
		ht.resize()
	}
	list.RemoveAt(idx)
	ht.size--
	return true
}

func (ht *HashTable) Each(iteratee func(interface{}, string)) {
	each(ht.bucket, iteratee)
}

func (ht *HashTable) Size() int {
	return ht.size
}

func (ht *HashTable) IsEmpty() bool {
	return ht.size == 0
}

func (ht *HashTable) Keys() []string {
	keys := make([]string, ht.size)
	each(ht.bucket, func(_ interface{}, key string) {
		keys = append(keys, key)
	})
	return keys
}

func (ht *HashTable) Values() []interface{} {
	values := make([]interface{}, ht.size)
	each(ht.bucket, func(val interface{}, _ string) {
		values = append(values, val)
	})
	return values
}

func find(list *linkedlist.LinkedList, key string) (interface{}, int) {
	found, idx := list.Find(func(val interface{}, i int) bool {
		return val.(*tableItem).key == key
	})
	return found.(*tableItem).val, idx
}

func each(bucket []*linkedlist.LinkedList, iteratee func(interface{}, string)) {
	for _, list := range bucket {
		if list != nil {
			list.Each(func(val interface{}, i int) {
				item := val.(*tableItem)
				iteratee(item.val, item.key)
			})
		}
	}
}
