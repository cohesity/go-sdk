# CompareADObjectsResultADObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeVec** | Pointer to [**[]CompareADObjectsResultADAttribute**](CompareADObjectsResultADAttribute.md) | Array of AD attributes of AD object. This will contain distinct attributes from source and destination objects. | [optional] 
**DestGuid** | Pointer to **NullableString** | Object guid from dest_server. If empty, compare could not find an AD object corresponding to &#39;source_guid&#39; even after looking up based on source_guid, source DN or source SAM account name. The SAM is applicable only for account type objects. | [optional] 
**DestPropCount** | Pointer to **NullableInt32** | Number of attributes in destination object including system properties compared. This count is useful for debugging. | [optional] 
**ExcludedPropCount** | Pointer to **NullableInt32** | Number of attributes not compared due to ADCompareOptionFlags.kExcludeSysProps. This count is useful for debugging. | [optional] 
**MismatchPropCount** | Pointer to **NullableInt32** | Number of AD attributes compared based on &#39;ADCompareOptionFlagsType&#39; flags and found to be mismatched. If this is non-zero, compared objects are different. If this is 0 ann&#39;dest_guid&#39; is empty, then object is missing. | [optional] 
**ObjectFlags** | Pointer to **NullableInt32** | Object result flags of type ADObjectFlags. | [optional] 
**SourceGuid** | Pointer to **NullableString** | Object guid from $SourceServer. Guid string with or without &#39;{}&#39; braces. | [optional] 
**SourcePropCount** | Pointer to **NullableInt32** | Number of attributes in source object including system properties compared. This count is useful for debugging. | [optional] 
**Status** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 

## Methods

### NewCompareADObjectsResultADObject

`func NewCompareADObjectsResultADObject() *CompareADObjectsResultADObject`

NewCompareADObjectsResultADObject instantiates a new CompareADObjectsResultADObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompareADObjectsResultADObjectWithDefaults

`func NewCompareADObjectsResultADObjectWithDefaults() *CompareADObjectsResultADObject`

NewCompareADObjectsResultADObjectWithDefaults instantiates a new CompareADObjectsResultADObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeVec

`func (o *CompareADObjectsResultADObject) GetAttributeVec() []CompareADObjectsResultADAttribute`

GetAttributeVec returns the AttributeVec field if non-nil, zero value otherwise.

### GetAttributeVecOk

`func (o *CompareADObjectsResultADObject) GetAttributeVecOk() (*[]CompareADObjectsResultADAttribute, bool)`

GetAttributeVecOk returns a tuple with the AttributeVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeVec

`func (o *CompareADObjectsResultADObject) SetAttributeVec(v []CompareADObjectsResultADAttribute)`

SetAttributeVec sets AttributeVec field to given value.

### HasAttributeVec

`func (o *CompareADObjectsResultADObject) HasAttributeVec() bool`

HasAttributeVec returns a boolean if a field has been set.

### SetAttributeVecNil

`func (o *CompareADObjectsResultADObject) SetAttributeVecNil(b bool)`

 SetAttributeVecNil sets the value for AttributeVec to be an explicit nil

### UnsetAttributeVec
`func (o *CompareADObjectsResultADObject) UnsetAttributeVec()`

UnsetAttributeVec ensures that no value is present for AttributeVec, not even an explicit nil
### GetDestGuid

`func (o *CompareADObjectsResultADObject) GetDestGuid() string`

GetDestGuid returns the DestGuid field if non-nil, zero value otherwise.

### GetDestGuidOk

`func (o *CompareADObjectsResultADObject) GetDestGuidOk() (*string, bool)`

GetDestGuidOk returns a tuple with the DestGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestGuid

`func (o *CompareADObjectsResultADObject) SetDestGuid(v string)`

SetDestGuid sets DestGuid field to given value.

### HasDestGuid

`func (o *CompareADObjectsResultADObject) HasDestGuid() bool`

HasDestGuid returns a boolean if a field has been set.

### SetDestGuidNil

`func (o *CompareADObjectsResultADObject) SetDestGuidNil(b bool)`

 SetDestGuidNil sets the value for DestGuid to be an explicit nil

### UnsetDestGuid
`func (o *CompareADObjectsResultADObject) UnsetDestGuid()`

UnsetDestGuid ensures that no value is present for DestGuid, not even an explicit nil
### GetDestPropCount

`func (o *CompareADObjectsResultADObject) GetDestPropCount() int32`

GetDestPropCount returns the DestPropCount field if non-nil, zero value otherwise.

### GetDestPropCountOk

`func (o *CompareADObjectsResultADObject) GetDestPropCountOk() (*int32, bool)`

GetDestPropCountOk returns a tuple with the DestPropCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPropCount

`func (o *CompareADObjectsResultADObject) SetDestPropCount(v int32)`

SetDestPropCount sets DestPropCount field to given value.

### HasDestPropCount

`func (o *CompareADObjectsResultADObject) HasDestPropCount() bool`

HasDestPropCount returns a boolean if a field has been set.

### SetDestPropCountNil

`func (o *CompareADObjectsResultADObject) SetDestPropCountNil(b bool)`

 SetDestPropCountNil sets the value for DestPropCount to be an explicit nil

### UnsetDestPropCount
`func (o *CompareADObjectsResultADObject) UnsetDestPropCount()`

UnsetDestPropCount ensures that no value is present for DestPropCount, not even an explicit nil
### GetExcludedPropCount

`func (o *CompareADObjectsResultADObject) GetExcludedPropCount() int32`

GetExcludedPropCount returns the ExcludedPropCount field if non-nil, zero value otherwise.

### GetExcludedPropCountOk

`func (o *CompareADObjectsResultADObject) GetExcludedPropCountOk() (*int32, bool)`

GetExcludedPropCountOk returns a tuple with the ExcludedPropCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPropCount

`func (o *CompareADObjectsResultADObject) SetExcludedPropCount(v int32)`

SetExcludedPropCount sets ExcludedPropCount field to given value.

### HasExcludedPropCount

`func (o *CompareADObjectsResultADObject) HasExcludedPropCount() bool`

HasExcludedPropCount returns a boolean if a field has been set.

### SetExcludedPropCountNil

`func (o *CompareADObjectsResultADObject) SetExcludedPropCountNil(b bool)`

 SetExcludedPropCountNil sets the value for ExcludedPropCount to be an explicit nil

### UnsetExcludedPropCount
`func (o *CompareADObjectsResultADObject) UnsetExcludedPropCount()`

UnsetExcludedPropCount ensures that no value is present for ExcludedPropCount, not even an explicit nil
### GetMismatchPropCount

`func (o *CompareADObjectsResultADObject) GetMismatchPropCount() int32`

GetMismatchPropCount returns the MismatchPropCount field if non-nil, zero value otherwise.

### GetMismatchPropCountOk

`func (o *CompareADObjectsResultADObject) GetMismatchPropCountOk() (*int32, bool)`

GetMismatchPropCountOk returns a tuple with the MismatchPropCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchPropCount

`func (o *CompareADObjectsResultADObject) SetMismatchPropCount(v int32)`

SetMismatchPropCount sets MismatchPropCount field to given value.

### HasMismatchPropCount

`func (o *CompareADObjectsResultADObject) HasMismatchPropCount() bool`

HasMismatchPropCount returns a boolean if a field has been set.

### SetMismatchPropCountNil

`func (o *CompareADObjectsResultADObject) SetMismatchPropCountNil(b bool)`

 SetMismatchPropCountNil sets the value for MismatchPropCount to be an explicit nil

### UnsetMismatchPropCount
`func (o *CompareADObjectsResultADObject) UnsetMismatchPropCount()`

UnsetMismatchPropCount ensures that no value is present for MismatchPropCount, not even an explicit nil
### GetObjectFlags

`func (o *CompareADObjectsResultADObject) GetObjectFlags() int32`

GetObjectFlags returns the ObjectFlags field if non-nil, zero value otherwise.

### GetObjectFlagsOk

`func (o *CompareADObjectsResultADObject) GetObjectFlagsOk() (*int32, bool)`

GetObjectFlagsOk returns a tuple with the ObjectFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFlags

`func (o *CompareADObjectsResultADObject) SetObjectFlags(v int32)`

SetObjectFlags sets ObjectFlags field to given value.

### HasObjectFlags

`func (o *CompareADObjectsResultADObject) HasObjectFlags() bool`

HasObjectFlags returns a boolean if a field has been set.

### SetObjectFlagsNil

`func (o *CompareADObjectsResultADObject) SetObjectFlagsNil(b bool)`

 SetObjectFlagsNil sets the value for ObjectFlags to be an explicit nil

### UnsetObjectFlags
`func (o *CompareADObjectsResultADObject) UnsetObjectFlags()`

UnsetObjectFlags ensures that no value is present for ObjectFlags, not even an explicit nil
### GetSourceGuid

`func (o *CompareADObjectsResultADObject) GetSourceGuid() string`

GetSourceGuid returns the SourceGuid field if non-nil, zero value otherwise.

### GetSourceGuidOk

`func (o *CompareADObjectsResultADObject) GetSourceGuidOk() (*string, bool)`

GetSourceGuidOk returns a tuple with the SourceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuid

`func (o *CompareADObjectsResultADObject) SetSourceGuid(v string)`

SetSourceGuid sets SourceGuid field to given value.

### HasSourceGuid

`func (o *CompareADObjectsResultADObject) HasSourceGuid() bool`

HasSourceGuid returns a boolean if a field has been set.

### SetSourceGuidNil

`func (o *CompareADObjectsResultADObject) SetSourceGuidNil(b bool)`

 SetSourceGuidNil sets the value for SourceGuid to be an explicit nil

### UnsetSourceGuid
`func (o *CompareADObjectsResultADObject) UnsetSourceGuid()`

UnsetSourceGuid ensures that no value is present for SourceGuid, not even an explicit nil
### GetSourcePropCount

`func (o *CompareADObjectsResultADObject) GetSourcePropCount() int32`

GetSourcePropCount returns the SourcePropCount field if non-nil, zero value otherwise.

### GetSourcePropCountOk

`func (o *CompareADObjectsResultADObject) GetSourcePropCountOk() (*int32, bool)`

GetSourcePropCountOk returns a tuple with the SourcePropCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePropCount

`func (o *CompareADObjectsResultADObject) SetSourcePropCount(v int32)`

SetSourcePropCount sets SourcePropCount field to given value.

### HasSourcePropCount

`func (o *CompareADObjectsResultADObject) HasSourcePropCount() bool`

HasSourcePropCount returns a boolean if a field has been set.

### SetSourcePropCountNil

`func (o *CompareADObjectsResultADObject) SetSourcePropCountNil(b bool)`

 SetSourcePropCountNil sets the value for SourcePropCount to be an explicit nil

### UnsetSourcePropCount
`func (o *CompareADObjectsResultADObject) UnsetSourcePropCount()`

UnsetSourcePropCount ensures that no value is present for SourcePropCount, not even an explicit nil
### GetStatus

`func (o *CompareADObjectsResultADObject) GetStatus() ErrorProto`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompareADObjectsResultADObject) GetStatusOk() (*ErrorProto, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompareADObjectsResultADObject) SetStatus(v ErrorProto)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompareADObjectsResultADObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


