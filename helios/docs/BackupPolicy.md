# BackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regular** | [**RegularBackupPolicy**](RegularBackupPolicy.md) |  | 
**Log** | Pointer to [**LogBackupPolicy**](LogBackupPolicy.md) |  | [optional] 
**Bmr** | Pointer to [**BmrBackupPolicy**](BmrBackupPolicy.md) |  | [optional] 
**Cdp** | Pointer to [**CdpBackupPolicy**](CdpBackupPolicy.md) |  | [optional] 

## Methods

### NewBackupPolicy

`func NewBackupPolicy(regular RegularBackupPolicy, ) *BackupPolicy`

NewBackupPolicy instantiates a new BackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPolicyWithDefaults

`func NewBackupPolicyWithDefaults() *BackupPolicy`

NewBackupPolicyWithDefaults instantiates a new BackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegular

`func (o *BackupPolicy) GetRegular() RegularBackupPolicy`

GetRegular returns the Regular field if non-nil, zero value otherwise.

### GetRegularOk

`func (o *BackupPolicy) GetRegularOk() (*RegularBackupPolicy, bool)`

GetRegularOk returns a tuple with the Regular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegular

`func (o *BackupPolicy) SetRegular(v RegularBackupPolicy)`

SetRegular sets Regular field to given value.


### GetLog

`func (o *BackupPolicy) GetLog() LogBackupPolicy`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *BackupPolicy) GetLogOk() (*LogBackupPolicy, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *BackupPolicy) SetLog(v LogBackupPolicy)`

SetLog sets Log field to given value.

### HasLog

`func (o *BackupPolicy) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetBmr

`func (o *BackupPolicy) GetBmr() BmrBackupPolicy`

GetBmr returns the Bmr field if non-nil, zero value otherwise.

### GetBmrOk

`func (o *BackupPolicy) GetBmrOk() (*BmrBackupPolicy, bool)`

GetBmrOk returns a tuple with the Bmr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmr

`func (o *BackupPolicy) SetBmr(v BmrBackupPolicy)`

SetBmr sets Bmr field to given value.

### HasBmr

`func (o *BackupPolicy) HasBmr() bool`

HasBmr returns a boolean if a field has been set.

### GetCdp

`func (o *BackupPolicy) GetCdp() CdpBackupPolicy`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *BackupPolicy) GetCdpOk() (*CdpBackupPolicy, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *BackupPolicy) SetCdp(v CdpBackupPolicy)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *BackupPolicy) HasCdp() bool`

HasCdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


