# S3AbacServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | **NullableString** | Specifies the path of URI for user requests. | 
**CaBundle** | **NullableString** | Specifies the intermediate certificates. | 
**Certificate** | **NullableString** | Specifies the client certificate. | 
**Hostname** | **NullableString** | Specifies the hostname of S3 ABAC server. | 
**Id** | **NullableInt64** | Specifies the ID of S3 ABAC server. | 
**Port** | **NullableInt64** | Specifies the port of S3 ABAC server. | 
**TenantId** | Pointer to **NullableString** | Specifies the tenant Id for S3 ABAC server. | [optional] 

## Methods

### NewS3AbacServerResponse

`func NewS3AbacServerResponse(basePath NullableString, caBundle NullableString, certificate NullableString, hostname NullableString, id NullableInt64, port NullableInt64, ) *S3AbacServerResponse`

NewS3AbacServerResponse instantiates a new S3AbacServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3AbacServerResponseWithDefaults

`func NewS3AbacServerResponseWithDefaults() *S3AbacServerResponse`

NewS3AbacServerResponseWithDefaults instantiates a new S3AbacServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *S3AbacServerResponse) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *S3AbacServerResponse) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *S3AbacServerResponse) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### SetBasePathNil

`func (o *S3AbacServerResponse) SetBasePathNil(b bool)`

 SetBasePathNil sets the value for BasePath to be an explicit nil

### UnsetBasePath
`func (o *S3AbacServerResponse) UnsetBasePath()`

UnsetBasePath ensures that no value is present for BasePath, not even an explicit nil
### GetCaBundle

`func (o *S3AbacServerResponse) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *S3AbacServerResponse) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *S3AbacServerResponse) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.


### SetCaBundleNil

`func (o *S3AbacServerResponse) SetCaBundleNil(b bool)`

 SetCaBundleNil sets the value for CaBundle to be an explicit nil

### UnsetCaBundle
`func (o *S3AbacServerResponse) UnsetCaBundle()`

UnsetCaBundle ensures that no value is present for CaBundle, not even an explicit nil
### GetCertificate

`func (o *S3AbacServerResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *S3AbacServerResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *S3AbacServerResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *S3AbacServerResponse) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *S3AbacServerResponse) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetHostname

`func (o *S3AbacServerResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *S3AbacServerResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *S3AbacServerResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### SetHostnameNil

`func (o *S3AbacServerResponse) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *S3AbacServerResponse) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetId

`func (o *S3AbacServerResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3AbacServerResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3AbacServerResponse) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *S3AbacServerResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *S3AbacServerResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPort

`func (o *S3AbacServerResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *S3AbacServerResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *S3AbacServerResponse) SetPort(v int64)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *S3AbacServerResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *S3AbacServerResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetTenantId

`func (o *S3AbacServerResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *S3AbacServerResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *S3AbacServerResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *S3AbacServerResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *S3AbacServerResponse) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *S3AbacServerResponse) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


