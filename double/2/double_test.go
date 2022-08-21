package double2_test

import (
	"fmt"
	"testing"

	double "double2"
)

func TestDouble2Returns4(t *testing.T) {
	t.Parallel()
	want := 4
	got := double.Double(2)
	if want != got {
		t.Errorf("Double(2): want %d, got %d", want, got)
	}
}

func ExampleDouble() {
	fmt.Println(double.Double(2))
	// Output:
	// 4
}

func ExampleDouble_with2() {
	fmt.Println(double.Double(2))
	// Output:
	// 4
}

func ExampleDouble_with3() {
	fmt.Println(double.Double(3))
	// Output:
	// 6
}
