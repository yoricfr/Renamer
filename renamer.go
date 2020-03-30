package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("./filenames.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
  // skip the first line (header)
  for _, row := range rows[1:] {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
