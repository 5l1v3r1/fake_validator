package api

import . "gopkg.in/check.v1"

type HealthCheckSuite struct{}

var _ = Suite(&HealthCheckSuite{})

func (s *HealthCheckSuite) TestHealthCheck(c *C) {
	result, err := HealthCheck()
	c.Assert(err, IsNil)
	c.Assert(result, Equals, "OK")
}
