package xlsx

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

// Xlsx struct provides a framework to set a .xlsx filename and sheets
type Xlsx struct {
	Filename string
	Sheets   []Sheet
}

// Sheet stores sheet name and data in a 2D string matrix
type Sheet struct {
	Name string
	Data [][]string
}

// WriteXlsx receives an Xlsx struct and writes is into an .xlsx file.
func WriteXlsx(x Xlsx) { // Output results to .xlsx file
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	for _, newSheet := range x.Sheets {
		sheet, err = file.AddSheet(newSheet.Name)
		if err != nil {
			fmt.Printf(err.Error())
		}
		for _, newRow := range newSheet.Data {
			row = sheet.AddRow()
			for _, cellVal := range newRow {
				cell = row.AddCell()
				cell.Value = cellVal
			}
		}

	}

	err = file.Save(x.Filename)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("Output file " + x.Filename + " created succesfully")
	}
}
