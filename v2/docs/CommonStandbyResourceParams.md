# CommonStandbyResourceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryPointObjectiveSecs** | Pointer to **NullableInt64** | Specifies the recovery point objective time user expects for this standby resource. | [optional] 

## Methods

### NewCommonStandbyResourceParams

`func NewCommonStandbyResourceParams() *CommonStandbyResourceParams`

NewCommonStandbyResourceParams instantiates a new CommonStandbyResourceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonStandbyResourceParamsWithDefaults

`func NewCommonStandbyResourceParamsWithDefaults() *CommonStandbyResourceParams`

NewCommonStandbyResourceParamsWithDefaults instantiates a new CommonStandbyResourceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryPointObjectiveSecs

`func (o *CommonStandbyResourceParams) GetRecoveryPointObjectiveSecs() int64`

GetRecoveryPointObjectiveSecs returns the RecoveryPointObjectiveSecs field if non-nil, zero value otherwise.

### GetRecoveryPointObjectiveSecsOk

`func (o *CommonStandbyResourceParams) GetRecoveryPointObjectiveSecsOk() (*int64, bool)`

GetRecoveryPointObjectiveSecsOk returns a tuple with the RecoveryPointObjectiveSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryPointObjectiveSecs

`func (o *CommonStandbyResourceParams) SetRecoveryPointObjectiveSecs(v int64)`

SetRecoveryPointObjectiveSecs sets RecoveryPointObjectiveSecs field to given value.

### HasRecoveryPointObjectiveSecs

`func (o *CommonStandbyResourceParams) HasRecoveryPointObjectiveSecs() bool`

HasRecoveryPointObjectiveSecs returns a boolean if a field has been set.

### SetRecoveryPointObjectiveSecsNil

`func (o *CommonStandbyResourceParams) SetRecoveryPointObjectiveSecsNil(b bool)`

 SetRecoveryPointObjectiveSecsNil sets the value for RecoveryPointObjectiveSecs to be an explicit nil

### UnsetRecoveryPointObjectiveSecs
`func (o *CommonStandbyResourceParams) UnsetRecoveryPointObjectiveSecs()`

UnsetRecoveryPointObjectiveSecs ensures that no value is present for RecoveryPointObjectiveSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


