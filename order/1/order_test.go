package order_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCompareMaps(t *testing.T) {
	t.Parallel()
	want := map[int]bool{1: true, 2: false}
	got := map[int]bool{2: false, 1: true}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCompareSlices(t *testing.T) {
	t.Parallel()
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCompareSortedSlices(t *testing.T) {
	t.Parallel()
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}
	sort := cmpopts.SortSlices(func(a, b int) bool {
		return a < b
	})
	if !cmp.Equal(want, got, sort) {
		t.Error(cmp.Diff(want, got, sort))
	}
}
