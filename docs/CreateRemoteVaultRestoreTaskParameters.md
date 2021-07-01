# CreateRemoteVaultRestoreTaskParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlacierRetrievalType** | Pointer to **NullableString** | Specifies the way data needs to be retrieved from the external target. This information will be filled in by Iris and Magneto will pass it along to the Icebox as it is to support bulk retrieval from Glacier. Specifies the type of Restore Task.  &#39;kStandard&#39; specifies retrievals that allow to access any of your archives within several hours. Standard retrievals typically complete within 3–5 hours. This is the default option for retrieval requests that do not specify the retrieval option. &#39;kBulk&#39; specifies retrievals that are Glacier’s lowest-cost retrieval option, which can be used to retrieve large amounts, even petabytes, of data inexpensively in a day. Bulk retrieval typically complete within 5–12 hours. &#39;kExpedited&#39; specifies retrievals that allows to quickly access your data when occasional urgent requests for a subset of archives are required. For all but the largest archives (250 MB+), data accessed using Expedited retrievals are typically made available within 1–5 minutes. | [optional] 
**RestoreObjects** | Pointer to [**[]IndexAndSnapshots**](IndexAndSnapshots.md) | Array of Restore Objects.  Specifies the list of Snapshots and the index to be restored from the remote Vault. The data on the remote Vault may have been originally archived from multiple remote Clusters. | [optional] 
**SearchJobUid** | [**NullableUniversalId**](UniversalId.md) | Specifies the unique id of the remote Vault search Job. | 
**TaskName** | **NullableString** | Specifies a name of the restore task. | 
**VaultId** | **NullableInt64** | Specifies the id of the Vault that contains the index and Snapshots to restore to the current Cluster. This is the id assigned by the Cohesity Cluster when Vault was registered as an External Target. | 

## Methods

### NewCreateRemoteVaultRestoreTaskParameters

`func NewCreateRemoteVaultRestoreTaskParameters(searchJobUid NullableUniversalId, taskName NullableString, vaultId NullableInt64, ) *CreateRemoteVaultRestoreTaskParameters`

NewCreateRemoteVaultRestoreTaskParameters instantiates a new CreateRemoteVaultRestoreTaskParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRemoteVaultRestoreTaskParametersWithDefaults

`func NewCreateRemoteVaultRestoreTaskParametersWithDefaults() *CreateRemoteVaultRestoreTaskParameters`

NewCreateRemoteVaultRestoreTaskParametersWithDefaults instantiates a new CreateRemoteVaultRestoreTaskParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlacierRetrievalType

`func (o *CreateRemoteVaultRestoreTaskParameters) GetGlacierRetrievalType() string`

GetGlacierRetrievalType returns the GlacierRetrievalType field if non-nil, zero value otherwise.

### GetGlacierRetrievalTypeOk

`func (o *CreateRemoteVaultRestoreTaskParameters) GetGlacierRetrievalTypeOk() (*string, bool)`

GetGlacierRetrievalTypeOk returns a tuple with the GlacierRetrievalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacierRetrievalType

`func (o *CreateRemoteVaultRestoreTaskParameters) SetGlacierRetrievalType(v string)`

SetGlacierRetrievalType sets GlacierRetrievalType field to given value.

### HasGlacierRetrievalType

`func (o *CreateRemoteVaultRestoreTaskParameters) HasGlacierRetrievalType() bool`

HasGlacierRetrievalType returns a boolean if a field has been set.

### SetGlacierRetrievalTypeNil

`func (o *CreateRemoteVaultRestoreTaskParameters) SetGlacierRetrievalTypeNil(b bool)`

 SetGlacierRetrievalTypeNil sets the value for GlacierRetrievalType to be an explicit nil

### UnsetGlacierRetrievalType
`func (o *CreateRemoteVaultRestoreTaskParameters) UnsetGlacierRetrievalType()`

UnsetGlacierRetrievalType ensures that no value is present for GlacierRetrievalType, not even an explicit nil
### GetRestoreObjects

`func (o *CreateRemoteVaultRestoreTaskParameters) GetRestoreObjects() []IndexAndSnapshots`

GetRestoreObjects returns the RestoreObjects field if non-nil, zero value otherwise.

### GetRestoreObjectsOk

`func (o *CreateRemoteVaultRestoreTaskParameters) GetRestoreObjectsOk() (*[]IndexAndSnapshots, bool)`

GetRestoreObjectsOk returns a tuple with the RestoreObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreObjects

`func (o *CreateRemoteVaultRestoreTaskParameters) SetRestoreObjects(v []IndexAndSnapshots)`

SetRestoreObjects sets RestoreObjects field to given value.

### HasRestoreObjects

`func (o *CreateRemoteVaultRestoreTaskParameters) HasRestoreObjects() bool`

HasRestoreObjects returns a boolean if a field has been set.

### SetRestoreObjectsNil

`func (o *CreateRemoteVaultRestoreTaskParameters) SetRestoreObjectsNil(b bool)`

 SetRestoreObjectsNil sets the value for RestoreObjects to be an explicit nil

### UnsetRestoreObjects
`func (o *CreateRemoteVaultRestoreTaskParameters) UnsetRestoreObjects()`

UnsetRestoreObjects ensures that no value is present for RestoreObjects, not even an explicit nil
### GetSearchJobUid

`func (o *CreateRemoteVaultRestoreTaskParameters) GetSearchJobUid() UniversalId`

GetSearchJobUid returns the SearchJobUid field if non-nil, zero value otherwise.

### GetSearchJobUidOk

`func (o *CreateRemoteVaultRestoreTaskParameters) GetSearchJobUidOk() (*UniversalId, bool)`

GetSearchJobUidOk returns a tuple with the SearchJobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchJobUid

`func (o *CreateRemoteVaultRestoreTaskParameters) SetSearchJobUid(v UniversalId)`

SetSearchJobUid sets SearchJobUid field to given value.


### SetSearchJobUidNil

`func (o *CreateRemoteVaultRestoreTaskParameters) SetSearchJobUidNil(b bool)`

 SetSearchJobUidNil sets the value for SearchJobUid to be an explicit nil

### UnsetSearchJobUid
`func (o *CreateRemoteVaultRestoreTaskParameters) UnsetSearchJobUid()`

UnsetSearchJobUid ensures that no value is present for SearchJobUid, not even an explicit nil
### GetTaskName

`func (o *CreateRemoteVaultRestoreTaskParameters) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *CreateRemoteVaultRestoreTaskParameters) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *CreateRemoteVaultRestoreTaskParameters) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### SetTaskNameNil

`func (o *CreateRemoteVaultRestoreTaskParameters) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *CreateRemoteVaultRestoreTaskParameters) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
### GetVaultId

`func (o *CreateRemoteVaultRestoreTaskParameters) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CreateRemoteVaultRestoreTaskParameters) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CreateRemoteVaultRestoreTaskParameters) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.


### SetVaultIdNil

`func (o *CreateRemoteVaultRestoreTaskParameters) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *CreateRemoteVaultRestoreTaskParameters) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


