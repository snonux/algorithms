package search

type GoMap map[int]int

func NewGoMap() GoMap {
	return make(GoMap)
}

func (m GoMap) Empty() bool {
	return m.Size() == 0
}

func (m GoMap) Size() int {
	return len(m)
}

func (m GoMap) Put(key, val int) {
	m[key] = val
}

func (m GoMap) Get(key int) (int, error) {
	val, ok := m[key]
	if !ok {
		return -1, NotFound
	}
	return val, nil
}

func (m GoMap) Del(key int) (int, error) {
	val, ok := m[key]
	if !ok {
		return -1, NotFound
	}
	delete(m, key)
	return val, nil
}
