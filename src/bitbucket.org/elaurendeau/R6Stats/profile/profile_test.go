package profile

import "testing"

func Test_process_valid(*testing.T) {

	_, err := Process(Request{Name: "minthok", Platform: "uplay"})

	if err != nil  {
		panic("Expected result shouldn't have an err")
	}

}

func Test_process_invalid(*testing.T) {

	_, err := Process(Request{Name: "A", Platform: "test"})

	if err == nil  {
		panic("Process shouldn't return an err")
	}

}