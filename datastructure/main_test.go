package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	res := NewHotelReservationSystem(append([]int{}, 1, 2, 3, 4, 5))
	assert.Equal(t, len(res.Map), 5)

	v := res.Iterate()

	assert.Equal(t, len(v), 5)

}

func TestBook(t *testing.T) {

	res := NewHotelReservationSystem(append([]int{}, 1, 2, 3, 4, 5))
	ok := res.Book(6)
	assert.False(t, ok)

	ok = res.Book(1)
	assert.True(t, ok)

	ok = res.Book(1)
	assert.False(t, ok)

	book, ok := res.Get(1)

	assert.True(t, ok)
	assert.True(t, book.IsBooked)
}
