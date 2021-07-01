# ExchangeDatabaseCopyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationPreferenceNumber** | Pointer to **NullableInt32** | Specifies the activation preference number assigned for this database copy. | [optional] 
**AppServerId** | Pointer to **NullableInt64** | Specifies the entity id of the Exchange Application Server which has this database copy. | [optional] 
**BackupSupported** | Pointer to **NullableBool** | Specifies if backup is supported for the Exchange database copy. | [optional] 
**BackupUnsupportedReasons** | Pointer to **[]string** | Specifies any reason(s) for Exchange database backup not supported. | [optional] 
**CopyGuid** | Pointer to **NullableString** | Specifies the Guid of the Exchange Database Copy. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the database is created in Unix epoch time in milliseconds. | [optional] 
**DatabaseState** | Pointer to **NullableString** | Specifies the state of the Exchange database copy. Specifies the state of Exchange Database Copy.   &#39;kUnknown&#39; indicates the status is not known. &#39;kMounted&#39; indicates the exchange database copy is mounted and healthy. &#39;kError&#39; indicates  the  exchange  database  copy  is unmounted or partially mounted or is in error state. | [optional] 
**DbSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the Exchange database copy in bytes. | [optional] 
**Dbguid** | Pointer to **NullableString** | Specifies the guid of the Exchange Database. | [optional] 
**Fqdn** | Pointer to **NullableString** | Specifies the fully qualified domain name of the Exchange Server on which the database copy is hosted. | [optional] 
**IsActiveCopy** | Pointer to **NullableBool** | Specifies if the Exchange database copy present on the Exchange Application server is active or passive. | [optional] 
**Name** | Pointer to **NullableString** | Specifes the name of the Exchange Database. | [optional] 
**OwnerId** | Pointer to **NullableInt64** | Specifies the owner entity id of the Exchange Application Server which has this database copy. | [optional] 
**ServerName** | Pointer to **NullableString** | Specifies the display name of the Exchange Application Server on which the database copy is hosted. | [optional] 
**UtcOffsetMin** | Pointer to **NullableInt64** | Specifies UTC time offset of database creation time. | [optional] 

## Methods

### NewExchangeDatabaseCopyInfo

`func NewExchangeDatabaseCopyInfo() *ExchangeDatabaseCopyInfo`

NewExchangeDatabaseCopyInfo instantiates a new ExchangeDatabaseCopyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeDatabaseCopyInfoWithDefaults

`func NewExchangeDatabaseCopyInfoWithDefaults() *ExchangeDatabaseCopyInfo`

NewExchangeDatabaseCopyInfoWithDefaults instantiates a new ExchangeDatabaseCopyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationPreferenceNumber

`func (o *ExchangeDatabaseCopyInfo) GetActivationPreferenceNumber() int32`

GetActivationPreferenceNumber returns the ActivationPreferenceNumber field if non-nil, zero value otherwise.

### GetActivationPreferenceNumberOk

`func (o *ExchangeDatabaseCopyInfo) GetActivationPreferenceNumberOk() (*int32, bool)`

GetActivationPreferenceNumberOk returns a tuple with the ActivationPreferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPreferenceNumber

`func (o *ExchangeDatabaseCopyInfo) SetActivationPreferenceNumber(v int32)`

SetActivationPreferenceNumber sets ActivationPreferenceNumber field to given value.

### HasActivationPreferenceNumber

`func (o *ExchangeDatabaseCopyInfo) HasActivationPreferenceNumber() bool`

HasActivationPreferenceNumber returns a boolean if a field has been set.

### SetActivationPreferenceNumberNil

`func (o *ExchangeDatabaseCopyInfo) SetActivationPreferenceNumberNil(b bool)`

 SetActivationPreferenceNumberNil sets the value for ActivationPreferenceNumber to be an explicit nil

### UnsetActivationPreferenceNumber
`func (o *ExchangeDatabaseCopyInfo) UnsetActivationPreferenceNumber()`

UnsetActivationPreferenceNumber ensures that no value is present for ActivationPreferenceNumber, not even an explicit nil
### GetAppServerId

`func (o *ExchangeDatabaseCopyInfo) GetAppServerId() int64`

GetAppServerId returns the AppServerId field if non-nil, zero value otherwise.

### GetAppServerIdOk

`func (o *ExchangeDatabaseCopyInfo) GetAppServerIdOk() (*int64, bool)`

GetAppServerIdOk returns a tuple with the AppServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerId

`func (o *ExchangeDatabaseCopyInfo) SetAppServerId(v int64)`

SetAppServerId sets AppServerId field to given value.

### HasAppServerId

`func (o *ExchangeDatabaseCopyInfo) HasAppServerId() bool`

HasAppServerId returns a boolean if a field has been set.

### SetAppServerIdNil

`func (o *ExchangeDatabaseCopyInfo) SetAppServerIdNil(b bool)`

 SetAppServerIdNil sets the value for AppServerId to be an explicit nil

### UnsetAppServerId
`func (o *ExchangeDatabaseCopyInfo) UnsetAppServerId()`

UnsetAppServerId ensures that no value is present for AppServerId, not even an explicit nil
### GetBackupSupported

`func (o *ExchangeDatabaseCopyInfo) GetBackupSupported() bool`

GetBackupSupported returns the BackupSupported field if non-nil, zero value otherwise.

### GetBackupSupportedOk

`func (o *ExchangeDatabaseCopyInfo) GetBackupSupportedOk() (*bool, bool)`

GetBackupSupportedOk returns a tuple with the BackupSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSupported

`func (o *ExchangeDatabaseCopyInfo) SetBackupSupported(v bool)`

SetBackupSupported sets BackupSupported field to given value.

### HasBackupSupported

`func (o *ExchangeDatabaseCopyInfo) HasBackupSupported() bool`

HasBackupSupported returns a boolean if a field has been set.

### SetBackupSupportedNil

`func (o *ExchangeDatabaseCopyInfo) SetBackupSupportedNil(b bool)`

 SetBackupSupportedNil sets the value for BackupSupported to be an explicit nil

### UnsetBackupSupported
`func (o *ExchangeDatabaseCopyInfo) UnsetBackupSupported()`

UnsetBackupSupported ensures that no value is present for BackupSupported, not even an explicit nil
### GetBackupUnsupportedReasons

`func (o *ExchangeDatabaseCopyInfo) GetBackupUnsupportedReasons() []string`

GetBackupUnsupportedReasons returns the BackupUnsupportedReasons field if non-nil, zero value otherwise.

### GetBackupUnsupportedReasonsOk

`func (o *ExchangeDatabaseCopyInfo) GetBackupUnsupportedReasonsOk() (*[]string, bool)`

GetBackupUnsupportedReasonsOk returns a tuple with the BackupUnsupportedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUnsupportedReasons

`func (o *ExchangeDatabaseCopyInfo) SetBackupUnsupportedReasons(v []string)`

SetBackupUnsupportedReasons sets BackupUnsupportedReasons field to given value.

### HasBackupUnsupportedReasons

`func (o *ExchangeDatabaseCopyInfo) HasBackupUnsupportedReasons() bool`

HasBackupUnsupportedReasons returns a boolean if a field has been set.

### SetBackupUnsupportedReasonsNil

`func (o *ExchangeDatabaseCopyInfo) SetBackupUnsupportedReasonsNil(b bool)`

 SetBackupUnsupportedReasonsNil sets the value for BackupUnsupportedReasons to be an explicit nil

### UnsetBackupUnsupportedReasons
`func (o *ExchangeDatabaseCopyInfo) UnsetBackupUnsupportedReasons()`

UnsetBackupUnsupportedReasons ensures that no value is present for BackupUnsupportedReasons, not even an explicit nil
### GetCopyGuid

`func (o *ExchangeDatabaseCopyInfo) GetCopyGuid() string`

GetCopyGuid returns the CopyGuid field if non-nil, zero value otherwise.

### GetCopyGuidOk

`func (o *ExchangeDatabaseCopyInfo) GetCopyGuidOk() (*string, bool)`

GetCopyGuidOk returns a tuple with the CopyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyGuid

`func (o *ExchangeDatabaseCopyInfo) SetCopyGuid(v string)`

SetCopyGuid sets CopyGuid field to given value.

### HasCopyGuid

`func (o *ExchangeDatabaseCopyInfo) HasCopyGuid() bool`

HasCopyGuid returns a boolean if a field has been set.

### SetCopyGuidNil

`func (o *ExchangeDatabaseCopyInfo) SetCopyGuidNil(b bool)`

 SetCopyGuidNil sets the value for CopyGuid to be an explicit nil

### UnsetCopyGuid
`func (o *ExchangeDatabaseCopyInfo) UnsetCopyGuid()`

UnsetCopyGuid ensures that no value is present for CopyGuid, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *ExchangeDatabaseCopyInfo) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *ExchangeDatabaseCopyInfo) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *ExchangeDatabaseCopyInfo) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *ExchangeDatabaseCopyInfo) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *ExchangeDatabaseCopyInfo) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *ExchangeDatabaseCopyInfo) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetDatabaseState

`func (o *ExchangeDatabaseCopyInfo) GetDatabaseState() string`

GetDatabaseState returns the DatabaseState field if non-nil, zero value otherwise.

### GetDatabaseStateOk

`func (o *ExchangeDatabaseCopyInfo) GetDatabaseStateOk() (*string, bool)`

GetDatabaseStateOk returns a tuple with the DatabaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseState

`func (o *ExchangeDatabaseCopyInfo) SetDatabaseState(v string)`

SetDatabaseState sets DatabaseState field to given value.

### HasDatabaseState

`func (o *ExchangeDatabaseCopyInfo) HasDatabaseState() bool`

HasDatabaseState returns a boolean if a field has been set.

### SetDatabaseStateNil

`func (o *ExchangeDatabaseCopyInfo) SetDatabaseStateNil(b bool)`

 SetDatabaseStateNil sets the value for DatabaseState to be an explicit nil

### UnsetDatabaseState
`func (o *ExchangeDatabaseCopyInfo) UnsetDatabaseState()`

UnsetDatabaseState ensures that no value is present for DatabaseState, not even an explicit nil
### GetDbSizeBytes

`func (o *ExchangeDatabaseCopyInfo) GetDbSizeBytes() int64`

GetDbSizeBytes returns the DbSizeBytes field if non-nil, zero value otherwise.

### GetDbSizeBytesOk

`func (o *ExchangeDatabaseCopyInfo) GetDbSizeBytesOk() (*int64, bool)`

GetDbSizeBytesOk returns a tuple with the DbSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbSizeBytes

`func (o *ExchangeDatabaseCopyInfo) SetDbSizeBytes(v int64)`

SetDbSizeBytes sets DbSizeBytes field to given value.

### HasDbSizeBytes

`func (o *ExchangeDatabaseCopyInfo) HasDbSizeBytes() bool`

HasDbSizeBytes returns a boolean if a field has been set.

### SetDbSizeBytesNil

`func (o *ExchangeDatabaseCopyInfo) SetDbSizeBytesNil(b bool)`

 SetDbSizeBytesNil sets the value for DbSizeBytes to be an explicit nil

### UnsetDbSizeBytes
`func (o *ExchangeDatabaseCopyInfo) UnsetDbSizeBytes()`

UnsetDbSizeBytes ensures that no value is present for DbSizeBytes, not even an explicit nil
### GetDbguid

`func (o *ExchangeDatabaseCopyInfo) GetDbguid() string`

GetDbguid returns the Dbguid field if non-nil, zero value otherwise.

### GetDbguidOk

`func (o *ExchangeDatabaseCopyInfo) GetDbguidOk() (*string, bool)`

GetDbguidOk returns a tuple with the Dbguid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbguid

`func (o *ExchangeDatabaseCopyInfo) SetDbguid(v string)`

SetDbguid sets Dbguid field to given value.

### HasDbguid

`func (o *ExchangeDatabaseCopyInfo) HasDbguid() bool`

HasDbguid returns a boolean if a field has been set.

### SetDbguidNil

`func (o *ExchangeDatabaseCopyInfo) SetDbguidNil(b bool)`

 SetDbguidNil sets the value for Dbguid to be an explicit nil

### UnsetDbguid
`func (o *ExchangeDatabaseCopyInfo) UnsetDbguid()`

UnsetDbguid ensures that no value is present for Dbguid, not even an explicit nil
### GetFqdn

`func (o *ExchangeDatabaseCopyInfo) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ExchangeDatabaseCopyInfo) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ExchangeDatabaseCopyInfo) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ExchangeDatabaseCopyInfo) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *ExchangeDatabaseCopyInfo) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *ExchangeDatabaseCopyInfo) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetIsActiveCopy

`func (o *ExchangeDatabaseCopyInfo) GetIsActiveCopy() bool`

GetIsActiveCopy returns the IsActiveCopy field if non-nil, zero value otherwise.

### GetIsActiveCopyOk

`func (o *ExchangeDatabaseCopyInfo) GetIsActiveCopyOk() (*bool, bool)`

GetIsActiveCopyOk returns a tuple with the IsActiveCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActiveCopy

`func (o *ExchangeDatabaseCopyInfo) SetIsActiveCopy(v bool)`

SetIsActiveCopy sets IsActiveCopy field to given value.

### HasIsActiveCopy

`func (o *ExchangeDatabaseCopyInfo) HasIsActiveCopy() bool`

HasIsActiveCopy returns a boolean if a field has been set.

### SetIsActiveCopyNil

`func (o *ExchangeDatabaseCopyInfo) SetIsActiveCopyNil(b bool)`

 SetIsActiveCopyNil sets the value for IsActiveCopy to be an explicit nil

### UnsetIsActiveCopy
`func (o *ExchangeDatabaseCopyInfo) UnsetIsActiveCopy()`

UnsetIsActiveCopy ensures that no value is present for IsActiveCopy, not even an explicit nil
### GetName

`func (o *ExchangeDatabaseCopyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeDatabaseCopyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeDatabaseCopyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeDatabaseCopyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExchangeDatabaseCopyInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExchangeDatabaseCopyInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *ExchangeDatabaseCopyInfo) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ExchangeDatabaseCopyInfo) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ExchangeDatabaseCopyInfo) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ExchangeDatabaseCopyInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *ExchangeDatabaseCopyInfo) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *ExchangeDatabaseCopyInfo) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetServerName

`func (o *ExchangeDatabaseCopyInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ExchangeDatabaseCopyInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ExchangeDatabaseCopyInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ExchangeDatabaseCopyInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *ExchangeDatabaseCopyInfo) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *ExchangeDatabaseCopyInfo) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetUtcOffsetMin

`func (o *ExchangeDatabaseCopyInfo) GetUtcOffsetMin() int64`

GetUtcOffsetMin returns the UtcOffsetMin field if non-nil, zero value otherwise.

### GetUtcOffsetMinOk

`func (o *ExchangeDatabaseCopyInfo) GetUtcOffsetMinOk() (*int64, bool)`

GetUtcOffsetMinOk returns a tuple with the UtcOffsetMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffsetMin

`func (o *ExchangeDatabaseCopyInfo) SetUtcOffsetMin(v int64)`

SetUtcOffsetMin sets UtcOffsetMin field to given value.

### HasUtcOffsetMin

`func (o *ExchangeDatabaseCopyInfo) HasUtcOffsetMin() bool`

HasUtcOffsetMin returns a boolean if a field has been set.

### SetUtcOffsetMinNil

`func (o *ExchangeDatabaseCopyInfo) SetUtcOffsetMinNil(b bool)`

 SetUtcOffsetMinNil sets the value for UtcOffsetMin to be an explicit nil

### UnsetUtcOffsetMin
`func (o *ExchangeDatabaseCopyInfo) UnsetUtcOffsetMin()`

UnsetUtcOffsetMin ensures that no value is present for UtcOffsetMin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


