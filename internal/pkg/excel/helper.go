package excel

import (
	"bytes"
	"io"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/xuri/excelize/v2"
)

type Helper struct {
	file   *excelize.File
	data   map[string][][]interface{}
	header map[string][]string
	logger *log.Helper
}

func (e *Helper) NewSheet(sheetname string) (int, error) {
	return e.file.NewSheet(sheetname)
}

func (e *Helper) SetActiveSheet(index int) {
	e.file.SetActiveSheet(index)
}

func (e *Helper) Header(sheetname string, header []string) *Helper {
	e.header[sheetname] = header
	return e
}

func (e *Helper) Data(sheetname string, data [][]interface{}) *Helper {
	e.data[sheetname] = data
	return e
}

func (e *Helper) WriteTo(w io.Writer) (n int64, err error) {

	row := 1
	for sheetname, data := range e.data {
		if header, ok := e.header[sheetname]; ok {
			for i, v := range header {
				k, err := ToExcelAxis(row, i+1)
				if err != nil {
					e.logger.Error(err)
					return 0, err
				}
				err = e.file.SetCellValue(sheetname, k, v)
				if err != nil {
					e.logger.Error(err)
					return 0, err
				}
			}
			row++
		}

		for _, columns := range data {
			for i, v := range columns {
				k, err := ToExcelAxis(row, i+1)
				if err != nil {
					e.logger.Error(err)
					return 0, err
				}
				err = e.file.SetCellValue(sheetname, k, v)
				if err != nil {
					e.logger.Error(err)
					return 0, err
				}
			}
			row++
		}
	}
	n, err = e.file.WriteTo(w)
	if err != nil {
		e.logger.Error(err)
	}
	return
}

func (e *Helper) WriteToBuffer() (*bytes.Buffer, error) {

	buf := make([]byte, 0)
	buffer := bytes.NewBuffer(buf)
	_, err := e.WriteTo(buffer)
	if err != nil {
		e.logger.Error(err)
		return nil, err
	}
	return buffer, nil
}

func (e *Helper) SaveAs(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			e.logger.Error(err)
		}
	}()
	_, err = e.WriteTo(file)
	if err != nil {
		e.logger.Error(err)
	}
	return err
}

func NewHelper(logger log.Logger) *Helper {
	return &Helper{
		file:   excelize.NewFile(),
		data:   make(map[string][][]interface{}, 0),
		header: make(map[string][]string),
		logger: log.NewHelper(logger),
	}
}

/** col 与 row 都是从1开始的*/
func ToExcelAxis(row int, col int) (string, error) {
	return excelize.CoordinatesToCellName(col, row)
}
