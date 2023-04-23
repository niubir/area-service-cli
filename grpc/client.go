package grpc

import (
	"context"
	"encoding/json"

	"github.com/niubir/area-service-cli/grpc/server"
	"github.com/niubir/area-service-cli/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	address string
	conn    *grpc.ClientConn
	client  server.AreaServiceClient
}

func NewClient(address string) (*Client, error) {
	return &Client{address: address}, nil
}

// 刷新区域
func (c *Client) FreshAreas() error {
	if err := c.prefunc(); err != nil {
		return err
	}
	if _, err := c.client.FreshAreas(context.Background(), &server.Empty{}); err != nil {
		return err
	}
	return nil
}

// 获取指定区域
func (c *Client) GetArea(code string) (*models.Area, error) {
	if err := c.prefunc(); err != nil {
		return nil, err
	}
	ack, err := c.client.GetArea(context.Background(), &server.GetAreaReq{Code: code})
	if err != nil {
		return nil, err
	}
	return c.areaExchange(ack.Area)
}

// 获取省份列表
func (c *Client) GetProvinces() (models.Areas, error) {
	if err := c.prefunc(); err != nil {
		return nil, err
	}
	ack, err := c.client.GetProvinces(context.Background(), &server.GetProvincesReq{})
	if err != nil {
		return nil, err
	}
	return c.areasExchange(ack.Provinces)
}

// 获取城市列表
func (c *Client) GetCities(provinceCode string) (models.Areas, error) {
	if err := c.prefunc(); err != nil {
		return nil, err
	}
	ack, err := c.client.GetCities(context.Background(), &server.GetCitiesReq{ProvinceCode: provinceCode})
	if err != nil {
		return nil, err
	}
	return c.areasExchange(ack.Cities)
}

// 获取区县列表
func (c *Client) GetDistricts(cityCode string) (models.Areas, error) {
	if err := c.prefunc(); err != nil {
		return nil, err
	}
	ack, err := c.client.GetDistricts(context.Background(), &server.GetDistrictsReq{CityCode: cityCode})
	if err != nil {
		return nil, err
	}
	return c.areasExchange(ack.Districts)
}

// 获取街道列表
func (c *Client) GetStreets(districtCode string) (models.Areas, error) {
	if err := c.prefunc(); err != nil {
		return nil, err
	}
	ack, err := c.client.GetStreets(context.Background(), &server.GetStreetsReq{DistrictCode: districtCode})
	if err != nil {
		return nil, err
	}
	return c.areasExchange(ack.Streets)
}

func (c *Client) prefunc() error {
	if c.conn == nil {
		conn, err := grpc.Dial(c.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}
		c.conn = conn
		c.client = server.NewAreaServiceClient(conn)
	}
	return nil
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *Client) areasExchange(areas []*server.Area) (models.Areas, error) {
	b, err := json.Marshal(areas)
	if err != nil {
		return nil, err
	}
	var resp models.Areas
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) areaExchange(area *server.Area) (*models.Area, error) {
	b, err := json.Marshal(area)
	if err != nil {
		return nil, err
	}
	var resp *models.Area
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
