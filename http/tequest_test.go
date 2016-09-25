package http

import (
	"net/url"
	"testing"
)

func TestGetURLString(t *testing.T) {
	params := url.Values{}
	params.Add("test", "test")
	params.Add("test1", "test1")

	urlString := GetURLString("test", params)

	expected := "test?test=test&test1=test1"
	if urlString != expected {
		t.Errorf(
			"Wrong url string\nexpected: %s\nactual: %s",
			expected,
			urlString,
		)
	}
}
