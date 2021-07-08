# BackupJobProtoBackupSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]EntityProto**](EntityProto.md) | Source entities. NOTE: Multiple sources can be specified here for non-leaf-level entities in the hierarchy. The sources obtained after expanding these will be intersected among each other to form the final set of sources. e.g. this can be used to backup only those VMs that have both the tags &#39;SQL&#39; and &#39;3hrs&#39;. | [optional] 

## Methods

### NewBackupJobProtoBackupSource

`func NewBackupJobProtoBackupSource() *BackupJobProtoBackupSource`

NewBackupJobProtoBackupSource instantiates a new BackupJobProtoBackupSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobProtoBackupSourceWithDefaults

`func NewBackupJobProtoBackupSourceWithDefaults() *BackupJobProtoBackupSource`

NewBackupJobProtoBackupSourceWithDefaults instantiates a new BackupJobProtoBackupSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *BackupJobProtoBackupSource) GetEntities() []EntityProto`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *BackupJobProtoBackupSource) GetEntitiesOk() (*[]EntityProto, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *BackupJobProtoBackupSource) SetEntities(v []EntityProto)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *BackupJobProtoBackupSource) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### SetEntitiesNil

`func (o *BackupJobProtoBackupSource) SetEntitiesNil(b bool)`

 SetEntitiesNil sets the value for Entities to be an explicit nil

### UnsetEntities
`func (o *BackupJobProtoBackupSource) UnsetEntities()`

UnsetEntities ensures that no value is present for Entities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


