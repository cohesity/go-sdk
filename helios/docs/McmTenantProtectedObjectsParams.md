# McmTenantProtectedObjectsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the ID of the tenant. | 
**StartIndex** | Pointer to **NullableInt64** | Specifies the index of the first object to be returned. | [optional] 
**NumResults** | Pointer to **NullableInt64** | Specifies the max number of objects to be returned. Default value is 100. | [optional] 

## Methods

### NewMcmTenantProtectedObjectsParams

`func NewMcmTenantProtectedObjectsParams(tenantId NullableString, ) *McmTenantProtectedObjectsParams`

NewMcmTenantProtectedObjectsParams instantiates a new McmTenantProtectedObjectsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmTenantProtectedObjectsParamsWithDefaults

`func NewMcmTenantProtectedObjectsParamsWithDefaults() *McmTenantProtectedObjectsParams`

NewMcmTenantProtectedObjectsParamsWithDefaults instantiates a new McmTenantProtectedObjectsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *McmTenantProtectedObjectsParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *McmTenantProtectedObjectsParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *McmTenantProtectedObjectsParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *McmTenantProtectedObjectsParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *McmTenantProtectedObjectsParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetStartIndex

`func (o *McmTenantProtectedObjectsParams) GetStartIndex() int64`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *McmTenantProtectedObjectsParams) GetStartIndexOk() (*int64, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *McmTenantProtectedObjectsParams) SetStartIndex(v int64)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *McmTenantProtectedObjectsParams) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### SetStartIndexNil

`func (o *McmTenantProtectedObjectsParams) SetStartIndexNil(b bool)`

 SetStartIndexNil sets the value for StartIndex to be an explicit nil

### UnsetStartIndex
`func (o *McmTenantProtectedObjectsParams) UnsetStartIndex()`

UnsetStartIndex ensures that no value is present for StartIndex, not even an explicit nil
### GetNumResults

`func (o *McmTenantProtectedObjectsParams) GetNumResults() int64`

GetNumResults returns the NumResults field if non-nil, zero value otherwise.

### GetNumResultsOk

`func (o *McmTenantProtectedObjectsParams) GetNumResultsOk() (*int64, bool)`

GetNumResultsOk returns a tuple with the NumResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumResults

`func (o *McmTenantProtectedObjectsParams) SetNumResults(v int64)`

SetNumResults sets NumResults field to given value.

### HasNumResults

`func (o *McmTenantProtectedObjectsParams) HasNumResults() bool`

HasNumResults returns a boolean if a field has been set.

### SetNumResultsNil

`func (o *McmTenantProtectedObjectsParams) SetNumResultsNil(b bool)`

 SetNumResultsNil sets the value for NumResults to be an explicit nil

### UnsetNumResults
`func (o *McmTenantProtectedObjectsParams) UnsetNumResults()`

UnsetNumResults ensures that no value is present for NumResults, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


