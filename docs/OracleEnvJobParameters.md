# OracleEnvJobParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersistMountpoints** | Pointer to **NullableBool** | Specifies whether the mountpoints created while backing up Oracle DBs should be persisted. Note: This parameter is for the entire Job. For overriding persistence of mountpoints for a subset of Oracle hosts within the job, refer OracleSourceParams. | [optional] 
**VlanParams** | Pointer to [**VlanParams**](VlanParams.md) |  | [optional] 

## Methods

### NewOracleEnvJobParameters

`func NewOracleEnvJobParameters() *OracleEnvJobParameters`

NewOracleEnvJobParameters instantiates a new OracleEnvJobParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleEnvJobParametersWithDefaults

`func NewOracleEnvJobParametersWithDefaults() *OracleEnvJobParameters`

NewOracleEnvJobParametersWithDefaults instantiates a new OracleEnvJobParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersistMountpoints

`func (o *OracleEnvJobParameters) GetPersistMountpoints() bool`

GetPersistMountpoints returns the PersistMountpoints field if non-nil, zero value otherwise.

### GetPersistMountpointsOk

`func (o *OracleEnvJobParameters) GetPersistMountpointsOk() (*bool, bool)`

GetPersistMountpointsOk returns a tuple with the PersistMountpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMountpoints

`func (o *OracleEnvJobParameters) SetPersistMountpoints(v bool)`

SetPersistMountpoints sets PersistMountpoints field to given value.

### HasPersistMountpoints

`func (o *OracleEnvJobParameters) HasPersistMountpoints() bool`

HasPersistMountpoints returns a boolean if a field has been set.

### SetPersistMountpointsNil

`func (o *OracleEnvJobParameters) SetPersistMountpointsNil(b bool)`

 SetPersistMountpointsNil sets the value for PersistMountpoints to be an explicit nil

### UnsetPersistMountpoints
`func (o *OracleEnvJobParameters) UnsetPersistMountpoints()`

UnsetPersistMountpoints ensures that no value is present for PersistMountpoints, not even an explicit nil
### GetVlanParams

`func (o *OracleEnvJobParameters) GetVlanParams() VlanParams`

GetVlanParams returns the VlanParams field if non-nil, zero value otherwise.

### GetVlanParamsOk

`func (o *OracleEnvJobParameters) GetVlanParamsOk() (*VlanParams, bool)`

GetVlanParamsOk returns a tuple with the VlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParams

`func (o *OracleEnvJobParameters) SetVlanParams(v VlanParams)`

SetVlanParams sets VlanParams field to given value.

### HasVlanParams

`func (o *OracleEnvJobParameters) HasVlanParams() bool`

HasVlanParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


