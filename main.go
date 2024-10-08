package main

import (
	"fmt"
)

type Data struct {
	reference []int
	frames    []int
}

// FIFO Algorithm
func fifo(reference []int, frames int) int {
	fmt.Println("** FIFO")
	memory := make([]int, 0, frames)
	pageFaults := 0

	for _, page := range reference {
		found := false
		for _, memPage := range memory {
			if memPage == page {
				found = true
				break
			}
		}

		if !found {
			if len(memory) < frames {
				memory = append(memory, page)
			} else {
				memory = memory[1:]
				memory = append(memory, page)
			}
			pageFaults++
			fmt.Println("\t- ", memory)
		}
	}
	return pageFaults
}

// OPT Algorithm
func opt(reference []int, frames int) int {
	fmt.Println("** OPT")

	memory := make([]int, 0, frames)
	pageFaults := 0

	for i, page := range reference {
		found := false
		for _, memPage := range memory {
			if memPage == page {
				found = true
				break
			}
		}

		if !found {
			if len(memory) < frames {
				memory = append(memory, page)
			} else {
				futureUse := make(map[int]int)
				for j := i + 1; j < len(reference); j++ {
					if _, found := futureUse[reference[j]]; !found && contains(memory, reference[j]) {
						futureUse[reference[j]] = j
					}
				}

				// Find the page to replace
				if len(futureUse) < len(memory) {
					for _, memPage := range memory {
						if _, found := futureUse[memPage]; !found {
							memory = remove(memory, memPage)
							break
						}
					}
				} else {
					farthest := -1
					pageToRemove := -1
					for page, index := range futureUse {
						if index > farthest {
							farthest = index
							pageToRemove = page
						}
					}
					memory = remove(memory, pageToRemove)
				}
				memory = append(memory, page)
			}
			pageFaults++
		}
	}
	return pageFaults
}

// LRU Algorithm
func lru(reference []int, frames int) int {
	memory := make([]int, 0, frames)
	pageFaults := 0
	usedRecently := make([]int, 0, frames)

	for _, page := range reference {
		found := false
		for _, memPage := range memory {
			if memPage == page {
				found = true
				break
			}
		}

		if !found {
			if len(memory) < frames {
				memory = append(memory, page)
			} else {
				lruPage := usedRecently[0]
				memory = remove(memory, lruPage)
				usedRecently = usedRecently[1:]
				memory = append(memory, page)
			}
			pageFaults++
		} else {
			usedRecently = remove(usedRecently, page)
		}
		usedRecently = append(usedRecently, page)
	}
	return pageFaults
}

// Helper function to check if an element is in a slice
func contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// Helper function to remove an element from a slice
func remove(slice []int, element int) []int {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	references := []Data{
		{
			reference: []int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2, 1, 2, 0, 1, 7, 0, 1},
			frames:    []int{3, 4},
		},
		{
			reference: []int{1, 2, 3, 4, 2, 1, 5, 6, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1},
			frames:    []int{3, 4},
		},
		{
			reference: []int{3, 2, 1, 4, 5, 2, 1, 3, 6, 7, 2, 1, 3, 4, 5, 6, 7, 8, 1, 2},
			frames:    []int{3, 4},
		},
	}

	for i, ints := range references {
		fmt.Println("* Soy el ", i+1, "")

		for _, frame := range ints.frames {

			fmt.Println("- Con el frame ", frame)
			fmt.Printf("\t FIFO Page Faults: %d\n", fifo(ints.reference, frame))
			fmt.Printf("\t OPT Page Faults: %d\n", opt(ints.reference, frame))
			fmt.Printf("\t LRU Page Faults: %d\n", lru(ints.reference, frame))
		}
	}

}
