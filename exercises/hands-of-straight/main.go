package main

import (
	"fmt"
	"sort"
)

func main() {
	group := []int{1, 2, 3, 6, 2, 3, 4, 7, 8, 0, 7, 9}
	isGroup := isNStraightHand(group, 3)
	fmt.Println("This is the result =>", isGroup)
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := make(map[int]int)
	for _, char := range hand {
		count[char]++
	}
	// make a slice of cards
	cards := make([]int, 0, len(count))
	for key := range count {
		cards = append(cards, key)
	}
	sort.Ints(cards)
	for _, evrCard := range cards {
		if count[evrCard] > 0 {
			for i := 1; i < groupSize; i++ {
				if count[evrCard+i] < count[evrCard] {
					return false
				}
				count[evrCard+i] -= count[evrCard]
			}
		}
	}

	fmt.Println("The length of the map is,", len(count))
	for key, val := range count {
		fmt.Printf("Card, %v appears, %v times\n", key, val)
	}
	return true
}
