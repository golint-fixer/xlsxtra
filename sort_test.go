package xlsxtra_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stanim/xlsxtra"
)

// ExampleSort demonstrates multi column sort
func ExampleSort() {
	sheet, err := xlsxtra.OpenSheet(
		"xlsxtra_test.xlsx", "sort_test.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	// multi column sort
	xlsxtra.Sort(sheet, 1, -1,
		3,  // last name
		-2, // first name (reverse order)
		7,  // test empty column
		6,  // ip address
	)
	for _, row := range sheet.Rows {
		fmt.Println(
			strings.Join(xlsxtra.ToString(row.Cells), ", "))
	}
	fmt.Println()
	// by header
	col := xlsxtra.NewCol(sheet, 1)
	err = xlsxtra.SortByHeaders(sheet, 1, -1, col,
		"-amount", // reverse order
		"first_name",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range sheet.Rows {
		fmt.Println(strings.Join(xlsxtra.ToString(row.Cells), ", "))
	}

	// Output:
	// id, first_name, last_name, email, gender, amount
	// 9, Donald, Bryant, lharper8@wunderground.com, Female, 100000000
	// 7, Donald, Bryant, dbryant6@redcross.org, Male, 3000000
	// 10, Donald, Bryant, hmarshall9@stumbleupon.com, Male, € 9
	// 4, Teresa, Hunter, thall3@arizona.edu, Female, 6000
	// 5, Joshua, Hunter, jstone4@google.cn, Male, 50000
	// 8, Jacqueline, Hunter, jfields7@dagondesign.com, Female, 20000000
	// 2, Harry, Hunter, hhunter1@webnode.com, Male, 80
	// 6, Rose, Spencer, rjohnson5@odnoklassniki.ru, Female, 400000
	// 1, Jimmy, Spencer, jspencer0@cnet.com, Male, 9
	// 3, Benjamin, Spencer, bmorgan2@unblog.fr, Male, 700
	//
	// id, first_name, last_name, email, gender, amount
	// 9, Donald, Bryant, lharper8@wunderground.com, Female, 100000000
	// 8, Jacqueline, Hunter, jfields7@dagondesign.com, Female, 20000000
	// 7, Donald, Bryant, dbryant6@redcross.org, Male, 3000000
	// 6, Rose, Spencer, rjohnson5@odnoklassniki.ru, Female, 400000
	// 5, Joshua, Hunter, jstone4@google.cn, Male, 50000
	// 4, Teresa, Hunter, thall3@arizona.edu, Female, 6000
	// 3, Benjamin, Spencer, bmorgan2@unblog.fr, Male, 700
	// 2, Harry, Hunter, hhunter1@webnode.com, Male, 80
	// 10, Donald, Bryant, hmarshall9@stumbleupon.com, Male, € 9
	// 1, Jimmy, Spencer, jspencer0@cnet.com, Male, 9
}

// TestSort checks not existing header
func TestSort(t *testing.T) {
	sheet, err := xlsxtra.OpenSheet(
		"xlsxtra_test.xlsx", "sort_test.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	col := xlsxtra.NewCol(sheet, 1)
	err = xlsxtra.SortByHeaders(sheet, 1, -1, col,
		"not existing",
	)
	if err == nil {
		t.Fatal("TestSort: expected error for SortByHeaders")
	}
}
