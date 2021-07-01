# CompareAdObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreTaskId** | **NullableInt64** | Specifies the Restore Task Id corresponding to which we need to compare the AD objects. | 
**AllowEmptyDestGuids** | Pointer to **NullableBool** | Specifies the option to get object attributes from Snapshot AD when destination guid is missing in GuidPair. This helps to show attributes of AD object from Snapshot AD when the object is missing in Production AD. | [optional] 
**ExcludeSysAttributes** | Pointer to **NullableBool** | Specifies the option to exclude AD system attributes when comparing two AD object attributes. If the objects have same guid, most of the system attributes would match.If the AD object was recovered through a restore, then many system attributes will be different. Default compares all attributes. | [optional] 
**FilterNullValueAttributes** | Pointer to **NullableBool** | Specifies the option to not return attributes where source and destination values are null values. This reduces noise of the properties in the objects returned. | [optional] 
**FilterSameValueAttributes** | Pointer to **NullableBool** | Specifies the option to not return attributes where source and destination values are same. Use this flag to return only values that are different. | [optional] 
**GuidPairs** | [**[]GuidPair**](GuidPair.md) | Specifies the GuidPair of the AD Objects which we want to compare from both Snapshot and Production AD. | 
**QuickCompare** | Pointer to **NullableBool** | Specifies the option to do quick compare of specified guid between Snapshot AD and Production AD. If at least one attribute mismatch is found, comparison stops and returns with AdObjectFlag kNotEqual. | [optional] 

## Methods

### NewCompareAdObjectsRequest

`func NewCompareAdObjectsRequest(restoreTaskId NullableInt64, guidPairs []GuidPair, ) *CompareAdObjectsRequest`

NewCompareAdObjectsRequest instantiates a new CompareAdObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompareAdObjectsRequestWithDefaults

`func NewCompareAdObjectsRequestWithDefaults() *CompareAdObjectsRequest`

NewCompareAdObjectsRequestWithDefaults instantiates a new CompareAdObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreTaskId

`func (o *CompareAdObjectsRequest) GetRestoreTaskId() int64`

GetRestoreTaskId returns the RestoreTaskId field if non-nil, zero value otherwise.

### GetRestoreTaskIdOk

`func (o *CompareAdObjectsRequest) GetRestoreTaskIdOk() (*int64, bool)`

GetRestoreTaskIdOk returns a tuple with the RestoreTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTaskId

`func (o *CompareAdObjectsRequest) SetRestoreTaskId(v int64)`

SetRestoreTaskId sets RestoreTaskId field to given value.


### SetRestoreTaskIdNil

`func (o *CompareAdObjectsRequest) SetRestoreTaskIdNil(b bool)`

 SetRestoreTaskIdNil sets the value for RestoreTaskId to be an explicit nil

### UnsetRestoreTaskId
`func (o *CompareAdObjectsRequest) UnsetRestoreTaskId()`

UnsetRestoreTaskId ensures that no value is present for RestoreTaskId, not even an explicit nil
### GetAllowEmptyDestGuids

`func (o *CompareAdObjectsRequest) GetAllowEmptyDestGuids() bool`

GetAllowEmptyDestGuids returns the AllowEmptyDestGuids field if non-nil, zero value otherwise.

### GetAllowEmptyDestGuidsOk

`func (o *CompareAdObjectsRequest) GetAllowEmptyDestGuidsOk() (*bool, bool)`

GetAllowEmptyDestGuidsOk returns a tuple with the AllowEmptyDestGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmptyDestGuids

`func (o *CompareAdObjectsRequest) SetAllowEmptyDestGuids(v bool)`

SetAllowEmptyDestGuids sets AllowEmptyDestGuids field to given value.

### HasAllowEmptyDestGuids

`func (o *CompareAdObjectsRequest) HasAllowEmptyDestGuids() bool`

HasAllowEmptyDestGuids returns a boolean if a field has been set.

### SetAllowEmptyDestGuidsNil

`func (o *CompareAdObjectsRequest) SetAllowEmptyDestGuidsNil(b bool)`

 SetAllowEmptyDestGuidsNil sets the value for AllowEmptyDestGuids to be an explicit nil

### UnsetAllowEmptyDestGuids
`func (o *CompareAdObjectsRequest) UnsetAllowEmptyDestGuids()`

UnsetAllowEmptyDestGuids ensures that no value is present for AllowEmptyDestGuids, not even an explicit nil
### GetExcludeSysAttributes

`func (o *CompareAdObjectsRequest) GetExcludeSysAttributes() bool`

GetExcludeSysAttributes returns the ExcludeSysAttributes field if non-nil, zero value otherwise.

### GetExcludeSysAttributesOk

`func (o *CompareAdObjectsRequest) GetExcludeSysAttributesOk() (*bool, bool)`

GetExcludeSysAttributesOk returns a tuple with the ExcludeSysAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSysAttributes

`func (o *CompareAdObjectsRequest) SetExcludeSysAttributes(v bool)`

SetExcludeSysAttributes sets ExcludeSysAttributes field to given value.

### HasExcludeSysAttributes

`func (o *CompareAdObjectsRequest) HasExcludeSysAttributes() bool`

HasExcludeSysAttributes returns a boolean if a field has been set.

### SetExcludeSysAttributesNil

`func (o *CompareAdObjectsRequest) SetExcludeSysAttributesNil(b bool)`

 SetExcludeSysAttributesNil sets the value for ExcludeSysAttributes to be an explicit nil

### UnsetExcludeSysAttributes
`func (o *CompareAdObjectsRequest) UnsetExcludeSysAttributes()`

UnsetExcludeSysAttributes ensures that no value is present for ExcludeSysAttributes, not even an explicit nil
### GetFilterNullValueAttributes

`func (o *CompareAdObjectsRequest) GetFilterNullValueAttributes() bool`

GetFilterNullValueAttributes returns the FilterNullValueAttributes field if non-nil, zero value otherwise.

### GetFilterNullValueAttributesOk

`func (o *CompareAdObjectsRequest) GetFilterNullValueAttributesOk() (*bool, bool)`

GetFilterNullValueAttributesOk returns a tuple with the FilterNullValueAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterNullValueAttributes

`func (o *CompareAdObjectsRequest) SetFilterNullValueAttributes(v bool)`

SetFilterNullValueAttributes sets FilterNullValueAttributes field to given value.

### HasFilterNullValueAttributes

`func (o *CompareAdObjectsRequest) HasFilterNullValueAttributes() bool`

HasFilterNullValueAttributes returns a boolean if a field has been set.

### SetFilterNullValueAttributesNil

`func (o *CompareAdObjectsRequest) SetFilterNullValueAttributesNil(b bool)`

 SetFilterNullValueAttributesNil sets the value for FilterNullValueAttributes to be an explicit nil

### UnsetFilterNullValueAttributes
`func (o *CompareAdObjectsRequest) UnsetFilterNullValueAttributes()`

UnsetFilterNullValueAttributes ensures that no value is present for FilterNullValueAttributes, not even an explicit nil
### GetFilterSameValueAttributes

`func (o *CompareAdObjectsRequest) GetFilterSameValueAttributes() bool`

GetFilterSameValueAttributes returns the FilterSameValueAttributes field if non-nil, zero value otherwise.

### GetFilterSameValueAttributesOk

`func (o *CompareAdObjectsRequest) GetFilterSameValueAttributesOk() (*bool, bool)`

GetFilterSameValueAttributesOk returns a tuple with the FilterSameValueAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSameValueAttributes

`func (o *CompareAdObjectsRequest) SetFilterSameValueAttributes(v bool)`

SetFilterSameValueAttributes sets FilterSameValueAttributes field to given value.

### HasFilterSameValueAttributes

`func (o *CompareAdObjectsRequest) HasFilterSameValueAttributes() bool`

HasFilterSameValueAttributes returns a boolean if a field has been set.

### SetFilterSameValueAttributesNil

`func (o *CompareAdObjectsRequest) SetFilterSameValueAttributesNil(b bool)`

 SetFilterSameValueAttributesNil sets the value for FilterSameValueAttributes to be an explicit nil

### UnsetFilterSameValueAttributes
`func (o *CompareAdObjectsRequest) UnsetFilterSameValueAttributes()`

UnsetFilterSameValueAttributes ensures that no value is present for FilterSameValueAttributes, not even an explicit nil
### GetGuidPairs

`func (o *CompareAdObjectsRequest) GetGuidPairs() []GuidPair`

GetGuidPairs returns the GuidPairs field if non-nil, zero value otherwise.

### GetGuidPairsOk

`func (o *CompareAdObjectsRequest) GetGuidPairsOk() (*[]GuidPair, bool)`

GetGuidPairsOk returns a tuple with the GuidPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidPairs

`func (o *CompareAdObjectsRequest) SetGuidPairs(v []GuidPair)`

SetGuidPairs sets GuidPairs field to given value.


### SetGuidPairsNil

`func (o *CompareAdObjectsRequest) SetGuidPairsNil(b bool)`

 SetGuidPairsNil sets the value for GuidPairs to be an explicit nil

### UnsetGuidPairs
`func (o *CompareAdObjectsRequest) UnsetGuidPairs()`

UnsetGuidPairs ensures that no value is present for GuidPairs, not even an explicit nil
### GetQuickCompare

`func (o *CompareAdObjectsRequest) GetQuickCompare() bool`

GetQuickCompare returns the QuickCompare field if non-nil, zero value otherwise.

### GetQuickCompareOk

`func (o *CompareAdObjectsRequest) GetQuickCompareOk() (*bool, bool)`

GetQuickCompareOk returns a tuple with the QuickCompare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickCompare

`func (o *CompareAdObjectsRequest) SetQuickCompare(v bool)`

SetQuickCompare sets QuickCompare field to given value.

### HasQuickCompare

`func (o *CompareAdObjectsRequest) HasQuickCompare() bool`

HasQuickCompare returns a boolean if a field has been set.

### SetQuickCompareNil

`func (o *CompareAdObjectsRequest) SetQuickCompareNil(b bool)`

 SetQuickCompareNil sets the value for QuickCompare to be an explicit nil

### UnsetQuickCompare
`func (o *CompareAdObjectsRequest) UnsetQuickCompare()`

UnsetQuickCompare ensures that no value is present for QuickCompare, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


