# CdpLocalBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeInUsecs** | Pointer to **NullableInt64** | Specifies the end time of the last local backup taken. | [optional] 
**StartTimeInUsecs** | Pointer to **NullableInt64** | Specifies the start time of the last local backup taken. | [optional] 

## Methods

### NewCdpLocalBackupInfo

`func NewCdpLocalBackupInfo() *CdpLocalBackupInfo`

NewCdpLocalBackupInfo instantiates a new CdpLocalBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdpLocalBackupInfoWithDefaults

`func NewCdpLocalBackupInfoWithDefaults() *CdpLocalBackupInfo`

NewCdpLocalBackupInfoWithDefaults instantiates a new CdpLocalBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeInUsecs

`func (o *CdpLocalBackupInfo) GetEndTimeInUsecs() int64`

GetEndTimeInUsecs returns the EndTimeInUsecs field if non-nil, zero value otherwise.

### GetEndTimeInUsecsOk

`func (o *CdpLocalBackupInfo) GetEndTimeInUsecsOk() (*int64, bool)`

GetEndTimeInUsecsOk returns a tuple with the EndTimeInUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeInUsecs

`func (o *CdpLocalBackupInfo) SetEndTimeInUsecs(v int64)`

SetEndTimeInUsecs sets EndTimeInUsecs field to given value.

### HasEndTimeInUsecs

`func (o *CdpLocalBackupInfo) HasEndTimeInUsecs() bool`

HasEndTimeInUsecs returns a boolean if a field has been set.

### SetEndTimeInUsecsNil

`func (o *CdpLocalBackupInfo) SetEndTimeInUsecsNil(b bool)`

 SetEndTimeInUsecsNil sets the value for EndTimeInUsecs to be an explicit nil

### UnsetEndTimeInUsecs
`func (o *CdpLocalBackupInfo) UnsetEndTimeInUsecs()`

UnsetEndTimeInUsecs ensures that no value is present for EndTimeInUsecs, not even an explicit nil
### GetStartTimeInUsecs

`func (o *CdpLocalBackupInfo) GetStartTimeInUsecs() int64`

GetStartTimeInUsecs returns the StartTimeInUsecs field if non-nil, zero value otherwise.

### GetStartTimeInUsecsOk

`func (o *CdpLocalBackupInfo) GetStartTimeInUsecsOk() (*int64, bool)`

GetStartTimeInUsecsOk returns a tuple with the StartTimeInUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeInUsecs

`func (o *CdpLocalBackupInfo) SetStartTimeInUsecs(v int64)`

SetStartTimeInUsecs sets StartTimeInUsecs field to given value.

### HasStartTimeInUsecs

`func (o *CdpLocalBackupInfo) HasStartTimeInUsecs() bool`

HasStartTimeInUsecs returns a boolean if a field has been set.

### SetStartTimeInUsecsNil

`func (o *CdpLocalBackupInfo) SetStartTimeInUsecsNil(b bool)`

 SetStartTimeInUsecsNil sets the value for StartTimeInUsecs to be an explicit nil

### UnsetStartTimeInUsecs
`func (o *CdpLocalBackupInfo) UnsetStartTimeInUsecs()`

UnsetStartTimeInUsecs ensures that no value is present for StartTimeInUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


