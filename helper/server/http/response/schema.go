package response

import (
	"reflect"
	"tronglv_upload_svc/helper/util"
)

type emptyStruct struct{}

// swagger:model SuccessResponse
type responseHttp struct {
	// Meta is the API response information
	// in: MetaResponse
	Meta metaResponse `json:"meta"`
	// Data is our data
	// in: DataResponse
	Data data `json:"data"`
}

// swagger:model MetaResponse
type metaResponse struct {
	// TradeId is the response trace_id
	// in: string
	TradeId string `json:"trace_id"`
	// Code is the response code
	// in: int
	Code string `json:"code"`
	// Message is the response message
	// in: string
	Message string `json:"message"`
	// Errors is the response message
	// in: string
	Errors any `json:"errors,omitempty"`
}

// swagger:model DataResponse
type data struct {
	Pagination any `json:"pagination,omitempty"`
	Records    any `json:"records,omitempty"`
	Record     any `json:"record,omitempty"`
}

func (i *data) SetData(result interface{}, paging any) data {
	isNil := util.IsZeroOfUnderlyingType(result)
	if isSlice := reflect.ValueOf(result).Kind() == reflect.Slice; isSlice {
		if isNil {
			result = []emptyStruct{}
		}
		i.Records = result
	} else {
		if isNil {
			result = emptyStruct{}
		}
		i.Record = result
	}
	i.Pagination = paging
	return *i
}
