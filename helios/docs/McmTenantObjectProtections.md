# McmTenantObjectProtections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectProtections** | Pointer to [**[]McmTenantObjectProtection**](McmTenantObjectProtection.md) |  | [optional] 
**LastIndex** | Pointer to **NullableInt64** | Specifies the last index returned if there are more results to be returned. | [optional] 

## Methods

### NewMcmTenantObjectProtections

`func NewMcmTenantObjectProtections() *McmTenantObjectProtections`

NewMcmTenantObjectProtections instantiates a new McmTenantObjectProtections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmTenantObjectProtectionsWithDefaults

`func NewMcmTenantObjectProtectionsWithDefaults() *McmTenantObjectProtections`

NewMcmTenantObjectProtectionsWithDefaults instantiates a new McmTenantObjectProtections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectProtections

`func (o *McmTenantObjectProtections) GetObjectProtections() []McmTenantObjectProtection`

GetObjectProtections returns the ObjectProtections field if non-nil, zero value otherwise.

### GetObjectProtectionsOk

`func (o *McmTenantObjectProtections) GetObjectProtectionsOk() (*[]McmTenantObjectProtection, bool)`

GetObjectProtectionsOk returns a tuple with the ObjectProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtections

`func (o *McmTenantObjectProtections) SetObjectProtections(v []McmTenantObjectProtection)`

SetObjectProtections sets ObjectProtections field to given value.

### HasObjectProtections

`func (o *McmTenantObjectProtections) HasObjectProtections() bool`

HasObjectProtections returns a boolean if a field has been set.

### GetLastIndex

`func (o *McmTenantObjectProtections) GetLastIndex() int64`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *McmTenantObjectProtections) GetLastIndexOk() (*int64, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *McmTenantObjectProtections) SetLastIndex(v int64)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *McmTenantObjectProtections) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### SetLastIndexNil

`func (o *McmTenantObjectProtections) SetLastIndexNil(b bool)`

 SetLastIndexNil sets the value for LastIndex to be an explicit nil

### UnsetLastIndex
`func (o *McmTenantObjectProtections) UnsetLastIndex()`

UnsetLastIndex ensures that no value is present for LastIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


