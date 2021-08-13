# SearchObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeleted** | Pointer to **NullableBool** | Specifies whether the object is deleted. Deleted objects can&#39;t be protected but can be recovered. | [optional] 
**SourceInfo** | Pointer to [**Object**](Object.md) | Specifies the Source Object information. | [optional] 
**ObjectProtectionInfos** | Pointer to [**[]ObjectProtectionInfo**](ObjectProtectionInfo.md) | Specifies the object info on each cluster. | [optional] 

## Methods

### NewSearchObjectAllOf

`func NewSearchObjectAllOf() *SearchObjectAllOf`

NewSearchObjectAllOf instantiates a new SearchObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchObjectAllOfWithDefaults

`func NewSearchObjectAllOfWithDefaults() *SearchObjectAllOf`

NewSearchObjectAllOfWithDefaults instantiates a new SearchObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeleted

`func (o *SearchObjectAllOf) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *SearchObjectAllOf) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *SearchObjectAllOf) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *SearchObjectAllOf) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### SetIsDeletedNil

`func (o *SearchObjectAllOf) SetIsDeletedNil(b bool)`

 SetIsDeletedNil sets the value for IsDeleted to be an explicit nil

### UnsetIsDeleted
`func (o *SearchObjectAllOf) UnsetIsDeleted()`

UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
### GetSourceInfo

`func (o *SearchObjectAllOf) GetSourceInfo() Object`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *SearchObjectAllOf) GetSourceInfoOk() (*Object, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *SearchObjectAllOf) SetSourceInfo(v Object)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *SearchObjectAllOf) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.

### GetObjectProtectionInfos

`func (o *SearchObjectAllOf) GetObjectProtectionInfos() []ObjectProtectionInfo`

GetObjectProtectionInfos returns the ObjectProtectionInfos field if non-nil, zero value otherwise.

### GetObjectProtectionInfosOk

`func (o *SearchObjectAllOf) GetObjectProtectionInfosOk() (*[]ObjectProtectionInfo, bool)`

GetObjectProtectionInfosOk returns a tuple with the ObjectProtectionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionInfos

`func (o *SearchObjectAllOf) SetObjectProtectionInfos(v []ObjectProtectionInfo)`

SetObjectProtectionInfos sets ObjectProtectionInfos field to given value.

### HasObjectProtectionInfos

`func (o *SearchObjectAllOf) HasObjectProtectionInfos() bool`

HasObjectProtectionInfos returns a boolean if a field has been set.

### SetObjectProtectionInfosNil

`func (o *SearchObjectAllOf) SetObjectProtectionInfosNil(b bool)`

 SetObjectProtectionInfosNil sets the value for ObjectProtectionInfos to be an explicit nil

### UnsetObjectProtectionInfos
`func (o *SearchObjectAllOf) UnsetObjectProtectionInfos()`

UnsetObjectProtectionInfos ensures that no value is present for ObjectProtectionInfos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


