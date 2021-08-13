/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package helios

import (
	"bytes"
	"reflect"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"regexp"
	. "github.com/cohesity/go-sdk/helios/models"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// KerberosProviderService KerberosProvider service
type KerberosProviderService service

type ApiGetKerberosProviderByIdRequest struct {
	ctx _context.Context
	ApiService *KerberosProviderService
	Id string
	AccessClusterId *int64
	RegionId *string
}


func (a *KerberosProviderService) GetKerberosProviderById(r ApiGetKerberosProviderByIdRequest) (KerberosProvider, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return KerberosProvider{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetKerberosProviderByIdExecute(r)
}


/*
 * Execute executes the request
 * @return KerberosProvider
 */
func (a *KerberosProviderService) GetKerberosProviderByIdExecute(r ApiGetKerberosProviderByIdRequest) (KerberosProvider, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  KerberosProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KerberosProviderService.GetKerberosProviderById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kerberos-providers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.Id, "")), -1)

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
	if r.AccessClusterId != nil {
		localVarHeaderParams["accessClusterId"] = parameterToString(*r.AccessClusterId, "")
	}
	if r.RegionId != nil {
		localVarHeaderParams["regionId"] = parameterToString(*r.RegionId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["apiKey"] = key
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

type ApiGetKerberosProvidersRequest struct {
	ctx _context.Context
	ApiService *KerberosProviderService
	AccessClusterId *int64
	RegionId *string
	RealmNames *[]string
	Ids *[]int64
	KdcServers *[]string
}


func (a *KerberosProviderService) GetKerberosProviders(r ApiGetKerberosProvidersRequest) (KerberosProviders, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return KerberosProviders{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetKerberosProvidersExecute(r)
}


/*
 * Execute executes the request
 * @return KerberosProviders
 */
func (a *KerberosProviderService) GetKerberosProvidersExecute(r ApiGetKerberosProvidersRequest) (KerberosProviders, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  KerberosProviders
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KerberosProviderService.GetKerberosProviders")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kerberos-providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.RealmNames != nil {
		localVarQueryParams.Add("realmNames", parameterToString(*r.RealmNames, "csv"))
	}
	if r.Ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.Ids, "csv"))
	}
	if r.KdcServers != nil {
		localVarQueryParams.Add("kdcServers", parameterToString(*r.KdcServers, "csv"))
	}
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
	if r.AccessClusterId != nil {
		localVarHeaderParams["accessClusterId"] = parameterToString(*r.AccessClusterId, "")
	}
	if r.RegionId != nil {
		localVarHeaderParams["regionId"] = parameterToString(*r.RegionId, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["apiKey"] = key
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

type ApiRegisterKerberosProviderRequest struct {
	ctx _context.Context
	ApiService *KerberosProviderService
	Body *RegisterOrUpdateKerberosProviderRequest
	AccessClusterId *int64
	RegionId *string
}


func (a *KerberosProviderService) RegisterKerberosProvider(r ApiRegisterKerberosProviderRequest) (KerberosProvider, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return KerberosProvider{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.RegisterKerberosProviderExecute(r)
}


/*
 * Execute executes the request
 * @return KerberosProvider
 */
func (a *KerberosProviderService) RegisterKerberosProviderExecute(r ApiRegisterKerberosProviderRequest) (KerberosProvider, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  KerberosProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KerberosProviderService.RegisterKerberosProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kerberos-providers/register"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.Body == nil {
		return localVarReturnValue, nil, reportError("Body is required and must be specified")
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
	if r.AccessClusterId != nil {
		localVarHeaderParams["accessClusterId"] = parameterToString(*r.AccessClusterId, "")
	}
	if r.RegionId != nil {
		localVarHeaderParams["regionId"] = parameterToString(*r.RegionId, "")
	}
	// body params
	localVarPostBody = r.Body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["apiKey"] = key
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

type ApiUnregisterKerberosProviderRequest struct {
	ctx _context.Context
	ApiService *KerberosProviderService
	Id string
	Body *UnregisterKerberosRequest
	AccessClusterId *int64
	RegionId *string
}


func (a *KerberosProviderService) UnregisterKerberosProvider(r ApiUnregisterKerberosProviderRequest) (UnregisterKerberosProvider, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return UnregisterKerberosProvider{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.UnregisterKerberosProviderExecute(r)
}


/*
 * Execute executes the request
 * @return UnregisterKerberosProvider
 */
func (a *KerberosProviderService) UnregisterKerberosProviderExecute(r ApiUnregisterKerberosProviderRequest) (UnregisterKerberosProvider, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UnregisterKerberosProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KerberosProviderService.UnregisterKerberosProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kerberos-providers/{id}/unregister"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.Id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.Body == nil {
		return localVarReturnValue, nil, reportError("Body is required and must be specified")
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
	if r.AccessClusterId != nil {
		localVarHeaderParams["accessClusterId"] = parameterToString(*r.AccessClusterId, "")
	}
	if r.RegionId != nil {
		localVarHeaderParams["regionId"] = parameterToString(*r.RegionId, "")
	}
	// body params
	localVarPostBody = r.Body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["apiKey"] = key
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

type ApiUpdateKerberosProviderRequest struct {
	ctx _context.Context
	ApiService *KerberosProviderService
	Id string
	Body *RegisterOrUpdateKerberosProviderRequest
	AccessClusterId *int64
	RegionId *string
}


func (a *KerberosProviderService) UpdateKerberosProvider(r ApiUpdateKerberosProviderRequest) (KerberosProvider, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return KerberosProvider{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.UpdateKerberosProviderExecute(r)
}


/*
 * Execute executes the request
 * @return KerberosProvider
 */
func (a *KerberosProviderService) UpdateKerberosProviderExecute(r ApiUpdateKerberosProviderRequest) (KerberosProvider, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  KerberosProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KerberosProviderService.UpdateKerberosProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kerberos-providers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.Id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.Body == nil {
		return localVarReturnValue, nil, reportError("Body is required and must be specified")
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
	if r.AccessClusterId != nil {
		localVarHeaderParams["accessClusterId"] = parameterToString(*r.AccessClusterId, "")
	}
	if r.RegionId != nil {
		localVarHeaderParams["regionId"] = parameterToString(*r.RegionId, "")
	}
	// body params
	localVarPostBody = r.Body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["APIKeyHeader"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["apiKey"] = key
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
