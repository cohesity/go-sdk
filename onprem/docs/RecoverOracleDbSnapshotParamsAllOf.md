# RecoverOracleDbSnapshotParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstantRecoveryInfo** | Pointer to [**NullableRecoveryOracleTaskInfo**](RecoveryOracleTaskInfo.md) | Specifies the info about instant recovery. This is only applicable for RecoverOracle. | [optional] 

## Methods

### NewRecoverOracleDbSnapshotParamsAllOf

`func NewRecoverOracleDbSnapshotParamsAllOf() *RecoverOracleDbSnapshotParamsAllOf`

NewRecoverOracleDbSnapshotParamsAllOf instantiates a new RecoverOracleDbSnapshotParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleDbSnapshotParamsAllOfWithDefaults

`func NewRecoverOracleDbSnapshotParamsAllOfWithDefaults() *RecoverOracleDbSnapshotParamsAllOf`

NewRecoverOracleDbSnapshotParamsAllOfWithDefaults instantiates a new RecoverOracleDbSnapshotParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstantRecoveryInfo

`func (o *RecoverOracleDbSnapshotParamsAllOf) GetInstantRecoveryInfo() RecoveryOracleTaskInfo`

GetInstantRecoveryInfo returns the InstantRecoveryInfo field if non-nil, zero value otherwise.

### GetInstantRecoveryInfoOk

`func (o *RecoverOracleDbSnapshotParamsAllOf) GetInstantRecoveryInfoOk() (*RecoveryOracleTaskInfo, bool)`

GetInstantRecoveryInfoOk returns a tuple with the InstantRecoveryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRecoveryInfo

`func (o *RecoverOracleDbSnapshotParamsAllOf) SetInstantRecoveryInfo(v RecoveryOracleTaskInfo)`

SetInstantRecoveryInfo sets InstantRecoveryInfo field to given value.

### HasInstantRecoveryInfo

`func (o *RecoverOracleDbSnapshotParamsAllOf) HasInstantRecoveryInfo() bool`

HasInstantRecoveryInfo returns a boolean if a field has been set.

### SetInstantRecoveryInfoNil

`func (o *RecoverOracleDbSnapshotParamsAllOf) SetInstantRecoveryInfoNil(b bool)`

 SetInstantRecoveryInfoNil sets the value for InstantRecoveryInfo to be an explicit nil

### UnsetInstantRecoveryInfo
`func (o *RecoverOracleDbSnapshotParamsAllOf) UnsetInstantRecoveryInfo()`

UnsetInstantRecoveryInfo ensures that no value is present for InstantRecoveryInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


