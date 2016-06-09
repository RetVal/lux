package tornago

import (
	"io/ioutil"
	"os"
	"testing"
)

// seriously this is just here to reach the 100% test coverage.
func TestCoverage100(t *testing.T) {
	printer = ioutil.Discard
	d()
	df("")
	printer = os.Stdout
}
