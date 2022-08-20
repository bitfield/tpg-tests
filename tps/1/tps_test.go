package tps_test

import (
	"os"
	"testing"

	"tps"

	"github.com/google/go-cmp/cmp"
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
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
