package reception

import "errors"
import "sync"

type hotelRoom struct {
	uuid string
	empty bool
	size int
}

type Reception struct {
	mu sync.Mutex
	emptyRooms []int
	occupiedRooms map[int]int
}

func NewReception(n int) *Reception {
	// Initialize rooms
	rooms := make([]int, n)
	for i := 0; i < n; i++ {
		rooms[i] = i+1
	}
	return &Reception{
		emptyRooms: rooms,
		occupiedRooms: make(map[int]int),
	}
}

func (r *Reception) CheckIn(guestId int) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.emptyRooms) <= 0{
		return -1, errors.New("hotel has full capacity")
	}
	// r.emptyRooms = r.emptyRooms[1:] // O(n) - inefficient cuz slice is copying
	// Choose the last room not the first one and use SWAP and POP method
	n := len(r.emptyRooms) - 1
	room_number := r.emptyRooms[n]
	// Remove last element - just change last index
	r.emptyRooms = r.emptyRooms[:n]
	// Set room number occupied
	r.occupiedRooms[guestId] = room_number
	return room_number, nil
}

func (r *Reception) CheckOut(guestId int) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.occupiedRooms[guestId]; !ok {
		return -1, errors.New("room number already checked out or invalid")
	}
	room_number := r.occupiedRooms[guestId]
	delete(r.occupiedRooms, guestId) // free the room 
	r.emptyRooms = append(r.emptyRooms, room_number) // O(1)

	return room_number, nil
}