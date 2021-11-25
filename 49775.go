package main

/*
#cgo LDFLAGS: -lcpp49775
#include <stdlib.h>
#include "49775.h"

void free_data(SomeData* data) {
    free(data->data);
	data->len = 0;
}

void free_float_arrary(float * arr)
{
	free(arr);
}
*/
import "C"
import "fmt"

func main() {

	_49775()

	fmt.Println("I am telling A truth!")
}

func _49775() {
	/*
		case 0: msvc malloc
		gcc free  X
		msvc free OK
	*/
	// var sd C.SomeData = C.new_data()

	// defer C.delete_data(&sd) // OK
	// defer C.free_data(&sd) // oops

	/*
		case 1: no malloc
		gcc free OK
		msvc free OK
	*/
	// var sd C.SomeData

	// defer C.delete_data(&sd) // OK
	// defer C.free_data(&sd) // OK
	/*
		case 2: gcc malloc
		gcc free OK
		msvc free X
	*/

	// var sd C.SomeData
	// sd.data = C.CString("i am telling a truth")

	// defer C.delete_data(&sd) // oops
	// defer C.free_data(&sd) // OK

	/*
		case 3: msvc malloc a non-struct data
		gcc free OK
		msvc free X
	*/

	var arr = C.new_float_arrary()
	// C.delete_float_arrary(arr)
	C.free_float_arrary(arr)

	var _ = arr
}
