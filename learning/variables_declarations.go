package learning

import "fmt"

func primitiveDeclarations() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Alisson"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}

func pointersDeclarations() {
	var firstName *string = new(string)
	*firstName = "Alisson"
	fmt.Println(*firstName)

	lastName := "Lima"
	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Carvalho"
	fmt.Println(ptr, *ptr) //Prints new value, but the memory locations is still the same
}

func constantsDeclarations() {
	//Constants must be declared and initialised at the same time
	const pi = 3.1415
	fmt.Println(pi)

	const c = 3
	//c doesn't have a type so it's implicitly declared at runtime
	fmt.Println(c + 1.2)
}

func arrayDeclarations() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//Slice is a kind of pointer, but is not a pointer
	slice := arr2[:]
	fmt.Println(slice)

	//Changes slice and arr2 because slice points to data arr2 is keeping
	slice[1] = 42
	fmt.Println(arr2)
	fmt.Println(slice)

	//Here slice doesn't have an array with defined size, so slice will
	//manage the array size
	sliceFlex := []int{1, 2, 3}
	fmt.Println(sliceFlex)

	//append creates a new slice with new values appended
	sliceFlex = append(sliceFlex, 4, 2, 3)
	fmt.Println(sliceFlex)

	//Creates a new slice starting at the index 1
	sliceFromSlice := sliceFlex[1:]
	fmt.Println(sliceFromSlice)

	//Creates a new slice starting at the beginning going up to index 2 not including it
	sliceFromSlice2 := sliceFlex[:2]
	fmt.Println(sliceFromSlice2)

}

func mapDeclarations() {
	//Between brackets is the type of map key, out brackets is the type of values
	m := map[string]int{"foo":42}
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

}
