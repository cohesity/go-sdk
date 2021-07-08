# OracleDBChannelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivelogKeepDays** | Pointer to **NullableInt32** | Archived log deletion policy for this unique Oracle database. 1: keep archived log forever 0: delete archived log immediately n&gt;0: delete archived log after n days | [optional] 
**DbUniqueName** | Pointer to **NullableString** | The unique name of the database. | [optional] 
**DbUuid** | Pointer to **NullableString** | Database id, internal field, is filled by magneto master based on corresponding app entity id. | [optional] 
**EnableDgPrimaryBackup** | Pointer to **NullableBool** | If set to false, and if the DG database role is primary, we will not allow the backup of that database. | [optional] 
**HostInfoVec** | Pointer to [**[]OracleDBChannelInfoHostInfo**](OracleDBChannelInfoHostInfo.md) | Vector of Oracle hosts from which we are allowed to take the backup/restore. In case of RAC database it may be more than one. | [optional] 
**MaxNumHost** | Pointer to **NullableInt32** | Maximum number of hosts from which we are allowed to take backup/restore parallely. This will be less than or equal to host_info_vec_size. If this is less than host_info_vec_size we will choose max_num_host from host_info_vec and take backup/restore from this number of host. | [optional] 
**NumChannels** | Pointer to **NullableInt32** | The default number of channels to use per host per db. This value is used on all hosts unless host_info_vec.num_channels is specified for that host. Default value for num_channels will be calculated as minimum number of nodes in cohesity cluster, and 2 * number of cpu on Oracle host. Preference order for number of channels per host for given db is: 1. If user has specified host_info_vec.num_channels for host we will use that. 2. If user has not specified host_info_vec.num_channels but specified num_channels we will use this. 3. If user has neither specified host_info_vec.num_channels nor num_channels we will calculate default channels with above formula. | [optional] 
**RmanBackupType** | Pointer to **NullableInt32** | Type of Oracle RMAN backup rquested (i.e ImageCopy, BackupSets). | [optional] 

## Methods

### NewOracleDBChannelInfo

`func NewOracleDBChannelInfo() *OracleDBChannelInfo`

NewOracleDBChannelInfo instantiates a new OracleDBChannelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDBChannelInfoWithDefaults

`func NewOracleDBChannelInfoWithDefaults() *OracleDBChannelInfo`

NewOracleDBChannelInfoWithDefaults instantiates a new OracleDBChannelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivelogKeepDays

`func (o *OracleDBChannelInfo) GetArchivelogKeepDays() int32`

GetArchivelogKeepDays returns the ArchivelogKeepDays field if non-nil, zero value otherwise.

### GetArchivelogKeepDaysOk

`func (o *OracleDBChannelInfo) GetArchivelogKeepDaysOk() (*int32, bool)`

GetArchivelogKeepDaysOk returns a tuple with the ArchivelogKeepDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivelogKeepDays

`func (o *OracleDBChannelInfo) SetArchivelogKeepDays(v int32)`

SetArchivelogKeepDays sets ArchivelogKeepDays field to given value.

### HasArchivelogKeepDays

`func (o *OracleDBChannelInfo) HasArchivelogKeepDays() bool`

HasArchivelogKeepDays returns a boolean if a field has been set.

### SetArchivelogKeepDaysNil

`func (o *OracleDBChannelInfo) SetArchivelogKeepDaysNil(b bool)`

 SetArchivelogKeepDaysNil sets the value for ArchivelogKeepDays to be an explicit nil

### UnsetArchivelogKeepDays
`func (o *OracleDBChannelInfo) UnsetArchivelogKeepDays()`

UnsetArchivelogKeepDays ensures that no value is present for ArchivelogKeepDays, not even an explicit nil
### GetDbUniqueName

`func (o *OracleDBChannelInfo) GetDbUniqueName() string`

GetDbUniqueName returns the DbUniqueName field if non-nil, zero value otherwise.

### GetDbUniqueNameOk

`func (o *OracleDBChannelInfo) GetDbUniqueNameOk() (*string, bool)`

GetDbUniqueNameOk returns a tuple with the DbUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUniqueName

`func (o *OracleDBChannelInfo) SetDbUniqueName(v string)`

SetDbUniqueName sets DbUniqueName field to given value.

### HasDbUniqueName

`func (o *OracleDBChannelInfo) HasDbUniqueName() bool`

HasDbUniqueName returns a boolean if a field has been set.

### SetDbUniqueNameNil

`func (o *OracleDBChannelInfo) SetDbUniqueNameNil(b bool)`

 SetDbUniqueNameNil sets the value for DbUniqueName to be an explicit nil

### UnsetDbUniqueName
`func (o *OracleDBChannelInfo) UnsetDbUniqueName()`

UnsetDbUniqueName ensures that no value is present for DbUniqueName, not even an explicit nil
### GetDbUuid

`func (o *OracleDBChannelInfo) GetDbUuid() string`

GetDbUuid returns the DbUuid field if non-nil, zero value otherwise.

### GetDbUuidOk

`func (o *OracleDBChannelInfo) GetDbUuidOk() (*string, bool)`

GetDbUuidOk returns a tuple with the DbUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUuid

`func (o *OracleDBChannelInfo) SetDbUuid(v string)`

SetDbUuid sets DbUuid field to given value.

### HasDbUuid

`func (o *OracleDBChannelInfo) HasDbUuid() bool`

HasDbUuid returns a boolean if a field has been set.

### SetDbUuidNil

`func (o *OracleDBChannelInfo) SetDbUuidNil(b bool)`

 SetDbUuidNil sets the value for DbUuid to be an explicit nil

### UnsetDbUuid
`func (o *OracleDBChannelInfo) UnsetDbUuid()`

UnsetDbUuid ensures that no value is present for DbUuid, not even an explicit nil
### GetEnableDgPrimaryBackup

`func (o *OracleDBChannelInfo) GetEnableDgPrimaryBackup() bool`

GetEnableDgPrimaryBackup returns the EnableDgPrimaryBackup field if non-nil, zero value otherwise.

### GetEnableDgPrimaryBackupOk

`func (o *OracleDBChannelInfo) GetEnableDgPrimaryBackupOk() (*bool, bool)`

GetEnableDgPrimaryBackupOk returns a tuple with the EnableDgPrimaryBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDgPrimaryBackup

`func (o *OracleDBChannelInfo) SetEnableDgPrimaryBackup(v bool)`

SetEnableDgPrimaryBackup sets EnableDgPrimaryBackup field to given value.

### HasEnableDgPrimaryBackup

`func (o *OracleDBChannelInfo) HasEnableDgPrimaryBackup() bool`

HasEnableDgPrimaryBackup returns a boolean if a field has been set.

### SetEnableDgPrimaryBackupNil

`func (o *OracleDBChannelInfo) SetEnableDgPrimaryBackupNil(b bool)`

 SetEnableDgPrimaryBackupNil sets the value for EnableDgPrimaryBackup to be an explicit nil

### UnsetEnableDgPrimaryBackup
`func (o *OracleDBChannelInfo) UnsetEnableDgPrimaryBackup()`

UnsetEnableDgPrimaryBackup ensures that no value is present for EnableDgPrimaryBackup, not even an explicit nil
### GetHostInfoVec

`func (o *OracleDBChannelInfo) GetHostInfoVec() []OracleDBChannelInfoHostInfo`

GetHostInfoVec returns the HostInfoVec field if non-nil, zero value otherwise.

### GetHostInfoVecOk

`func (o *OracleDBChannelInfo) GetHostInfoVecOk() (*[]OracleDBChannelInfoHostInfo, bool)`

GetHostInfoVecOk returns a tuple with the HostInfoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfoVec

`func (o *OracleDBChannelInfo) SetHostInfoVec(v []OracleDBChannelInfoHostInfo)`

SetHostInfoVec sets HostInfoVec field to given value.

### HasHostInfoVec

`func (o *OracleDBChannelInfo) HasHostInfoVec() bool`

HasHostInfoVec returns a boolean if a field has been set.

### SetHostInfoVecNil

`func (o *OracleDBChannelInfo) SetHostInfoVecNil(b bool)`

 SetHostInfoVecNil sets the value for HostInfoVec to be an explicit nil

### UnsetHostInfoVec
`func (o *OracleDBChannelInfo) UnsetHostInfoVec()`

UnsetHostInfoVec ensures that no value is present for HostInfoVec, not even an explicit nil
### GetMaxNumHost

`func (o *OracleDBChannelInfo) GetMaxNumHost() int32`

GetMaxNumHost returns the MaxNumHost field if non-nil, zero value otherwise.

### GetMaxNumHostOk

`func (o *OracleDBChannelInfo) GetMaxNumHostOk() (*int32, bool)`

GetMaxNumHostOk returns a tuple with the MaxNumHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumHost

`func (o *OracleDBChannelInfo) SetMaxNumHost(v int32)`

SetMaxNumHost sets MaxNumHost field to given value.

### HasMaxNumHost

`func (o *OracleDBChannelInfo) HasMaxNumHost() bool`

HasMaxNumHost returns a boolean if a field has been set.

### SetMaxNumHostNil

`func (o *OracleDBChannelInfo) SetMaxNumHostNil(b bool)`

 SetMaxNumHostNil sets the value for MaxNumHost to be an explicit nil

### UnsetMaxNumHost
`func (o *OracleDBChannelInfo) UnsetMaxNumHost()`

UnsetMaxNumHost ensures that no value is present for MaxNumHost, not even an explicit nil
### GetNumChannels

`func (o *OracleDBChannelInfo) GetNumChannels() int32`

GetNumChannels returns the NumChannels field if non-nil, zero value otherwise.

### GetNumChannelsOk

`func (o *OracleDBChannelInfo) GetNumChannelsOk() (*int32, bool)`

GetNumChannelsOk returns a tuple with the NumChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumChannels

`func (o *OracleDBChannelInfo) SetNumChannels(v int32)`

SetNumChannels sets NumChannels field to given value.

### HasNumChannels

`func (o *OracleDBChannelInfo) HasNumChannels() bool`

HasNumChannels returns a boolean if a field has been set.

### SetNumChannelsNil

`func (o *OracleDBChannelInfo) SetNumChannelsNil(b bool)`

 SetNumChannelsNil sets the value for NumChannels to be an explicit nil

### UnsetNumChannels
`func (o *OracleDBChannelInfo) UnsetNumChannels()`

UnsetNumChannels ensures that no value is present for NumChannels, not even an explicit nil
### GetRmanBackupType

`func (o *OracleDBChannelInfo) GetRmanBackupType() int32`

GetRmanBackupType returns the RmanBackupType field if non-nil, zero value otherwise.

### GetRmanBackupTypeOk

`func (o *OracleDBChannelInfo) GetRmanBackupTypeOk() (*int32, bool)`

GetRmanBackupTypeOk returns a tuple with the RmanBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmanBackupType

`func (o *OracleDBChannelInfo) SetRmanBackupType(v int32)`

SetRmanBackupType sets RmanBackupType field to given value.

### HasRmanBackupType

`func (o *OracleDBChannelInfo) HasRmanBackupType() bool`

HasRmanBackupType returns a boolean if a field has been set.

### SetRmanBackupTypeNil

`func (o *OracleDBChannelInfo) SetRmanBackupTypeNil(b bool)`

 SetRmanBackupTypeNil sets the value for RmanBackupType to be an explicit nil

### UnsetRmanBackupType
`func (o *OracleDBChannelInfo) UnsetRmanBackupType()`

UnsetRmanBackupType ensures that no value is present for RmanBackupType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


