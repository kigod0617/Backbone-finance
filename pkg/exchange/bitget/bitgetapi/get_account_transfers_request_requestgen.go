// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Data -url /api/spot/v1/account/transferRecords -type GetAccountTransfersRequest -responseDataType []Transfer"; DO NOT EDIT.

package bitgetapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (g *GetAccountTransfersRequest) CoinId(coinId int) *GetAccountTransfersRequest {
	g.coinId = coinId
	return g
}

func (g *GetAccountTransfersRequest) FromType(fromType AccountType) *GetAccountTransfersRequest {
	g.fromType = fromType
	return g
}

func (g *GetAccountTransfersRequest) After(after string) *GetAccountTransfersRequest {
	g.after = after
	return g
}

func (g *GetAccountTransfersRequest) Before(before string) *GetAccountTransfersRequest {
	g.before = before
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetAccountTransfersRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetAccountTransfersRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check coinId field -> json key coinId
	coinId := g.coinId

	// assign parameter of coinId
	params["coinId"] = coinId
	// check fromType field -> json key fromType
	fromType := g.fromType

	// TEMPLATE check-valid-values
	switch fromType {
	case AccountExchange, AccountContract:
		params["fromType"] = fromType

	default:
		return nil, fmt.Errorf("fromType value %v is invalid", fromType)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of fromType
	params["fromType"] = fromType
	// check after field -> json key after
	after := g.after

	// assign parameter of after
	params["after"] = after
	// check before field -> json key before
	before := g.before

	// assign parameter of before
	params["before"] = before

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetAccountTransfersRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetAccountTransfersRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetAccountTransfersRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetAccountTransfersRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetAccountTransfersRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetAccountTransfersRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetAccountTransfersRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (g *GetAccountTransfersRequest) Do(ctx context.Context) ([]Transfer, error) {

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "/api/spot/v1/account/transferRecords"

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	var data []Transfer
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return data, nil
}
