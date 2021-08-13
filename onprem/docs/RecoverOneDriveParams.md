# RecoverOneDriveParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]ObjectOneDriveParam**](ObjectOneDriveParam.md) | Specifies a list of OneDrive params associated with the objects to recover. | 
**TargetDrive** | Pointer to [**TargetOneDriveParam**](TargetOneDriveParam.md) | Specifies the target OneDrive to recover to. If not specified, the objects will be recovered to original location. | [optional] 
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other OneDrive items if one of items failed to recover. Default value is false. | [optional] 

## Methods

### NewRecoverOneDriveParams

`func NewRecoverOneDriveParams(objects []ObjectOneDriveParam, ) *RecoverOneDriveParams`

NewRecoverOneDriveParams instantiates a new RecoverOneDriveParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOneDriveParamsWithDefaults

`func NewRecoverOneDriveParamsWithDefaults() *RecoverOneDriveParams`

NewRecoverOneDriveParamsWithDefaults instantiates a new RecoverOneDriveParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *RecoverOneDriveParams) GetObjects() []ObjectOneDriveParam`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverOneDriveParams) GetObjectsOk() (*[]ObjectOneDriveParam, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverOneDriveParams) SetObjects(v []ObjectOneDriveParam)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverOneDriveParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverOneDriveParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetTargetDrive

`func (o *RecoverOneDriveParams) GetTargetDrive() TargetOneDriveParam`

GetTargetDrive returns the TargetDrive field if non-nil, zero value otherwise.

### GetTargetDriveOk

`func (o *RecoverOneDriveParams) GetTargetDriveOk() (*TargetOneDriveParam, bool)`

GetTargetDriveOk returns a tuple with the TargetDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDrive

`func (o *RecoverOneDriveParams) SetTargetDrive(v TargetOneDriveParam)`

SetTargetDrive sets TargetDrive field to given value.

### HasTargetDrive

`func (o *RecoverOneDriveParams) HasTargetDrive() bool`

HasTargetDrive returns a boolean if a field has been set.

### GetContinueOnError

`func (o *RecoverOneDriveParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverOneDriveParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverOneDriveParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverOneDriveParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverOneDriveParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverOneDriveParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


