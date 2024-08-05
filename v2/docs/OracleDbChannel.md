# OracleDbChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveLogRetentionDays** | Pointer to **NullableInt32** | Specifies the number of days archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after n days, set this to n. | [optional] 
**ArchiveLogRetentionHours** | Pointer to **NullableInt32** | Specifies the number of hours archive log should be stored. For keeping the archived log forever, set this to -1. For deleting the archived log immediately, set this to 0. For deleting the archived log after k hours, set this to k. | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**DatabaseNodeList** | Pointer to [**[]OracleDatabaseHost**](OracleDatabaseHost.md) | Specifies the Node info from where we are allowed to take the backup/restore. | [optional] 
**DatabaseUniqueName** | Pointer to **NullableString** | Specifies the unique Name of the database. | [optional] 
**DatabaseUuid** | Pointer to **NullableString** | Specifies the database unique id. This is an internal field and is filled by magneto master based on corresponding app entity id. | [optional] 
**DefaultChannelCount** | Pointer to **NullableInt32** | Specifies the default number of channels to use per node per database. This value is used on all Oracle Database Nodes unless databaseNodeList item&#39;s channelCount is specified for the node. Default value for the number of channels will be calculated as the minimum of number of nodes in Cohesity cluster and 2 * number of CPU on the host. If the number of channels is unspecified here and unspecified within databaseNodeList, the above formula will be used to determine the same. | [optional] 
**EnableDgPrimaryBackup** | Pointer to **NullableBool** | Specifies whether the database having the Primary role within Data Guard configuration is to be backed up. | [optional] 
**MaxHostCount** | Pointer to **NullableInt32** | Specifies the maximum number of hosts from which backup/restore is allowed in parallel. This will be less than or equal to the number of databaseNode specified within databaseNodeList. | [optional] 
**RmanBackupType** | Pointer to **string** | Specifies the type of Oracle RMAN backup requested | [optional] 

## Methods

### NewOracleDbChannel

`func NewOracleDbChannel() *OracleDbChannel`

NewOracleDbChannel instantiates a new OracleDbChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDbChannelWithDefaults

`func NewOracleDbChannelWithDefaults() *OracleDbChannel`

NewOracleDbChannelWithDefaults instantiates a new OracleDbChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveLogRetentionDays

`func (o *OracleDbChannel) GetArchiveLogRetentionDays() int32`

GetArchiveLogRetentionDays returns the ArchiveLogRetentionDays field if non-nil, zero value otherwise.

### GetArchiveLogRetentionDaysOk

`func (o *OracleDbChannel) GetArchiveLogRetentionDaysOk() (*int32, bool)`

GetArchiveLogRetentionDaysOk returns a tuple with the ArchiveLogRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogRetentionDays

`func (o *OracleDbChannel) SetArchiveLogRetentionDays(v int32)`

SetArchiveLogRetentionDays sets ArchiveLogRetentionDays field to given value.

### HasArchiveLogRetentionDays

`func (o *OracleDbChannel) HasArchiveLogRetentionDays() bool`

HasArchiveLogRetentionDays returns a boolean if a field has been set.

### SetArchiveLogRetentionDaysNil

`func (o *OracleDbChannel) SetArchiveLogRetentionDaysNil(b bool)`

 SetArchiveLogRetentionDaysNil sets the value for ArchiveLogRetentionDays to be an explicit nil

### UnsetArchiveLogRetentionDays
`func (o *OracleDbChannel) UnsetArchiveLogRetentionDays()`

UnsetArchiveLogRetentionDays ensures that no value is present for ArchiveLogRetentionDays, not even an explicit nil
### GetArchiveLogRetentionHours

`func (o *OracleDbChannel) GetArchiveLogRetentionHours() int32`

GetArchiveLogRetentionHours returns the ArchiveLogRetentionHours field if non-nil, zero value otherwise.

### GetArchiveLogRetentionHoursOk

`func (o *OracleDbChannel) GetArchiveLogRetentionHoursOk() (*int32, bool)`

GetArchiveLogRetentionHoursOk returns a tuple with the ArchiveLogRetentionHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogRetentionHours

`func (o *OracleDbChannel) SetArchiveLogRetentionHours(v int32)`

SetArchiveLogRetentionHours sets ArchiveLogRetentionHours field to given value.

### HasArchiveLogRetentionHours

`func (o *OracleDbChannel) HasArchiveLogRetentionHours() bool`

HasArchiveLogRetentionHours returns a boolean if a field has been set.

### SetArchiveLogRetentionHoursNil

`func (o *OracleDbChannel) SetArchiveLogRetentionHoursNil(b bool)`

 SetArchiveLogRetentionHoursNil sets the value for ArchiveLogRetentionHours to be an explicit nil

### UnsetArchiveLogRetentionHours
`func (o *OracleDbChannel) UnsetArchiveLogRetentionHours()`

UnsetArchiveLogRetentionHours ensures that no value is present for ArchiveLogRetentionHours, not even an explicit nil
### GetCredentials

`func (o *OracleDbChannel) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OracleDbChannel) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OracleDbChannel) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OracleDbChannel) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDatabaseNodeList

`func (o *OracleDbChannel) GetDatabaseNodeList() []OracleDatabaseHost`

GetDatabaseNodeList returns the DatabaseNodeList field if non-nil, zero value otherwise.

### GetDatabaseNodeListOk

`func (o *OracleDbChannel) GetDatabaseNodeListOk() (*[]OracleDatabaseHost, bool)`

GetDatabaseNodeListOk returns a tuple with the DatabaseNodeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseNodeList

`func (o *OracleDbChannel) SetDatabaseNodeList(v []OracleDatabaseHost)`

SetDatabaseNodeList sets DatabaseNodeList field to given value.

### HasDatabaseNodeList

`func (o *OracleDbChannel) HasDatabaseNodeList() bool`

HasDatabaseNodeList returns a boolean if a field has been set.

### SetDatabaseNodeListNil

`func (o *OracleDbChannel) SetDatabaseNodeListNil(b bool)`

 SetDatabaseNodeListNil sets the value for DatabaseNodeList to be an explicit nil

### UnsetDatabaseNodeList
`func (o *OracleDbChannel) UnsetDatabaseNodeList()`

UnsetDatabaseNodeList ensures that no value is present for DatabaseNodeList, not even an explicit nil
### GetDatabaseUniqueName

`func (o *OracleDbChannel) GetDatabaseUniqueName() string`

GetDatabaseUniqueName returns the DatabaseUniqueName field if non-nil, zero value otherwise.

### GetDatabaseUniqueNameOk

`func (o *OracleDbChannel) GetDatabaseUniqueNameOk() (*string, bool)`

GetDatabaseUniqueNameOk returns a tuple with the DatabaseUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUniqueName

`func (o *OracleDbChannel) SetDatabaseUniqueName(v string)`

SetDatabaseUniqueName sets DatabaseUniqueName field to given value.

### HasDatabaseUniqueName

`func (o *OracleDbChannel) HasDatabaseUniqueName() bool`

HasDatabaseUniqueName returns a boolean if a field has been set.

### SetDatabaseUniqueNameNil

`func (o *OracleDbChannel) SetDatabaseUniqueNameNil(b bool)`

 SetDatabaseUniqueNameNil sets the value for DatabaseUniqueName to be an explicit nil

### UnsetDatabaseUniqueName
`func (o *OracleDbChannel) UnsetDatabaseUniqueName()`

UnsetDatabaseUniqueName ensures that no value is present for DatabaseUniqueName, not even an explicit nil
### GetDatabaseUuid

`func (o *OracleDbChannel) GetDatabaseUuid() string`

GetDatabaseUuid returns the DatabaseUuid field if non-nil, zero value otherwise.

### GetDatabaseUuidOk

`func (o *OracleDbChannel) GetDatabaseUuidOk() (*string, bool)`

GetDatabaseUuidOk returns a tuple with the DatabaseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUuid

`func (o *OracleDbChannel) SetDatabaseUuid(v string)`

SetDatabaseUuid sets DatabaseUuid field to given value.

### HasDatabaseUuid

`func (o *OracleDbChannel) HasDatabaseUuid() bool`

HasDatabaseUuid returns a boolean if a field has been set.

### SetDatabaseUuidNil

`func (o *OracleDbChannel) SetDatabaseUuidNil(b bool)`

 SetDatabaseUuidNil sets the value for DatabaseUuid to be an explicit nil

### UnsetDatabaseUuid
`func (o *OracleDbChannel) UnsetDatabaseUuid()`

UnsetDatabaseUuid ensures that no value is present for DatabaseUuid, not even an explicit nil
### GetDefaultChannelCount

`func (o *OracleDbChannel) GetDefaultChannelCount() int32`

GetDefaultChannelCount returns the DefaultChannelCount field if non-nil, zero value otherwise.

### GetDefaultChannelCountOk

`func (o *OracleDbChannel) GetDefaultChannelCountOk() (*int32, bool)`

GetDefaultChannelCountOk returns a tuple with the DefaultChannelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelCount

`func (o *OracleDbChannel) SetDefaultChannelCount(v int32)`

SetDefaultChannelCount sets DefaultChannelCount field to given value.

### HasDefaultChannelCount

`func (o *OracleDbChannel) HasDefaultChannelCount() bool`

HasDefaultChannelCount returns a boolean if a field has been set.

### SetDefaultChannelCountNil

`func (o *OracleDbChannel) SetDefaultChannelCountNil(b bool)`

 SetDefaultChannelCountNil sets the value for DefaultChannelCount to be an explicit nil

### UnsetDefaultChannelCount
`func (o *OracleDbChannel) UnsetDefaultChannelCount()`

UnsetDefaultChannelCount ensures that no value is present for DefaultChannelCount, not even an explicit nil
### GetEnableDgPrimaryBackup

`func (o *OracleDbChannel) GetEnableDgPrimaryBackup() bool`

GetEnableDgPrimaryBackup returns the EnableDgPrimaryBackup field if non-nil, zero value otherwise.

### GetEnableDgPrimaryBackupOk

`func (o *OracleDbChannel) GetEnableDgPrimaryBackupOk() (*bool, bool)`

GetEnableDgPrimaryBackupOk returns a tuple with the EnableDgPrimaryBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDgPrimaryBackup

`func (o *OracleDbChannel) SetEnableDgPrimaryBackup(v bool)`

SetEnableDgPrimaryBackup sets EnableDgPrimaryBackup field to given value.

### HasEnableDgPrimaryBackup

`func (o *OracleDbChannel) HasEnableDgPrimaryBackup() bool`

HasEnableDgPrimaryBackup returns a boolean if a field has been set.

### SetEnableDgPrimaryBackupNil

`func (o *OracleDbChannel) SetEnableDgPrimaryBackupNil(b bool)`

 SetEnableDgPrimaryBackupNil sets the value for EnableDgPrimaryBackup to be an explicit nil

### UnsetEnableDgPrimaryBackup
`func (o *OracleDbChannel) UnsetEnableDgPrimaryBackup()`

UnsetEnableDgPrimaryBackup ensures that no value is present for EnableDgPrimaryBackup, not even an explicit nil
### GetMaxHostCount

`func (o *OracleDbChannel) GetMaxHostCount() int32`

GetMaxHostCount returns the MaxHostCount field if non-nil, zero value otherwise.

### GetMaxHostCountOk

`func (o *OracleDbChannel) GetMaxHostCountOk() (*int32, bool)`

GetMaxHostCountOk returns a tuple with the MaxHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHostCount

`func (o *OracleDbChannel) SetMaxHostCount(v int32)`

SetMaxHostCount sets MaxHostCount field to given value.

### HasMaxHostCount

`func (o *OracleDbChannel) HasMaxHostCount() bool`

HasMaxHostCount returns a boolean if a field has been set.

### SetMaxHostCountNil

`func (o *OracleDbChannel) SetMaxHostCountNil(b bool)`

 SetMaxHostCountNil sets the value for MaxHostCount to be an explicit nil

### UnsetMaxHostCount
`func (o *OracleDbChannel) UnsetMaxHostCount()`

UnsetMaxHostCount ensures that no value is present for MaxHostCount, not even an explicit nil
### GetRmanBackupType

`func (o *OracleDbChannel) GetRmanBackupType() string`

GetRmanBackupType returns the RmanBackupType field if non-nil, zero value otherwise.

### GetRmanBackupTypeOk

`func (o *OracleDbChannel) GetRmanBackupTypeOk() (*string, bool)`

GetRmanBackupTypeOk returns a tuple with the RmanBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmanBackupType

`func (o *OracleDbChannel) SetRmanBackupType(v string)`

SetRmanBackupType sets RmanBackupType field to given value.

### HasRmanBackupType

`func (o *OracleDbChannel) HasRmanBackupType() bool`

HasRmanBackupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


