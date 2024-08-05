# S3AbacServerCreateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | **NullableString** | Specifies the path of URI for user requests. | 
**CaBundle** | **NullableString** | Specifies the intermediate certificates. | 
**Certificate** | **NullableString** | Specifies the client certificate. | 
**Key** | Pointer to **NullableString** | Specifies the RSA private key. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant Id for S3 ABAC server. | [optional] 
**Hostname** | **NullableString** | Specifies the hostname of S3 ABAC server. | 
**Port** | **NullableInt64** | Specifies the port of S3 ABAC server. | 

## Methods

### NewS3AbacServerCreateRequestParams

`func NewS3AbacServerCreateRequestParams(basePath NullableString, caBundle NullableString, certificate NullableString, hostname NullableString, port NullableInt64, ) *S3AbacServerCreateRequestParams`

NewS3AbacServerCreateRequestParams instantiates a new S3AbacServerCreateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3AbacServerCreateRequestParamsWithDefaults

`func NewS3AbacServerCreateRequestParamsWithDefaults() *S3AbacServerCreateRequestParams`

NewS3AbacServerCreateRequestParamsWithDefaults instantiates a new S3AbacServerCreateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *S3AbacServerCreateRequestParams) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *S3AbacServerCreateRequestParams) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *S3AbacServerCreateRequestParams) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### SetBasePathNil

`func (o *S3AbacServerCreateRequestParams) SetBasePathNil(b bool)`

 SetBasePathNil sets the value for BasePath to be an explicit nil

### UnsetBasePath
`func (o *S3AbacServerCreateRequestParams) UnsetBasePath()`

UnsetBasePath ensures that no value is present for BasePath, not even an explicit nil
### GetCaBundle

`func (o *S3AbacServerCreateRequestParams) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *S3AbacServerCreateRequestParams) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *S3AbacServerCreateRequestParams) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.


### SetCaBundleNil

`func (o *S3AbacServerCreateRequestParams) SetCaBundleNil(b bool)`

 SetCaBundleNil sets the value for CaBundle to be an explicit nil

### UnsetCaBundle
`func (o *S3AbacServerCreateRequestParams) UnsetCaBundle()`

UnsetCaBundle ensures that no value is present for CaBundle, not even an explicit nil
### GetCertificate

`func (o *S3AbacServerCreateRequestParams) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *S3AbacServerCreateRequestParams) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *S3AbacServerCreateRequestParams) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *S3AbacServerCreateRequestParams) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *S3AbacServerCreateRequestParams) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetKey

`func (o *S3AbacServerCreateRequestParams) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3AbacServerCreateRequestParams) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3AbacServerCreateRequestParams) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3AbacServerCreateRequestParams) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *S3AbacServerCreateRequestParams) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *S3AbacServerCreateRequestParams) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetTenantId

`func (o *S3AbacServerCreateRequestParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *S3AbacServerCreateRequestParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *S3AbacServerCreateRequestParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *S3AbacServerCreateRequestParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *S3AbacServerCreateRequestParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *S3AbacServerCreateRequestParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetHostname

`func (o *S3AbacServerCreateRequestParams) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *S3AbacServerCreateRequestParams) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *S3AbacServerCreateRequestParams) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### SetHostnameNil

`func (o *S3AbacServerCreateRequestParams) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *S3AbacServerCreateRequestParams) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetPort

`func (o *S3AbacServerCreateRequestParams) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *S3AbacServerCreateRequestParams) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *S3AbacServerCreateRequestParams) SetPort(v int64)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *S3AbacServerCreateRequestParams) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *S3AbacServerCreateRequestParams) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


