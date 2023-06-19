package sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMap(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("aaa", 10)

	v, _ := s.Get("aaa")
	assert.Equal(t, v, 10)

	s.Add("bbb", 20)
	a, _ := s.Get("bbb")

	assert.Equal(t, a, 20)

	assert.Equal(t, "aaa", s.Arr[0].Key)

	assert.Equal(t, "bbb", s.Arr[1].Key)

}

func TestOverlaped(t *testing.T) {

	var s SortedMap[string, int]

	s.Add("aaa", 10)
	v, _ := s.Get("aaa")
	assert.Equal(t, 10, v)

	s.Add("aaa", 20)
	v, _ = s.Get("aaa")

	assert.Equal(t, 20, v)
	assert.Equal(t, 1, len(s.Arr))
}

func TestSortedGetEmpty(t *testing.T) {

	var s SortedMap[string, int]
	s.Add("bbb", 10)
	_, ok := s.Get("aaa")

	assert.Equal(t, ok, false)
}

func TestSortedGetEmptyAsc(t *testing.T) {

	var s SortedMap[string, int]
	s.Add("aaa", 10)
	_, ok := s.Get("bbb")

	assert.Equal(t, ok, false)
}

func TestRmove(t *testing.T) {

	var s SortedMap[string, int]
	s.Add("aaa", 10)
	_, ok := s.Get("aaa")

	assert.True(t, ok)

	s.Remove("aaa")

	v, removed := s.Get("aaa")

	assert.Equal(t, v, 0)

	assert.False(t, removed)

	notOk := s.Remove("bbb")
	assert.False(t, notOk)
}
