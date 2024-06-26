package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type UserDto struct {
	UserID       string `json:"userId"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Gender       int    `json:"gender"`
	DocumentType int    `json:"documentType"`
	DocumentNum  string `json:"documentNum"`
	Email        string `json:"email"`
}

type UserResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		UserId string `json:"userId"`
	} `json:"data"`
}

func (s *SvcImpl) AddUser(user *UserDto) (*UserResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminuserservice/users", user)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result UserResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) UpdateUser(user *UserDto) (*UserResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminuserservice/users", user)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result UserResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) DeleteUser(userId string) (*UserResponse, error) {
	url := fmt.Sprintf("%s/api/v1/adminuserservice/users/%s", s.BaseUrl, userId)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result UserResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) GetAllUsers() ([]UserDto, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminuserservice/users", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users []UserDto
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
