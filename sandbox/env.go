package sandbox

import (
	"fmt"
	"os"
)

// ListEnvs all env vars
func ListEnvs() {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
