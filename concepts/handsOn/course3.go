package main

import "fmt"

func main() {
	slicey := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{56, 57, 58, 59, 60}

	for i, v := range slicey {
		fmt.Printf("At index %d we have %d\n", i, v)
	}
	fmt.Printf("Slice is of type %T\n", slicey)

	fmt.Println(slicey[:5])
	fmt.Println(slicey[5:])
	fmt.Println(slicey[2:7])
	fmt.Println(slicey[1:6])

	slicey = append(slicey, 52)
	slicey = append(slicey, 53, 54, 55)
	slicey = append(slicey, y...)
	fmt.Println(slicey)

	slicey = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slicey = append(slicey[:3], slicey[6:]...)
	fmt.Println(slicey)

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Hellooooo, James"}

	records := [][]string{jb, mp}

	for i, v := range records {
		fmt.Printf("Index %d is %v\n", i, v)
		for j, val := range v {
			fmt.Printf("\tAt index %d is %s\n", j, val)
		}
	}

	jMap := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Chocolate"},
		"no_dr":           []string{"Being Evel", "Ice Cream", "Sunsets"},
	}

	jMap["powers_austin"] = []string{"pees a lot", "bad teeth"}

	delete(jMap, "powers_austin")
	for i, v := range jMap {
		fmt.Println("Record: ", i)
		for j, val := range v {
			fmt.Printf("\tIndex %d is %s\n", j, val)
		}
	}
}
