/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onprem

import (
	"bytes"
	"reflect"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"regexp"
	. "github.com/cohesity/go-sdk/onprem/models"
)

// Linger please
var (
	_ _context.Context
)

// RegistrationService Registration service
type RegistrationService service

type ApiGetHeliosRegConfigRequest struct {
	ctx _context.Context
	ApiService *RegistrationService
}


func (a *RegistrationService) GetHeliosRegConfig(r ApiGetHeliosRegConfigRequest) (HeliosRegConfig, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return HeliosRegConfig{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetHeliosRegConfigExecute(r)
}


/*
 * Execute executes the request
 * @return HeliosRegConfig
 */
func (a *RegistrationService) GetHeliosRegConfigExecute(r ApiGetHeliosRegConfigRequest) (HeliosRegConfig, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  HeliosRegConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrationService.GetHeliosRegConfig")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helios-registration-config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		if match, matchErr := regexp.MatchString(`Client.Timeout exceeded`, err.Error()); matchErr == nil && match {
			serName := reflect.TypeOf(r).Name()
			serName = serName[3:len(serName)-7] // remove Api prefix, and Request suffix
			err = GenericOpenAPIError{
				error: "Network timeout when making a request of " + serName +
					". Consider increase the request timeout in the client config." ,
			}
		}
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHeliosClaimRequest struct {
	ctx _context.Context
	ApiService *RegistrationService
	Body *HeliosClaimRequest
}


func (a *RegistrationService) HeliosClaim(r ApiHeliosClaimRequest) (*_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.HeliosClaimExecute(r)
}


/*
 * Execute executes the request
 */
func (a *RegistrationService) HeliosClaimExecute(r ApiHeliosClaimRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrationService.HeliosClaim")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helios-registration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.Body == nil {
		return nil, reportError("Body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.Body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["TokenHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		if match, matchErr := regexp.MatchString(`Client.Timeout exceeded`, err.Error()); matchErr == nil && match {
			serName := reflect.TypeOf(r).Name()
			serName = serName[3:len(serName)-7] // remove Api prefix, and Request suffix
			err = GenericOpenAPIError{
				error: "Network timeout when making a request of " + serName +
					". Consider increase the request timeout in the client config." ,
			}
		}
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
