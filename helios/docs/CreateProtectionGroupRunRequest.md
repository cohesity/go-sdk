# CreateProtectionGroupRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunType** | **NullableString** | Type of protection run. &#39;kRegular&#39; indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. &#39;kFull&#39; indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. &#39;kLog&#39; indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. &#39;kSystem&#39; indicates system volume backup. It produces an image for bare metal recovery. | 
**Objects** | Pointer to [**[]RunObject**](RunObject.md) | Specifies the list of objects to be protected by this Protection Group run. These can be leaf objects or non-leaf objects in the protection hierarchy. This must be specified only if a subset of objects from the Protection Groups needs to be protected. | [optional] 
**TargetsConfig** | Pointer to [**RunTargetsConfiguration**](RunTargetsConfiguration.md) |  | [optional] 

## Methods

### NewCreateProtectionGroupRunRequest

`func NewCreateProtectionGroupRunRequest(runType NullableString, ) *CreateProtectionGroupRunRequest`

NewCreateProtectionGroupRunRequest instantiates a new CreateProtectionGroupRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProtectionGroupRunRequestWithDefaults

`func NewCreateProtectionGroupRunRequestWithDefaults() *CreateProtectionGroupRunRequest`

NewCreateProtectionGroupRunRequestWithDefaults instantiates a new CreateProtectionGroupRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunType

`func (o *CreateProtectionGroupRunRequest) GetRunType() string`

GetRunType returns the RunType field if non-nil, zero value otherwise.

### GetRunTypeOk

`func (o *CreateProtectionGroupRunRequest) GetRunTypeOk() (*string, bool)`

GetRunTypeOk returns a tuple with the RunType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunType

`func (o *CreateProtectionGroupRunRequest) SetRunType(v string)`

SetRunType sets RunType field to given value.


### SetRunTypeNil

`func (o *CreateProtectionGroupRunRequest) SetRunTypeNil(b bool)`

 SetRunTypeNil sets the value for RunType to be an explicit nil

### UnsetRunType
`func (o *CreateProtectionGroupRunRequest) UnsetRunType()`

UnsetRunType ensures that no value is present for RunType, not even an explicit nil
### GetObjects

`func (o *CreateProtectionGroupRunRequest) GetObjects() []RunObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CreateProtectionGroupRunRequest) GetObjectsOk() (*[]RunObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CreateProtectionGroupRunRequest) SetObjects(v []RunObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *CreateProtectionGroupRunRequest) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetTargetsConfig

`func (o *CreateProtectionGroupRunRequest) GetTargetsConfig() RunTargetsConfiguration`

GetTargetsConfig returns the TargetsConfig field if non-nil, zero value otherwise.

### GetTargetsConfigOk

`func (o *CreateProtectionGroupRunRequest) GetTargetsConfigOk() (*RunTargetsConfiguration, bool)`

GetTargetsConfigOk returns a tuple with the TargetsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsConfig

`func (o *CreateProtectionGroupRunRequest) SetTargetsConfig(v RunTargetsConfiguration)`

SetTargetsConfig sets TargetsConfig field to given value.

### HasTargetsConfig

`func (o *CreateProtectionGroupRunRequest) HasTargetsConfig() bool`

HasTargetsConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


