# CreateRemoteVaultSearchJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterMatchString** | Pointer to **NullableString** | Filter by specifying a Cluster name prefix string. Only Clusters with names that start with this prefix string are returned in the search result. If not set, all Clusters archiving data to the Vault are returned in the search results. | [optional] 
**EncryptionKeys** | Pointer to [**[]VaultEncryptionKey**](VaultEncryptionKey.md) | Array of Encryption Keys.  Specifies an optional list of encryption keys that may be needed to search and restore data that was archived to a remote Vault. Archived data cannot be searched or restored without the corresponding encryption key used by the original Cluster to archive the data. Encryption keys can be uploaded using the REST API public/remoteVaults/encryptionKeys operation. If the key is already uploaded, this field does not need to be specified. | [optional] 
**EndTimeUsecs** | Pointer to **NullableInt64** | Filter by a end time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that completed before the specified end time are included in the search results. | [optional] 
**JobMatchString** | Pointer to **NullableString** | Filter by specifying a Protection Job name prefix string. Only Protection Jobs with names that start with this prefix string are returned in the search result. If not set, all Protection Jobs archiving data to the Vault are returned in the search results. | [optional] 
**SearchJobName** | **NullableString** | Specifies the search Job name. | 
**StartTimeUsecs** | Pointer to **NullableInt64** | Filter by a start time specified as a Unix epoch Timestamp (in microseconds). Only Job Runs that started after the specified time are included in the search results. | [optional] 
**VaultId** | **NullableInt64** | Specifies the id of the Vault to search. This id was assigned by the local Cohesity Cluster when Vault was registered as an External Target. | 

## Methods

### NewCreateRemoteVaultSearchJobParameters

`func NewCreateRemoteVaultSearchJobParameters(searchJobName NullableString, vaultId NullableInt64, ) *CreateRemoteVaultSearchJobParameters`

NewCreateRemoteVaultSearchJobParameters instantiates a new CreateRemoteVaultSearchJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRemoteVaultSearchJobParametersWithDefaults

`func NewCreateRemoteVaultSearchJobParametersWithDefaults() *CreateRemoteVaultSearchJobParameters`

NewCreateRemoteVaultSearchJobParametersWithDefaults instantiates a new CreateRemoteVaultSearchJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterMatchString

`func (o *CreateRemoteVaultSearchJobParameters) GetClusterMatchString() string`

GetClusterMatchString returns the ClusterMatchString field if non-nil, zero value otherwise.

### GetClusterMatchStringOk

`func (o *CreateRemoteVaultSearchJobParameters) GetClusterMatchStringOk() (*string, bool)`

GetClusterMatchStringOk returns a tuple with the ClusterMatchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMatchString

`func (o *CreateRemoteVaultSearchJobParameters) SetClusterMatchString(v string)`

SetClusterMatchString sets ClusterMatchString field to given value.

### HasClusterMatchString

`func (o *CreateRemoteVaultSearchJobParameters) HasClusterMatchString() bool`

HasClusterMatchString returns a boolean if a field has been set.

### SetClusterMatchStringNil

`func (o *CreateRemoteVaultSearchJobParameters) SetClusterMatchStringNil(b bool)`

 SetClusterMatchStringNil sets the value for ClusterMatchString to be an explicit nil

### UnsetClusterMatchString
`func (o *CreateRemoteVaultSearchJobParameters) UnsetClusterMatchString()`

UnsetClusterMatchString ensures that no value is present for ClusterMatchString, not even an explicit nil
### GetEncryptionKeys

`func (o *CreateRemoteVaultSearchJobParameters) GetEncryptionKeys() []VaultEncryptionKey`

GetEncryptionKeys returns the EncryptionKeys field if non-nil, zero value otherwise.

### GetEncryptionKeysOk

`func (o *CreateRemoteVaultSearchJobParameters) GetEncryptionKeysOk() (*[]VaultEncryptionKey, bool)`

GetEncryptionKeysOk returns a tuple with the EncryptionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeys

`func (o *CreateRemoteVaultSearchJobParameters) SetEncryptionKeys(v []VaultEncryptionKey)`

SetEncryptionKeys sets EncryptionKeys field to given value.

### HasEncryptionKeys

`func (o *CreateRemoteVaultSearchJobParameters) HasEncryptionKeys() bool`

HasEncryptionKeys returns a boolean if a field has been set.

### SetEncryptionKeysNil

`func (o *CreateRemoteVaultSearchJobParameters) SetEncryptionKeysNil(b bool)`

 SetEncryptionKeysNil sets the value for EncryptionKeys to be an explicit nil

### UnsetEncryptionKeys
`func (o *CreateRemoteVaultSearchJobParameters) UnsetEncryptionKeys()`

UnsetEncryptionKeys ensures that no value is present for EncryptionKeys, not even an explicit nil
### GetEndTimeUsecs

`func (o *CreateRemoteVaultSearchJobParameters) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *CreateRemoteVaultSearchJobParameters) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *CreateRemoteVaultSearchJobParameters) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *CreateRemoteVaultSearchJobParameters) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *CreateRemoteVaultSearchJobParameters) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *CreateRemoteVaultSearchJobParameters) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetJobMatchString

`func (o *CreateRemoteVaultSearchJobParameters) GetJobMatchString() string`

GetJobMatchString returns the JobMatchString field if non-nil, zero value otherwise.

### GetJobMatchStringOk

`func (o *CreateRemoteVaultSearchJobParameters) GetJobMatchStringOk() (*string, bool)`

GetJobMatchStringOk returns a tuple with the JobMatchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMatchString

`func (o *CreateRemoteVaultSearchJobParameters) SetJobMatchString(v string)`

SetJobMatchString sets JobMatchString field to given value.

### HasJobMatchString

`func (o *CreateRemoteVaultSearchJobParameters) HasJobMatchString() bool`

HasJobMatchString returns a boolean if a field has been set.

### SetJobMatchStringNil

`func (o *CreateRemoteVaultSearchJobParameters) SetJobMatchStringNil(b bool)`

 SetJobMatchStringNil sets the value for JobMatchString to be an explicit nil

### UnsetJobMatchString
`func (o *CreateRemoteVaultSearchJobParameters) UnsetJobMatchString()`

UnsetJobMatchString ensures that no value is present for JobMatchString, not even an explicit nil
### GetSearchJobName

`func (o *CreateRemoteVaultSearchJobParameters) GetSearchJobName() string`

GetSearchJobName returns the SearchJobName field if non-nil, zero value otherwise.

### GetSearchJobNameOk

`func (o *CreateRemoteVaultSearchJobParameters) GetSearchJobNameOk() (*string, bool)`

GetSearchJobNameOk returns a tuple with the SearchJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchJobName

`func (o *CreateRemoteVaultSearchJobParameters) SetSearchJobName(v string)`

SetSearchJobName sets SearchJobName field to given value.


### SetSearchJobNameNil

`func (o *CreateRemoteVaultSearchJobParameters) SetSearchJobNameNil(b bool)`

 SetSearchJobNameNil sets the value for SearchJobName to be an explicit nil

### UnsetSearchJobName
`func (o *CreateRemoteVaultSearchJobParameters) UnsetSearchJobName()`

UnsetSearchJobName ensures that no value is present for SearchJobName, not even an explicit nil
### GetStartTimeUsecs

`func (o *CreateRemoteVaultSearchJobParameters) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *CreateRemoteVaultSearchJobParameters) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *CreateRemoteVaultSearchJobParameters) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *CreateRemoteVaultSearchJobParameters) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *CreateRemoteVaultSearchJobParameters) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *CreateRemoteVaultSearchJobParameters) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetVaultId

`func (o *CreateRemoteVaultSearchJobParameters) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CreateRemoteVaultSearchJobParameters) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CreateRemoteVaultSearchJobParameters) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.


### SetVaultIdNil

`func (o *CreateRemoteVaultSearchJobParameters) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *CreateRemoteVaultSearchJobParameters) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


