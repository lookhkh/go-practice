package sorted

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Element[TKey constraints.Ordered, TValue any] struct {
	Key   TKey
	Value TValue
}

type SortedMap[TKey constraints.Ordered, TValue any] struct {
	Arr []Element[TKey, TValue]
}

func (s *SortedMap[TKey, TValue]) Add(key TKey, value TValue) {

	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr[idx].Value = value
		return
	}

	s.Arr = append(s.Arr,
		append([]Element[TKey, TValue]{{Key: key, Value: value}},
			s.Arr[idx:]...)...)

}

func (s *SortedMap[TKey, TValue]) Get(key TKey) (value TValue, ok bool) {

	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	var t TValue

	if len(s.Arr) <= idx {
		return t, false
	}

	if s.Arr[idx].Key == key {
		return s.Arr[idx].Value, true
	}

	return t, false
}

func (s *SortedMap[TKey, TValue]) Remove(key TKey) (ok bool) {

	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if len(s.Arr) <= idx {
		return false
	}

	if s.Arr[idx].Key == key {
		s.Arr = append(s.Arr[:idx], s.Arr[idx+1:]...)
		return true
	}

	return false
}
