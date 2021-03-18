// Code generated by go-swagger; DO NOT EDIT.

package records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetRecordsURL generates an URL for the get records operation
type GetRecordsURL struct {
	City      *string
	DrugName  *string
	Specialty *string
	State     *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRecordsURL) WithBasePath(bp string) *GetRecordsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRecordsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetRecordsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/records"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/go/api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var cityQ string
	if o.City != nil {
		cityQ = *o.City
	}
	if cityQ != "" {
		qs.Set("city", cityQ)
	}

	var drugNameQ string
	if o.DrugName != nil {
		drugNameQ = *o.DrugName
	}
	if drugNameQ != "" {
		qs.Set("drugName", drugNameQ)
	}

	var specialtyQ string
	if o.Specialty != nil {
		specialtyQ = *o.Specialty
	}
	if specialtyQ != "" {
		qs.Set("specialty", specialtyQ)
	}

	var stateQ string
	if o.State != nil {
		stateQ = *o.State
	}
	if stateQ != "" {
		qs.Set("state", stateQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetRecordsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetRecordsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetRecordsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetRecordsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetRecordsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetRecordsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
