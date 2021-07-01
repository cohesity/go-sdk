# BackupJobProtoDRToCloudParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeedToFailOver** | Pointer to **NullableBool** | Whether the objects in this job will be failed over to cloud. In case of VMs, we need to fetch information about the logical volumes present on the VM. Magneto might fail backup of a VM in case volume information can not be fetched (maybe because the agent is not installed or if the VM is turned off etc.).  The VM will be backed up using the physical agent when it is running in the cloud. We might choose to backup the VM in the cloud using native API at a later point.  This flag makes sense when configuring a job to backup on-prem VMs. | [optional] 

## Methods

### NewBackupJobProtoDRToCloudParams

`func NewBackupJobProtoDRToCloudParams() *BackupJobProtoDRToCloudParams`

NewBackupJobProtoDRToCloudParams instantiates a new BackupJobProtoDRToCloudParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobProtoDRToCloudParamsWithDefaults

`func NewBackupJobProtoDRToCloudParamsWithDefaults() *BackupJobProtoDRToCloudParams`

NewBackupJobProtoDRToCloudParamsWithDefaults instantiates a new BackupJobProtoDRToCloudParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeedToFailOver

`func (o *BackupJobProtoDRToCloudParams) GetNeedToFailOver() bool`

GetNeedToFailOver returns the NeedToFailOver field if non-nil, zero value otherwise.

### GetNeedToFailOverOk

`func (o *BackupJobProtoDRToCloudParams) GetNeedToFailOverOk() (*bool, bool)`

GetNeedToFailOverOk returns a tuple with the NeedToFailOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedToFailOver

`func (o *BackupJobProtoDRToCloudParams) SetNeedToFailOver(v bool)`

SetNeedToFailOver sets NeedToFailOver field to given value.

### HasNeedToFailOver

`func (o *BackupJobProtoDRToCloudParams) HasNeedToFailOver() bool`

HasNeedToFailOver returns a boolean if a field has been set.

### SetNeedToFailOverNil

`func (o *BackupJobProtoDRToCloudParams) SetNeedToFailOverNil(b bool)`

 SetNeedToFailOverNil sets the value for NeedToFailOver to be an explicit nil

### UnsetNeedToFailOver
`func (o *BackupJobProtoDRToCloudParams) UnsetNeedToFailOver()`

UnsetNeedToFailOver ensures that no value is present for NeedToFailOver, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


