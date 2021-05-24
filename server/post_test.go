package server

import "testing"

func TestCreatePostID(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "first post, hello world", want: "first-post-hello-world"},
		{in: "first post!hello world", want: "first-posthello-world"},
		{in: "first-post   !@#hello_world", want: "first-post-hello_world"},
	}

	for _, test := range tests {
		if got := idFromTitle(test.in); got != test.want {
			t.Errorf("got: %v; want: %v", got, test.want)
		}
	}
}
