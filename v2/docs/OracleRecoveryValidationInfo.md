# OracleRecoveryValidationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateDummyInstance** | Pointer to **NullableBool** | Specifies whether we will be creating a dummy oracle instance to run the validations. Generally if source and target location are different we create a dummy oracle instance else we use the source db. | [optional] 

## Methods

### NewOracleRecoveryValidationInfo

`func NewOracleRecoveryValidationInfo() *OracleRecoveryValidationInfo`

NewOracleRecoveryValidationInfo instantiates a new OracleRecoveryValidationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleRecoveryValidationInfoWithDefaults

`func NewOracleRecoveryValidationInfoWithDefaults() *OracleRecoveryValidationInfo`

NewOracleRecoveryValidationInfoWithDefaults instantiates a new OracleRecoveryValidationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateDummyInstance

`func (o *OracleRecoveryValidationInfo) GetCreateDummyInstance() bool`

GetCreateDummyInstance returns the CreateDummyInstance field if non-nil, zero value otherwise.

### GetCreateDummyInstanceOk

`func (o *OracleRecoveryValidationInfo) GetCreateDummyInstanceOk() (*bool, bool)`

GetCreateDummyInstanceOk returns a tuple with the CreateDummyInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDummyInstance

`func (o *OracleRecoveryValidationInfo) SetCreateDummyInstance(v bool)`

SetCreateDummyInstance sets CreateDummyInstance field to given value.

### HasCreateDummyInstance

`func (o *OracleRecoveryValidationInfo) HasCreateDummyInstance() bool`

HasCreateDummyInstance returns a boolean if a field has been set.

### SetCreateDummyInstanceNil

`func (o *OracleRecoveryValidationInfo) SetCreateDummyInstanceNil(b bool)`

 SetCreateDummyInstanceNil sets the value for CreateDummyInstance to be an explicit nil

### UnsetCreateDummyInstance
`func (o *OracleRecoveryValidationInfo) UnsetCreateDummyInstance()`

UnsetCreateDummyInstance ensures that no value is present for CreateDummyInstance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


