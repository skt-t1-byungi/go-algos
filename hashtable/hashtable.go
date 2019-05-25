package hashtable

import "github.com/skt-t1-byungi/go-algos/linkedlist"

type hashTableItem struct {
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
	oldBucket := ht.bucket
	ht.bucket = make([]*linkedlist.LinkedList, len(oldBucket)*2)
	ht.size = 0
	eachItem(oldBucket, func(val interface{}, key string) {
		ht.Set(key, val)
	})
}

func (ht *HashTable) Set(key string, val interface{}) {
	pos := ht.pos(key)
	list := ht.bucket[pos]
	if list == nil {
		list = linkedlist.New()
		ht.bucket[pos] = list
	}

	if _, idx := search(list, key); idx > -1 {
		list.RemoveAt(idx)
	} else {
		ht.size++
	}

	list.Append(&hashTableItem{key, val})

	if ht.size == len(ht.bucket)+2 {
		ht.resize()
	}
}

func (ht *HashTable) pos(key string) int {
	hash := 0
	for _, char := range key {
		hash = hash<<4 + int(char)
	}
	return hash % len(ht.bucket)
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	list := ht.bucket[ht.pos(key)]
	if list == nil {
		return nil, false
	}
	found, idx := search(list, key)
	return found, idx > -1
}

func (ht *HashTable) Delete(key string) bool {
	list := ht.bucket[ht.pos(key)]
	if list == nil {
		return false
	}

	_, idx := search(list, key)
	if idx == -1 {
		return false
	}
	list.RemoveAt(idx)
	ht.size--

	if ht.size == len(ht.bucket)-2 {
		ht.resize()
	}

	return true
}

func (ht *HashTable) Each(iteratee func(val interface{}, key string)) {
	eachItem(ht.bucket, iteratee)
}

func (ht *HashTable) Size() int {
	return ht.size
}

func (ht *HashTable) IsEmpty() bool {
	return ht.size == 0
}

func (ht *HashTable) Keys() []string {
	keys := make([]string, ht.size)
	i := 0
	eachItem(ht.bucket, func(_ interface{}, key string) {
		keys[i] = key
		i++
	})
	return keys
}

func (ht *HashTable) Values() []interface{} {
	values := make([]interface{}, ht.size)
	i := 0
	eachItem(ht.bucket, func(val interface{}, _ string) {
		values[i] = val
		i++
	})
	return values
}

func search(list *linkedlist.LinkedList, key string) (interface{}, int) {
	found, idx := list.Find(func(val interface{}, i int) bool {
		return val.(*hashTableItem).key == key
	})

	if idx > -1 {
		return found.(*hashTableItem).val, idx
	} else {
		return nil, -1
	}
}

func eachItem(bucket []*linkedlist.LinkedList, iteratee func(val interface{}, key string)) {
	for _, list := range bucket {
		if list != nil {
			list.Each(func(val interface{}, i int) {
				item := val.(*hashTableItem)
				iteratee(item.val, item.key)
			})
		}
	}
}
