package main

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

func mockData() [][]string {
	//시간, 사용처, 금액, 비고
	return [][]string{
		{"2024-01-04 11:66", "하늘나라", "12000", "신용카드"},
		{"2024-01-03 11:66", "우리나라", "8000", "체크카드"},
		{"2023-06-02 11:66", "재미나라", "7000", "현금"},
		{"2023-03-01 11:66", "김밥나라", "13000", "신용카드"},
	}

}

func main() {
	f := excelize.NewFile()

	sheetName := "시트 네임"
	sheet, err := f.NewSheet(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	exampleData := mockData()

	// 해더 설정 하기
	f.SetCellValue(sheetName, "A1", "시간")
	f.SetCellValue(sheetName, "B1", "나라")
	f.SetCellValue(sheetName, "C1", "돈")
	f.SetCellValue(sheetName, "D1", "이용방법")

	//이후 데이터 뿌려 주기
	for i, row := range exampleData {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), row[0])
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), row[1])
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), row[2])
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), row[3])
	}

	filename := fmt.Sprintf("%s.xlsx", time.Now().Format("2006-01-02-1504"))

	f.SetActiveSheet(sheet)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
	}
	fmt.Println("생성 완료")
}
