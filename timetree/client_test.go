package timetree

import (
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	httpCli := new(http.Client)
	token := "test token"
	cli := NewClient(httpCli, token)

	if cli.Client() != httpCli {
		t.Fatalf("unexpected result. actual=%#v, expect=%#v", cli.Client(), httpCli)
	}
}
