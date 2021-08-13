# PureProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]PureProtectionGroupObjectParams**](PureProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | 
**MaxSnapshotsOnPrimary** | Pointer to **NullableInt64** | Specifies the number of snapshots to retain on the primary environment. If not specified, then snapshots will not be deleted from the primary environment. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 
**PrePostScript** | Pointer to [**HostBasedBackupScriptParams**](HostBasedBackupScriptParams.md) |  | [optional] 

## Methods

### NewPureProtectionGroupParams

`func NewPureProtectionGroupParams(objects []PureProtectionGroupObjectParams, ) *PureProtectionGroupParams`

NewPureProtectionGroupParams instantiates a new PureProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPureProtectionGroupParamsWithDefaults

`func NewPureProtectionGroupParamsWithDefaults() *PureProtectionGroupParams`

NewPureProtectionGroupParamsWithDefaults instantiates a new PureProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *PureProtectionGroupParams) GetObjects() []PureProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *PureProtectionGroupParams) GetObjectsOk() (*[]PureProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *PureProtectionGroupParams) SetObjects(v []PureProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetMaxSnapshotsOnPrimary

`func (o *PureProtectionGroupParams) GetMaxSnapshotsOnPrimary() int64`

GetMaxSnapshotsOnPrimary returns the MaxSnapshotsOnPrimary field if non-nil, zero value otherwise.

### GetMaxSnapshotsOnPrimaryOk

`func (o *PureProtectionGroupParams) GetMaxSnapshotsOnPrimaryOk() (*int64, bool)`

GetMaxSnapshotsOnPrimaryOk returns a tuple with the MaxSnapshotsOnPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSnapshotsOnPrimary

`func (o *PureProtectionGroupParams) SetMaxSnapshotsOnPrimary(v int64)`

SetMaxSnapshotsOnPrimary sets MaxSnapshotsOnPrimary field to given value.

### HasMaxSnapshotsOnPrimary

`func (o *PureProtectionGroupParams) HasMaxSnapshotsOnPrimary() bool`

HasMaxSnapshotsOnPrimary returns a boolean if a field has been set.

### SetMaxSnapshotsOnPrimaryNil

`func (o *PureProtectionGroupParams) SetMaxSnapshotsOnPrimaryNil(b bool)`

 SetMaxSnapshotsOnPrimaryNil sets the value for MaxSnapshotsOnPrimary to be an explicit nil

### UnsetMaxSnapshotsOnPrimary
`func (o *PureProtectionGroupParams) UnsetMaxSnapshotsOnPrimary()`

UnsetMaxSnapshotsOnPrimary ensures that no value is present for MaxSnapshotsOnPrimary, not even an explicit nil
### GetSourceId

`func (o *PureProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PureProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PureProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *PureProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *PureProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *PureProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *PureProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *PureProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *PureProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *PureProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *PureProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *PureProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetPrePostScript

`func (o *PureProtectionGroupParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *PureProtectionGroupParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *PureProtectionGroupParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *PureProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


