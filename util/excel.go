package util

import (
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	uuid "github.com/satori/go.uuid"
)

func XLSXGen(datamap []map[string]string) string {
	f := excelize.NewFile()
	f.SetColWidth("Sheet1", "A", "D", 20)
	f.SetCellValue("Sheet1", "A1", "编号")
	f.SetCellValue("Sheet1", "B1", "页面")
	f.SetCellValue("Sheet1", "C1", "用户")
	f.SetCellValue("Sheet1", "D1", "访问时间")
	style, _ := f.NewStyle(`{
		"border": [
			{
				"type": "left",
				"color": "000000",
				"style": 2
			},
			{
				"type": "right",
				"color": "000000",
				"style": 2
			},
			{
				"type": "top",
				"color": "000000",
				"style": 2
			},
			{
				"type": "bottom",
				"color": "000000",
				"style": 2
			}
		]
	}`)
	f.SetCellStyle("Sheet1", "A1", "D1", style)
	for i := 0; i < len(datamap); i++ {
		col := strconv.Itoa(i + 2)
		f.SetCellValue("Sheet1", "A"+col, i+1)
		f.SetCellValue("Sheet1", "B"+col, datamap[i]["page"])
		f.SetCellValue("Sheet1", "C"+col, datamap[i]["user"])
		f.SetCellValue("Sheet1", "D"+col, datamap[i]["visit_time"])
	}
	uid := uuid.NewV4().String()
	filename := uid + ".xlsx"
	_ = os.Mkdir("./temp", os.ModePerm)
	f.SaveAs("./temp/" + filename)
	return filename
}
