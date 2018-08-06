package main

import (
	"reflect"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{
			arg:  "Brian",
			want: "Hello Brian!",
		}, {
			arg:  "Mary",
			want: "Hello David!",
		},
	}
	for _, test := range tests {
		if got := Greet(test.arg); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Test failed: got %+v, want %+v", got, test.want)
		}
	}
}
