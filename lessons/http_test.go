package lessons

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandle(t *testing.T) {
	p := indexPage{
		name:  "Tester",
		title: "Sir",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	indexHandle(w, r, p)

	body, _ := ioutil.ReadAll(w.Body)

	e := "Hello Sir Tester\n"

	if string(body) != e {
		t.Errorf("expected %q got %q", e, body)
	}
}
