package parallel

import (
	"sync"
)


//Process is the type of a function that takes a string as an argument and returns a structure
type Process func(string) struct{}

//Process is a function that takes a slice of string and a function. It will apply the function to each string and return a slice of struct {}
func Process(dataSlice []string, f Process) []struct{} {

	in := groupData(dataSlice)

	out1 := split(in)
	out2 := split(in)
	out3 := split(in)

	resultDataSlice := make([]struct{}, 0)

	for i := range merge(out1, out2, out3) {
		resultDataSlice = append(resultDataSlice, i)
	}

	return resultDataSlice
}

func groupData(dataSlice []string) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for _, data := range dataSlice {
			out <- data
		}
	}()

	return out
}

func split(in <- chan interface{}) <- chan interface{} {
	out := make(chan interface{})

	go func() {
		for data := range in {
			out <- data
		}

		close(out)
	}()

	return out
}

func merge(in  ...<- chan interface{}) <- chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup

	output := func(channel <- chan interface{}) () {
		for data := range channel {
			out <- data
		}

		wg.Done()
	}

	for _, channel := range in {
		go output(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	wg.Add(len(in))

	return out
}
