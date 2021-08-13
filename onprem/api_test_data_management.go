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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// TestDataManagementService TestDataManagement service
type TestDataManagementService service

type ApiCreateTdmTaskRequest struct {
	ctx _context.Context
	ApiService *TestDataManagementService
	Body *CreateTdmTaskRequest
}


func (a *TestDataManagementService) CreateTdmTask(r ApiCreateTdmTaskRequest) (TdmTask, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return TdmTask{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.CreateTdmTaskExecute(r)
}


/*
 * Execute executes the request
 * @return TdmTask
 */
func (a *TestDataManagementService) CreateTdmTaskExecute(r ApiCreateTdmTaskRequest) (TdmTask, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TdmTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestDataManagementService.CreateTdmTask")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tdm/tasks"

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

type ApiDeleteTdmSnapshotByIdRequest struct {
	ctx _context.Context
	ApiService *TestDataManagementService
	Id string
}


func (a *TestDataManagementService) DeleteTdmSnapshotById(r ApiDeleteTdmSnapshotByIdRequest) (*_nethttp.Response, error) {
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

	return r.ApiService.DeleteTdmSnapshotByIdExecute(r)
}


/*
 * Execute executes the request
 */
func (a *TestDataManagementService) DeleteTdmSnapshotByIdExecute(r ApiDeleteTdmSnapshotByIdRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestDataManagementService.DeleteTdmSnapshotById")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tdm/snapshots/{id}"
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

type ApiGetTdmObjectByIdRequest struct {
	ctx _context.Context
	ApiService *TestDataManagementService
	Id string
}


func (a *TestDataManagementService) GetTdmObjectById(r ApiGetTdmObjectByIdRequest) (TdmObject, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return TdmObject{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetTdmObjectByIdExecute(r)
}


/*
 * Execute executes the request
 * @return TdmObject
 */
func (a *TestDataManagementService) GetTdmObjectByIdExecute(r ApiGetTdmObjectByIdRequest) (TdmObject, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TdmObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestDataManagementService.GetTdmObjectById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tdm/objects/{id}"
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

type ApiGetTdmObjectsRequest struct {
	ctx _context.Context
	ApiService *TestDataManagementService
	Ids *[]string
	Environments *[]string
	Name *string
	TaskIds *[]string
	Statuses *[]string
	PaginationCookie *string
}


func (a *TestDataManagementService) GetTdmObjects(r ApiGetTdmObjectsRequest) (TdmObjects, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return TdmObjects{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetTdmObjectsExecute(r)
}


/*
 * Execute executes the request
 * @return TdmObjects
 */
func (a *TestDataManagementService) GetTdmObjectsExecute(r ApiGetTdmObjectsRequest) (TdmObjects, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TdmObjects
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestDataManagementService.GetTdmObjects")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tdm/objects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.Ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.Ids, "csv"))
	}
	if r.Environments != nil {
		localVarQueryParams.Add("environments", parameterToString(*r.Environments, "csv"))
	}
	if r.Name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.Name, ""))
	}
	if r.TaskIds != nil {
		localVarQueryParams.Add("taskIds", parameterToString(*r.TaskIds, "csv"))
	}
	if r.Statuses != nil {
		localVarQueryParams.Add("statuses", parameterToString(*r.Statuses, "csv"))
	}
	if r.PaginationCookie != nil {
		localVarQueryParams.Add("paginationCookie", parameterToString(*r.PaginationCookie, ""))
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

type ApiGetTdmTaskByIdRequest struct {
	ctx _context.Context
	ApiService *TestDataManagementService
	Id string
}


func (a *TestDataManagementService) GetTdmTaskById(r ApiGetTdmTaskByIdRequest) (TdmTask, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return TdmTask{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetTdmTaskByIdExecute(r)
}


/*
 * Execute executes the request
 * @return TdmTask
 */
func (a *TestDataManagementService) GetTdmTaskByIdExecute(r ApiGetTdmTaskByIdRequest) (TdmTask, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TdmTask
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestDataManagementService.GetTdmTaskById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tdm/tasks/{id}"
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

type ApiGetTdmTasksRequest struct {
	ctx _context.Context
	ApiService *TestDataManagementService
	Ids *[]string
	Actions *[]string
	Environments *[]string
	CreatedAfterUsecs *int64
	CreatedBeforeUsecs *int64
	Statuses *[]string
	ObjectIds *[]string
	PaginationCookie *string
}


func (a *TestDataManagementService) GetTdmTasks(r ApiGetTdmTasksRequest) (TdmTasks, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return TdmTasks{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetTdmTasksExecute(r)
}


/*
 * Execute executes the request
 * @return TdmTasks
 */
func (a *TestDataManagementService) GetTdmTasksExecute(r ApiGetTdmTasksRequest) (TdmTasks, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TdmTasks
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestDataManagementService.GetTdmTasks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tdm/tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.Ids != nil {
		localVarQueryParams.Add("ids", parameterToString(*r.Ids, "csv"))
	}
	if r.Actions != nil {
		localVarQueryParams.Add("actions", parameterToString(*r.Actions, "csv"))
	}
	if r.Environments != nil {
		localVarQueryParams.Add("environments", parameterToString(*r.Environments, "csv"))
	}
	if r.CreatedAfterUsecs != nil {
		localVarQueryParams.Add("createdAfterUsecs", parameterToString(*r.CreatedAfterUsecs, ""))
	}
	if r.CreatedBeforeUsecs != nil {
		localVarQueryParams.Add("createdBeforeUsecs", parameterToString(*r.CreatedBeforeUsecs, ""))
	}
	if r.Statuses != nil {
		localVarQueryParams.Add("statuses", parameterToString(*r.Statuses, "csv"))
	}
	if r.ObjectIds != nil {
		localVarQueryParams.Add("objectIds", parameterToString(*r.ObjectIds, "csv"))
	}
	if r.PaginationCookie != nil {
		localVarQueryParams.Add("paginationCookie", parameterToString(*r.PaginationCookie, ""))
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

type ApiGetTdmTimelineEventsByObjectIdRequest struct {
	ctx _context.Context
	ApiService *TestDataManagementService
	Id string
	CreatedAfter *int64
	CreatedBefore *int64
}


func (a *TestDataManagementService) GetTdmTimelineEventsByObjectId(r ApiGetTdmTimelineEventsByObjectIdRequest) (TdmObjectTimelineEvents, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return TdmObjectTimelineEvents{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetTdmTimelineEventsByObjectIdExecute(r)
}


/*
 * Execute executes the request
 * @return TdmObjectTimelineEvents
 */
func (a *TestDataManagementService) GetTdmTimelineEventsByObjectIdExecute(r ApiGetTdmTimelineEventsByObjectIdRequest) (TdmObjectTimelineEvents, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TdmObjectTimelineEvents
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestDataManagementService.GetTdmTimelineEventsByObjectId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tdm/objects/{id}/timeline-events"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.Id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.CreatedAfter != nil {
		localVarQueryParams.Add("createdAfter", parameterToString(*r.CreatedAfter, ""))
	}
	if r.CreatedBefore != nil {
		localVarQueryParams.Add("createdBefore", parameterToString(*r.CreatedBefore, ""))
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

type ApiUpdateTdmSnapshotByIdRequest struct {
	ctx _context.Context
	ApiService *TestDataManagementService
	Id string
	Body *UpdateTdmSnapshotRequest
}


func (a *TestDataManagementService) UpdateTdmSnapshotById(r ApiUpdateTdmSnapshotByIdRequest) (TdmSnapshot, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokenService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return TdmSnapshot{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.UpdateTdmSnapshotByIdExecute(r)
}


/*
 * Execute executes the request
 * @return TdmSnapshot
 */
func (a *TestDataManagementService) UpdateTdmSnapshotByIdExecute(r ApiUpdateTdmSnapshotByIdRequest) (TdmSnapshot, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TdmSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestDataManagementService.UpdateTdmSnapshotById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tdm/snapshots/{id}"
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
