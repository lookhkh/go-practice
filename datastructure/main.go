package main

import (
	"fmt"
	"io"
	"os"
)

type Room struct {
	RoomNum  int
	IsBooked bool
}

func (r *Room) Book() bool {

	if r.IsBooked {
		return false
	}

	r.IsBooked = true

	return true
}

type HotelReservation struct {
	Map map[int]*Room
}

func NewHotelReservationSystem(room []int) *HotelReservation {
	r := &HotelReservation{
		Map: make(map[int]*Room),
	}

	for _, v := range room {
		r.Map[v] = &Room{
			RoomNum:  v,
			IsBooked: false,
		}
	}

	return r

}

func (h *HotelReservation) Iterate() []Room {

	rooms := make([]Room, 0)

	for _, v := range h.Map {
		rooms = append(rooms, *v)
	}

	return rooms
}

func (h *HotelReservation) Book(num int) bool {
	if v, ok := h.Map[num]; ok {
		if v.IsBooked {
			return false
		}

		v.IsBooked = true
		return true
	}

	return false
}

func (h *HotelReservation) Get(num int) (*Room, bool) {
	if v, ok := h.Map[num]; ok {
		return v, true
	}

	return &Room{}, false

}

func PrintAll(w io.Writer, arr []Room) {

	for _, v := range arr {
		fmt.Fprintln(w, v)
	}

}

func main() {

	fmt.Println("Hello World")
	res := NewHotelReservationSystem(append([]int{}, 1, 2, 3, 4, 5))
	PrintAll(os.Stdout, res.Iterate())

}
