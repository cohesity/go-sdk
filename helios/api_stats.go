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
)

// Linger please
var (
	_ _context.Context
)

// StatsService Stats service
type StatsService service

type ApiGetProtectionRunsStatsRequest struct {
	ctx _context.Context
	ApiService *StatsService
	AccessClusterId *int64
	RegionId *string
	StartTimeUsecs *int64
	EndTimeUsecs *int64
	RunStatus *[]string
}


func (a *StatsService) GetProtectionRunsStats(r ApiGetProtectionRunsStatsRequest) (GetProtectionRunsStatusResponseBody, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return GetProtectionRunsStatusResponseBody{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetProtectionRunsStatsExecute(r)
}


/*
 * Execute executes the request
 * @return GetProtectionRunsStatusResponseBody
 */
func (a *StatsService) GetProtectionRunsStatsExecute(r ApiGetProtectionRunsStatsRequest) (GetProtectionRunsStatusResponseBody, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetProtectionRunsStatusResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsService.GetProtectionRunsStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stats/protection-runs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.StartTimeUsecs != nil {
		localVarQueryParams.Add("startTimeUsecs", parameterToString(*r.StartTimeUsecs, ""))
	}
	if r.EndTimeUsecs != nil {
		localVarQueryParams.Add("endTimeUsecs", parameterToString(*r.EndTimeUsecs, ""))
	}
	if r.RunStatus != nil {
		localVarQueryParams.Add("runStatus", parameterToString(*r.RunStatus, "csv"))
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

type ApiMcmGetPolicyLastRunStatsRequest struct {
	ctx _context.Context
	ApiService *StatsService
	RegionId *string
	RegionIds *[]string
}


func (a *StatsService) McmGetPolicyLastRunStats(r ApiMcmGetPolicyLastRunStatsRequest) (McmGetPolicyLastRunStatsResponseBody, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return McmGetPolicyLastRunStatsResponseBody{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.McmGetPolicyLastRunStatsExecute(r)
}


/*
 * Execute executes the request
 * @return McmGetPolicyLastRunStatsResponseBody
 */
func (a *StatsService) McmGetPolicyLastRunStatsExecute(r ApiMcmGetPolicyLastRunStatsRequest) (McmGetPolicyLastRunStatsResponseBody, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  McmGetPolicyLastRunStatsResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsService.McmGetPolicyLastRunStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mcm/stats/policies/last-run"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.RegionIds != nil {
		localVarQueryParams.Add("regionIds", parameterToString(*r.RegionIds, "csv"))
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

type ApiMcmGetProtectionRunLastRunStatsRequest struct {
	ctx _context.Context
	ApiService *StatsService
	RegionId *string
	RegionIds *[]string
}


func (a *StatsService) McmGetProtectionRunLastRunStats(r ApiMcmGetProtectionRunLastRunStatsRequest) (McmGetProtectionLastRunStatsResponseBody, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return McmGetProtectionLastRunStatsResponseBody{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.McmGetProtectionRunLastRunStatsExecute(r)
}


/*
 * Execute executes the request
 * @return McmGetProtectionLastRunStatsResponseBody
 */
func (a *StatsService) McmGetProtectionRunLastRunStatsExecute(r ApiMcmGetProtectionRunLastRunStatsRequest) (McmGetProtectionLastRunStatsResponseBody, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  McmGetProtectionLastRunStatsResponseBody
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatsService.McmGetProtectionRunLastRunStats")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mcm/stats/protection-runs/last-run"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.RegionIds != nil {
		localVarQueryParams.Add("regionIds", parameterToString(*r.RegionIds, "csv"))
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
