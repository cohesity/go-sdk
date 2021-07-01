# UnprotectObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteSnapshots** | Pointer to **NullableBool** | Specifies whether to delete the snapshots of the Protection Object. | [optional] 
**ProtectionSourceId** | **NullableInt64** | Specifies the id of the Protection Source to be unprotected. | 
**RpoPolicyId** | **NullableString** | Specifies the id of the Rpo Policy from which to unprotect the object. | 

## Methods

### NewUnprotectObjectParams

`func NewUnprotectObjectParams(protectionSourceId NullableInt64, rpoPolicyId NullableString, ) *UnprotectObjectParams`

NewUnprotectObjectParams instantiates a new UnprotectObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnprotectObjectParamsWithDefaults

`func NewUnprotectObjectParamsWithDefaults() *UnprotectObjectParams`

NewUnprotectObjectParamsWithDefaults instantiates a new UnprotectObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteSnapshots

`func (o *UnprotectObjectParams) GetDeleteSnapshots() bool`

GetDeleteSnapshots returns the DeleteSnapshots field if non-nil, zero value otherwise.

### GetDeleteSnapshotsOk

`func (o *UnprotectObjectParams) GetDeleteSnapshotsOk() (*bool, bool)`

GetDeleteSnapshotsOk returns a tuple with the DeleteSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSnapshots

`func (o *UnprotectObjectParams) SetDeleteSnapshots(v bool)`

SetDeleteSnapshots sets DeleteSnapshots field to given value.

### HasDeleteSnapshots

`func (o *UnprotectObjectParams) HasDeleteSnapshots() bool`

HasDeleteSnapshots returns a boolean if a field has been set.

### SetDeleteSnapshotsNil

`func (o *UnprotectObjectParams) SetDeleteSnapshotsNil(b bool)`

 SetDeleteSnapshotsNil sets the value for DeleteSnapshots to be an explicit nil

### UnsetDeleteSnapshots
`func (o *UnprotectObjectParams) UnsetDeleteSnapshots()`

UnsetDeleteSnapshots ensures that no value is present for DeleteSnapshots, not even an explicit nil
### GetProtectionSourceId

`func (o *UnprotectObjectParams) GetProtectionSourceId() int64`

GetProtectionSourceId returns the ProtectionSourceId field if non-nil, zero value otherwise.

### GetProtectionSourceIdOk

`func (o *UnprotectObjectParams) GetProtectionSourceIdOk() (*int64, bool)`

GetProtectionSourceIdOk returns a tuple with the ProtectionSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceId

`func (o *UnprotectObjectParams) SetProtectionSourceId(v int64)`

SetProtectionSourceId sets ProtectionSourceId field to given value.


### SetProtectionSourceIdNil

`func (o *UnprotectObjectParams) SetProtectionSourceIdNil(b bool)`

 SetProtectionSourceIdNil sets the value for ProtectionSourceId to be an explicit nil

### UnsetProtectionSourceId
`func (o *UnprotectObjectParams) UnsetProtectionSourceId()`

UnsetProtectionSourceId ensures that no value is present for ProtectionSourceId, not even an explicit nil
### GetRpoPolicyId

`func (o *UnprotectObjectParams) GetRpoPolicyId() string`

GetRpoPolicyId returns the RpoPolicyId field if non-nil, zero value otherwise.

### GetRpoPolicyIdOk

`func (o *UnprotectObjectParams) GetRpoPolicyIdOk() (*string, bool)`

GetRpoPolicyIdOk returns a tuple with the RpoPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoPolicyId

`func (o *UnprotectObjectParams) SetRpoPolicyId(v string)`

SetRpoPolicyId sets RpoPolicyId field to given value.


### SetRpoPolicyIdNil

`func (o *UnprotectObjectParams) SetRpoPolicyIdNil(b bool)`

 SetRpoPolicyIdNil sets the value for RpoPolicyId to be an explicit nil

### UnsetRpoPolicyId
`func (o *UnprotectObjectParams) UnsetRpoPolicyId()`

UnsetRpoPolicyId ensures that no value is present for RpoPolicyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


