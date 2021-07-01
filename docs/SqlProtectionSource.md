# SqlProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAvailableForVssBackup** | Pointer to **NullableBool** | Specifies whether the database is marked as available for backup according to the SQL Server VSS writer. This may be false if either the state of the databases is not online, or if the VSS writer is not online. This field is set only for type &#39;kDatabase&#39;. | [optional] 
**CreatedTimestamp** | Pointer to **NullableString** | Specifies the time when the database was created. It is displayed in the timezone of the SQL server on which this database is running. | [optional] 
**DatabaseName** | Pointer to **NullableString** | Specifies the database name of the SQL Protection Source, if the type is database. | [optional] 
**DbAagEntityId** | Pointer to **NullableInt64** | Specifies the AAG entity id if the database is part of an AAG. This field is set only for type &#39;kDatabase&#39;. | [optional] 
**DbAagName** | Pointer to **NullableString** | Specifies the name of the AAG if the database is part of an AAG. This field is set only for type &#39;kDatabase&#39;. | [optional] 
**DbCompatibilityLevel** | Pointer to **NullableInt64** | Specifies the versions of SQL server that the database is compatible with. | [optional] 
**DbFileGroups** | Pointer to **[]string** | Specifies the information about the set of file groups for this db on the host. This is only set if the type is kDatabase. | [optional] 
**DbFiles** | Pointer to [**[]DbFileInfo**](DbFileInfo.md) | Specifies the last known information about the set of database files on the host. This field is set only for type &#39;kDatabase&#39;. | [optional] 
**DbOwnerUsername** | Pointer to **NullableString** | Specifies the name of the database owner. | [optional] 
**Id** | Pointer to [**SqlSourceId**](SqlSourceId.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies the instance name of the SQL Protection Source | [optional] 
**OwnerId** | Pointer to **NullableInt64** | Specifies the id of the container VM for the SQL Protection Source. | [optional] 
**RecoveryModel** | Pointer to **NullableString** | Specifies the Recovery Model for the database in SQL environment. Only meaningful for the &#39;kDatabase&#39; SQL Protection Source. Specifies the Recovery Model set for the Microsoft SQL Server. &#39;kSimpleRecoveryModel&#39; indicates the Simple SQL Recovery Model which does not utilize log backups. &#39;kFullRecoveryModel&#39; indicates the Full SQL Recovery Model which requires log backups and allows recovery to a single point in time. &#39;kBulkLoggedRecoveryModel&#39; indicates the Bulk Logged SQL Recovery Model which requires log backups and allows high-performance bulk copy operations. | [optional] 
**SqlServerDbState** | Pointer to **NullableString** | The state of the database as returned by SQL Server. Indicates the state of the database. The values correspond to the &#39;state&#39; field in the system table sys.databases. See https://goo.gl/P66XqM. &#39;kOnline&#39; indicates that database is in online state. &#39;kRestoring&#39; indicates that database is in restore state. &#39;kRecovering&#39; indicates that database is in recovery state. &#39;kRecoveryPending&#39; indicates that database recovery is in pending state. &#39;kSuspect&#39; indicates that primary filegroup is suspect and may be damaged. &#39;kEmergency&#39; indicates that manually forced emergency state. &#39;kOffline&#39; indicates that database is in offline state. &#39;kCopying&#39; indicates that database is in copying state. &#39;kOfflineSecondary&#39; indicates that secondary database is in offline state. | [optional] 
**SqlServerInstanceVersion** | Pointer to [**SQLServerInstanceVersion**](SQLServerInstanceVersion.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed Object in a SQL Protection Source. Examples of SQL Objects include &#39;kInstance&#39; and &#39;kDatabase&#39;. &#39;kInstance&#39; indicates that SQL server instance is being protected. &#39;kDatabase&#39; indicates that SQL server database is being protected. &#39;kAAG&#39; indicates that SQL AAG (AlwaysOn Availability Group) is being protected. &#39;kAAGRootContainer&#39; indicates that SQL AAG&#39;s root container is being protected. &#39;kRootContainer&#39; indicates root container for SQL sources. | [optional] 

## Methods

### NewSqlProtectionSource

`func NewSqlProtectionSource() *SqlProtectionSource`

NewSqlProtectionSource instantiates a new SqlProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlProtectionSourceWithDefaults

`func NewSqlProtectionSourceWithDefaults() *SqlProtectionSource`

NewSqlProtectionSourceWithDefaults instantiates a new SqlProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAvailableForVssBackup

`func (o *SqlProtectionSource) GetIsAvailableForVssBackup() bool`

GetIsAvailableForVssBackup returns the IsAvailableForVssBackup field if non-nil, zero value otherwise.

### GetIsAvailableForVssBackupOk

`func (o *SqlProtectionSource) GetIsAvailableForVssBackupOk() (*bool, bool)`

GetIsAvailableForVssBackupOk returns a tuple with the IsAvailableForVssBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableForVssBackup

`func (o *SqlProtectionSource) SetIsAvailableForVssBackup(v bool)`

SetIsAvailableForVssBackup sets IsAvailableForVssBackup field to given value.

### HasIsAvailableForVssBackup

`func (o *SqlProtectionSource) HasIsAvailableForVssBackup() bool`

HasIsAvailableForVssBackup returns a boolean if a field has been set.

### SetIsAvailableForVssBackupNil

`func (o *SqlProtectionSource) SetIsAvailableForVssBackupNil(b bool)`

 SetIsAvailableForVssBackupNil sets the value for IsAvailableForVssBackup to be an explicit nil

### UnsetIsAvailableForVssBackup
`func (o *SqlProtectionSource) UnsetIsAvailableForVssBackup()`

UnsetIsAvailableForVssBackup ensures that no value is present for IsAvailableForVssBackup, not even an explicit nil
### GetCreatedTimestamp

`func (o *SqlProtectionSource) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *SqlProtectionSource) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *SqlProtectionSource) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *SqlProtectionSource) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### SetCreatedTimestampNil

`func (o *SqlProtectionSource) SetCreatedTimestampNil(b bool)`

 SetCreatedTimestampNil sets the value for CreatedTimestamp to be an explicit nil

### UnsetCreatedTimestamp
`func (o *SqlProtectionSource) UnsetCreatedTimestamp()`

UnsetCreatedTimestamp ensures that no value is present for CreatedTimestamp, not even an explicit nil
### GetDatabaseName

`func (o *SqlProtectionSource) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *SqlProtectionSource) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *SqlProtectionSource) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *SqlProtectionSource) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *SqlProtectionSource) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *SqlProtectionSource) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetDbAagEntityId

`func (o *SqlProtectionSource) GetDbAagEntityId() int64`

GetDbAagEntityId returns the DbAagEntityId field if non-nil, zero value otherwise.

### GetDbAagEntityIdOk

`func (o *SqlProtectionSource) GetDbAagEntityIdOk() (*int64, bool)`

GetDbAagEntityIdOk returns a tuple with the DbAagEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbAagEntityId

`func (o *SqlProtectionSource) SetDbAagEntityId(v int64)`

SetDbAagEntityId sets DbAagEntityId field to given value.

### HasDbAagEntityId

`func (o *SqlProtectionSource) HasDbAagEntityId() bool`

HasDbAagEntityId returns a boolean if a field has been set.

### SetDbAagEntityIdNil

`func (o *SqlProtectionSource) SetDbAagEntityIdNil(b bool)`

 SetDbAagEntityIdNil sets the value for DbAagEntityId to be an explicit nil

### UnsetDbAagEntityId
`func (o *SqlProtectionSource) UnsetDbAagEntityId()`

UnsetDbAagEntityId ensures that no value is present for DbAagEntityId, not even an explicit nil
### GetDbAagName

`func (o *SqlProtectionSource) GetDbAagName() string`

GetDbAagName returns the DbAagName field if non-nil, zero value otherwise.

### GetDbAagNameOk

`func (o *SqlProtectionSource) GetDbAagNameOk() (*string, bool)`

GetDbAagNameOk returns a tuple with the DbAagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbAagName

`func (o *SqlProtectionSource) SetDbAagName(v string)`

SetDbAagName sets DbAagName field to given value.

### HasDbAagName

`func (o *SqlProtectionSource) HasDbAagName() bool`

HasDbAagName returns a boolean if a field has been set.

### SetDbAagNameNil

`func (o *SqlProtectionSource) SetDbAagNameNil(b bool)`

 SetDbAagNameNil sets the value for DbAagName to be an explicit nil

### UnsetDbAagName
`func (o *SqlProtectionSource) UnsetDbAagName()`

UnsetDbAagName ensures that no value is present for DbAagName, not even an explicit nil
### GetDbCompatibilityLevel

`func (o *SqlProtectionSource) GetDbCompatibilityLevel() int64`

GetDbCompatibilityLevel returns the DbCompatibilityLevel field if non-nil, zero value otherwise.

### GetDbCompatibilityLevelOk

`func (o *SqlProtectionSource) GetDbCompatibilityLevelOk() (*int64, bool)`

GetDbCompatibilityLevelOk returns a tuple with the DbCompatibilityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbCompatibilityLevel

`func (o *SqlProtectionSource) SetDbCompatibilityLevel(v int64)`

SetDbCompatibilityLevel sets DbCompatibilityLevel field to given value.

### HasDbCompatibilityLevel

`func (o *SqlProtectionSource) HasDbCompatibilityLevel() bool`

HasDbCompatibilityLevel returns a boolean if a field has been set.

### SetDbCompatibilityLevelNil

`func (o *SqlProtectionSource) SetDbCompatibilityLevelNil(b bool)`

 SetDbCompatibilityLevelNil sets the value for DbCompatibilityLevel to be an explicit nil

### UnsetDbCompatibilityLevel
`func (o *SqlProtectionSource) UnsetDbCompatibilityLevel()`

UnsetDbCompatibilityLevel ensures that no value is present for DbCompatibilityLevel, not even an explicit nil
### GetDbFileGroups

`func (o *SqlProtectionSource) GetDbFileGroups() []string`

GetDbFileGroups returns the DbFileGroups field if non-nil, zero value otherwise.

### GetDbFileGroupsOk

`func (o *SqlProtectionSource) GetDbFileGroupsOk() (*[]string, bool)`

GetDbFileGroupsOk returns a tuple with the DbFileGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbFileGroups

`func (o *SqlProtectionSource) SetDbFileGroups(v []string)`

SetDbFileGroups sets DbFileGroups field to given value.

### HasDbFileGroups

`func (o *SqlProtectionSource) HasDbFileGroups() bool`

HasDbFileGroups returns a boolean if a field has been set.

### SetDbFileGroupsNil

`func (o *SqlProtectionSource) SetDbFileGroupsNil(b bool)`

 SetDbFileGroupsNil sets the value for DbFileGroups to be an explicit nil

### UnsetDbFileGroups
`func (o *SqlProtectionSource) UnsetDbFileGroups()`

UnsetDbFileGroups ensures that no value is present for DbFileGroups, not even an explicit nil
### GetDbFiles

`func (o *SqlProtectionSource) GetDbFiles() []DbFileInfo`

GetDbFiles returns the DbFiles field if non-nil, zero value otherwise.

### GetDbFilesOk

`func (o *SqlProtectionSource) GetDbFilesOk() (*[]DbFileInfo, bool)`

GetDbFilesOk returns a tuple with the DbFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbFiles

`func (o *SqlProtectionSource) SetDbFiles(v []DbFileInfo)`

SetDbFiles sets DbFiles field to given value.

### HasDbFiles

`func (o *SqlProtectionSource) HasDbFiles() bool`

HasDbFiles returns a boolean if a field has been set.

### SetDbFilesNil

`func (o *SqlProtectionSource) SetDbFilesNil(b bool)`

 SetDbFilesNil sets the value for DbFiles to be an explicit nil

### UnsetDbFiles
`func (o *SqlProtectionSource) UnsetDbFiles()`

UnsetDbFiles ensures that no value is present for DbFiles, not even an explicit nil
### GetDbOwnerUsername

`func (o *SqlProtectionSource) GetDbOwnerUsername() string`

GetDbOwnerUsername returns the DbOwnerUsername field if non-nil, zero value otherwise.

### GetDbOwnerUsernameOk

`func (o *SqlProtectionSource) GetDbOwnerUsernameOk() (*string, bool)`

GetDbOwnerUsernameOk returns a tuple with the DbOwnerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbOwnerUsername

`func (o *SqlProtectionSource) SetDbOwnerUsername(v string)`

SetDbOwnerUsername sets DbOwnerUsername field to given value.

### HasDbOwnerUsername

`func (o *SqlProtectionSource) HasDbOwnerUsername() bool`

HasDbOwnerUsername returns a boolean if a field has been set.

### SetDbOwnerUsernameNil

`func (o *SqlProtectionSource) SetDbOwnerUsernameNil(b bool)`

 SetDbOwnerUsernameNil sets the value for DbOwnerUsername to be an explicit nil

### UnsetDbOwnerUsername
`func (o *SqlProtectionSource) UnsetDbOwnerUsername()`

UnsetDbOwnerUsername ensures that no value is present for DbOwnerUsername, not even an explicit nil
### GetId

`func (o *SqlProtectionSource) GetId() SqlSourceId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SqlProtectionSource) GetIdOk() (*SqlSourceId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SqlProtectionSource) SetId(v SqlSourceId)`

SetId sets Id field to given value.

### HasId

`func (o *SqlProtectionSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SqlProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SqlProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SqlProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SqlProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *SqlProtectionSource) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SqlProtectionSource) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SqlProtectionSource) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SqlProtectionSource) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *SqlProtectionSource) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *SqlProtectionSource) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetRecoveryModel

`func (o *SqlProtectionSource) GetRecoveryModel() string`

GetRecoveryModel returns the RecoveryModel field if non-nil, zero value otherwise.

### GetRecoveryModelOk

`func (o *SqlProtectionSource) GetRecoveryModelOk() (*string, bool)`

GetRecoveryModelOk returns a tuple with the RecoveryModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryModel

`func (o *SqlProtectionSource) SetRecoveryModel(v string)`

SetRecoveryModel sets RecoveryModel field to given value.

### HasRecoveryModel

`func (o *SqlProtectionSource) HasRecoveryModel() bool`

HasRecoveryModel returns a boolean if a field has been set.

### SetRecoveryModelNil

`func (o *SqlProtectionSource) SetRecoveryModelNil(b bool)`

 SetRecoveryModelNil sets the value for RecoveryModel to be an explicit nil

### UnsetRecoveryModel
`func (o *SqlProtectionSource) UnsetRecoveryModel()`

UnsetRecoveryModel ensures that no value is present for RecoveryModel, not even an explicit nil
### GetSqlServerDbState

`func (o *SqlProtectionSource) GetSqlServerDbState() string`

GetSqlServerDbState returns the SqlServerDbState field if non-nil, zero value otherwise.

### GetSqlServerDbStateOk

`func (o *SqlProtectionSource) GetSqlServerDbStateOk() (*string, bool)`

GetSqlServerDbStateOk returns a tuple with the SqlServerDbState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlServerDbState

`func (o *SqlProtectionSource) SetSqlServerDbState(v string)`

SetSqlServerDbState sets SqlServerDbState field to given value.

### HasSqlServerDbState

`func (o *SqlProtectionSource) HasSqlServerDbState() bool`

HasSqlServerDbState returns a boolean if a field has been set.

### SetSqlServerDbStateNil

`func (o *SqlProtectionSource) SetSqlServerDbStateNil(b bool)`

 SetSqlServerDbStateNil sets the value for SqlServerDbState to be an explicit nil

### UnsetSqlServerDbState
`func (o *SqlProtectionSource) UnsetSqlServerDbState()`

UnsetSqlServerDbState ensures that no value is present for SqlServerDbState, not even an explicit nil
### GetSqlServerInstanceVersion

`func (o *SqlProtectionSource) GetSqlServerInstanceVersion() SQLServerInstanceVersion`

GetSqlServerInstanceVersion returns the SqlServerInstanceVersion field if non-nil, zero value otherwise.

### GetSqlServerInstanceVersionOk

`func (o *SqlProtectionSource) GetSqlServerInstanceVersionOk() (*SQLServerInstanceVersion, bool)`

GetSqlServerInstanceVersionOk returns a tuple with the SqlServerInstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlServerInstanceVersion

`func (o *SqlProtectionSource) SetSqlServerInstanceVersion(v SQLServerInstanceVersion)`

SetSqlServerInstanceVersion sets SqlServerInstanceVersion field to given value.

### HasSqlServerInstanceVersion

`func (o *SqlProtectionSource) HasSqlServerInstanceVersion() bool`

HasSqlServerInstanceVersion returns a boolean if a field has been set.

### GetType

`func (o *SqlProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SqlProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SqlProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SqlProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SqlProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SqlProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


