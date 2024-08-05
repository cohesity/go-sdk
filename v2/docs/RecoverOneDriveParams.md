# RecoverOneDriveParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other OneDrive items if one of items failed to recover. Default value is false. | [optional] 
**Objects** | [**[]ObjectOneDriveParam**](ObjectOneDriveParam.md) | Specifies a list of OneDrive params associated with the objects to recover. These parameters allow overriding the request level &#39;recoverUserDefaultDrive&#39; parameter for each object specified here. | 
**RecoverPreservationHoldLibrary** | Pointer to **NullableBool** | Specifies whether to recover Preservation Hold Library associated with the OneDrives selected for restore. Default value is false. | [optional] 
**RecoverUserDefaultDrive** | Pointer to **NullableBool** | Specifies whether to recover default drives associated with the OneDrives selected for restore. Default value is true. This setting can be overridden for each object selected for recovery, by specifying &#39;recoverEntireDrive&#39; for the desired drive within &#39;oneDriveParams&#39;. Granular recovery is still allowed even if this value is set to true. | [optional] 
**TargetDrive** | Pointer to [**RecoverOneDriveParamsTargetDrive**](RecoverOneDriveParamsTargetDrive.md) |  | [optional] 

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
### GetRecoverPreservationHoldLibrary

`func (o *RecoverOneDriveParams) GetRecoverPreservationHoldLibrary() bool`

GetRecoverPreservationHoldLibrary returns the RecoverPreservationHoldLibrary field if non-nil, zero value otherwise.

### GetRecoverPreservationHoldLibraryOk

`func (o *RecoverOneDriveParams) GetRecoverPreservationHoldLibraryOk() (*bool, bool)`

GetRecoverPreservationHoldLibraryOk returns a tuple with the RecoverPreservationHoldLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverPreservationHoldLibrary

`func (o *RecoverOneDriveParams) SetRecoverPreservationHoldLibrary(v bool)`

SetRecoverPreservationHoldLibrary sets RecoverPreservationHoldLibrary field to given value.

### HasRecoverPreservationHoldLibrary

`func (o *RecoverOneDriveParams) HasRecoverPreservationHoldLibrary() bool`

HasRecoverPreservationHoldLibrary returns a boolean if a field has been set.

### SetRecoverPreservationHoldLibraryNil

`func (o *RecoverOneDriveParams) SetRecoverPreservationHoldLibraryNil(b bool)`

 SetRecoverPreservationHoldLibraryNil sets the value for RecoverPreservationHoldLibrary to be an explicit nil

### UnsetRecoverPreservationHoldLibrary
`func (o *RecoverOneDriveParams) UnsetRecoverPreservationHoldLibrary()`

UnsetRecoverPreservationHoldLibrary ensures that no value is present for RecoverPreservationHoldLibrary, not even an explicit nil
### GetRecoverUserDefaultDrive

`func (o *RecoverOneDriveParams) GetRecoverUserDefaultDrive() bool`

GetRecoverUserDefaultDrive returns the RecoverUserDefaultDrive field if non-nil, zero value otherwise.

### GetRecoverUserDefaultDriveOk

`func (o *RecoverOneDriveParams) GetRecoverUserDefaultDriveOk() (*bool, bool)`

GetRecoverUserDefaultDriveOk returns a tuple with the RecoverUserDefaultDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverUserDefaultDrive

`func (o *RecoverOneDriveParams) SetRecoverUserDefaultDrive(v bool)`

SetRecoverUserDefaultDrive sets RecoverUserDefaultDrive field to given value.

### HasRecoverUserDefaultDrive

`func (o *RecoverOneDriveParams) HasRecoverUserDefaultDrive() bool`

HasRecoverUserDefaultDrive returns a boolean if a field has been set.

### SetRecoverUserDefaultDriveNil

`func (o *RecoverOneDriveParams) SetRecoverUserDefaultDriveNil(b bool)`

 SetRecoverUserDefaultDriveNil sets the value for RecoverUserDefaultDrive to be an explicit nil

### UnsetRecoverUserDefaultDrive
`func (o *RecoverOneDriveParams) UnsetRecoverUserDefaultDrive()`

UnsetRecoverUserDefaultDrive ensures that no value is present for RecoverUserDefaultDrive, not even an explicit nil
### GetTargetDrive

`func (o *RecoverOneDriveParams) GetTargetDrive() RecoverOneDriveParamsTargetDrive`

GetTargetDrive returns the TargetDrive field if non-nil, zero value otherwise.

### GetTargetDriveOk

`func (o *RecoverOneDriveParams) GetTargetDriveOk() (*RecoverOneDriveParamsTargetDrive, bool)`

GetTargetDriveOk returns a tuple with the TargetDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDrive

`func (o *RecoverOneDriveParams) SetTargetDrive(v RecoverOneDriveParamsTargetDrive)`

SetTargetDrive sets TargetDrive field to given value.

### HasTargetDrive

`func (o *RecoverOneDriveParams) HasTargetDrive() bool`

HasTargetDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


