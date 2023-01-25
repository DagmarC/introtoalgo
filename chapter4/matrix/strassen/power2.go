package strassen

import (
	"errors"
)

type PowersOf2 map[int]int

var powers2 PowersOf2 = map[int]int{
	1:    0,
	2:    1,
	4:    2,
	8:    3,
	16:   4,
	32:   5,
	64:   6,
	128:  7,
	256:  8,
	512:  9,
	1024: 10,
}
var powers2arr []int = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

func IsPower(n int) bool {
	if _, ok := powers2[n]; ok {
		return true
	}
	return false
}

func NearestPower(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("negative power is not supported")
	}
	if n == 0 {
		return 1, nil
	}

	// the only case when 1 is returned from the NearestPower is when you are looking for nearest power where n=0 -> 2ˆ0=1 is the nearest power
	// n=3 is the next possible case where you are looking for 2ˆ2=4, so it is enough to start from there
	for i := 2; i < len(powers2arr); i++ {
		ub := powers2arr[i]
		if ub > n {
			return ub, nil // First match - it is nearest power of 2
		}
	}
	return -1, errors.New("nearest power not found")

}
