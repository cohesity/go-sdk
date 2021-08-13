# McmTenantObjectIdsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the ID of the tenant. | 
**ObjectHashes** | **[]string** | Specifies a list of object hashes of objects to fetch. | 

## Methods

### NewMcmTenantObjectIdsParams

`func NewMcmTenantObjectIdsParams(tenantId NullableString, objectHashes []string, ) *McmTenantObjectIdsParams`

NewMcmTenantObjectIdsParams instantiates a new McmTenantObjectIdsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmTenantObjectIdsParamsWithDefaults

`func NewMcmTenantObjectIdsParamsWithDefaults() *McmTenantObjectIdsParams`

NewMcmTenantObjectIdsParamsWithDefaults instantiates a new McmTenantObjectIdsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *McmTenantObjectIdsParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *McmTenantObjectIdsParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *McmTenantObjectIdsParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *McmTenantObjectIdsParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *McmTenantObjectIdsParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetObjectHashes

`func (o *McmTenantObjectIdsParams) GetObjectHashes() []string`

GetObjectHashes returns the ObjectHashes field if non-nil, zero value otherwise.

### GetObjectHashesOk

`func (o *McmTenantObjectIdsParams) GetObjectHashesOk() (*[]string, bool)`

GetObjectHashesOk returns a tuple with the ObjectHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHashes

`func (o *McmTenantObjectIdsParams) SetObjectHashes(v []string)`

SetObjectHashes sets ObjectHashes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


