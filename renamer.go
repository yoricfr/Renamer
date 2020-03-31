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
  var data [][2]string
  // skip the first line (header)
  for _, row := range rows[1:] {
    line := [2]string{"",""}
		for i, colCell := range row {
      line[i] = colCell
		}
    // skip the last empry line
    if line[0] != "" && line[1] != "" {
      data = append(data, line)
    }
	}
  fmt.Println(data)
}
