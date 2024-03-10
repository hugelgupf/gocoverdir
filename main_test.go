package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestXxx(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "bin")
	c := exec.Command("go", "build", "-covermode=atomic", "-o", bin, ".")
	if out, err := c.CombinedOutput(); err != nil {
		t.Logf("out: %s", out)
		t.Fatal(err)
	}

	t.Logf("gocoverdir: %s", os.Getenv("GOCOVERDIR"))

	c = exec.Command(bin)
	out, err := c.CombinedOutput()
	t.Logf("out: %s", out)
	if err != nil {
		t.Fatal(err)
	}
}
