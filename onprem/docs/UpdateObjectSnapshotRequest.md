# UpdateObjectSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetLegalHold** | Pointer to **NullableBool** | Whether to set the snapshot on legal hold. If set to true, the run cannot be deleted during the retention period. | [optional] 
**DataLockType** | Pointer to **NullableString** | Specifies the snapshot data lock type. | [optional] 
**ExpiryTimeSecs** | Pointer to **NullableInt32** | Specifies the expiry time of the snapshot in Unix timestamp epoch in seconds. | [optional] 

## Methods

### NewUpdateObjectSnapshotRequest

`func NewUpdateObjectSnapshotRequest() *UpdateObjectSnapshotRequest`

NewUpdateObjectSnapshotRequest instantiates a new UpdateObjectSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateObjectSnapshotRequestWithDefaults

`func NewUpdateObjectSnapshotRequestWithDefaults() *UpdateObjectSnapshotRequest`

NewUpdateObjectSnapshotRequestWithDefaults instantiates a new UpdateObjectSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetLegalHold

`func (o *UpdateObjectSnapshotRequest) GetSetLegalHold() bool`

GetSetLegalHold returns the SetLegalHold field if non-nil, zero value otherwise.

### GetSetLegalHoldOk

`func (o *UpdateObjectSnapshotRequest) GetSetLegalHoldOk() (*bool, bool)`

GetSetLegalHoldOk returns a tuple with the SetLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetLegalHold

`func (o *UpdateObjectSnapshotRequest) SetSetLegalHold(v bool)`

SetSetLegalHold sets SetLegalHold field to given value.

### HasSetLegalHold

`func (o *UpdateObjectSnapshotRequest) HasSetLegalHold() bool`

HasSetLegalHold returns a boolean if a field has been set.

### SetSetLegalHoldNil

`func (o *UpdateObjectSnapshotRequest) SetSetLegalHoldNil(b bool)`

 SetSetLegalHoldNil sets the value for SetLegalHold to be an explicit nil

### UnsetSetLegalHold
`func (o *UpdateObjectSnapshotRequest) UnsetSetLegalHold()`

UnsetSetLegalHold ensures that no value is present for SetLegalHold, not even an explicit nil
### GetDataLockType

`func (o *UpdateObjectSnapshotRequest) GetDataLockType() string`

GetDataLockType returns the DataLockType field if non-nil, zero value otherwise.

### GetDataLockTypeOk

`func (o *UpdateObjectSnapshotRequest) GetDataLockTypeOk() (*string, bool)`

GetDataLockTypeOk returns a tuple with the DataLockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockType

`func (o *UpdateObjectSnapshotRequest) SetDataLockType(v string)`

SetDataLockType sets DataLockType field to given value.

### HasDataLockType

`func (o *UpdateObjectSnapshotRequest) HasDataLockType() bool`

HasDataLockType returns a boolean if a field has been set.

### SetDataLockTypeNil

`func (o *UpdateObjectSnapshotRequest) SetDataLockTypeNil(b bool)`

 SetDataLockTypeNil sets the value for DataLockType to be an explicit nil

### UnsetDataLockType
`func (o *UpdateObjectSnapshotRequest) UnsetDataLockType()`

UnsetDataLockType ensures that no value is present for DataLockType, not even an explicit nil
### GetExpiryTimeSecs

`func (o *UpdateObjectSnapshotRequest) GetExpiryTimeSecs() int32`

GetExpiryTimeSecs returns the ExpiryTimeSecs field if non-nil, zero value otherwise.

### GetExpiryTimeSecsOk

`func (o *UpdateObjectSnapshotRequest) GetExpiryTimeSecsOk() (*int32, bool)`

GetExpiryTimeSecsOk returns a tuple with the ExpiryTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeSecs

`func (o *UpdateObjectSnapshotRequest) SetExpiryTimeSecs(v int32)`

SetExpiryTimeSecs sets ExpiryTimeSecs field to given value.

### HasExpiryTimeSecs

`func (o *UpdateObjectSnapshotRequest) HasExpiryTimeSecs() bool`

HasExpiryTimeSecs returns a boolean if a field has been set.

### SetExpiryTimeSecsNil

`func (o *UpdateObjectSnapshotRequest) SetExpiryTimeSecsNil(b bool)`

 SetExpiryTimeSecsNil sets the value for ExpiryTimeSecs to be an explicit nil

### UnsetExpiryTimeSecs
`func (o *UpdateObjectSnapshotRequest) UnsetExpiryTimeSecs()`

UnsetExpiryTimeSecs ensures that no value is present for ExpiryTimeSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


