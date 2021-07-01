# OutlookBackupEnvParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilteringPolicy** | Pointer to [**FilteringPolicyProto**](FilteringPolicyProto.md) |  | [optional] 
**ShouldBackupMailbox** | Pointer to **NullableBool** | Specifies whether the mailbox for all the Office365 Users present in the protection job should be backed up. | [optional] 

## Methods

### NewOutlookBackupEnvParams

`func NewOutlookBackupEnvParams() *OutlookBackupEnvParams`

NewOutlookBackupEnvParams instantiates a new OutlookBackupEnvParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlookBackupEnvParamsWithDefaults

`func NewOutlookBackupEnvParamsWithDefaults() *OutlookBackupEnvParams`

NewOutlookBackupEnvParamsWithDefaults instantiates a new OutlookBackupEnvParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilteringPolicy

`func (o *OutlookBackupEnvParams) GetFilteringPolicy() FilteringPolicyProto`

GetFilteringPolicy returns the FilteringPolicy field if non-nil, zero value otherwise.

### GetFilteringPolicyOk

`func (o *OutlookBackupEnvParams) GetFilteringPolicyOk() (*FilteringPolicyProto, bool)`

GetFilteringPolicyOk returns a tuple with the FilteringPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringPolicy

`func (o *OutlookBackupEnvParams) SetFilteringPolicy(v FilteringPolicyProto)`

SetFilteringPolicy sets FilteringPolicy field to given value.

### HasFilteringPolicy

`func (o *OutlookBackupEnvParams) HasFilteringPolicy() bool`

HasFilteringPolicy returns a boolean if a field has been set.

### GetShouldBackupMailbox

`func (o *OutlookBackupEnvParams) GetShouldBackupMailbox() bool`

GetShouldBackupMailbox returns the ShouldBackupMailbox field if non-nil, zero value otherwise.

### GetShouldBackupMailboxOk

`func (o *OutlookBackupEnvParams) GetShouldBackupMailboxOk() (*bool, bool)`

GetShouldBackupMailboxOk returns a tuple with the ShouldBackupMailbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBackupMailbox

`func (o *OutlookBackupEnvParams) SetShouldBackupMailbox(v bool)`

SetShouldBackupMailbox sets ShouldBackupMailbox field to given value.

### HasShouldBackupMailbox

`func (o *OutlookBackupEnvParams) HasShouldBackupMailbox() bool`

HasShouldBackupMailbox returns a boolean if a field has been set.

### SetShouldBackupMailboxNil

`func (o *OutlookBackupEnvParams) SetShouldBackupMailboxNil(b bool)`

 SetShouldBackupMailboxNil sets the value for ShouldBackupMailbox to be an explicit nil

### UnsetShouldBackupMailbox
`func (o *OutlookBackupEnvParams) UnsetShouldBackupMailbox()`

UnsetShouldBackupMailbox ensures that no value is present for ShouldBackupMailbox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


