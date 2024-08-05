# ImportCertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaChainPem** | **[]string** | Full ca certificate chain in pem format. | 
**CertPem** | **string** | Certificate (pem) to be imported | 
**PrivateKey** | **string** | Private key | 
**ServiceType** | Pointer to **NullableString** | Specifies the service that this certificate/key material is used. | [optional] [default to "kAll"]
**TenantId** | Pointer to **NullableString** | Specifies the tenant id | [optional] 

## Methods

### NewImportCertRequest

`func NewImportCertRequest(caChainPem []string, certPem string, privateKey string, ) *ImportCertRequest`

NewImportCertRequest instantiates a new ImportCertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportCertRequestWithDefaults

`func NewImportCertRequestWithDefaults() *ImportCertRequest`

NewImportCertRequestWithDefaults instantiates a new ImportCertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaChainPem

`func (o *ImportCertRequest) GetCaChainPem() []string`

GetCaChainPem returns the CaChainPem field if non-nil, zero value otherwise.

### GetCaChainPemOk

`func (o *ImportCertRequest) GetCaChainPemOk() (*[]string, bool)`

GetCaChainPemOk returns a tuple with the CaChainPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaChainPem

`func (o *ImportCertRequest) SetCaChainPem(v []string)`

SetCaChainPem sets CaChainPem field to given value.


### GetCertPem

`func (o *ImportCertRequest) GetCertPem() string`

GetCertPem returns the CertPem field if non-nil, zero value otherwise.

### GetCertPemOk

`func (o *ImportCertRequest) GetCertPemOk() (*string, bool)`

GetCertPemOk returns a tuple with the CertPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPem

`func (o *ImportCertRequest) SetCertPem(v string)`

SetCertPem sets CertPem field to given value.


### GetPrivateKey

`func (o *ImportCertRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ImportCertRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ImportCertRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetServiceType

`func (o *ImportCertRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ImportCertRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ImportCertRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ImportCertRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *ImportCertRequest) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *ImportCertRequest) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetTenantId

`func (o *ImportCertRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ImportCertRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ImportCertRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ImportCertRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ImportCertRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ImportCertRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


