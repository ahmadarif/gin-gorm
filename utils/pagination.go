package utils

import (
	"math"
	"reflect"
)

type Pagination struct {
	CurrentPage int         `json:"current_page"`
	PerPage     int         `json:"per_page"`
	LastPage    int         `json:"last_page"`
	Total       int         `json:"total"`
	Data        interface{} `json:"data"`
}

func Paginate(data interface{}, page, limit int) Pagination {
	paging := Pagination{}
	paging.CurrentPage = page
	paging.PerPage = limit

	arrData := reflect.ValueOf(data)
	paging.Total = arrData.Len()
	paging.LastPage = int(math.Ceil(float64(paging.Total) / float64(paging.PerPage)))
	if paging.LastPage < 0 {
		paging.LastPage = 1
	}

	// kalkulasi untuk proses pemotongan array
	idxStart := paging.PerPage * (paging.CurrentPage - 1)
	idxFinish := idxStart + paging.PerPage
	if idxFinish > paging.Total {
		idxFinish = paging.Total
	}

	// kondisi sebelum melakukan pagination
	if paging.CurrentPage <= 0 || idxStart > idxFinish {
		return paging
	}

	// paging sesuai data
	arrData = arrData.Slice(idxStart, idxFinish)

	// append array to data
	var source []interface{}
	for i := 0; i < arrData.Len(); i++ {
		dataTmp := make(map[string]interface{})
		for j := 0; j < arrData.Index(i).NumField(); j++ {
			dataTmp[arrData.Index(i).Type().Field(j).Name] = arrData.Index(i).Field(j).Interface()
		}
		source = append(source, dataTmp)
	}
	paging.Data = source

	return paging
}

func CalcOffset(perPage, page int) int {
	return perPage * (page - 1)
}
