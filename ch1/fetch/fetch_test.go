package fetch

import "testing"

func TestHelloName(t *testing.T) {
	Fetch([]string{"reqres.in/api/users"})
}
