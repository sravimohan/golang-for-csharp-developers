package main

import "fmt"

func main() {
	status := "ok"

	// if
	if status == "ok" {
		fmt.Println("If", status)
	}

	// if else
	status = "bad"
	if status == "ok" {
		fmt.Println("If", status)
	} else {
		fmt.Println("If else", status)
	}

	// switch
	// - break is implicit in the case block
	// - break statement will break out of switch
	// - fallthrough
	id := 0
	fmt.Print("switch")
	switch id {
	case 0:
		fmt.Print(" : ", 0)
		fallthrough
	case 1:
		fmt.Print(" : ", 1)
	case 2, 3: // multiple case in a single block
		fmt.Print(" : ", 2, 3)
	default:
		fmt.Print(" : ", "invalid")
	}

	// for
	fmt.Print("\nfor")
	for i := 0; i < 10; i++ {
		fmt.Print(" : ", i)
	}

	// nested for with lable
	// - break
	// - labeled break - used for nested loops
	fmt.Print("\nnested for")
iloop:
	for i := 0; i < 5; i++ {
		if i > 1 {
			fmt.Print(" : ", i, " : ", "break i")
			break
		}
		for j := 0; j < 5; j++ {

			if j > 2 {
				fmt.Print(" : ", i, j, " : ", "break j")
				break iloop
			}

			fmt.Print(" : ", i, j)
		}
	}

	// for each
	ids := [3]int{1, 2, 3}
	fmt.Print("\nfor each")
	for _, id := range ids {
		fmt.Print(" : ", id)
	}

	// while with for
	w := 0
	fmt.Print("\nwhile with for")
	for w < 3 {
		fmt.Print(" : ", w)
		w++
	}

	// do while with for
	w = 0
	fmt.Print("\ndo while with for")
	for {
		fmt.Print(" : ", w)
		w++

		if w > 3 {
			break
		}
	}

	fmt.Print("\n")
}
