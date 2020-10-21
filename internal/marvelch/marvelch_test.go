package marvelch

import "testing"

func TestServiceRoutes(t *testing.T) {
	got := struct {
		Placeholder string
	}{
		Placeholder: "got",
	}
	want := struct {
		Placeholder string
	}{
		Placeholder: "want",
	}

	// TODO:
	// init test case data
	// create service
	// send http request to service
	// get response
	// compare response to testcase

	if want != got {
		t.Errorf("want %+v got %+v | fail\n", want, got)
	}
}
