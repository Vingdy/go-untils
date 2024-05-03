package file

import (
	"github.com/tealeg/xlsx"
	"go_untils/untilsError"
	"go_untils/valid"
)

type SheetData struct {
	SheetName string
	Data      [][]string
}

// GetXlsxFileExtract
// return: XlsxData
// - sleetName 	string
// - Data	    [Row][Col]string
// 结果为长方体，无数据部分为空
func GetXlsxFileExtract(excelFileName string) (sheetData []*SheetData, err error) {
	err = valid.ExcelFileNameValid(excelFileName)
	if err != nil {
		return nil, err
	}
	// 打开 xlsx 文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return nil, untilsError.NewError("Error opening file: %v", err)
	}

	// 遍历所有工作表
	for _, sheet := range xlFile.Sheets {

		oneSheet := &SheetData{
			SheetName: sheet.Name,
		}

		oneSheetData := [][]string{}
		maxCol := sheet.MaxCol
		// 遍历所有行
		for i, row := range sheet.Rows {
			oneSheetData = append(oneSheetData, make([]string, maxCol))
			// 遍历行中的所有单元格
			for j, cell := range row.Cells {
				text := cell.String()
				oneSheetData[i][j] = text
			}
		}

		oneSheet.Data = oneSheetData
		sheetData = append(sheetData, oneSheet)
	}
	return sheetData, nil
}
