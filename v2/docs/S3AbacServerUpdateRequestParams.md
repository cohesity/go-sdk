# S3AbacServerUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePath** | **NullableString** | Specifies the path of URI for user requests. | 
**CaBundle** | **NullableString** | Specifies the intermediate certificates. | 
**Certificate** | **NullableString** | Specifies the client certificate. | 
**Key** | Pointer to **NullableString** | Specifies the RSA private key. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant Id for S3 ABAC server. | [optional] 

## Methods

### NewS3AbacServerUpdateRequestParams

`func NewS3AbacServerUpdateRequestParams(basePath NullableString, caBundle NullableString, certificate NullableString, ) *S3AbacServerUpdateRequestParams`

NewS3AbacServerUpdateRequestParams instantiates a new S3AbacServerUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3AbacServerUpdateRequestParamsWithDefaults

`func NewS3AbacServerUpdateRequestParamsWithDefaults() *S3AbacServerUpdateRequestParams`

NewS3AbacServerUpdateRequestParamsWithDefaults instantiates a new S3AbacServerUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePath

`func (o *S3AbacServerUpdateRequestParams) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *S3AbacServerUpdateRequestParams) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *S3AbacServerUpdateRequestParams) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.


### SetBasePathNil

`func (o *S3AbacServerUpdateRequestParams) SetBasePathNil(b bool)`

 SetBasePathNil sets the value for BasePath to be an explicit nil

### UnsetBasePath
`func (o *S3AbacServerUpdateRequestParams) UnsetBasePath()`

UnsetBasePath ensures that no value is present for BasePath, not even an explicit nil
### GetCaBundle

`func (o *S3AbacServerUpdateRequestParams) GetCaBundle() string`

GetCaBundle returns the CaBundle field if non-nil, zero value otherwise.

### GetCaBundleOk

`func (o *S3AbacServerUpdateRequestParams) GetCaBundleOk() (*string, bool)`

GetCaBundleOk returns a tuple with the CaBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaBundle

`func (o *S3AbacServerUpdateRequestParams) SetCaBundle(v string)`

SetCaBundle sets CaBundle field to given value.


### SetCaBundleNil

`func (o *S3AbacServerUpdateRequestParams) SetCaBundleNil(b bool)`

 SetCaBundleNil sets the value for CaBundle to be an explicit nil

### UnsetCaBundle
`func (o *S3AbacServerUpdateRequestParams) UnsetCaBundle()`

UnsetCaBundle ensures that no value is present for CaBundle, not even an explicit nil
### GetCertificate

`func (o *S3AbacServerUpdateRequestParams) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *S3AbacServerUpdateRequestParams) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *S3AbacServerUpdateRequestParams) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *S3AbacServerUpdateRequestParams) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *S3AbacServerUpdateRequestParams) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetKey

`func (o *S3AbacServerUpdateRequestParams) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *S3AbacServerUpdateRequestParams) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *S3AbacServerUpdateRequestParams) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *S3AbacServerUpdateRequestParams) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *S3AbacServerUpdateRequestParams) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *S3AbacServerUpdateRequestParams) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetTenantId

`func (o *S3AbacServerUpdateRequestParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *S3AbacServerUpdateRequestParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *S3AbacServerUpdateRequestParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *S3AbacServerUpdateRequestParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *S3AbacServerUpdateRequestParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *S3AbacServerUpdateRequestParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


