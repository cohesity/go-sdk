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

// CertificatesApiService CertificatesApi service
type CertificatesApiService service

type ApiDeleteWebServerCertificateRequest struct {
	ctx _context.Context
	ApiService *CertificatesApiService
}

/*

*/

/*
func (r ApiDeleteWebServerCertificateRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteWebServerCertificateExecute(r)
}

 * DeleteWebServerCertificate Delete the SSL Certificate in the Cluster.
 * Returns delete status upon completion.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDeleteWebServerCertificateRequest
 
func (a *CertificatesApiService) DeleteWebServerCertificate(ctx _context.Context) ApiDeleteWebServerCertificateRequest {
	return ApiDeleteWebServerCertificateRequest{
		ApiService: a,
		ctx: ctx,
	}
}
*/

func (a *CertificatesApiService) DeleteWebServerCertificate(r ApiDeleteWebServerCertificateRequest) (*_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokensApiService" {
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

	return r.ApiService.DeleteWebServerCertificateExecute(r)
}


/*
 * Execute executes the request
 */
func (a *CertificatesApiService) DeleteWebServerCertificateExecute(r ApiDeleteWebServerCertificateRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.DeleteWebServerCertificate")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/certificates/webServer"

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
			var v RequestError
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

type ApiDeployHostCertificateRequest struct {
	ctx _context.Context
	ApiService *CertificatesApiService
	Body *DeployCertParameters
}

/*

func (r ApiDeployHostCertificateRequest) Body(body DeployCertParameters) ApiDeployHostCertificateRequest {
	r.body = &body
	return r
}
*/

/*
func (r ApiDeployHostCertificateRequest) Execute() (CertificateDetails, *_nethttp.Response, error) {
	return r.ApiService.DeployHostCertificateExecute(r)
}

 * DeployHostCertificate Generate and deploy certificate for a single or multiple hosts.
 * Returns the global certificate for a single or multiple hosts.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDeployHostCertificateRequest
 
func (a *CertificatesApiService) DeployHostCertificate(ctx _context.Context) ApiDeployHostCertificateRequest {
	return ApiDeployHostCertificateRequest{
		ApiService: a,
		ctx: ctx,
	}
}
*/

func (a *CertificatesApiService) DeployHostCertificate(r ApiDeployHostCertificateRequest) (CertificateDetails, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokensApiService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return CertificateDetails{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.DeployHostCertificateExecute(r)
}


/*
 * Execute executes the request
 * @return CertificateDetails
 */
func (a *CertificatesApiService) DeployHostCertificateExecute(r ApiDeployHostCertificateRequest) (CertificateDetails, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CertificateDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.DeployHostCertificate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/certificates/global"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiGetCertificateListRequest struct {
	ctx _context.Context
	ApiService *CertificatesApiService
}

/*

*/

/*
func (r ApiGetCertificateListRequest) Execute() (ListCertResponse, *_nethttp.Response, error) {
	return r.ApiService.GetCertificateListExecute(r)
}

 * GetCertificateList List the certificates generated and deployed on hosts.
 * Returns the all certificate and their details generated from this cluster.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetCertificateListRequest
 
func (a *CertificatesApiService) GetCertificateList(ctx _context.Context) ApiGetCertificateListRequest {
	return ApiGetCertificateListRequest{
		ApiService: a,
		ctx: ctx,
	}
}
*/

func (a *CertificatesApiService) GetCertificateList(r ApiGetCertificateListRequest) (ListCertResponse, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokensApiService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return ListCertResponse{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetCertificateListExecute(r)
}


/*
 * Execute executes the request
 * @return ListCertResponse
 */
func (a *CertificatesApiService) GetCertificateListExecute(r ApiGetCertificateListRequest) (ListCertResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListCertResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.GetCertificateList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/certificates/global"

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

type ApiGetWebServerCertificateRequest struct {
	ctx _context.Context
	ApiService *CertificatesApiService
}

/*

*/

/*
func (r ApiGetWebServerCertificateRequest) Execute() (SslCertificateConfig, *_nethttp.Response, error) {
	return r.ApiService.GetWebServerCertificateExecute(r)
}

 * GetWebServerCertificate Get the Server Certificate configured on the Cluster.
 * Returns the Server Certificate configured on the cluster.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetWebServerCertificateRequest
 
func (a *CertificatesApiService) GetWebServerCertificate(ctx _context.Context) ApiGetWebServerCertificateRequest {
	return ApiGetWebServerCertificateRequest{
		ApiService: a,
		ctx: ctx,
	}
}
*/

func (a *CertificatesApiService) GetWebServerCertificate(r ApiGetWebServerCertificateRequest) (SslCertificateConfig, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokensApiService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return SslCertificateConfig{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.GetWebServerCertificateExecute(r)
}


/*
 * Execute executes the request
 * @return SslCertificateConfig
 */
func (a *CertificatesApiService) GetWebServerCertificateExecute(r ApiGetWebServerCertificateRequest) (SslCertificateConfig, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SslCertificateConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.GetWebServerCertificate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/certificates/webServer"

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

type ApiUpdateWebServerCertificateRequest struct {
	ctx _context.Context
	ApiService *CertificatesApiService
	Body *SslCertificateConfig
}

/*

func (r ApiUpdateWebServerCertificateRequest) Body(body SslCertificateConfig) ApiUpdateWebServerCertificateRequest {
	r.body = &body
	return r
}
*/

/*
func (r ApiUpdateWebServerCertificateRequest) Execute() (SslCertificateConfig, *_nethttp.Response, error) {
	return r.ApiService.UpdateWebServerCertificateExecute(r)
}

 * UpdateWebServerCertificate Update the Web Server Certificate on the Cluster.
 * Returns the updated Web Server Certificate on the cluster.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiUpdateWebServerCertificateRequest
 
func (a *CertificatesApiService) UpdateWebServerCertificate(ctx _context.Context) ApiUpdateWebServerCertificateRequest {
	return ApiUpdateWebServerCertificateRequest{
		ApiService: a,
		ctx: ctx,
	}
}
*/

func (a *CertificatesApiService) UpdateWebServerCertificate(r ApiUpdateWebServerCertificateRequest) (SslCertificateConfig, *_nethttp.Response, error) {
	if reflect.TypeOf(*a).Name() != "AccessTokensApiService" {
		token, err := GetTokenHelper(a.client)

		if err != nil {
			return SslCertificateConfig{}, nil, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		r.ctx = _context.WithValue(_context.Background(), ContextAPIKeys, map[string]APIKey{"TokenHeader": {Key: token, Prefix: "Bearer"}})
	} else {
		r.ctx = _context.Background()
	}
	r.ApiService = a

	return r.ApiService.UpdateWebServerCertificateExecute(r)
}


/*
 * Execute executes the request
 * @return SslCertificateConfig
 */
func (a *CertificatesApiService) UpdateWebServerCertificateExecute(r ApiUpdateWebServerCertificateRequest) (SslCertificateConfig, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SslCertificateConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesApiService.UpdateWebServerCertificate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/public/certificates/webServer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
