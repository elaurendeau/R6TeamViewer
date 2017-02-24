package profile

import "testing"

func TestValidProcess(t *testing.T) {

	_, err := Process(Request{Name: "minthok", Platform: "uplay"})

	if err != nil  {
		t.Fail()
	}

}

func TestInvalidProcess(t *testing.T) {

	_, err := Process(Request{Name: "A", Platform: "test"})

	if err == nil  {
		t.Fail()
	}

}