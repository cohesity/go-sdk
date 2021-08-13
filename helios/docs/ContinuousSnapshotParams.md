# ContinuousSnapshotParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **NullableBool** | Specifies whether source snapshots should be taken even if there is a pending run. | 
**MaxAllowedSnapshots** | Pointer to **NullableInt32** | Specifies the maximum number of source snapshots allowed for a given object in a protection group. This is only applicable if isContinuousSnapshottingEnabled is set to true. | [optional] 

## Methods

### NewContinuousSnapshotParams

`func NewContinuousSnapshotParams(isEnabled NullableBool, ) *ContinuousSnapshotParams`

NewContinuousSnapshotParams instantiates a new ContinuousSnapshotParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuousSnapshotParamsWithDefaults

`func NewContinuousSnapshotParamsWithDefaults() *ContinuousSnapshotParams`

NewContinuousSnapshotParamsWithDefaults instantiates a new ContinuousSnapshotParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ContinuousSnapshotParams) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ContinuousSnapshotParams) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ContinuousSnapshotParams) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### SetIsEnabledNil

`func (o *ContinuousSnapshotParams) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *ContinuousSnapshotParams) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetMaxAllowedSnapshots

`func (o *ContinuousSnapshotParams) GetMaxAllowedSnapshots() int32`

GetMaxAllowedSnapshots returns the MaxAllowedSnapshots field if non-nil, zero value otherwise.

### GetMaxAllowedSnapshotsOk

`func (o *ContinuousSnapshotParams) GetMaxAllowedSnapshotsOk() (*int32, bool)`

GetMaxAllowedSnapshotsOk returns a tuple with the MaxAllowedSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedSnapshots

`func (o *ContinuousSnapshotParams) SetMaxAllowedSnapshots(v int32)`

SetMaxAllowedSnapshots sets MaxAllowedSnapshots field to given value.

### HasMaxAllowedSnapshots

`func (o *ContinuousSnapshotParams) HasMaxAllowedSnapshots() bool`

HasMaxAllowedSnapshots returns a boolean if a field has been set.

### SetMaxAllowedSnapshotsNil

`func (o *ContinuousSnapshotParams) SetMaxAllowedSnapshotsNil(b bool)`

 SetMaxAllowedSnapshotsNil sets the value for MaxAllowedSnapshots to be an explicit nil

### UnsetMaxAllowedSnapshots
`func (o *ContinuousSnapshotParams) UnsetMaxAllowedSnapshots()`

UnsetMaxAllowedSnapshots ensures that no value is present for MaxAllowedSnapshots, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


