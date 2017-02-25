package infrastructure

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"github.com/stretchr/testify/assert"
	"time"
	"strings"
)

func TestValidHttpRequest(t *testing.T) {

	expectedContent := "premade http body content"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedContent)
	}))
	defer ts.Close()

	handler := HttpHandler{}

	httpContent, err := handler.get(ts.URL)
	fmt.Println(ts.URL)

	assert.Nil(t, err)

	assert.Nil(t, err)
	assert.Equal(t, expectedContent , httpContent.Content)

}

func TestInvalidHttpRequest(t *testing.T) {

	handler := HttpHandler{}

	_, err := handler.get("")

	assert.Error(t, err, "error is expected")
}


func TestTimeoutHttpRequest(t *testing.T) {

	expectedContent := "premade http body content"
	handler := HttpHandler{HttpTimeout: 1}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(handler.HttpTimeout * 2 * time.Second)
		fmt.Fprint(w, expectedContent)
	}))
	defer ts.Close()



	_, err := handler.get(ts.URL)

	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "Client.Timeout exceeded while awaiting headers"))


}

