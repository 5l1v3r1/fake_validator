package api

import (
	. "gopkg.in/check.v1"
)

type NumberValidatorSuite struct{}

var _ = Suite(&NumberValidatorSuite{})

type ValidatorDataProvider struct {
	Number string
	Expected bool
	Error bool
}

func (s *NumberValidatorSuite) TestNumberValidator(c *C) {
	dataProvider := []ValidatorDataProvider{
		ValidatorDataProvider{Number: "23", Expected: true, Error: false},
		ValidatorDataProvider{Number: "0", Expected: true, Error: false},
		ValidatorDataProvider{Number: "100", Expected: true, Error: false},
		ValidatorDataProvider{Number: "123", Expected: false, Error: true},
		ValidatorDataProvider{Number: "-10", Expected: false, Error: true},
		ValidatorDataProvider{Number: "four", Expected: false, Error: true},
		ValidatorDataProvider{Number: "", Expected: false, Error: true},
	}

	for _, data := range dataProvider {
		c.Log(data.Number)
		validator := NewNumberValidator(data.Number)
		result, err := validator.Validate()
		if data.Error {
			c.Assert(err, NotNil)
		} else {
			c.Assert(err, IsNil)
		}
		c.Assert(result, Equals, data.Expected)
	}
}
