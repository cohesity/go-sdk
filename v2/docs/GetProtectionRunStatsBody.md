# GetProtectionRunStatsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalRun** | Pointer to [**[]ArchivalTargetStatsInfo**](ArchivalTargetStatsInfo.md) | Stats for the archival run. | [optional] 
**LocalRun** | Pointer to [**BackupRunStatsInfo**](BackupRunStatsInfo.md) |  | [optional] 

## Methods

### NewGetProtectionRunStatsBody

`func NewGetProtectionRunStatsBody() *GetProtectionRunStatsBody`

NewGetProtectionRunStatsBody instantiates a new GetProtectionRunStatsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProtectionRunStatsBodyWithDefaults

`func NewGetProtectionRunStatsBodyWithDefaults() *GetProtectionRunStatsBody`

NewGetProtectionRunStatsBodyWithDefaults instantiates a new GetProtectionRunStatsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalRun

`func (o *GetProtectionRunStatsBody) GetArchivalRun() []ArchivalTargetStatsInfo`

GetArchivalRun returns the ArchivalRun field if non-nil, zero value otherwise.

### GetArchivalRunOk

`func (o *GetProtectionRunStatsBody) GetArchivalRunOk() (*[]ArchivalTargetStatsInfo, bool)`

GetArchivalRunOk returns a tuple with the ArchivalRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalRun

`func (o *GetProtectionRunStatsBody) SetArchivalRun(v []ArchivalTargetStatsInfo)`

SetArchivalRun sets ArchivalRun field to given value.

### HasArchivalRun

`func (o *GetProtectionRunStatsBody) HasArchivalRun() bool`

HasArchivalRun returns a boolean if a field has been set.

### SetArchivalRunNil

`func (o *GetProtectionRunStatsBody) SetArchivalRunNil(b bool)`

 SetArchivalRunNil sets the value for ArchivalRun to be an explicit nil

### UnsetArchivalRun
`func (o *GetProtectionRunStatsBody) UnsetArchivalRun()`

UnsetArchivalRun ensures that no value is present for ArchivalRun, not even an explicit nil
### GetLocalRun

`func (o *GetProtectionRunStatsBody) GetLocalRun() BackupRunStatsInfo`

GetLocalRun returns the LocalRun field if non-nil, zero value otherwise.

### GetLocalRunOk

`func (o *GetProtectionRunStatsBody) GetLocalRunOk() (*BackupRunStatsInfo, bool)`

GetLocalRunOk returns a tuple with the LocalRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalRun

`func (o *GetProtectionRunStatsBody) SetLocalRun(v BackupRunStatsInfo)`

SetLocalRun sets LocalRun field to given value.

### HasLocalRun

`func (o *GetProtectionRunStatsBody) HasLocalRun() bool`

HasLocalRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


