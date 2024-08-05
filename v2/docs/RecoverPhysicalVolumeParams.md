# RecoverPhysicalVolumeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhysicalTargetParams** | Pointer to [**NullableRecoverPhysicalVolumeParamsPhysicalTargetParams**](RecoverPhysicalVolumeParamsPhysicalTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverPhysicalVolumeParams

`func NewRecoverPhysicalVolumeParams(targetEnvironment string, ) *RecoverPhysicalVolumeParams`

NewRecoverPhysicalVolumeParams instantiates a new RecoverPhysicalVolumeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPhysicalVolumeParamsWithDefaults

`func NewRecoverPhysicalVolumeParamsWithDefaults() *RecoverPhysicalVolumeParams`

NewRecoverPhysicalVolumeParamsWithDefaults instantiates a new RecoverPhysicalVolumeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhysicalTargetParams

`func (o *RecoverPhysicalVolumeParams) GetPhysicalTargetParams() RecoverPhysicalVolumeParamsPhysicalTargetParams`

GetPhysicalTargetParams returns the PhysicalTargetParams field if non-nil, zero value otherwise.

### GetPhysicalTargetParamsOk

`func (o *RecoverPhysicalVolumeParams) GetPhysicalTargetParamsOk() (*RecoverPhysicalVolumeParamsPhysicalTargetParams, bool)`

GetPhysicalTargetParamsOk returns a tuple with the PhysicalTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalTargetParams

`func (o *RecoverPhysicalVolumeParams) SetPhysicalTargetParams(v RecoverPhysicalVolumeParamsPhysicalTargetParams)`

SetPhysicalTargetParams sets PhysicalTargetParams field to given value.

### HasPhysicalTargetParams

`func (o *RecoverPhysicalVolumeParams) HasPhysicalTargetParams() bool`

HasPhysicalTargetParams returns a boolean if a field has been set.

### SetPhysicalTargetParamsNil

`func (o *RecoverPhysicalVolumeParams) SetPhysicalTargetParamsNil(b bool)`

 SetPhysicalTargetParamsNil sets the value for PhysicalTargetParams to be an explicit nil

### UnsetPhysicalTargetParams
`func (o *RecoverPhysicalVolumeParams) UnsetPhysicalTargetParams()`

UnsetPhysicalTargetParams ensures that no value is present for PhysicalTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverPhysicalVolumeParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverPhysicalVolumeParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverPhysicalVolumeParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


