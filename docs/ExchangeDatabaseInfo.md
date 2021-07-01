# ExchangeDatabaseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppServerId** | Pointer to **NullableInt64** | Specifies the entity id of the Exchange Application Server which has this database copy. | [optional] 
**BackupSupported** | Pointer to **NullableBool** | Specifies if backup is supported for the Exchange database copy. | [optional] 
**BackupUnsupportedReasons** | Pointer to **[]string** | Specifies any reason(s) for Exchange database backup not supported. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the database is created in Unix epoch time in milliseconds. | [optional] 
**DatabaseState** | Pointer to **NullableString** | Specifies the state of the Exchange database copy. Specifies the state of Exchange Database Copy.   &#39;kUnknown&#39; indicates the status is not known. &#39;kMounted&#39; indicates the exchange database copy is mounted and healthy. &#39;kError&#39; indicates  the  exchange  database  copy  is unmounted or partially mounted or is in error state. | [optional] 
**DbSizeBytes** | Pointer to **NullableInt64** | Specifies the size of the Exchange database copy in bytes. | [optional] 
**Dbguid** | Pointer to **NullableString** | Specifies the guid of the Exchange Database. | [optional] 
**Name** | Pointer to **NullableString** | Specifes the name of the Exchange Database. | [optional] 
**OwnerId** | Pointer to **NullableInt64** | Specifies the owner entity id of the Exchange Application Server which has this database copy. | [optional] 
**UtcOffsetMin** | Pointer to **NullableInt64** | Specifies UTC time offset of database creation time. | [optional] 

## Methods

### NewExchangeDatabaseInfo

`func NewExchangeDatabaseInfo() *ExchangeDatabaseInfo`

NewExchangeDatabaseInfo instantiates a new ExchangeDatabaseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeDatabaseInfoWithDefaults

`func NewExchangeDatabaseInfoWithDefaults() *ExchangeDatabaseInfo`

NewExchangeDatabaseInfoWithDefaults instantiates a new ExchangeDatabaseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppServerId

`func (o *ExchangeDatabaseInfo) GetAppServerId() int64`

GetAppServerId returns the AppServerId field if non-nil, zero value otherwise.

### GetAppServerIdOk

`func (o *ExchangeDatabaseInfo) GetAppServerIdOk() (*int64, bool)`

GetAppServerIdOk returns a tuple with the AppServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServerId

`func (o *ExchangeDatabaseInfo) SetAppServerId(v int64)`

SetAppServerId sets AppServerId field to given value.

### HasAppServerId

`func (o *ExchangeDatabaseInfo) HasAppServerId() bool`

HasAppServerId returns a boolean if a field has been set.

### SetAppServerIdNil

`func (o *ExchangeDatabaseInfo) SetAppServerIdNil(b bool)`

 SetAppServerIdNil sets the value for AppServerId to be an explicit nil

### UnsetAppServerId
`func (o *ExchangeDatabaseInfo) UnsetAppServerId()`

UnsetAppServerId ensures that no value is present for AppServerId, not even an explicit nil
### GetBackupSupported

`func (o *ExchangeDatabaseInfo) GetBackupSupported() bool`

GetBackupSupported returns the BackupSupported field if non-nil, zero value otherwise.

### GetBackupSupportedOk

`func (o *ExchangeDatabaseInfo) GetBackupSupportedOk() (*bool, bool)`

GetBackupSupportedOk returns a tuple with the BackupSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSupported

`func (o *ExchangeDatabaseInfo) SetBackupSupported(v bool)`

SetBackupSupported sets BackupSupported field to given value.

### HasBackupSupported

`func (o *ExchangeDatabaseInfo) HasBackupSupported() bool`

HasBackupSupported returns a boolean if a field has been set.

### SetBackupSupportedNil

`func (o *ExchangeDatabaseInfo) SetBackupSupportedNil(b bool)`

 SetBackupSupportedNil sets the value for BackupSupported to be an explicit nil

### UnsetBackupSupported
`func (o *ExchangeDatabaseInfo) UnsetBackupSupported()`

UnsetBackupSupported ensures that no value is present for BackupSupported, not even an explicit nil
### GetBackupUnsupportedReasons

`func (o *ExchangeDatabaseInfo) GetBackupUnsupportedReasons() []string`

GetBackupUnsupportedReasons returns the BackupUnsupportedReasons field if non-nil, zero value otherwise.

### GetBackupUnsupportedReasonsOk

`func (o *ExchangeDatabaseInfo) GetBackupUnsupportedReasonsOk() (*[]string, bool)`

GetBackupUnsupportedReasonsOk returns a tuple with the BackupUnsupportedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUnsupportedReasons

`func (o *ExchangeDatabaseInfo) SetBackupUnsupportedReasons(v []string)`

SetBackupUnsupportedReasons sets BackupUnsupportedReasons field to given value.

### HasBackupUnsupportedReasons

`func (o *ExchangeDatabaseInfo) HasBackupUnsupportedReasons() bool`

HasBackupUnsupportedReasons returns a boolean if a field has been set.

### SetBackupUnsupportedReasonsNil

`func (o *ExchangeDatabaseInfo) SetBackupUnsupportedReasonsNil(b bool)`

 SetBackupUnsupportedReasonsNil sets the value for BackupUnsupportedReasons to be an explicit nil

### UnsetBackupUnsupportedReasons
`func (o *ExchangeDatabaseInfo) UnsetBackupUnsupportedReasons()`

UnsetBackupUnsupportedReasons ensures that no value is present for BackupUnsupportedReasons, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *ExchangeDatabaseInfo) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *ExchangeDatabaseInfo) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *ExchangeDatabaseInfo) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *ExchangeDatabaseInfo) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *ExchangeDatabaseInfo) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *ExchangeDatabaseInfo) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetDatabaseState

`func (o *ExchangeDatabaseInfo) GetDatabaseState() string`

GetDatabaseState returns the DatabaseState field if non-nil, zero value otherwise.

### GetDatabaseStateOk

`func (o *ExchangeDatabaseInfo) GetDatabaseStateOk() (*string, bool)`

GetDatabaseStateOk returns a tuple with the DatabaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseState

`func (o *ExchangeDatabaseInfo) SetDatabaseState(v string)`

SetDatabaseState sets DatabaseState field to given value.

### HasDatabaseState

`func (o *ExchangeDatabaseInfo) HasDatabaseState() bool`

HasDatabaseState returns a boolean if a field has been set.

### SetDatabaseStateNil

`func (o *ExchangeDatabaseInfo) SetDatabaseStateNil(b bool)`

 SetDatabaseStateNil sets the value for DatabaseState to be an explicit nil

### UnsetDatabaseState
`func (o *ExchangeDatabaseInfo) UnsetDatabaseState()`

UnsetDatabaseState ensures that no value is present for DatabaseState, not even an explicit nil
### GetDbSizeBytes

`func (o *ExchangeDatabaseInfo) GetDbSizeBytes() int64`

GetDbSizeBytes returns the DbSizeBytes field if non-nil, zero value otherwise.

### GetDbSizeBytesOk

`func (o *ExchangeDatabaseInfo) GetDbSizeBytesOk() (*int64, bool)`

GetDbSizeBytesOk returns a tuple with the DbSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbSizeBytes

`func (o *ExchangeDatabaseInfo) SetDbSizeBytes(v int64)`

SetDbSizeBytes sets DbSizeBytes field to given value.

### HasDbSizeBytes

`func (o *ExchangeDatabaseInfo) HasDbSizeBytes() bool`

HasDbSizeBytes returns a boolean if a field has been set.

### SetDbSizeBytesNil

`func (o *ExchangeDatabaseInfo) SetDbSizeBytesNil(b bool)`

 SetDbSizeBytesNil sets the value for DbSizeBytes to be an explicit nil

### UnsetDbSizeBytes
`func (o *ExchangeDatabaseInfo) UnsetDbSizeBytes()`

UnsetDbSizeBytes ensures that no value is present for DbSizeBytes, not even an explicit nil
### GetDbguid

`func (o *ExchangeDatabaseInfo) GetDbguid() string`

GetDbguid returns the Dbguid field if non-nil, zero value otherwise.

### GetDbguidOk

`func (o *ExchangeDatabaseInfo) GetDbguidOk() (*string, bool)`

GetDbguidOk returns a tuple with the Dbguid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbguid

`func (o *ExchangeDatabaseInfo) SetDbguid(v string)`

SetDbguid sets Dbguid field to given value.

### HasDbguid

`func (o *ExchangeDatabaseInfo) HasDbguid() bool`

HasDbguid returns a boolean if a field has been set.

### SetDbguidNil

`func (o *ExchangeDatabaseInfo) SetDbguidNil(b bool)`

 SetDbguidNil sets the value for Dbguid to be an explicit nil

### UnsetDbguid
`func (o *ExchangeDatabaseInfo) UnsetDbguid()`

UnsetDbguid ensures that no value is present for Dbguid, not even an explicit nil
### GetName

`func (o *ExchangeDatabaseInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExchangeDatabaseInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExchangeDatabaseInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExchangeDatabaseInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExchangeDatabaseInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExchangeDatabaseInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *ExchangeDatabaseInfo) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ExchangeDatabaseInfo) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ExchangeDatabaseInfo) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ExchangeDatabaseInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *ExchangeDatabaseInfo) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *ExchangeDatabaseInfo) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetUtcOffsetMin

`func (o *ExchangeDatabaseInfo) GetUtcOffsetMin() int64`

GetUtcOffsetMin returns the UtcOffsetMin field if non-nil, zero value otherwise.

### GetUtcOffsetMinOk

`func (o *ExchangeDatabaseInfo) GetUtcOffsetMinOk() (*int64, bool)`

GetUtcOffsetMinOk returns a tuple with the UtcOffsetMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffsetMin

`func (o *ExchangeDatabaseInfo) SetUtcOffsetMin(v int64)`

SetUtcOffsetMin sets UtcOffsetMin field to given value.

### HasUtcOffsetMin

`func (o *ExchangeDatabaseInfo) HasUtcOffsetMin() bool`

HasUtcOffsetMin returns a boolean if a field has been set.

### SetUtcOffsetMinNil

`func (o *ExchangeDatabaseInfo) SetUtcOffsetMinNil(b bool)`

 SetUtcOffsetMinNil sets the value for UtcOffsetMin to be an explicit nil

### UnsetUtcOffsetMin
`func (o *ExchangeDatabaseInfo) UnsetUtcOffsetMin()`

UnsetUtcOffsetMin ensures that no value is present for UtcOffsetMin, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


