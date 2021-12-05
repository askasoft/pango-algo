package sortx

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	f := func(data Interface) {
		ShellSort(data)
	}
	testSortFunc(t, "ShellSort", f)
}

func TestShellKunthSort(t *testing.T) {
	f := func(data Interface) {
		ShellKunthSort(data)
	}
	testSortFunc(t, "ShellKunthSort", f)
}
