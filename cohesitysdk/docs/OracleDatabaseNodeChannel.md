# OracleDatabaseNodeChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveLogKeepDays** | Pointer to **NullableInt32** | Specifies the number of days archive log should be stored. | [optional] 
**DatabaseNodeList** | Pointer to [**[]OracleDatabaseNode**](OracleDatabaseNode.md) | Array of nodes of a database.  Specifies the Node info from where we are allowed to take the backup/restore. | [optional] 
**DatabaseUniqueName** | Pointer to **NullableString** | Specifies the unique Name of the database. | [optional] 
**DatabaseUuid** | Pointer to **NullableString** | Specifies the database unique id. This is an internal field and is filled by magneto master based on corresponding app entity id. | [optional] 
**DefaultChannelCount** | Pointer to **NullableInt32** | Specifies the default number of channels to use per node per database. The default number of channels to use per host per db. This value is used on all OracleDatabaseNode&#39;s unless databaseNodeList item&#39;s channelCount is specified for the node. | [optional] 
**EnableDgPrimaryBackup** | Pointer to **NullableBool** | Specifies whether the database having the Primary role within Data Guard configuration is to be backed up. | [optional] 
**MaxNodeCount** | Pointer to **NullableInt32** | Specifies the maximum number of nodes from which we are allowed to take backup/restore. | [optional] 
**RmanBackupType** | Pointer to **NullableInt32** | Specifies the type of Oracle RMAN backup. | [optional] 

## Methods

### NewOracleDatabaseNodeChannel

`func NewOracleDatabaseNodeChannel() *OracleDatabaseNodeChannel`

NewOracleDatabaseNodeChannel instantiates a new OracleDatabaseNodeChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDatabaseNodeChannelWithDefaults

`func NewOracleDatabaseNodeChannelWithDefaults() *OracleDatabaseNodeChannel`

NewOracleDatabaseNodeChannelWithDefaults instantiates a new OracleDatabaseNodeChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveLogKeepDays

`func (o *OracleDatabaseNodeChannel) GetArchiveLogKeepDays() int32`

GetArchiveLogKeepDays returns the ArchiveLogKeepDays field if non-nil, zero value otherwise.

### GetArchiveLogKeepDaysOk

`func (o *OracleDatabaseNodeChannel) GetArchiveLogKeepDaysOk() (*int32, bool)`

GetArchiveLogKeepDaysOk returns a tuple with the ArchiveLogKeepDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogKeepDays

`func (o *OracleDatabaseNodeChannel) SetArchiveLogKeepDays(v int32)`

SetArchiveLogKeepDays sets ArchiveLogKeepDays field to given value.

### HasArchiveLogKeepDays

`func (o *OracleDatabaseNodeChannel) HasArchiveLogKeepDays() bool`

HasArchiveLogKeepDays returns a boolean if a field has been set.

### SetArchiveLogKeepDaysNil

`func (o *OracleDatabaseNodeChannel) SetArchiveLogKeepDaysNil(b bool)`

 SetArchiveLogKeepDaysNil sets the value for ArchiveLogKeepDays to be an explicit nil

### UnsetArchiveLogKeepDays
`func (o *OracleDatabaseNodeChannel) UnsetArchiveLogKeepDays()`

UnsetArchiveLogKeepDays ensures that no value is present for ArchiveLogKeepDays, not even an explicit nil
### GetDatabaseNodeList

`func (o *OracleDatabaseNodeChannel) GetDatabaseNodeList() []OracleDatabaseNode`

GetDatabaseNodeList returns the DatabaseNodeList field if non-nil, zero value otherwise.

### GetDatabaseNodeListOk

`func (o *OracleDatabaseNodeChannel) GetDatabaseNodeListOk() (*[]OracleDatabaseNode, bool)`

GetDatabaseNodeListOk returns a tuple with the DatabaseNodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseNodeList

`func (o *OracleDatabaseNodeChannel) SetDatabaseNodeList(v []OracleDatabaseNode)`

SetDatabaseNodeList sets DatabaseNodeList field to given value.

### HasDatabaseNodeList

`func (o *OracleDatabaseNodeChannel) HasDatabaseNodeList() bool`

HasDatabaseNodeList returns a boolean if a field has been set.

### SetDatabaseNodeListNil

`func (o *OracleDatabaseNodeChannel) SetDatabaseNodeListNil(b bool)`

 SetDatabaseNodeListNil sets the value for DatabaseNodeList to be an explicit nil

### UnsetDatabaseNodeList
`func (o *OracleDatabaseNodeChannel) UnsetDatabaseNodeList()`

UnsetDatabaseNodeList ensures that no value is present for DatabaseNodeList, not even an explicit nil
### GetDatabaseUniqueName

`func (o *OracleDatabaseNodeChannel) GetDatabaseUniqueName() string`

GetDatabaseUniqueName returns the DatabaseUniqueName field if non-nil, zero value otherwise.

### GetDatabaseUniqueNameOk

`func (o *OracleDatabaseNodeChannel) GetDatabaseUniqueNameOk() (*string, bool)`

GetDatabaseUniqueNameOk returns a tuple with the DatabaseUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUniqueName

`func (o *OracleDatabaseNodeChannel) SetDatabaseUniqueName(v string)`

SetDatabaseUniqueName sets DatabaseUniqueName field to given value.

### HasDatabaseUniqueName

`func (o *OracleDatabaseNodeChannel) HasDatabaseUniqueName() bool`

HasDatabaseUniqueName returns a boolean if a field has been set.

### SetDatabaseUniqueNameNil

`func (o *OracleDatabaseNodeChannel) SetDatabaseUniqueNameNil(b bool)`

 SetDatabaseUniqueNameNil sets the value for DatabaseUniqueName to be an explicit nil

### UnsetDatabaseUniqueName
`func (o *OracleDatabaseNodeChannel) UnsetDatabaseUniqueName()`

UnsetDatabaseUniqueName ensures that no value is present for DatabaseUniqueName, not even an explicit nil
### GetDatabaseUuid

`func (o *OracleDatabaseNodeChannel) GetDatabaseUuid() string`

GetDatabaseUuid returns the DatabaseUuid field if non-nil, zero value otherwise.

### GetDatabaseUuidOk

`func (o *OracleDatabaseNodeChannel) GetDatabaseUuidOk() (*string, bool)`

GetDatabaseUuidOk returns a tuple with the DatabaseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUuid

`func (o *OracleDatabaseNodeChannel) SetDatabaseUuid(v string)`

SetDatabaseUuid sets DatabaseUuid field to given value.

### HasDatabaseUuid

`func (o *OracleDatabaseNodeChannel) HasDatabaseUuid() bool`

HasDatabaseUuid returns a boolean if a field has been set.

### SetDatabaseUuidNil

`func (o *OracleDatabaseNodeChannel) SetDatabaseUuidNil(b bool)`

 SetDatabaseUuidNil sets the value for DatabaseUuid to be an explicit nil

### UnsetDatabaseUuid
`func (o *OracleDatabaseNodeChannel) UnsetDatabaseUuid()`

UnsetDatabaseUuid ensures that no value is present for DatabaseUuid, not even an explicit nil
### GetDefaultChannelCount

`func (o *OracleDatabaseNodeChannel) GetDefaultChannelCount() int32`

GetDefaultChannelCount returns the DefaultChannelCount field if non-nil, zero value otherwise.

### GetDefaultChannelCountOk

`func (o *OracleDatabaseNodeChannel) GetDefaultChannelCountOk() (*int32, bool)`

GetDefaultChannelCountOk returns a tuple with the DefaultChannelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelCount

`func (o *OracleDatabaseNodeChannel) SetDefaultChannelCount(v int32)`

SetDefaultChannelCount sets DefaultChannelCount field to given value.

### HasDefaultChannelCount

`func (o *OracleDatabaseNodeChannel) HasDefaultChannelCount() bool`

HasDefaultChannelCount returns a boolean if a field has been set.

### SetDefaultChannelCountNil

`func (o *OracleDatabaseNodeChannel) SetDefaultChannelCountNil(b bool)`

 SetDefaultChannelCountNil sets the value for DefaultChannelCount to be an explicit nil

### UnsetDefaultChannelCount
`func (o *OracleDatabaseNodeChannel) UnsetDefaultChannelCount()`

UnsetDefaultChannelCount ensures that no value is present for DefaultChannelCount, not even an explicit nil
### GetEnableDgPrimaryBackup

`func (o *OracleDatabaseNodeChannel) GetEnableDgPrimaryBackup() bool`

GetEnableDgPrimaryBackup returns the EnableDgPrimaryBackup field if non-nil, zero value otherwise.

### GetEnableDgPrimaryBackupOk

`func (o *OracleDatabaseNodeChannel) GetEnableDgPrimaryBackupOk() (*bool, bool)`

GetEnableDgPrimaryBackupOk returns a tuple with the EnableDgPrimaryBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDgPrimaryBackup

`func (o *OracleDatabaseNodeChannel) SetEnableDgPrimaryBackup(v bool)`

SetEnableDgPrimaryBackup sets EnableDgPrimaryBackup field to given value.

### HasEnableDgPrimaryBackup

`func (o *OracleDatabaseNodeChannel) HasEnableDgPrimaryBackup() bool`

HasEnableDgPrimaryBackup returns a boolean if a field has been set.

### SetEnableDgPrimaryBackupNil

`func (o *OracleDatabaseNodeChannel) SetEnableDgPrimaryBackupNil(b bool)`

 SetEnableDgPrimaryBackupNil sets the value for EnableDgPrimaryBackup to be an explicit nil

### UnsetEnableDgPrimaryBackup
`func (o *OracleDatabaseNodeChannel) UnsetEnableDgPrimaryBackup()`

UnsetEnableDgPrimaryBackup ensures that no value is present for EnableDgPrimaryBackup, not even an explicit nil
### GetMaxNodeCount

`func (o *OracleDatabaseNodeChannel) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *OracleDatabaseNodeChannel) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *OracleDatabaseNodeChannel) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *OracleDatabaseNodeChannel) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *OracleDatabaseNodeChannel) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *OracleDatabaseNodeChannel) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetRmanBackupType

`func (o *OracleDatabaseNodeChannel) GetRmanBackupType() int32`

GetRmanBackupType returns the RmanBackupType field if non-nil, zero value otherwise.

### GetRmanBackupTypeOk

`func (o *OracleDatabaseNodeChannel) GetRmanBackupTypeOk() (*int32, bool)`

GetRmanBackupTypeOk returns a tuple with the RmanBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmanBackupType

`func (o *OracleDatabaseNodeChannel) SetRmanBackupType(v int32)`

SetRmanBackupType sets RmanBackupType field to given value.

### HasRmanBackupType

`func (o *OracleDatabaseNodeChannel) HasRmanBackupType() bool`

HasRmanBackupType returns a boolean if a field has been set.

### SetRmanBackupTypeNil

`func (o *OracleDatabaseNodeChannel) SetRmanBackupTypeNil(b bool)`

 SetRmanBackupTypeNil sets the value for RmanBackupType to be an explicit nil

### UnsetRmanBackupType
`func (o *OracleDatabaseNodeChannel) UnsetRmanBackupType()`

UnsetRmanBackupType ensures that no value is present for RmanBackupType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


