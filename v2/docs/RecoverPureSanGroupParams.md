# RecoverPureSanGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PureTargetParams** | Pointer to [**NullableRecoverPureSanGroupParamsPureTargetParams**](RecoverPureSanGroupParamsPureTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding target params must be filled out. | 

## Methods

### NewRecoverPureSanGroupParams

`func NewRecoverPureSanGroupParams(targetEnvironment string, ) *RecoverPureSanGroupParams`

NewRecoverPureSanGroupParams instantiates a new RecoverPureSanGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPureSanGroupParamsWithDefaults

`func NewRecoverPureSanGroupParamsWithDefaults() *RecoverPureSanGroupParams`

NewRecoverPureSanGroupParamsWithDefaults instantiates a new RecoverPureSanGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPureTargetParams

`func (o *RecoverPureSanGroupParams) GetPureTargetParams() RecoverPureSanGroupParamsPureTargetParams`

GetPureTargetParams returns the PureTargetParams field if non-nil, zero value otherwise.

### GetPureTargetParamsOk

`func (o *RecoverPureSanGroupParams) GetPureTargetParamsOk() (*RecoverPureSanGroupParamsPureTargetParams, bool)`

GetPureTargetParamsOk returns a tuple with the PureTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureTargetParams

`func (o *RecoverPureSanGroupParams) SetPureTargetParams(v RecoverPureSanGroupParamsPureTargetParams)`

SetPureTargetParams sets PureTargetParams field to given value.

### HasPureTargetParams

`func (o *RecoverPureSanGroupParams) HasPureTargetParams() bool`

HasPureTargetParams returns a boolean if a field has been set.

### SetPureTargetParamsNil

`func (o *RecoverPureSanGroupParams) SetPureTargetParamsNil(b bool)`

 SetPureTargetParamsNil sets the value for PureTargetParams to be an explicit nil

### UnsetPureTargetParams
`func (o *RecoverPureSanGroupParams) UnsetPureTargetParams()`

UnsetPureTargetParams ensures that no value is present for PureTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverPureSanGroupParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverPureSanGroupParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverPureSanGroupParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


