# ADObjectRestoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestGuid** | Pointer to **NullableString** | Destination guid string of the AD object that is newly created on production AD corresponding to &#39;source_guid&#39;. If the object was restored from production AD recycle Bin, this value can be empty or set to same value as &#39;source_guid&#39;. If this value is non-empty and is different from source_guid, implying production AD object is a new object created in production AD as part of restore. | [optional] 
**ObjectFlags** | Pointer to **NullableInt32** | Object result flags of type ADObjectFlags. | [optional] 
**PropertyStatusVec** | Pointer to [**[]ADObjectRestoreStatusADAttributeRestoreStatus**](ADObjectRestoreStatusADAttributeRestoreStatus.md) | AD object attribute(property) restore status vector. | [optional] 
**SourceGuid** | Pointer to **NullableString** | Source guid of AD object that was restored. This will not be empty. This is populated from the source of request argument. | [optional] 
**Status** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**TimetakenMs** | Pointer to **NullableInt32** | Time taken in milliseconds to restore the individual object or attribute update. If this object restore was part of a batch, it shows the time taken once the operation was dispatched to AD for the object. This time can be useful in answering why some objects took long time to restore. Note that this time is not the elapsed time when the request was made from Magneto. | [optional] 

## Methods

### NewADObjectRestoreStatus

`func NewADObjectRestoreStatus() *ADObjectRestoreStatus`

NewADObjectRestoreStatus instantiates a new ADObjectRestoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADObjectRestoreStatusWithDefaults

`func NewADObjectRestoreStatusWithDefaults() *ADObjectRestoreStatus`

NewADObjectRestoreStatusWithDefaults instantiates a new ADObjectRestoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestGuid

`func (o *ADObjectRestoreStatus) GetDestGuid() string`

GetDestGuid returns the DestGuid field if non-nil, zero value otherwise.

### GetDestGuidOk

`func (o *ADObjectRestoreStatus) GetDestGuidOk() (*string, bool)`

GetDestGuidOk returns a tuple with the DestGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestGuid

`func (o *ADObjectRestoreStatus) SetDestGuid(v string)`

SetDestGuid sets DestGuid field to given value.

### HasDestGuid

`func (o *ADObjectRestoreStatus) HasDestGuid() bool`

HasDestGuid returns a boolean if a field has been set.

### SetDestGuidNil

`func (o *ADObjectRestoreStatus) SetDestGuidNil(b bool)`

 SetDestGuidNil sets the value for DestGuid to be an explicit nil

### UnsetDestGuid
`func (o *ADObjectRestoreStatus) UnsetDestGuid()`

UnsetDestGuid ensures that no value is present for DestGuid, not even an explicit nil
### GetObjectFlags

`func (o *ADObjectRestoreStatus) GetObjectFlags() int32`

GetObjectFlags returns the ObjectFlags field if non-nil, zero value otherwise.

### GetObjectFlagsOk

`func (o *ADObjectRestoreStatus) GetObjectFlagsOk() (*int32, bool)`

GetObjectFlagsOk returns a tuple with the ObjectFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFlags

`func (o *ADObjectRestoreStatus) SetObjectFlags(v int32)`

SetObjectFlags sets ObjectFlags field to given value.

### HasObjectFlags

`func (o *ADObjectRestoreStatus) HasObjectFlags() bool`

HasObjectFlags returns a boolean if a field has been set.

### SetObjectFlagsNil

`func (o *ADObjectRestoreStatus) SetObjectFlagsNil(b bool)`

 SetObjectFlagsNil sets the value for ObjectFlags to be an explicit nil

### UnsetObjectFlags
`func (o *ADObjectRestoreStatus) UnsetObjectFlags()`

UnsetObjectFlags ensures that no value is present for ObjectFlags, not even an explicit nil
### GetPropertyStatusVec

`func (o *ADObjectRestoreStatus) GetPropertyStatusVec() []ADObjectRestoreStatusADAttributeRestoreStatus`

GetPropertyStatusVec returns the PropertyStatusVec field if non-nil, zero value otherwise.

### GetPropertyStatusVecOk

`func (o *ADObjectRestoreStatus) GetPropertyStatusVecOk() (*[]ADObjectRestoreStatusADAttributeRestoreStatus, bool)`

GetPropertyStatusVecOk returns a tuple with the PropertyStatusVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyStatusVec

`func (o *ADObjectRestoreStatus) SetPropertyStatusVec(v []ADObjectRestoreStatusADAttributeRestoreStatus)`

SetPropertyStatusVec sets PropertyStatusVec field to given value.

### HasPropertyStatusVec

`func (o *ADObjectRestoreStatus) HasPropertyStatusVec() bool`

HasPropertyStatusVec returns a boolean if a field has been set.

### SetPropertyStatusVecNil

`func (o *ADObjectRestoreStatus) SetPropertyStatusVecNil(b bool)`

 SetPropertyStatusVecNil sets the value for PropertyStatusVec to be an explicit nil

### UnsetPropertyStatusVec
`func (o *ADObjectRestoreStatus) UnsetPropertyStatusVec()`

UnsetPropertyStatusVec ensures that no value is present for PropertyStatusVec, not even an explicit nil
### GetSourceGuid

`func (o *ADObjectRestoreStatus) GetSourceGuid() string`

GetSourceGuid returns the SourceGuid field if non-nil, zero value otherwise.

### GetSourceGuidOk

`func (o *ADObjectRestoreStatus) GetSourceGuidOk() (*string, bool)`

GetSourceGuidOk returns a tuple with the SourceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuid

`func (o *ADObjectRestoreStatus) SetSourceGuid(v string)`

SetSourceGuid sets SourceGuid field to given value.

### HasSourceGuid

`func (o *ADObjectRestoreStatus) HasSourceGuid() bool`

HasSourceGuid returns a boolean if a field has been set.

### SetSourceGuidNil

`func (o *ADObjectRestoreStatus) SetSourceGuidNil(b bool)`

 SetSourceGuidNil sets the value for SourceGuid to be an explicit nil

### UnsetSourceGuid
`func (o *ADObjectRestoreStatus) UnsetSourceGuid()`

UnsetSourceGuid ensures that no value is present for SourceGuid, not even an explicit nil
### GetStatus

`func (o *ADObjectRestoreStatus) GetStatus() ErrorProto`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ADObjectRestoreStatus) GetStatusOk() (*ErrorProto, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ADObjectRestoreStatus) SetStatus(v ErrorProto)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ADObjectRestoreStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimetakenMs

`func (o *ADObjectRestoreStatus) GetTimetakenMs() int32`

GetTimetakenMs returns the TimetakenMs field if non-nil, zero value otherwise.

### GetTimetakenMsOk

`func (o *ADObjectRestoreStatus) GetTimetakenMsOk() (*int32, bool)`

GetTimetakenMsOk returns a tuple with the TimetakenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimetakenMs

`func (o *ADObjectRestoreStatus) SetTimetakenMs(v int32)`

SetTimetakenMs sets TimetakenMs field to given value.

### HasTimetakenMs

`func (o *ADObjectRestoreStatus) HasTimetakenMs() bool`

HasTimetakenMs returns a boolean if a field has been set.

### SetTimetakenMsNil

`func (o *ADObjectRestoreStatus) SetTimetakenMsNil(b bool)`

 SetTimetakenMsNil sets the value for TimetakenMs to be an explicit nil

### UnsetTimetakenMs
`func (o *ADObjectRestoreStatus) UnsetTimetakenMs()`

UnsetTimetakenMs ensures that no value is present for TimetakenMs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


