package tps_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/bitfield/tps"
)

func TestWriteReportFile_ProducesCorrectOutputFile(t *testing.T) {
	t.Parallel()
	output := t.TempDir() + "/" + t.Name()
	tps.WriteReportFile(output)
	want, err := os.ReadFile("testdata/output.golden")
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}
