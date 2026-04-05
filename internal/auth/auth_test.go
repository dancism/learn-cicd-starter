package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}

	tests := []test{
		{input: http.Header{"Authorization": {"ApiKey asfsdfsfsfergergrg"}}, want: "asfsdfsfsfergergrg"},
		{input: http.Header{"Authorization": {"ApiKey sadasdasd"}}, want: "sadasdasd"},
		{input: http.Header{"Authorization": {"ApiKey"}}, want: ""},
		{input: http.Header{"Authorization": {"asdasd sadasdasd"}}, want: ""},
	}

	for i, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if tc.want != got {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}
