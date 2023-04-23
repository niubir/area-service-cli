package cli

import "testing"

var (
	httpCli, grpcCli Client
)

func initHTTPTest() {
	var err error
	httpCli, err = NewClient(HTTP, DEFALUT_HTTP_ADDRESS)
	if err != nil {
		panic(err)
	}
}

func initGRPCTest() {
	var err error
	grpcCli, err = NewClient(GRPC, DEFALUT_GRPC_ADDRESS)
	if err != nil {
		panic(err)
	}
}

func TestHTTPFreshAreas(t *testing.T) {
	initHTTPTest()
	if err := httpCli.FreshAreas(); err != nil {
		t.Fatal(err)
	}
}

func TestHTTPGetArea(t *testing.T) {
	initHTTPTest()
	resp, err := httpCli.GetArea("110000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestHTTPGetProvinces(t *testing.T) {
	initHTTPTest()
	resp, err := httpCli.GetProvinces()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestHTTPGetCities(t *testing.T) {
	initHTTPTest()
	resp, err := httpCli.GetCities("110000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestHTTPGetDistricts(t *testing.T) {
	initHTTPTest()
	resp, err := httpCli.GetDistricts("110100")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestHTTPGetStreets(t *testing.T) {
	initHTTPTest()
	resp, err := httpCli.GetStreets("110101")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGRPCFreshAreas(t *testing.T) {
	initGRPCTest()
	if err := grpcCli.FreshAreas(); err != nil {
		t.Fatal(err)
	}
}

func TestGRPCGetArea(t *testing.T) {
	initGRPCTest()
	resp, err := grpcCli.GetArea("110000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGRPCGetProvinces(t *testing.T) {
	initGRPCTest()
	resp, err := grpcCli.GetProvinces()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGRPCGetCities(t *testing.T) {
	initGRPCTest()
	resp, err := grpcCli.GetCities("110000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGRPCGetDistricts(t *testing.T) {
	initGRPCTest()
	resp, err := grpcCli.GetDistricts("110100")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGRPCGetStreets(t *testing.T) {
	initGRPCTest()
	resp, err := grpcCli.GetStreets("110101")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
