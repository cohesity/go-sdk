# OracleBackupJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersistMountpoints** | Pointer to **NullableBool** | Indicates whether the mountpoints created while backing up Oracle DBs should be persisted. If this is set to &#39;false&#39; all Oracle views mounted to the hosts will be unmounted at the end. Note: This parameter is for the entire Job. For overriding persistence of mountpoints for a subset of Oracle hosts within the job, refer OracleSourceParams. | [optional] 
**VlanParams** | Pointer to [**VlanParams**](VlanParams.md) |  | [optional] 

## Methods

### NewOracleBackupJobParams

`func NewOracleBackupJobParams() *OracleBackupJobParams`

NewOracleBackupJobParams instantiates a new OracleBackupJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleBackupJobParamsWithDefaults

`func NewOracleBackupJobParamsWithDefaults() *OracleBackupJobParams`

NewOracleBackupJobParamsWithDefaults instantiates a new OracleBackupJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersistMountpoints

`func (o *OracleBackupJobParams) GetPersistMountpoints() bool`

GetPersistMountpoints returns the PersistMountpoints field if non-nil, zero value otherwise.

### GetPersistMountpointsOk

`func (o *OracleBackupJobParams) GetPersistMountpointsOk() (*bool, bool)`

GetPersistMountpointsOk returns a tuple with the PersistMountpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMountpoints

`func (o *OracleBackupJobParams) SetPersistMountpoints(v bool)`

SetPersistMountpoints sets PersistMountpoints field to given value.

### HasPersistMountpoints

`func (o *OracleBackupJobParams) HasPersistMountpoints() bool`

HasPersistMountpoints returns a boolean if a field has been set.

### SetPersistMountpointsNil

`func (o *OracleBackupJobParams) SetPersistMountpointsNil(b bool)`

 SetPersistMountpointsNil sets the value for PersistMountpoints to be an explicit nil

### UnsetPersistMountpoints
`func (o *OracleBackupJobParams) UnsetPersistMountpoints()`

UnsetPersistMountpoints ensures that no value is present for PersistMountpoints, not even an explicit nil
### GetVlanParams

`func (o *OracleBackupJobParams) GetVlanParams() VlanParams`

GetVlanParams returns the VlanParams field if non-nil, zero value otherwise.

### GetVlanParamsOk

`func (o *OracleBackupJobParams) GetVlanParamsOk() (*VlanParams, bool)`

GetVlanParamsOk returns a tuple with the VlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParams

`func (o *OracleBackupJobParams) SetVlanParams(v VlanParams)`

SetVlanParams sets VlanParams field to given value.

### HasVlanParams

`func (o *OracleBackupJobParams) HasVlanParams() bool`

HasVlanParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


