# OracleProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveLogEnabled** | Pointer to **NullableBool** | Specifies whether the database is running in ARCHIVELOG mode. It enables the redo of log files into archived redo log files. | [optional] 
**BctEnabled** | Pointer to **NullableBool** | Specifies whether the Block Change Tracking is enabled. BCT improves the performance of incremental backups by recording changed blocks into the block change tracking file. RMAN then uses this file to identify changed blocks to be backed up. | [optional] 
**ContainerDatabaseInfo** | Pointer to [**OracleContainerDatabaseInfo**](OracleContainerDatabaseInfo.md) |  | [optional] 
**DataGuardInfo** | Pointer to [**OracleDataGuardInfo**](OracleDataGuardInfo.md) |  | [optional] 
**DatabaseUniqueName** | Pointer to **NullableString** | Specifies the unique name of the Oracle entity. | [optional] 
**DbType** | Pointer to **NullableString** | Specifies the type of the database in Oracle Protection Source. &#39;kRACDatabase&#39; indicates the database is a RAC DB. &#39;kSingleInstance&#39; indicates that the database is single instance. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the Oracle DB Domain. | [optional] 
**FraSize** | Pointer to **NullableInt64** | Specifies Flash/Fast Recovery area size for the current DB entity. | [optional] 
**Hosts** | Pointer to [**[]OracleHost**](OracleHost.md) | Specifies the list of hosts for the current DB entity. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the instance name of the Oracle entity. | [optional] 
**OwnerId** | Pointer to **NullableInt64** | Specifies the entity id of the owner entity (such as a VM). This is only set if type is kDatabase. | [optional] 
**SgaTargetSize** | Pointer to **NullableString** | Specifies System Global Area size for the current DB entity. A system global area (SGA) is a group of shared memory structures that contain data and control information for one Oracle database. | [optional] 
**SharedPoolSize** | Pointer to **NullableString** | Specifies Shared Pool Size for the current DB entity. | [optional] 
**Size** | Pointer to **NullableInt64** | Specifies database size. | [optional] 
**TdeEncryptedTsCount** | Pointer to **NullableInt64** | Specifies the number of TDE encrypted tablespaces found in the database. | [optional] 
**TempFilesCount** | Pointer to **NullableInt64** | Specifies number of temporary files for the current DB entity. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed Object in Oracle Protection Source. &#39;kRACRootContainer&#39; indicates the entity is a root container to an Oracle Real Application clusters(Oracle RAC). &#39;kRootContainer&#39; indicates the entity is a root container to an Oracle standalone server. &#39;kHost&#39; indicates the entity is an Oracle host. &#39;kDatabase&#39; indicates the entity is an Oracle Database. &#39;kTableSpace&#39; indicates the entity is an Oracle table space. &#39;kTable&#39; indicates the entity is an Oracle table. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the Oracle entity. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the Oracle database instance version. | [optional] 

## Methods

### NewOracleProtectionSource

`func NewOracleProtectionSource() *OracleProtectionSource`

NewOracleProtectionSource instantiates a new OracleProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleProtectionSourceWithDefaults

`func NewOracleProtectionSourceWithDefaults() *OracleProtectionSource`

NewOracleProtectionSourceWithDefaults instantiates a new OracleProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveLogEnabled

`func (o *OracleProtectionSource) GetArchiveLogEnabled() bool`

GetArchiveLogEnabled returns the ArchiveLogEnabled field if non-nil, zero value otherwise.

### GetArchiveLogEnabledOk

`func (o *OracleProtectionSource) GetArchiveLogEnabledOk() (*bool, bool)`

GetArchiveLogEnabledOk returns a tuple with the ArchiveLogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogEnabled

`func (o *OracleProtectionSource) SetArchiveLogEnabled(v bool)`

SetArchiveLogEnabled sets ArchiveLogEnabled field to given value.

### HasArchiveLogEnabled

`func (o *OracleProtectionSource) HasArchiveLogEnabled() bool`

HasArchiveLogEnabled returns a boolean if a field has been set.

### SetArchiveLogEnabledNil

`func (o *OracleProtectionSource) SetArchiveLogEnabledNil(b bool)`

 SetArchiveLogEnabledNil sets the value for ArchiveLogEnabled to be an explicit nil

### UnsetArchiveLogEnabled
`func (o *OracleProtectionSource) UnsetArchiveLogEnabled()`

UnsetArchiveLogEnabled ensures that no value is present for ArchiveLogEnabled, not even an explicit nil
### GetBctEnabled

`func (o *OracleProtectionSource) GetBctEnabled() bool`

GetBctEnabled returns the BctEnabled field if non-nil, zero value otherwise.

### GetBctEnabledOk

`func (o *OracleProtectionSource) GetBctEnabledOk() (*bool, bool)`

GetBctEnabledOk returns a tuple with the BctEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBctEnabled

`func (o *OracleProtectionSource) SetBctEnabled(v bool)`

SetBctEnabled sets BctEnabled field to given value.

### HasBctEnabled

`func (o *OracleProtectionSource) HasBctEnabled() bool`

HasBctEnabled returns a boolean if a field has been set.

### SetBctEnabledNil

`func (o *OracleProtectionSource) SetBctEnabledNil(b bool)`

 SetBctEnabledNil sets the value for BctEnabled to be an explicit nil

### UnsetBctEnabled
`func (o *OracleProtectionSource) UnsetBctEnabled()`

UnsetBctEnabled ensures that no value is present for BctEnabled, not even an explicit nil
### GetContainerDatabaseInfo

`func (o *OracleProtectionSource) GetContainerDatabaseInfo() OracleContainerDatabaseInfo`

GetContainerDatabaseInfo returns the ContainerDatabaseInfo field if non-nil, zero value otherwise.

### GetContainerDatabaseInfoOk

`func (o *OracleProtectionSource) GetContainerDatabaseInfoOk() (*OracleContainerDatabaseInfo, bool)`

GetContainerDatabaseInfoOk returns a tuple with the ContainerDatabaseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDatabaseInfo

`func (o *OracleProtectionSource) SetContainerDatabaseInfo(v OracleContainerDatabaseInfo)`

SetContainerDatabaseInfo sets ContainerDatabaseInfo field to given value.

### HasContainerDatabaseInfo

`func (o *OracleProtectionSource) HasContainerDatabaseInfo() bool`

HasContainerDatabaseInfo returns a boolean if a field has been set.

### GetDataGuardInfo

`func (o *OracleProtectionSource) GetDataGuardInfo() OracleDataGuardInfo`

GetDataGuardInfo returns the DataGuardInfo field if non-nil, zero value otherwise.

### GetDataGuardInfoOk

`func (o *OracleProtectionSource) GetDataGuardInfoOk() (*OracleDataGuardInfo, bool)`

GetDataGuardInfoOk returns a tuple with the DataGuardInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataGuardInfo

`func (o *OracleProtectionSource) SetDataGuardInfo(v OracleDataGuardInfo)`

SetDataGuardInfo sets DataGuardInfo field to given value.

### HasDataGuardInfo

`func (o *OracleProtectionSource) HasDataGuardInfo() bool`

HasDataGuardInfo returns a boolean if a field has been set.

### GetDatabaseUniqueName

`func (o *OracleProtectionSource) GetDatabaseUniqueName() string`

GetDatabaseUniqueName returns the DatabaseUniqueName field if non-nil, zero value otherwise.

### GetDatabaseUniqueNameOk

`func (o *OracleProtectionSource) GetDatabaseUniqueNameOk() (*string, bool)`

GetDatabaseUniqueNameOk returns a tuple with the DatabaseUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUniqueName

`func (o *OracleProtectionSource) SetDatabaseUniqueName(v string)`

SetDatabaseUniqueName sets DatabaseUniqueName field to given value.

### HasDatabaseUniqueName

`func (o *OracleProtectionSource) HasDatabaseUniqueName() bool`

HasDatabaseUniqueName returns a boolean if a field has been set.

### SetDatabaseUniqueNameNil

`func (o *OracleProtectionSource) SetDatabaseUniqueNameNil(b bool)`

 SetDatabaseUniqueNameNil sets the value for DatabaseUniqueName to be an explicit nil

### UnsetDatabaseUniqueName
`func (o *OracleProtectionSource) UnsetDatabaseUniqueName()`

UnsetDatabaseUniqueName ensures that no value is present for DatabaseUniqueName, not even an explicit nil
### GetDbType

`func (o *OracleProtectionSource) GetDbType() string`

GetDbType returns the DbType field if non-nil, zero value otherwise.

### GetDbTypeOk

`func (o *OracleProtectionSource) GetDbTypeOk() (*string, bool)`

GetDbTypeOk returns a tuple with the DbType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbType

`func (o *OracleProtectionSource) SetDbType(v string)`

SetDbType sets DbType field to given value.

### HasDbType

`func (o *OracleProtectionSource) HasDbType() bool`

HasDbType returns a boolean if a field has been set.

### SetDbTypeNil

`func (o *OracleProtectionSource) SetDbTypeNil(b bool)`

 SetDbTypeNil sets the value for DbType to be an explicit nil

### UnsetDbType
`func (o *OracleProtectionSource) UnsetDbType()`

UnsetDbType ensures that no value is present for DbType, not even an explicit nil
### GetDomain

`func (o *OracleProtectionSource) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *OracleProtectionSource) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *OracleProtectionSource) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *OracleProtectionSource) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *OracleProtectionSource) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *OracleProtectionSource) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetFraSize

`func (o *OracleProtectionSource) GetFraSize() int64`

GetFraSize returns the FraSize field if non-nil, zero value otherwise.

### GetFraSizeOk

`func (o *OracleProtectionSource) GetFraSizeOk() (*int64, bool)`

GetFraSizeOk returns a tuple with the FraSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraSize

`func (o *OracleProtectionSource) SetFraSize(v int64)`

SetFraSize sets FraSize field to given value.

### HasFraSize

`func (o *OracleProtectionSource) HasFraSize() bool`

HasFraSize returns a boolean if a field has been set.

### SetFraSizeNil

`func (o *OracleProtectionSource) SetFraSizeNil(b bool)`

 SetFraSizeNil sets the value for FraSize to be an explicit nil

### UnsetFraSize
`func (o *OracleProtectionSource) UnsetFraSize()`

UnsetFraSize ensures that no value is present for FraSize, not even an explicit nil
### GetHosts

`func (o *OracleProtectionSource) GetHosts() []OracleHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *OracleProtectionSource) GetHostsOk() (*[]OracleHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *OracleProtectionSource) SetHosts(v []OracleHost)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *OracleProtectionSource) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### SetHostsNil

`func (o *OracleProtectionSource) SetHostsNil(b bool)`

 SetHostsNil sets the value for Hosts to be an explicit nil

### UnsetHosts
`func (o *OracleProtectionSource) UnsetHosts()`

UnsetHosts ensures that no value is present for Hosts, not even an explicit nil
### GetName

`func (o *OracleProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OracleProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OracleProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OracleProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OracleProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OracleProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOwnerId

`func (o *OracleProtectionSource) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *OracleProtectionSource) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *OracleProtectionSource) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *OracleProtectionSource) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *OracleProtectionSource) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *OracleProtectionSource) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetSgaTargetSize

`func (o *OracleProtectionSource) GetSgaTargetSize() string`

GetSgaTargetSize returns the SgaTargetSize field if non-nil, zero value otherwise.

### GetSgaTargetSizeOk

`func (o *OracleProtectionSource) GetSgaTargetSizeOk() (*string, bool)`

GetSgaTargetSizeOk returns a tuple with the SgaTargetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgaTargetSize

`func (o *OracleProtectionSource) SetSgaTargetSize(v string)`

SetSgaTargetSize sets SgaTargetSize field to given value.

### HasSgaTargetSize

`func (o *OracleProtectionSource) HasSgaTargetSize() bool`

HasSgaTargetSize returns a boolean if a field has been set.

### SetSgaTargetSizeNil

`func (o *OracleProtectionSource) SetSgaTargetSizeNil(b bool)`

 SetSgaTargetSizeNil sets the value for SgaTargetSize to be an explicit nil

### UnsetSgaTargetSize
`func (o *OracleProtectionSource) UnsetSgaTargetSize()`

UnsetSgaTargetSize ensures that no value is present for SgaTargetSize, not even an explicit nil
### GetSharedPoolSize

`func (o *OracleProtectionSource) GetSharedPoolSize() string`

GetSharedPoolSize returns the SharedPoolSize field if non-nil, zero value otherwise.

### GetSharedPoolSizeOk

`func (o *OracleProtectionSource) GetSharedPoolSizeOk() (*string, bool)`

GetSharedPoolSizeOk returns a tuple with the SharedPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPoolSize

`func (o *OracleProtectionSource) SetSharedPoolSize(v string)`

SetSharedPoolSize sets SharedPoolSize field to given value.

### HasSharedPoolSize

`func (o *OracleProtectionSource) HasSharedPoolSize() bool`

HasSharedPoolSize returns a boolean if a field has been set.

### SetSharedPoolSizeNil

`func (o *OracleProtectionSource) SetSharedPoolSizeNil(b bool)`

 SetSharedPoolSizeNil sets the value for SharedPoolSize to be an explicit nil

### UnsetSharedPoolSize
`func (o *OracleProtectionSource) UnsetSharedPoolSize()`

UnsetSharedPoolSize ensures that no value is present for SharedPoolSize, not even an explicit nil
### GetSize

`func (o *OracleProtectionSource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OracleProtectionSource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OracleProtectionSource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *OracleProtectionSource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *OracleProtectionSource) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *OracleProtectionSource) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetTdeEncryptedTsCount

`func (o *OracleProtectionSource) GetTdeEncryptedTsCount() int64`

GetTdeEncryptedTsCount returns the TdeEncryptedTsCount field if non-nil, zero value otherwise.

### GetTdeEncryptedTsCountOk

`func (o *OracleProtectionSource) GetTdeEncryptedTsCountOk() (*int64, bool)`

GetTdeEncryptedTsCountOk returns a tuple with the TdeEncryptedTsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeEncryptedTsCount

`func (o *OracleProtectionSource) SetTdeEncryptedTsCount(v int64)`

SetTdeEncryptedTsCount sets TdeEncryptedTsCount field to given value.

### HasTdeEncryptedTsCount

`func (o *OracleProtectionSource) HasTdeEncryptedTsCount() bool`

HasTdeEncryptedTsCount returns a boolean if a field has been set.

### SetTdeEncryptedTsCountNil

`func (o *OracleProtectionSource) SetTdeEncryptedTsCountNil(b bool)`

 SetTdeEncryptedTsCountNil sets the value for TdeEncryptedTsCount to be an explicit nil

### UnsetTdeEncryptedTsCount
`func (o *OracleProtectionSource) UnsetTdeEncryptedTsCount()`

UnsetTdeEncryptedTsCount ensures that no value is present for TdeEncryptedTsCount, not even an explicit nil
### GetTempFilesCount

`func (o *OracleProtectionSource) GetTempFilesCount() int64`

GetTempFilesCount returns the TempFilesCount field if non-nil, zero value otherwise.

### GetTempFilesCountOk

`func (o *OracleProtectionSource) GetTempFilesCountOk() (*int64, bool)`

GetTempFilesCountOk returns a tuple with the TempFilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempFilesCount

`func (o *OracleProtectionSource) SetTempFilesCount(v int64)`

SetTempFilesCount sets TempFilesCount field to given value.

### HasTempFilesCount

`func (o *OracleProtectionSource) HasTempFilesCount() bool`

HasTempFilesCount returns a boolean if a field has been set.

### SetTempFilesCountNil

`func (o *OracleProtectionSource) SetTempFilesCountNil(b bool)`

 SetTempFilesCountNil sets the value for TempFilesCount to be an explicit nil

### UnsetTempFilesCount
`func (o *OracleProtectionSource) UnsetTempFilesCount()`

UnsetTempFilesCount ensures that no value is present for TempFilesCount, not even an explicit nil
### GetType

`func (o *OracleProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OracleProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OracleProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OracleProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OracleProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OracleProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *OracleProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OracleProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OracleProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OracleProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *OracleProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *OracleProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *OracleProtectionSource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OracleProtectionSource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OracleProtectionSource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OracleProtectionSource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *OracleProtectionSource) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *OracleProtectionSource) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


