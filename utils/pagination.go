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
	paging.LastPage = CalcLastPage(paging.Total, paging.PerPage)

	// kalkulasi untuk proses pemotongan array
	offset := CalcOffset(paging.PerPage, paging.CurrentPage)
	offsetEnd := offset + paging.PerPage
	if offsetEnd > paging.Total {
		offsetEnd = paging.Total
	}

	// kondisi sebelum melakukan pagination
	if paging.CurrentPage <= 0 || offset > offsetEnd {
		return paging
	}

	// paging sesuai data
	arrData = arrData.Slice(offset, offsetEnd)

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

func CalcLastPage(total, perPage int) int {
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	if lastPage < 0 {
		lastPage = 1
	}
	return lastPage
}
