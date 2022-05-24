// Code generated by "requestgen -method POST -url /api/v3/wallet/:walletType/orders -type WalletCreateOrderRequest -responseType .Order -debug"; DO NOT EDIT.

package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/c9s/bbgo/pkg/exchange/max/maxapi"
	"net/url"
	"reflect"
	"regexp"
)

func (w *WalletCreateOrderRequest) Market(market string) *WalletCreateOrderRequest {
	w.market = market
	return w
}

func (w *WalletCreateOrderRequest) Side(side string) *WalletCreateOrderRequest {
	w.side = side
	return w
}

func (w *WalletCreateOrderRequest) Volume(volume string) *WalletCreateOrderRequest {
	w.volume = volume
	return w
}

func (w *WalletCreateOrderRequest) OrderType(orderType string) *WalletCreateOrderRequest {
	w.orderType = orderType
	return w
}

func (w *WalletCreateOrderRequest) Price(price string) *WalletCreateOrderRequest {
	w.price = &price
	return w
}

func (w *WalletCreateOrderRequest) StopPrice(stopPrice string) *WalletCreateOrderRequest {
	w.stopPrice = &stopPrice
	return w
}

func (w *WalletCreateOrderRequest) ClientOrderID(clientOrderID string) *WalletCreateOrderRequest {
	w.clientOrderID = &clientOrderID
	return w
}

func (w *WalletCreateOrderRequest) GroupID(groupID string) *WalletCreateOrderRequest {
	w.groupID = &groupID
	return w
}

func (w *WalletCreateOrderRequest) WalletType(walletType WalletType) *WalletCreateOrderRequest {
	w.walletType = walletType
	return w
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (w *WalletCreateOrderRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (w *WalletCreateOrderRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check market field -> json key market
	market := w.market

	// TEMPLATE check-required
	if len(market) == 0 {
		return nil, fmt.Errorf("market is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of market
	params["market"] = market
	// check side field -> json key side
	side := w.side

	// TEMPLATE check-required
	if len(side) == 0 {
		return nil, fmt.Errorf("side is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of side
	params["side"] = side
	// check volume field -> json key volume
	volume := w.volume

	// TEMPLATE check-required
	if len(volume) == 0 {
		return nil, fmt.Errorf("volume is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of volume
	params["volume"] = volume
	// check orderType field -> json key ord_type
	orderType := w.orderType

	// assign parameter of orderType
	params["ord_type"] = orderType
	// check price field -> json key price
	if w.price != nil {
		price := *w.price

		// assign parameter of price
		params["price"] = price
	} else {
	}
	// check stopPrice field -> json key stop_price
	if w.stopPrice != nil {
		stopPrice := *w.stopPrice

		// assign parameter of stopPrice
		params["stop_price"] = stopPrice
	} else {
	}
	// check clientOrderID field -> json key client_oid
	if w.clientOrderID != nil {
		clientOrderID := *w.clientOrderID

		// assign parameter of clientOrderID
		params["client_oid"] = clientOrderID
	} else {
	}
	// check groupID field -> json key group_id
	if w.groupID != nil {
		groupID := *w.groupID

		// assign parameter of groupID
		params["group_id"] = groupID
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (w *WalletCreateOrderRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := w.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if w.isVarSlice(_v) {
			w.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (w *WalletCreateOrderRequest) GetParametersJSON() ([]byte, error) {
	params, err := w.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (w *WalletCreateOrderRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check walletType field -> json key walletType
	walletType := w.walletType

	// TEMPLATE check-required
	if len(walletType) == 0 {
		return nil, fmt.Errorf("walletType is required, empty string given")
	}
	// END TEMPLATE check-required

	// TEMPLATE check-valid-values
	switch walletType {
	case WalletTypeSpot, WalletTypeMargin:
		params["walletType"] = walletType

	default:
		return nil, fmt.Errorf("walletType value %v is invalid", walletType)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of walletType
	params["walletType"] = walletType

	return params, nil
}

func (w *WalletCreateOrderRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (w *WalletCreateOrderRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (w *WalletCreateOrderRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (w *WalletCreateOrderRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := w.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (w *WalletCreateOrderRequest) Do(ctx context.Context) (*max.Order, error) {

	params, err := w.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	apiURL := "/api/v3/wallet/:walletType/orders"
	slugs, err := w.GetSlugsMap()
	if err != nil {
		return nil, err
	}

	apiURL = w.applySlugsToUrl(apiURL, slugs)

	req, err := w.client.NewAuthenticatedRequest(ctx, "POST", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := w.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse max.Order
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
