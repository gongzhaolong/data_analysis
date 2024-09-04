package generate_number

import (
	"data_analysis/constant"
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ReadData() [][]float64 {
	// 打开已有的 Excel 文件
	filePath := "data.xlsx" // 你的已有文件路径
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open Excel file: %v", err)
	}
	// 确保工作表存在
	sheetName := "data_real"
	// 获取所有行列的数据
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to get rows from sheet %s: %v", sheetName, err)
	}

	// 检查是否有数据
	if len(rows) < 2 {
		log.Fatalf("Sheet %s does not have enough data", sheetName)
	}

	// 按列读取数据，除去表头和第一列
	columns := make([][]float64, len(rows[0])-1)
	for _, row := range rows[2:] {
		for colIndex := 1; colIndex < len(row); colIndex++ {
			value, err := strconv.ParseFloat(row[colIndex], 64)
			if err != nil {
				log.Fatalf("Failed to convert cell value to int: %v", err)
			}
			columns[colIndex-1] = append(columns[colIndex-1], value)
		}
	}

	return columns
}
func UpdateSheet(column []float64) {
	// 平滑数据
	filePath := "data.xlsx" // 你的已有文件路径
	// 确保工作表存在
	sheetName := "data_smooth"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, sample := range column {
		cell := fmt.Sprintf("E%d", i+2) // A列，从第1行开始
		f.SetCellValue(sheetName, cell, sample)
	}
	// 保存 Excel 文件
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Data written to", filePath, "successfully.")
	}
}

func WriteToExcel(origin_data []float64, number, grade int, metric string, catgory int) {
	// 打开已有的Excel文件
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 指定工作表名称
	sheetName := constant.Sheetname

	// 获取工作表中已有的最大列数
	colCount, err := f.GetCols(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	newCol := string('A' + len(colCount)) // 新列的字母表示
	//写入表头
	cell := fmt.Sprintf("%s%d", newCol, 1) // 生成单元格地址，如 B1, B2, ...
	f.SetCellValue(sheetName, cell, strconv.Itoa(number)+metric)
	// 将数据写入新的列
	for i, v := range origin_data {
		cell = fmt.Sprintf("%s%d", newCol, i+3) // 生成单元格地址，如 B1, B2, ...
		f.SetCellValue(sheetName, cell, v)
	}
	cell = fmt.Sprintf("%s%d", newCol, 48) // 生成单元格地址，如 B1, B2, ...
	f.SetCellValue(sheetName, cell, grade)

	cell = fmt.Sprintf("%s%d", newCol, 49)
	f.SetCellValue(sheetName, cell, catgory)
	// 保存Excel文件
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("数据已成功写入Excel文件")

}
