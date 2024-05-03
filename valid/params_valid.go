package valid

import (
	"errors"
	"go_untils/untilsError"
	"strings"
)

var (
	excelFileSuffix = []string{"xls", "xlsx", "xlsb", "xlsm", "scv"}
)

func PositiveValid(input int) (bool, error) {
	if input < 0 {
		return false, errors.New("input parameter must be an integer")
	}
	return true, nil
}

// ExcelFileNameValid
// Excel 文件类型包含 xls xlsx xlsb xlsm scv
func ExcelFileNameValid(input string) error {
	for _, suffix := range excelFileSuffix {
		res := strings.HasSuffix(input, suffix)
		if res {
			return nil
		}
	}
	return untilsError.NewString("input file name: %s suffix is illegal", input)
}
