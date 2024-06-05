package testing_init

import (
	"os"
	"path"
	"runtime"
)

// Make root the default directory while testing

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
