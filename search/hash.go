package search

type Hash struct {
	buckets  []*Elementary
	capacity int
	size     int
}

func NewHash(capacity int) *Hash {
	return &Hash{
		buckets:  make([]*Elementary, capacity),
		capacity: capacity,
	}
}

func (h *Hash) Empty() bool {
	return h.Size() == 0
}

func (h *Hash) Size() int {
	return h.size
}

func (h *Hash) hash(key int) int {
	i := key + key*2 + key<<10 + key>>2
	if i < 0 {
		i = -i
	}
	return i % h.capacity
}

func (h *Hash) Put(key int, val int) {
	i := h.hash(key)

	if h.buckets[i] == nil {
		elem := NewElementary()
		elem.Put(key, val)
		h.buckets[i] = elem
		return
	}

	h.buckets[i].Put(key, val)
}

func (h *Hash) Get(key int) (int, error) {
	i := h.hash(key)

	if h.buckets[i] == nil {
		return -1, NotFound
	}

	return h.buckets[i].Get(key)
}

func (h *Hash) Del(key int) (int, error) {
	i := h.hash(key)

	if h.buckets[i] == nil {
		return -1, NotFound
	}

	return h.buckets[i].Del(key)
}
