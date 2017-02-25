package parallel

import (
	"sync"
)



//Process is a function that takes a slice of interface{} and a function. It will apply the function to each string and return a slice of struct {}
func Process(dataSlice []interface{}, f func(interface{}) interface{}) <-chan interface{} {

	in := groupData(dataSlice)

	out1 := split(in, f)
	out2 := split(in, f)
	out3 := split(in, f)

	return merge(out1, out2, out3)
}

func groupData(dataSlice []interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		for _, data := range dataSlice {
			out <- data
		}
	}()

	return out
}

func split(in <- chan interface{}, f func(interface{}) interface{}) <- chan interface{} {
	out := make(chan interface{})

	go func() {
		for data := range in {
			out <- f(data)
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
