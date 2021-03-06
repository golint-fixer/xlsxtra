package xlsxtra_test

import (
	"fmt"
	"testing"

	"github.com/stanim/xlsxtra"
	"github.com/tealeg/xlsx"
)

func ExampleColRange() {
	fmt.Println(xlsxtra.ColRange("X", "AD"))
	fmt.Println(xlsxtra.ColRange("1", "AD"))
	// Output:
	// [X Y Z AA AB AC AD]
	// []
}

func TestSplitCoord(t *testing.T) {
	column, row, err := xlsxtra.SplitCoord("AA11")
	if err != nil {
		t.Fatal(err)
	}
	if column != "AA" || row != 11 {
		t.Fatalf("expected \"AA\" and 11; got %q and %d",
			column, row)
	}
	_, _, err = xlsxtra.SplitCoord("A0")
	if err == nil {
		t.Fatal("expected error")
	}
}

func Example() {
	fmt.Println(xlsxtra.ColStr[26], xlsxtra.StrCol["AA"])
	// Output: Z 27
}

func ExampleAbs() {
	fmt.Println(xlsxtra.Abs("B12"))
	fmt.Println(xlsxtra.Abs("C5:G20"))
	fmt.Println(xlsxtra.Abs("Invalid"))
	// Output:
	// $B$12
	// $C$5:$G$20
	// Invalid
}

func ExampleRangeBounds() {
	fmt.Println(xlsxtra.RangeBounds("A1:E6"))
	fmt.Println(xlsxtra.RangeBounds("$A$1:$E$6"))
	fmt.Println(xlsxtra.RangeBounds("A1"))
	// invalid: no column name given
	fmt.Println(xlsxtra.RangeBounds("11:E6"))
	// invalid: row zero does not exist
	fmt.Println(xlsxtra.RangeBounds("A0:E6"))
	// Output:
	// 1 1 5 6 <nil>
	// 1 1 5 6 <nil>
	// 1 1 1 1 <nil>
	// 0 0 0 0 Invalid range "11:E6"
	// 0 0 0 0 Invalid range "A0:E6"
}

func TestTranspose(t *testing.T) {
	cells := [][]*xlsx.Cell{}
	if len(xlsxtra.Transpose(cells)) != 0 {
		t.Fatalf("Expected transpose of empty cell slice to be empty")
	}
}

func ExampleCoord() {
	fmt.Println(xlsxtra.Coord(2, 12))
	fmt.Println(xlsxtra.Coord(0, 12))
	// Output:
	// B12
	// ?12
}

func ExampleCoordAbs() {
	fmt.Println(xlsxtra.CoordAbs(2, 12))
	fmt.Println(xlsxtra.CoordAbs(0, 12))
	// Output:
	// $B$12
	// ?12
}
