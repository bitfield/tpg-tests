package double_test

import (
	"double"
	"fmt"
	"testing"
)

func TestDouble2Returns4(t *testing.T) {
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

func ExampleDouble_with_2() {
	fmt.Println(double.Double(2))
	// Output:
	// 4
}

func ExampleDouble_with_3() {
	fmt.Println(double.Double(3))
	// Output:
	// 6
}
