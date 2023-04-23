package grpc

import (
	"testing"
)

var (
	c *Client
)

func initTest() {
	var err error
	c, err = NewClient("127.0.0.1:10012")
	if err != nil {
		panic(err)
	}
}

func TestFreshAreas(t *testing.T) {
	initTest()
	if err := c.FreshAreas(); err != nil {
		t.Fatal(err)
	}
}

func TestGetArea(t *testing.T) {
	initTest()
	resp, err := c.GetArea("110000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGetProvinces(t *testing.T) {
	initTest()
	resp, err := c.GetProvinces()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGetCities(t *testing.T) {
	initTest()
	resp, err := c.GetCities("110000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGetDistricts(t *testing.T) {
	initTest()
	resp, err := c.GetDistricts("110100")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGetStreets(t *testing.T) {
	initTest()
	resp, err := c.GetStreets("110101")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
