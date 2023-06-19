package set

import (
	"golang.org/x/exp/constraints"
)

type Element[Tkey constraints.Ordered, Tvalue any] struct {
	Key   Tkey
	Value any
}

type SparseSet[Tkey constraints.Ordered, Tvalue any] struct {
	dense  []Element[Tkey, Tvalue]
	sparse map[Tkey]int
}

type Iterator[Tkey constraints.Ordered, Tvalue any] struct {
	dense []Element[Tkey, Tvalue]
	idx   int
}

func NewSpareSet[Tkey constraints.Ordered, Tvalue any]() *SparseSet[Tkey, Tvalue] {

	return &SparseSet[Tkey, Tvalue]{
		sparse: make(map[Tkey]int),
	}
}

func (s *SparseSet[Tkey, Tvalue]) Iterator() *Iterator[Tkey, Tvalue] {

	return &Iterator[Tkey, Tvalue]{
		dense: s.dense,
		idx:   0,
	}
}

func (i *Iterator[Tkey, Tvalue]) IsEnd() bool {
	return i.idx >= len(i.dense)
}

func (i *Iterator[Tkey, Tvalue]) Next() {
	i.idx += 1
}

func (i *Iterator[Tkey, Tvalue]) Get() Element[Tkey, Tvalue] {
	return i.dense[i.idx]
}

func (s *SparseSet[Tkey, Tvalue]) Add(key Tkey, value Tvalue) {

	if idx, ok := s.sparse[key]; ok {
		s.dense[idx].Value = value
		return
	}

	s.dense = append(s.dense, Element[Tkey, Tvalue]{
		Key:   key,
		Value: value,
	})

	s.sparse[key] = len(s.dense) - 1
}

func (s *SparseSet[Tkey, Tvalue]) Get(key Tkey) (Tvalue, bool) {

	if idx, ok := s.sparse[key]; ok {
		value := s.dense[idx].Value
		return value.(Tvalue), true
	}
	var t Tvalue
	return t, false
}

func (s *SparseSet[Tkey, Tvalue]) Remove(key Tkey) bool {

	if idx, ok := s.sparse[key]; ok {
		last := len(s.dense) - 1

		if idx < last {
			s.dense[idx] = s.dense[last]
			s.sparse[s.dense[idx].Key] = idx

		}

		s.dense = s.dense[:last]
		delete(s.sparse, key)
		return true

	}

	return false
}
