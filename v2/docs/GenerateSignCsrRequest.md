# GenerateSignCsrRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsrPem** | **string** | Certificate signing request (csr) in pem format | 
**Expiry** | Pointer to **string** | Duration(e.g. 100h) of the certificate | [optional] 
**SanList** | Pointer to **[]string** | Specifies an alternative subject name component to be included in the certificate. It is used to identify the ways the Cluster will be accessed. It is given as a comma separated list of FQDNs. The default value is the Cluster&#39;s VIP hostname. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id | [optional] 

## Methods

### NewGenerateSignCsrRequest

`func NewGenerateSignCsrRequest(csrPem string, ) *GenerateSignCsrRequest`

NewGenerateSignCsrRequest instantiates a new GenerateSignCsrRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSignCsrRequestWithDefaults

`func NewGenerateSignCsrRequestWithDefaults() *GenerateSignCsrRequest`

NewGenerateSignCsrRequestWithDefaults instantiates a new GenerateSignCsrRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsrPem

`func (o *GenerateSignCsrRequest) GetCsrPem() string`

GetCsrPem returns the CsrPem field if non-nil, zero value otherwise.

### GetCsrPemOk

`func (o *GenerateSignCsrRequest) GetCsrPemOk() (*string, bool)`

GetCsrPemOk returns a tuple with the CsrPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrPem

`func (o *GenerateSignCsrRequest) SetCsrPem(v string)`

SetCsrPem sets CsrPem field to given value.


### GetExpiry

`func (o *GenerateSignCsrRequest) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GenerateSignCsrRequest) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GenerateSignCsrRequest) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GenerateSignCsrRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSanList

`func (o *GenerateSignCsrRequest) GetSanList() []string`

GetSanList returns the SanList field if non-nil, zero value otherwise.

### GetSanListOk

`func (o *GenerateSignCsrRequest) GetSanListOk() (*[]string, bool)`

GetSanListOk returns a tuple with the SanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanList

`func (o *GenerateSignCsrRequest) SetSanList(v []string)`

SetSanList sets SanList field to given value.

### HasSanList

`func (o *GenerateSignCsrRequest) HasSanList() bool`

HasSanList returns a boolean if a field has been set.

### SetSanListNil

`func (o *GenerateSignCsrRequest) SetSanListNil(b bool)`

 SetSanListNil sets the value for SanList to be an explicit nil

### UnsetSanList
`func (o *GenerateSignCsrRequest) UnsetSanList()`

UnsetSanList ensures that no value is present for SanList, not even an explicit nil
### GetTenantId

`func (o *GenerateSignCsrRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GenerateSignCsrRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GenerateSignCsrRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *GenerateSignCsrRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *GenerateSignCsrRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *GenerateSignCsrRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


