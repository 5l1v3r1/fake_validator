package api

// Validator interface
type Validator interface {
	Validate() (bool, error)
}
