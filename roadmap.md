# This file describe step by step every new feature that have been added

## Parsing the Excel file (.xlsx)

The purpose is to use a library (written in Go), like ["Excelize"](https://github.com/360EntSecGroup-Skylar/excelize), then Read the Excel file and create a slice out of it:

    data [][]string
