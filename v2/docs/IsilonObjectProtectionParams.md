# IsilonObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousSnapshots** | Pointer to [**ContinuousSnapshotParams**](ContinuousSnapshotParams.md) |  | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the protocol of the NAS device being backed up. | [optional] 
**UseChangelist** | Pointer to **NullableBool** | Specify whether to use the Isilon Changelist API to directly discover changed files/directories for faster incremental backup. Cohesity will keep an extra snapshot which will be deleted by the next successful backup. | [optional] 

## Methods

### NewIsilonObjectProtectionParams

`func NewIsilonObjectProtectionParams() *IsilonObjectProtectionParams`

NewIsilonObjectProtectionParams instantiates a new IsilonObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonObjectProtectionParamsWithDefaults

`func NewIsilonObjectProtectionParamsWithDefaults() *IsilonObjectProtectionParams`

NewIsilonObjectProtectionParamsWithDefaults instantiates a new IsilonObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuousSnapshots

`func (o *IsilonObjectProtectionParams) GetContinuousSnapshots() ContinuousSnapshotParams`

GetContinuousSnapshots returns the ContinuousSnapshots field if non-nil, zero value otherwise.

### GetContinuousSnapshotsOk

`func (o *IsilonObjectProtectionParams) GetContinuousSnapshotsOk() (*ContinuousSnapshotParams, bool)`

GetContinuousSnapshotsOk returns a tuple with the ContinuousSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousSnapshots

`func (o *IsilonObjectProtectionParams) SetContinuousSnapshots(v ContinuousSnapshotParams)`

SetContinuousSnapshots sets ContinuousSnapshots field to given value.

### HasContinuousSnapshots

`func (o *IsilonObjectProtectionParams) HasContinuousSnapshots() bool`

HasContinuousSnapshots returns a boolean if a field has been set.

### GetProtocol

`func (o *IsilonObjectProtectionParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IsilonObjectProtectionParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IsilonObjectProtectionParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IsilonObjectProtectionParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *IsilonObjectProtectionParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *IsilonObjectProtectionParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetUseChangelist

`func (o *IsilonObjectProtectionParams) GetUseChangelist() bool`

GetUseChangelist returns the UseChangelist field if non-nil, zero value otherwise.

### GetUseChangelistOk

`func (o *IsilonObjectProtectionParams) GetUseChangelistOk() (*bool, bool)`

GetUseChangelistOk returns a tuple with the UseChangelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseChangelist

`func (o *IsilonObjectProtectionParams) SetUseChangelist(v bool)`

SetUseChangelist sets UseChangelist field to given value.

### HasUseChangelist

`func (o *IsilonObjectProtectionParams) HasUseChangelist() bool`

HasUseChangelist returns a boolean if a field has been set.

### SetUseChangelistNil

`func (o *IsilonObjectProtectionParams) SetUseChangelistNil(b bool)`

 SetUseChangelistNil sets the value for UseChangelist to be an explicit nil

### UnsetUseChangelist
`func (o *IsilonObjectProtectionParams) UnsetUseChangelist()`

UnsetUseChangelist ensures that no value is present for UseChangelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


