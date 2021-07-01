# FullSnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreInfo** | Pointer to [**RestoreInfo**](RestoreInfo.md) |  | [optional] 
**SnapshotTarget** | Pointer to [**[]SnapshotTargetSettings**](SnapshotTargetSettings.md) | Specifies the location holding snapshot copies that may be used for restore. | [optional] 

## Methods

### NewFullSnapshotInfo

`func NewFullSnapshotInfo() *FullSnapshotInfo`

NewFullSnapshotInfo instantiates a new FullSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullSnapshotInfoWithDefaults

`func NewFullSnapshotInfoWithDefaults() *FullSnapshotInfo`

NewFullSnapshotInfoWithDefaults instantiates a new FullSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreInfo

`func (o *FullSnapshotInfo) GetRestoreInfo() RestoreInfo`

GetRestoreInfo returns the RestoreInfo field if non-nil, zero value otherwise.

### GetRestoreInfoOk

`func (o *FullSnapshotInfo) GetRestoreInfoOk() (*RestoreInfo, bool)`

GetRestoreInfoOk returns a tuple with the RestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreInfo

`func (o *FullSnapshotInfo) SetRestoreInfo(v RestoreInfo)`

SetRestoreInfo sets RestoreInfo field to given value.

### HasRestoreInfo

`func (o *FullSnapshotInfo) HasRestoreInfo() bool`

HasRestoreInfo returns a boolean if a field has been set.

### GetSnapshotTarget

`func (o *FullSnapshotInfo) GetSnapshotTarget() []SnapshotTargetSettings`

GetSnapshotTarget returns the SnapshotTarget field if non-nil, zero value otherwise.

### GetSnapshotTargetOk

`func (o *FullSnapshotInfo) GetSnapshotTargetOk() (*[]SnapshotTargetSettings, bool)`

GetSnapshotTargetOk returns a tuple with the SnapshotTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTarget

`func (o *FullSnapshotInfo) SetSnapshotTarget(v []SnapshotTargetSettings)`

SetSnapshotTarget sets SnapshotTarget field to given value.

### HasSnapshotTarget

`func (o *FullSnapshotInfo) HasSnapshotTarget() bool`

HasSnapshotTarget returns a boolean if a field has been set.

### SetSnapshotTargetNil

`func (o *FullSnapshotInfo) SetSnapshotTargetNil(b bool)`

 SetSnapshotTargetNil sets the value for SnapshotTarget to be an explicit nil

### UnsetSnapshotTarget
`func (o *FullSnapshotInfo) UnsetSnapshotTarget()`

UnsetSnapshotTarget ensures that no value is present for SnapshotTarget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


