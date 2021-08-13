# UnprotectActionObjectLevelParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAllSnapshots** | Pointer to **NullableBool** | Specifies whether to delete all snapshots along with unprotecting object protection. If set to true, all snapshots will be deleted and no more recoverable. | [optional] 
**ForceUnprotect** | Pointer to **NullableBool** | If specified as true then it will cancel the ongoing runs and immediatly unprotect. | [optional] 

## Methods

### NewUnprotectActionObjectLevelParamsAllOf

`func NewUnprotectActionObjectLevelParamsAllOf() *UnprotectActionObjectLevelParamsAllOf`

NewUnprotectActionObjectLevelParamsAllOf instantiates a new UnprotectActionObjectLevelParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnprotectActionObjectLevelParamsAllOfWithDefaults

`func NewUnprotectActionObjectLevelParamsAllOfWithDefaults() *UnprotectActionObjectLevelParamsAllOf`

NewUnprotectActionObjectLevelParamsAllOfWithDefaults instantiates a new UnprotectActionObjectLevelParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAllSnapshots

`func (o *UnprotectActionObjectLevelParamsAllOf) GetDeleteAllSnapshots() bool`

GetDeleteAllSnapshots returns the DeleteAllSnapshots field if non-nil, zero value otherwise.

### GetDeleteAllSnapshotsOk

`func (o *UnprotectActionObjectLevelParamsAllOf) GetDeleteAllSnapshotsOk() (*bool, bool)`

GetDeleteAllSnapshotsOk returns a tuple with the DeleteAllSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllSnapshots

`func (o *UnprotectActionObjectLevelParamsAllOf) SetDeleteAllSnapshots(v bool)`

SetDeleteAllSnapshots sets DeleteAllSnapshots field to given value.

### HasDeleteAllSnapshots

`func (o *UnprotectActionObjectLevelParamsAllOf) HasDeleteAllSnapshots() bool`

HasDeleteAllSnapshots returns a boolean if a field has been set.

### SetDeleteAllSnapshotsNil

`func (o *UnprotectActionObjectLevelParamsAllOf) SetDeleteAllSnapshotsNil(b bool)`

 SetDeleteAllSnapshotsNil sets the value for DeleteAllSnapshots to be an explicit nil

### UnsetDeleteAllSnapshots
`func (o *UnprotectActionObjectLevelParamsAllOf) UnsetDeleteAllSnapshots()`

UnsetDeleteAllSnapshots ensures that no value is present for DeleteAllSnapshots, not even an explicit nil
### GetForceUnprotect

`func (o *UnprotectActionObjectLevelParamsAllOf) GetForceUnprotect() bool`

GetForceUnprotect returns the ForceUnprotect field if non-nil, zero value otherwise.

### GetForceUnprotectOk

`func (o *UnprotectActionObjectLevelParamsAllOf) GetForceUnprotectOk() (*bool, bool)`

GetForceUnprotectOk returns a tuple with the ForceUnprotect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnprotect

`func (o *UnprotectActionObjectLevelParamsAllOf) SetForceUnprotect(v bool)`

SetForceUnprotect sets ForceUnprotect field to given value.

### HasForceUnprotect

`func (o *UnprotectActionObjectLevelParamsAllOf) HasForceUnprotect() bool`

HasForceUnprotect returns a boolean if a field has been set.

### SetForceUnprotectNil

`func (o *UnprotectActionObjectLevelParamsAllOf) SetForceUnprotectNil(b bool)`

 SetForceUnprotectNil sets the value for ForceUnprotect to be an explicit nil

### UnsetForceUnprotect
`func (o *UnprotectActionObjectLevelParamsAllOf) UnsetForceUnprotect()`

UnsetForceUnprotect ensures that no value is present for ForceUnprotect, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


