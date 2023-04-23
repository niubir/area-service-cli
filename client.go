package cli

import (
	"errors"

	"github.com/niubir/area-service-cli/grpc"
	"github.com/niubir/area-service-cli/http"
	"github.com/niubir/area-service-cli/models"
)

const (
	GRPC clientType = iota
	HTTP

	DEFAULT_HTTP_PORT    = 10011
	DEFALUT_HTTP_ADDRESS = "127.0.0.1:10011"

	DEFAULT_GRPC_PORT    = 10021
	DEFALUT_GRPC_ADDRESS = "http:127.0.0.1:10012"
)

type clientType int

type Client interface {
	FreshAreas() error
	GetArea(code string) (*models.Area, error)
	GetProvinces() (models.Areas, error)
	GetCities(provinceCode string) (models.Areas, error)
	GetDistricts(cityCode string) (models.Areas, error)
	GetStreets(districtCode string) (models.Areas, error)
}

func NewClient(clientType clientType, address string) (Client, error) {
	switch clientType {
	case GRPC:
		return grpc.NewClient(address)
	case HTTP:
		return http.NewClient(address)
	default:
		return nil, errors.New("invalid client type")
	}
}
