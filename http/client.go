package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/niubir/area-service-cli/models"
)

type Client struct {
	c       *http.Client
	address string
}

func NewClient(address string) (*Client, error) {
	return &Client{c: &http.Client{}, address: address}, nil
}

// 刷新区域
func (c *Client) FreshAreas() error {
	req, err := c.postRequest("", nil)
	if err != nil {
		return err
	}
	if _, err := c.c.Do(req); err != nil {
		return err
	}
	return nil
}

// 获取指定区域
func (c *Client) GetArea(code string) (*models.Area, error) {
	req, err := c.getRequest("", map[string]string{"code": code})
	if err != nil {
		return nil, err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	var area models.Area
	if err := c.parseResponse(resp, &area); err != nil {
		return nil, err
	}
	return &area, nil
}

// 获取省份列表
func (c *Client) GetProvinces() (models.Areas, error) {
	req, err := c.getRequest("/provinces", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	var areas models.Areas
	if err := c.parseResponse(resp, &areas); err != nil {
		return nil, err
	}
	return areas, nil
}

// 获取城市列表
func (c *Client) GetCities(provinceCode string) (models.Areas, error) {
	req, err := c.getRequest("/cities", map[string]string{"provinceCode": provinceCode})
	if err != nil {
		return nil, err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	var areas models.Areas
	if err := c.parseResponse(resp, &areas); err != nil {
		return nil, err
	}
	return areas, nil
}

// 获取区县列表
func (c *Client) GetDistricts(cityCode string) (models.Areas, error) {
	req, err := c.getRequest("/districts", map[string]string{"cityCode": cityCode})
	if err != nil {
		return nil, err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	var areas models.Areas
	if err := c.parseResponse(resp, &areas); err != nil {
		return nil, err
	}
	return areas, nil
}

// 获取街道列表
func (c *Client) GetStreets(districtCode string) (models.Areas, error) {
	req, err := c.getRequest("/streets", map[string]string{"districtCode": districtCode})
	if err != nil {
		return nil, err
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	var areas models.Areas
	if err := c.parseResponse(resp, &areas); err != nil {
		return nil, err
	}
	return areas, nil
}

func (c *Client) postRequest(path string, data interface{}) (*http.Request, error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s%s", c.address, path),
		bytes.NewReader(bytesData),
	)
}

func (c *Client) getRequest(path string, data map[string]string) (*http.Request, error) {
	v := url.Values{}
	for key, value := range data {
		v.Add(key, value)
	}
	return http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s%s?%s", c.address, path, v.Encode()),
		nil,
	)
}

func (c *Client) parseResponse(resp *http.Response, data interface{}) error {
	if data == nil {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &data)
}
