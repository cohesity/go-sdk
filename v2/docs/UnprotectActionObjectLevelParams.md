# UnprotectActionObjectLevelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 
**DeleteAllSnapshots** | Pointer to **NullableBool** | Specifies whether to delete all snapshots along with unprotecting object protection. If set to true, all snapshots will be deleted and no more recoverable. | [optional] 
**ForceUnprotect** | Pointer to **NullableBool** | If specified as true then it will cancel the ongoing runs and immediatly unprotect. | [optional] 

## Methods

### NewUnprotectActionObjectLevelParams

`func NewUnprotectActionObjectLevelParams(id NullableInt64, ) *UnprotectActionObjectLevelParams`

NewUnprotectActionObjectLevelParams instantiates a new UnprotectActionObjectLevelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnprotectActionObjectLevelParamsWithDefaults

`func NewUnprotectActionObjectLevelParamsWithDefaults() *UnprotectActionObjectLevelParams`

NewUnprotectActionObjectLevelParamsWithDefaults instantiates a new UnprotectActionObjectLevelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnprotectActionObjectLevelParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnprotectActionObjectLevelParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnprotectActionObjectLevelParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *UnprotectActionObjectLevelParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UnprotectActionObjectLevelParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *UnprotectActionObjectLevelParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnprotectActionObjectLevelParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnprotectActionObjectLevelParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnprotectActionObjectLevelParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UnprotectActionObjectLevelParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UnprotectActionObjectLevelParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeleteAllSnapshots

`func (o *UnprotectActionObjectLevelParams) GetDeleteAllSnapshots() bool`

GetDeleteAllSnapshots returns the DeleteAllSnapshots field if non-nil, zero value otherwise.

### GetDeleteAllSnapshotsOk

`func (o *UnprotectActionObjectLevelParams) GetDeleteAllSnapshotsOk() (*bool, bool)`

GetDeleteAllSnapshotsOk returns a tuple with the DeleteAllSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAllSnapshots

`func (o *UnprotectActionObjectLevelParams) SetDeleteAllSnapshots(v bool)`

SetDeleteAllSnapshots sets DeleteAllSnapshots field to given value.

### HasDeleteAllSnapshots

`func (o *UnprotectActionObjectLevelParams) HasDeleteAllSnapshots() bool`

HasDeleteAllSnapshots returns a boolean if a field has been set.

### SetDeleteAllSnapshotsNil

`func (o *UnprotectActionObjectLevelParams) SetDeleteAllSnapshotsNil(b bool)`

 SetDeleteAllSnapshotsNil sets the value for DeleteAllSnapshots to be an explicit nil

### UnsetDeleteAllSnapshots
`func (o *UnprotectActionObjectLevelParams) UnsetDeleteAllSnapshots()`

UnsetDeleteAllSnapshots ensures that no value is present for DeleteAllSnapshots, not even an explicit nil
### GetForceUnprotect

`func (o *UnprotectActionObjectLevelParams) GetForceUnprotect() bool`

GetForceUnprotect returns the ForceUnprotect field if non-nil, zero value otherwise.

### GetForceUnprotectOk

`func (o *UnprotectActionObjectLevelParams) GetForceUnprotectOk() (*bool, bool)`

GetForceUnprotectOk returns a tuple with the ForceUnprotect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnprotect

`func (o *UnprotectActionObjectLevelParams) SetForceUnprotect(v bool)`

SetForceUnprotect sets ForceUnprotect field to given value.

### HasForceUnprotect

`func (o *UnprotectActionObjectLevelParams) HasForceUnprotect() bool`

HasForceUnprotect returns a boolean if a field has been set.

### SetForceUnprotectNil

`func (o *UnprotectActionObjectLevelParams) SetForceUnprotectNil(b bool)`

 SetForceUnprotectNil sets the value for ForceUnprotect to be an explicit nil

### UnsetForceUnprotect
`func (o *UnprotectActionObjectLevelParams) UnsetForceUnprotect()`

UnsetForceUnprotect ensures that no value is present for ForceUnprotect, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


