package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var a [5]int
	fmt.Println("array:", a)

	a[0] = 1
	fmt.Println("after change array:", a)
	fmt.Println("array length:", len(a))

	// declare & initialize array
	b := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println("new array:", b)

	// two dimensional arrays
	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("two dimensional array: ", c)

	// Slices
	fmt.Println("Slices")

	d := make([]string, 3)
	fmt.Println(d)

	d[0] = "a"
	d[1] = "b"
	d[2] = "c"
	fmt.Println(d)

	fmt.Println("Slice operations")

	fmt.Println("Append operation")
	d = append(d, "d")
	d = append(d, "e")
	fmt.Println(d)

	e := make([]string, len(d))
	// copy into c from s
	fmt.Println("Copy operation")
	copy(e, d)
	fmt.Println("New copied slice", e)
	// slice from 1 included to 4 excluded
	f := e[1:4]
	fmt.Println("Sliced from 1 till 4th element:", f)

	f = e[3:]
	fmt.Println("Sliced from third element till the end", f)

	// remember: excluded!
	// It will show elements before(!) 4
	f = e[:4]
	fmt.Println("Sliced from the beginning till the 4th excluded element", f)

	fmt.Println("Two dimensional slices")
	g := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerCounter := i + 1
		g[i] = make([]int, innerCounter)
		for j := 0; j < innerCounter; j++ {
			g[i][j] = i + j
		}
	}
	fmt.Println("Two dimensional result slice: ", g)

	fmt.Println("Worling with Maps")
	// short desc: make(map[keyType]valueType)
	firstMap := make(map[string]int)
	firstMap["key1"] = 3
	firstMap["key2"] = 4
	fmt.Println("Created and filled map:", firstMap)

	firstMap = map[string]int{"key1": 1, "key2": 2}
	fmt.Println("Declare & initialize syntax for map:", firstMap)

	fmt.Printf("Get value from map by key1: %v, with length of map: %v\n", firstMap["key1"], len(firstMap))

	fmt.Println("Map before deletion:", firstMap)
	delete(firstMap, "key2")
	fmt.Println("Map after deletion:", firstMap)

	// The optional second return value when getting a value from
	// a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys
	// with zero values like 0 or "". Here we didn’t need the value itself,
	// so we ignored it with the blank identifier _.
	_, deletedValue := firstMap["key2"]
	fmt.Println("value:", deletedValue)

	fmt.Println("Ranges")
	firstRange := []int{1, 2, 3}
	sum := 0
	for _, num := range firstRange {
		sum += num
	}
	fmt.Println("Sum is", sum)

	for i, num := range firstRange {
		if num == 2 {
			fmt.Printf("Index of current num: %v is %v\n", num, i)
		}
	}

	// range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	keyValuePair := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range keyValuePair {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// you can go through unicode sequence using plain symbols
	for i, c := range "go" {
		fmt.Printf("Index %v value is %v\n", i, c)
	}

}
