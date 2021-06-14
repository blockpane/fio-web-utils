// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetFeeVotesProducerURL generates an URL for the get fee votes producer operation
type GetFeeVotesProducerURL struct {
	Producer string

	ApplyMultiplier *bool
	Format          *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetFeeVotesProducerURL) WithBasePath(bp string) *GetFeeVotesProducerURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetFeeVotesProducerURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetFeeVotesProducerURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/fee/votes/{producer}"

	producer := o.Producer
	if producer != "" {
		_path = strings.Replace(_path, "{producer}", producer, -1)
	} else {
		return nil, errors.New("producer is required on GetFeeVotesProducerURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var applyMultiplierQ string
	if o.ApplyMultiplier != nil {
		applyMultiplierQ = swag.FormatBool(*o.ApplyMultiplier)
	}
	if applyMultiplierQ != "" {
		qs.Set("apply_multiplier", applyMultiplierQ)
	}

	var formatQ string
	if o.Format != nil {
		formatQ = *o.Format
	}
	if formatQ != "" {
		qs.Set("format", formatQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetFeeVotesProducerURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetFeeVotesProducerURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetFeeVotesProducerURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetFeeVotesProducerURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetFeeVotesProducerURL")
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
func (o *GetFeeVotesProducerURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
