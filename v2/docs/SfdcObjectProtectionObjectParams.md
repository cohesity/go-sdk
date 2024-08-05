# SfdcObjectProtectionObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the ids of the objects to be excluded in the Object Protection. | [optional] 
**FieldExclusionList** | Pointer to [**[]SfdcObjectFieldExclusion**](SfdcObjectFieldExclusion.md) | Specifies the list of field names to be excluded in an Sfdc object. A user can specify multiple such entries in this list. | [optional] 
**Id** | **int64** | Specifies the id of the Sfdc Org being protected. This cannot be the id of a leaf level object. By default, the Sfdc Org is auto-protected. | 

## Methods

### NewSfdcObjectProtectionObjectParams

`func NewSfdcObjectProtectionObjectParams(id int64, ) *SfdcObjectProtectionObjectParams`

NewSfdcObjectProtectionObjectParams instantiates a new SfdcObjectProtectionObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcObjectProtectionObjectParamsWithDefaults

`func NewSfdcObjectProtectionObjectParamsWithDefaults() *SfdcObjectProtectionObjectParams`

NewSfdcObjectProtectionObjectParamsWithDefaults instantiates a new SfdcObjectProtectionObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeObjectIds

`func (o *SfdcObjectProtectionObjectParams) GetExcludeObjectIds() []int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *SfdcObjectProtectionObjectParams) GetExcludeObjectIdsOk() (*[]int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *SfdcObjectProtectionObjectParams) SetExcludeObjectIds(v []int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *SfdcObjectProtectionObjectParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### SetExcludeObjectIdsNil

`func (o *SfdcObjectProtectionObjectParams) SetExcludeObjectIdsNil(b bool)`

 SetExcludeObjectIdsNil sets the value for ExcludeObjectIds to be an explicit nil

### UnsetExcludeObjectIds
`func (o *SfdcObjectProtectionObjectParams) UnsetExcludeObjectIds()`

UnsetExcludeObjectIds ensures that no value is present for ExcludeObjectIds, not even an explicit nil
### GetFieldExclusionList

`func (o *SfdcObjectProtectionObjectParams) GetFieldExclusionList() []SfdcObjectFieldExclusion`

GetFieldExclusionList returns the FieldExclusionList field if non-nil, zero value otherwise.

### GetFieldExclusionListOk

`func (o *SfdcObjectProtectionObjectParams) GetFieldExclusionListOk() (*[]SfdcObjectFieldExclusion, bool)`

GetFieldExclusionListOk returns a tuple with the FieldExclusionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldExclusionList

`func (o *SfdcObjectProtectionObjectParams) SetFieldExclusionList(v []SfdcObjectFieldExclusion)`

SetFieldExclusionList sets FieldExclusionList field to given value.

### HasFieldExclusionList

`func (o *SfdcObjectProtectionObjectParams) HasFieldExclusionList() bool`

HasFieldExclusionList returns a boolean if a field has been set.

### SetFieldExclusionListNil

`func (o *SfdcObjectProtectionObjectParams) SetFieldExclusionListNil(b bool)`

 SetFieldExclusionListNil sets the value for FieldExclusionList to be an explicit nil

### UnsetFieldExclusionList
`func (o *SfdcObjectProtectionObjectParams) UnsetFieldExclusionList()`

UnsetFieldExclusionList ensures that no value is present for FieldExclusionList, not even an explicit nil
### GetId

`func (o *SfdcObjectProtectionObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SfdcObjectProtectionObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SfdcObjectProtectionObjectParams) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


