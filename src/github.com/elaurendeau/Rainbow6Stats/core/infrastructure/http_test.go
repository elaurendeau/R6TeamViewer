package infrastructure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestValidHttpRequest(t *testing.T) {

	expectedContent := "premade http body content"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedContent)
	}))
	defer ts.Close()

	handler := HttpHandler{}

	httpContent, err := handler.Get(ts.URL)
	fmt.Println(ts.URL)

	assert.Nil(t, err)

	assert.Nil(t, err)
	assert.Equal(t, expectedContent, httpContent.Content)
	assert.Equal(t, "200 OK", httpContent.Status)
	assert.Equal(t, 200, httpContent.StatusCode)

}

func TestInvalidHttpRequest(t *testing.T) {

	handler := HttpHandler{}

	_, err := handler.Get("")

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

	_, err := handler.Get(ts.URL)

	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "Client.Timeout exceeded while awaiting headers"))

}
