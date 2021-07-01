/*
 * Cohesity REST API
 *
 * This API list provides operations for interfacing with the Cohesity Cluster.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesitysdk

import (
	"bytes"
	"reflect"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"regexp"
)

// Linger please
var (
	_ _context.Context
)

// CloneRefreshTasksApiService CloneRefreshTasksApi service
type CloneRefreshTasksApiService service

type ApiCreateCloneRefreshTaskRequest struct {
	ctx _context.Context
	ApiService *CloneRefreshTasksApiService
	Body *CloneRefreshRequest
}

/*

func (r ApiCreateCloneRefreshTaskRequest) Body(body CloneRefreshRequest) ApiCreateCloneRefreshTaskRequest {
	r.body = &body
	return r
}
*/

/*
func (r ApiCreateCloneRefreshTaskRequest) Execute() (RestoreTaskWrapper, *_nethttp.Response, error) {
	return r.ApiService.CreateCloneRefreshTaskExecute(r)
}

 * CreateCloneRefreshTask Create a Clone Refresh Task to refresh the clone with the new data.
 * Returns the created Clone Refresh Task which refreshes the clone with specified
data.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateCloneRefreshTaskRequest
 
func (a *CloneRefreshTasksApiService) CreateCloneRefreshTask(ctx _context.Context) ApiCreateCloneRefreshTaskRequest {
	return ApiCreateCloneRefreshTaskRequest{
		ApiService: a,
		ctx: ctx,
	}
}
*/

func (a *CloneRefreshTasksApiService) CreateCloneRefreshTask(r ApiCreateCloneRefreshTaskRequest) (RestoreTaskWrapper, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokensApiService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return RestoreTaskWrapper{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.CreateCloneRefreshTaskExecute(r)
}


/*
 * Execute executes the request
 * @return RestoreTaskWrapper
 */
func (a *CloneRefreshTasksApiService) CreateCloneRefreshTaskExecute(r ApiCreateCloneRefreshTaskRequest) (RestoreTaskWrapper, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RestoreTaskWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloneRefreshTasksApiService.CreateCloneRefreshTask")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/restore/applicationsClone/refresh"

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
			var v RequestError
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
