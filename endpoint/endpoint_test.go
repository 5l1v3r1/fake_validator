package endpoint

import (
	"testing"
	. "gopkg.in/check.v1"
	"net/http/httptest"
	"github.com/fake_validator/endpoint"
	"github.com/fake_validator/router"
	"net/http"
	"strings"
	"io/ioutil"
)

// DataProvider type
type DataProvider struct {
	RequestedURL       string
	RequestMethod      string
	RequestBody        string
	ExpectedBody       string
	ExpectedHeader     map[string]string
	ExpectedReturnCode int
}

type EndpointSuite struct{}

var _ = Suite(&EndpointSuite{})

// Run all test suites
func TestAllSuite(t *testing.T) {
	TestingT(t)
}

func (s *EndpointSuite) TestEndpointSuite(c *C) {
	router := router.NewRouter(endpoint.EndpointRoutes)
	server := httptest.NewServer(router)
	defer server.Close()

	dataProvider := []DataProvider{
		DataProvider{
			RequestedURL: "/ping",
			RequestMethod: "GET",
			RequestBody: "",
			ExpectedBody: "PONG",
			ExpectedReturnCode: http.StatusOK,
		},
		DataProvider{
			RequestedURL: "/health",
			RequestMethod: "GET",
			RequestBody: "",
			ExpectedBody: "OK",
			ExpectedReturnCode: http.StatusOK,
		},
		DataProvider{
			RequestedURL: "/validate/number/57",
			RequestMethod: "GET",
			RequestBody: "",
			ExpectedBody: "true",
			ExpectedReturnCode: http.StatusOK,
		},
		DataProvider{
			RequestedURL: "/validate/number/189",
			RequestMethod: "GET",
			RequestBody: "",
			ExpectedBody: "Number out of range",
			ExpectedReturnCode: http.StatusOK,
		},
		DataProvider{
			RequestedURL: "/validate/number/five",
			RequestMethod: "GET",
			RequestBody: "",
			ExpectedBody: "Wrong input",
			ExpectedReturnCode: http.StatusOK,
		},
	}

	for _, data := range dataProvider {
		req, err := http.NewRequest(
			data.RequestMethod,
			server.URL+data.RequestedURL,
			strings.NewReader(data.RequestBody),
		)
		c.Assert(err, IsNil)

		resp, err := http.DefaultClient.Do(req)
		c.Assert(err, IsNil)
		c.Assert(resp.StatusCode, Equals, data.ExpectedReturnCode)

		actual, err := ioutil.ReadAll(resp.Body)
		c.Assert(err, IsNil)

		obtained := strings.Trim(string(actual), "\n")
		c.Assert(obtained, Equals, data.ExpectedBody)
	}
}
