# RecoverPureParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]CommonRecoverObjectSnapshotParams**](CommonRecoverObjectSnapshotParams.md) | Specifies the list of recover object parameters. | 
**RecoverSanGroupParams** | Pointer to [**NullableRecoverPureParamsRecoverSanGroupParams**](RecoverPureParamsRecoverSanGroupParams.md) |  | [optional] 
**RecoverSanVolumeParams** | Pointer to [**NullableRecoverPureParamsRecoverSanVolumeParams**](RecoverPureParamsRecoverSanVolumeParams.md) |  | [optional] 
**RecoveryAction** | **string** | Specifies the type of recovery action to be performed. The corresponding recovery action params must be filled out. | 

## Methods

### NewRecoverPureParams

`func NewRecoverPureParams(objects []CommonRecoverObjectSnapshotParams, recoveryAction string, ) *RecoverPureParams`

NewRecoverPureParams instantiates a new RecoverPureParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPureParamsWithDefaults

`func NewRecoverPureParamsWithDefaults() *RecoverPureParams`

NewRecoverPureParamsWithDefaults instantiates a new RecoverPureParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *RecoverPureParams) GetObjects() []CommonRecoverObjectSnapshotParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverPureParams) GetObjectsOk() (*[]CommonRecoverObjectSnapshotParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverPureParams) SetObjects(v []CommonRecoverObjectSnapshotParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverPureParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverPureParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverSanGroupParams

`func (o *RecoverPureParams) GetRecoverSanGroupParams() RecoverPureParamsRecoverSanGroupParams`

GetRecoverSanGroupParams returns the RecoverSanGroupParams field if non-nil, zero value otherwise.

### GetRecoverSanGroupParamsOk

`func (o *RecoverPureParams) GetRecoverSanGroupParamsOk() (*RecoverPureParamsRecoverSanGroupParams, bool)`

GetRecoverSanGroupParamsOk returns a tuple with the RecoverSanGroupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverSanGroupParams

`func (o *RecoverPureParams) SetRecoverSanGroupParams(v RecoverPureParamsRecoverSanGroupParams)`

SetRecoverSanGroupParams sets RecoverSanGroupParams field to given value.

### HasRecoverSanGroupParams

`func (o *RecoverPureParams) HasRecoverSanGroupParams() bool`

HasRecoverSanGroupParams returns a boolean if a field has been set.

### SetRecoverSanGroupParamsNil

`func (o *RecoverPureParams) SetRecoverSanGroupParamsNil(b bool)`

 SetRecoverSanGroupParamsNil sets the value for RecoverSanGroupParams to be an explicit nil

### UnsetRecoverSanGroupParams
`func (o *RecoverPureParams) UnsetRecoverSanGroupParams()`

UnsetRecoverSanGroupParams ensures that no value is present for RecoverSanGroupParams, not even an explicit nil
### GetRecoverSanVolumeParams

`func (o *RecoverPureParams) GetRecoverSanVolumeParams() RecoverPureParamsRecoverSanVolumeParams`

GetRecoverSanVolumeParams returns the RecoverSanVolumeParams field if non-nil, zero value otherwise.

### GetRecoverSanVolumeParamsOk

`func (o *RecoverPureParams) GetRecoverSanVolumeParamsOk() (*RecoverPureParamsRecoverSanVolumeParams, bool)`

GetRecoverSanVolumeParamsOk returns a tuple with the RecoverSanVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverSanVolumeParams

`func (o *RecoverPureParams) SetRecoverSanVolumeParams(v RecoverPureParamsRecoverSanVolumeParams)`

SetRecoverSanVolumeParams sets RecoverSanVolumeParams field to given value.

### HasRecoverSanVolumeParams

`func (o *RecoverPureParams) HasRecoverSanVolumeParams() bool`

HasRecoverSanVolumeParams returns a boolean if a field has been set.

### SetRecoverSanVolumeParamsNil

`func (o *RecoverPureParams) SetRecoverSanVolumeParamsNil(b bool)`

 SetRecoverSanVolumeParamsNil sets the value for RecoverSanVolumeParams to be an explicit nil

### UnsetRecoverSanVolumeParams
`func (o *RecoverPureParams) UnsetRecoverSanVolumeParams()`

UnsetRecoverSanVolumeParams ensures that no value is present for RecoverSanVolumeParams, not even an explicit nil
### GetRecoveryAction

`func (o *RecoverPureParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverPureParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverPureParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


