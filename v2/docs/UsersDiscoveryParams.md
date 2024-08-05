# UsersDiscoveryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowChatsBackup** | Pointer to **NullableBool** | Specifies whether users&#39; chats should be backed up or not. If this is false or not specified users&#39; chats backup will not be done. | [optional] 
**DiscoverUsersWithMailbox** | Pointer to **NullableBool** | Specifies if office 365 users with valid mailboxes should be discovered or not. | [optional] 
**DiscoverUsersWithOnedrive** | Pointer to **NullableBool** | Specifies if office 365 users with valid Onedrives should be discovered or not. | [optional] 
**FetchMailboxInfo** | Pointer to **NullableBool** | Specifies whether users&#39; mailbox info including the provisioning status, mailbox type &amp; in-place archival support will be fetched and processed. | [optional] 
**FetchOneDriveInfo** | Pointer to **NullableBool** | Specifies whether users&#39; onedrive info including the provisioning status &amp; storage quota will be fetched and processed. | [optional] 
**SkipUsersWithoutMySite** | Pointer to **NullableBool** | Specifies whether to skip processing user who have uninitialized OneDrive or are without MySite. | [optional] 

## Methods

### NewUsersDiscoveryParams

`func NewUsersDiscoveryParams() *UsersDiscoveryParams`

NewUsersDiscoveryParams instantiates a new UsersDiscoveryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersDiscoveryParamsWithDefaults

`func NewUsersDiscoveryParamsWithDefaults() *UsersDiscoveryParams`

NewUsersDiscoveryParamsWithDefaults instantiates a new UsersDiscoveryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowChatsBackup

`func (o *UsersDiscoveryParams) GetAllowChatsBackup() bool`

GetAllowChatsBackup returns the AllowChatsBackup field if non-nil, zero value otherwise.

### GetAllowChatsBackupOk

`func (o *UsersDiscoveryParams) GetAllowChatsBackupOk() (*bool, bool)`

GetAllowChatsBackupOk returns a tuple with the AllowChatsBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChatsBackup

`func (o *UsersDiscoveryParams) SetAllowChatsBackup(v bool)`

SetAllowChatsBackup sets AllowChatsBackup field to given value.

### HasAllowChatsBackup

`func (o *UsersDiscoveryParams) HasAllowChatsBackup() bool`

HasAllowChatsBackup returns a boolean if a field has been set.

### SetAllowChatsBackupNil

`func (o *UsersDiscoveryParams) SetAllowChatsBackupNil(b bool)`

 SetAllowChatsBackupNil sets the value for AllowChatsBackup to be an explicit nil

### UnsetAllowChatsBackup
`func (o *UsersDiscoveryParams) UnsetAllowChatsBackup()`

UnsetAllowChatsBackup ensures that no value is present for AllowChatsBackup, not even an explicit nil
### GetDiscoverUsersWithMailbox

`func (o *UsersDiscoveryParams) GetDiscoverUsersWithMailbox() bool`

GetDiscoverUsersWithMailbox returns the DiscoverUsersWithMailbox field if non-nil, zero value otherwise.

### GetDiscoverUsersWithMailboxOk

`func (o *UsersDiscoveryParams) GetDiscoverUsersWithMailboxOk() (*bool, bool)`

GetDiscoverUsersWithMailboxOk returns a tuple with the DiscoverUsersWithMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverUsersWithMailbox

`func (o *UsersDiscoveryParams) SetDiscoverUsersWithMailbox(v bool)`

SetDiscoverUsersWithMailbox sets DiscoverUsersWithMailbox field to given value.

### HasDiscoverUsersWithMailbox

`func (o *UsersDiscoveryParams) HasDiscoverUsersWithMailbox() bool`

HasDiscoverUsersWithMailbox returns a boolean if a field has been set.

### SetDiscoverUsersWithMailboxNil

`func (o *UsersDiscoveryParams) SetDiscoverUsersWithMailboxNil(b bool)`

 SetDiscoverUsersWithMailboxNil sets the value for DiscoverUsersWithMailbox to be an explicit nil

### UnsetDiscoverUsersWithMailbox
`func (o *UsersDiscoveryParams) UnsetDiscoverUsersWithMailbox()`

UnsetDiscoverUsersWithMailbox ensures that no value is present for DiscoverUsersWithMailbox, not even an explicit nil
### GetDiscoverUsersWithOnedrive

`func (o *UsersDiscoveryParams) GetDiscoverUsersWithOnedrive() bool`

GetDiscoverUsersWithOnedrive returns the DiscoverUsersWithOnedrive field if non-nil, zero value otherwise.

### GetDiscoverUsersWithOnedriveOk

`func (o *UsersDiscoveryParams) GetDiscoverUsersWithOnedriveOk() (*bool, bool)`

GetDiscoverUsersWithOnedriveOk returns a tuple with the DiscoverUsersWithOnedrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverUsersWithOnedrive

`func (o *UsersDiscoveryParams) SetDiscoverUsersWithOnedrive(v bool)`

SetDiscoverUsersWithOnedrive sets DiscoverUsersWithOnedrive field to given value.

### HasDiscoverUsersWithOnedrive

`func (o *UsersDiscoveryParams) HasDiscoverUsersWithOnedrive() bool`

HasDiscoverUsersWithOnedrive returns a boolean if a field has been set.

### SetDiscoverUsersWithOnedriveNil

`func (o *UsersDiscoveryParams) SetDiscoverUsersWithOnedriveNil(b bool)`

 SetDiscoverUsersWithOnedriveNil sets the value for DiscoverUsersWithOnedrive to be an explicit nil

### UnsetDiscoverUsersWithOnedrive
`func (o *UsersDiscoveryParams) UnsetDiscoverUsersWithOnedrive()`

UnsetDiscoverUsersWithOnedrive ensures that no value is present for DiscoverUsersWithOnedrive, not even an explicit nil
### GetFetchMailboxInfo

`func (o *UsersDiscoveryParams) GetFetchMailboxInfo() bool`

GetFetchMailboxInfo returns the FetchMailboxInfo field if non-nil, zero value otherwise.

### GetFetchMailboxInfoOk

`func (o *UsersDiscoveryParams) GetFetchMailboxInfoOk() (*bool, bool)`

GetFetchMailboxInfoOk returns a tuple with the FetchMailboxInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchMailboxInfo

`func (o *UsersDiscoveryParams) SetFetchMailboxInfo(v bool)`

SetFetchMailboxInfo sets FetchMailboxInfo field to given value.

### HasFetchMailboxInfo

`func (o *UsersDiscoveryParams) HasFetchMailboxInfo() bool`

HasFetchMailboxInfo returns a boolean if a field has been set.

### SetFetchMailboxInfoNil

`func (o *UsersDiscoveryParams) SetFetchMailboxInfoNil(b bool)`

 SetFetchMailboxInfoNil sets the value for FetchMailboxInfo to be an explicit nil

### UnsetFetchMailboxInfo
`func (o *UsersDiscoveryParams) UnsetFetchMailboxInfo()`

UnsetFetchMailboxInfo ensures that no value is present for FetchMailboxInfo, not even an explicit nil
### GetFetchOneDriveInfo

`func (o *UsersDiscoveryParams) GetFetchOneDriveInfo() bool`

GetFetchOneDriveInfo returns the FetchOneDriveInfo field if non-nil, zero value otherwise.

### GetFetchOneDriveInfoOk

`func (o *UsersDiscoveryParams) GetFetchOneDriveInfoOk() (*bool, bool)`

GetFetchOneDriveInfoOk returns a tuple with the FetchOneDriveInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchOneDriveInfo

`func (o *UsersDiscoveryParams) SetFetchOneDriveInfo(v bool)`

SetFetchOneDriveInfo sets FetchOneDriveInfo field to given value.

### HasFetchOneDriveInfo

`func (o *UsersDiscoveryParams) HasFetchOneDriveInfo() bool`

HasFetchOneDriveInfo returns a boolean if a field has been set.

### SetFetchOneDriveInfoNil

`func (o *UsersDiscoveryParams) SetFetchOneDriveInfoNil(b bool)`

 SetFetchOneDriveInfoNil sets the value for FetchOneDriveInfo to be an explicit nil

### UnsetFetchOneDriveInfo
`func (o *UsersDiscoveryParams) UnsetFetchOneDriveInfo()`

UnsetFetchOneDriveInfo ensures that no value is present for FetchOneDriveInfo, not even an explicit nil
### GetSkipUsersWithoutMySite

`func (o *UsersDiscoveryParams) GetSkipUsersWithoutMySite() bool`

GetSkipUsersWithoutMySite returns the SkipUsersWithoutMySite field if non-nil, zero value otherwise.

### GetSkipUsersWithoutMySiteOk

`func (o *UsersDiscoveryParams) GetSkipUsersWithoutMySiteOk() (*bool, bool)`

GetSkipUsersWithoutMySiteOk returns a tuple with the SkipUsersWithoutMySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipUsersWithoutMySite

`func (o *UsersDiscoveryParams) SetSkipUsersWithoutMySite(v bool)`

SetSkipUsersWithoutMySite sets SkipUsersWithoutMySite field to given value.

### HasSkipUsersWithoutMySite

`func (o *UsersDiscoveryParams) HasSkipUsersWithoutMySite() bool`

HasSkipUsersWithoutMySite returns a boolean if a field has been set.

### SetSkipUsersWithoutMySiteNil

`func (o *UsersDiscoveryParams) SetSkipUsersWithoutMySiteNil(b bool)`

 SetSkipUsersWithoutMySiteNil sets the value for SkipUsersWithoutMySite to be an explicit nil

### UnsetSkipUsersWithoutMySite
`func (o *UsersDiscoveryParams) UnsetSkipUsersWithoutMySite()`

UnsetSkipUsersWithoutMySite ensures that no value is present for SkipUsersWithoutMySite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


