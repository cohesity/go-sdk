# OutlookEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePathFilter** | Pointer to [**FilePathFilter**](FilePathFilter.md) |  | [optional] 
**ShouldBackupMailbox** | Pointer to **NullableBool** | Specifies whether mailbox of each Office365 Users/Groups within the job, should be backed up or not. | [optional] 

## Methods

### NewOutlookEnvJobParameters

`func NewOutlookEnvJobParameters() *OutlookEnvJobParameters`

NewOutlookEnvJobParameters instantiates a new OutlookEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlookEnvJobParametersWithDefaults

`func NewOutlookEnvJobParametersWithDefaults() *OutlookEnvJobParameters`

NewOutlookEnvJobParametersWithDefaults instantiates a new OutlookEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePathFilter

`func (o *OutlookEnvJobParameters) GetFilePathFilter() FilePathFilter`

GetFilePathFilter returns the FilePathFilter field if non-nil, zero value otherwise.

### GetFilePathFilterOk

`func (o *OutlookEnvJobParameters) GetFilePathFilterOk() (*FilePathFilter, bool)`

GetFilePathFilterOk returns a tuple with the FilePathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePathFilter

`func (o *OutlookEnvJobParameters) SetFilePathFilter(v FilePathFilter)`

SetFilePathFilter sets FilePathFilter field to given value.

### HasFilePathFilter

`func (o *OutlookEnvJobParameters) HasFilePathFilter() bool`

HasFilePathFilter returns a boolean if a field has been set.

### GetShouldBackupMailbox

`func (o *OutlookEnvJobParameters) GetShouldBackupMailbox() bool`

GetShouldBackupMailbox returns the ShouldBackupMailbox field if non-nil, zero value otherwise.

### GetShouldBackupMailboxOk

`func (o *OutlookEnvJobParameters) GetShouldBackupMailboxOk() (*bool, bool)`

GetShouldBackupMailboxOk returns a tuple with the ShouldBackupMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBackupMailbox

`func (o *OutlookEnvJobParameters) SetShouldBackupMailbox(v bool)`

SetShouldBackupMailbox sets ShouldBackupMailbox field to given value.

### HasShouldBackupMailbox

`func (o *OutlookEnvJobParameters) HasShouldBackupMailbox() bool`

HasShouldBackupMailbox returns a boolean if a field has been set.

### SetShouldBackupMailboxNil

`func (o *OutlookEnvJobParameters) SetShouldBackupMailboxNil(b bool)`

 SetShouldBackupMailboxNil sets the value for ShouldBackupMailbox to be an explicit nil

### UnsetShouldBackupMailbox
`func (o *OutlookEnvJobParameters) UnsetShouldBackupMailbox()`

UnsetShouldBackupMailbox ensures that no value is present for ShouldBackupMailbox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


