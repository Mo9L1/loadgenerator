package service

import (
	"encoding/json"
	"io"
)

func (s *SvcImpl) ReqGetAllOrders() (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminorderservice/adminorder", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderArrResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqAddOrder(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminorderservice/adminorder", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqUpdateOrder(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminorderservice/adminorder", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqDeleteOrder(orderId string, trainNumber string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+"/api/v1/adminorderservice/adminorder/"+orderId+"/"+trainNumber, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
