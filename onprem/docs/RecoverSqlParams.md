# RecoverSqlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**RecoverAppParams** | Pointer to [**[]RecoverSqlAppParams**](RecoverSqlAppParams.md) | Specifies the parameters to recover Sql databases. | [optional] 
**VlanConfig** | Pointer to [**NullableRecoveryVlanConfig**](RecoveryVlanConfig.md) | Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client&#39;s (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery. | [optional] 

## Methods

### NewRecoverSqlParams

`func NewRecoverSqlParams(recoveryAction string, ) *RecoverSqlParams`

NewRecoverSqlParams instantiates a new RecoverSqlParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverSqlParamsWithDefaults

`func NewRecoverSqlParamsWithDefaults() *RecoverSqlParams`

NewRecoverSqlParamsWithDefaults instantiates a new RecoverSqlParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryAction

`func (o *RecoverSqlParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverSqlParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverSqlParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.


### GetRecoverAppParams

`func (o *RecoverSqlParams) GetRecoverAppParams() []RecoverSqlAppParams`

GetRecoverAppParams returns the RecoverAppParams field if non-nil, zero value otherwise.

### GetRecoverAppParamsOk

`func (o *RecoverSqlParams) GetRecoverAppParamsOk() (*[]RecoverSqlAppParams, bool)`

GetRecoverAppParamsOk returns a tuple with the RecoverAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverAppParams

`func (o *RecoverSqlParams) SetRecoverAppParams(v []RecoverSqlAppParams)`

SetRecoverAppParams sets RecoverAppParams field to given value.

### HasRecoverAppParams

`func (o *RecoverSqlParams) HasRecoverAppParams() bool`

HasRecoverAppParams returns a boolean if a field has been set.

### SetRecoverAppParamsNil

`func (o *RecoverSqlParams) SetRecoverAppParamsNil(b bool)`

 SetRecoverAppParamsNil sets the value for RecoverAppParams to be an explicit nil

### UnsetRecoverAppParams
`func (o *RecoverSqlParams) UnsetRecoverAppParams()`

UnsetRecoverAppParams ensures that no value is present for RecoverAppParams, not even an explicit nil
### GetVlanConfig

`func (o *RecoverSqlParams) GetVlanConfig() RecoveryVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverSqlParams) GetVlanConfigOk() (*RecoveryVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverSqlParams) SetVlanConfig(v RecoveryVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverSqlParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverSqlParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverSqlParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


