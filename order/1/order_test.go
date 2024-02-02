package order_test

import (
	"maps"
	"slices"
	"testing"
)

func TestMapsHaveSameContents(t *testing.T) {
	t.Parallel()
	want := map[int]bool{1: true, 2: false}
	got := map[int]bool{2: false, 1: true}
	if !maps.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSlicesHaveSameElementsInSameOrder(t *testing.T) {
	t.Parallel()
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}
	if !slices.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSlicesHaveSameElementsInAnyOrder(t *testing.T) {
	t.Parallel()
	want := []int{1, 2, 3}
	got := []int{3, 2, 1}
	slices.Sort(want)
	slices.Sort(got)
	if !slices.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
