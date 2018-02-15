package utils

import (
	"fmt"
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

	// kalkulasi untuk proses pemotongan array
	idxStart := paging.PerPage * (paging.CurrentPage - 1)
	idxFinish := idxStart + paging.PerPage
	if idxFinish > paging.Total {
		idxFinish = paging.Total
	}

	// kondisi sebelum melakukan pagination
	if paging.PerPage > paging.Total || paging.CurrentPage <= 0 || idxStart > idxFinish {
		return paging
	}

	// paging sesuai data
	arrData = arrData.Slice(idxStart, idxFinish)

	var source []interface{}
	if page > 0 && page <= paging.LastPage {
		for i := 0; i < arrData.Len(); i++ {
			dataTmp := make(map[string]interface{})
			for j := 0; j < arrData.Index(i).NumField(); j++ {
				dataTmp[arrData.Index(i).Type().Field(j).Name] = arrData.Index(i).Field(j).Interface()
			}
			source = append(source, dataTmp)
		}

	}
	paging.Data = source

	return paging
}
