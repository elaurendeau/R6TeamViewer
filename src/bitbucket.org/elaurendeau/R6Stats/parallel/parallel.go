package parallel

import (
	"sync"
)


//ProcessData is the type of a function that takes a struct as an argument and returns a structure
type ProcessData func(struct{}) struct{}

//Process is a function that takes a slice of struct{} and a function. It will apply the function to each string and return a slice of struct {}
func Process(dataSlice []struct{}, f ProcessData) <-chan struct{} {

	in := groupData(dataSlice)

	out1 := split(in, f)
	out2 := split(in, f)
	out3 := split(in, f)

	return merge(out1, out2, out3)
}

func groupData(dataSlice []struct{}) <-chan struct{} {
	out := make(chan struct{})

	go func() {
		defer close(out)
		for _, data := range dataSlice {
			out <- data
		}
	}()

	return out
}

func split(in <- chan struct{}, f ProcessData) <- chan struct{} {
	out := make(chan struct{})

	go func() {
		for data := range in {
			out <- f(data)
		}

		close(out)
	}()

	return out
}

func merge(in  ...<- chan struct{}) <- chan struct{} {
	out := make(chan struct{})
	var wg sync.WaitGroup

	output := func(channel <- chan struct{}) () {
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
