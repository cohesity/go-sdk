# MigrateCloneParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelaySecs** | Pointer to **NullableInt64** | Specifies when the migration of the oracle instance should be started after successful recovery. | [optional] 
**TargetPathVec** | Pointer to **[]string** | Specifies the target paths to be used for DB migration. | [optional] 

## Methods

### NewMigrateCloneParams

`func NewMigrateCloneParams() *MigrateCloneParams`

NewMigrateCloneParams instantiates a new MigrateCloneParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrateCloneParamsWithDefaults

`func NewMigrateCloneParamsWithDefaults() *MigrateCloneParams`

NewMigrateCloneParamsWithDefaults instantiates a new MigrateCloneParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelaySecs

`func (o *MigrateCloneParams) GetDelaySecs() int64`

GetDelaySecs returns the DelaySecs field if non-nil, zero value otherwise.

### GetDelaySecsOk

`func (o *MigrateCloneParams) GetDelaySecsOk() (*int64, bool)`

GetDelaySecsOk returns a tuple with the DelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelaySecs

`func (o *MigrateCloneParams) SetDelaySecs(v int64)`

SetDelaySecs sets DelaySecs field to given value.

### HasDelaySecs

`func (o *MigrateCloneParams) HasDelaySecs() bool`

HasDelaySecs returns a boolean if a field has been set.

### SetDelaySecsNil

`func (o *MigrateCloneParams) SetDelaySecsNil(b bool)`

 SetDelaySecsNil sets the value for DelaySecs to be an explicit nil

### UnsetDelaySecs
`func (o *MigrateCloneParams) UnsetDelaySecs()`

UnsetDelaySecs ensures that no value is present for DelaySecs, not even an explicit nil
### GetTargetPathVec

`func (o *MigrateCloneParams) GetTargetPathVec() []string`

GetTargetPathVec returns the TargetPathVec field if non-nil, zero value otherwise.

### GetTargetPathVecOk

`func (o *MigrateCloneParams) GetTargetPathVecOk() (*[]string, bool)`

GetTargetPathVecOk returns a tuple with the TargetPathVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPathVec

`func (o *MigrateCloneParams) SetTargetPathVec(v []string)`

SetTargetPathVec sets TargetPathVec field to given value.

### HasTargetPathVec

`func (o *MigrateCloneParams) HasTargetPathVec() bool`

HasTargetPathVec returns a boolean if a field has been set.

### SetTargetPathVecNil

`func (o *MigrateCloneParams) SetTargetPathVecNil(b bool)`

 SetTargetPathVecNil sets the value for TargetPathVec to be an explicit nil

### UnsetTargetPathVec
`func (o *MigrateCloneParams) UnsetTargetPathVec()`

UnsetTargetPathVec ensures that no value is present for TargetPathVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


