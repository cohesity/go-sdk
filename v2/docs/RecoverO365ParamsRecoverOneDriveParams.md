# RecoverO365ParamsRecoverOneDriveParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other OneDrive items if one of items failed to recover. Default value is false. | [optional] 
**Objects** | [**[]ObjectOneDriveParam**](ObjectOneDriveParam.md) | Specifies a list of OneDrive params associated with the objects to recover. These parameters allow overriding the request level &#39;recoverUserDefaultDrive&#39; parameter for each object specified here. | 
**RecoverPreservationHoldLibrary** | Pointer to **NullableBool** | Specifies whether to recover Preservation Hold Library associated with the OneDrives selected for restore. Default value is false. | [optional] 
**RecoverUserDefaultDrive** | Pointer to **NullableBool** | Specifies whether to recover default drives associated with the OneDrives selected for restore. Default value is true. This setting can be overridden for each object selected for recovery, by specifying &#39;recoverEntireDrive&#39; for the desired drive within &#39;oneDriveParams&#39;. Granular recovery is still allowed even if this value is set to true. | [optional] 
**TargetDrive** | Pointer to [**RecoverOneDriveParamsTargetDrive**](RecoverOneDriveParamsTargetDrive.md) |  | [optional] 

## Methods

### NewRecoverO365ParamsRecoverOneDriveParams

`func NewRecoverO365ParamsRecoverOneDriveParams(objects []ObjectOneDriveParam, ) *RecoverO365ParamsRecoverOneDriveParams`

NewRecoverO365ParamsRecoverOneDriveParams instantiates a new RecoverO365ParamsRecoverOneDriveParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverO365ParamsRecoverOneDriveParamsWithDefaults

`func NewRecoverO365ParamsRecoverOneDriveParamsWithDefaults() *RecoverO365ParamsRecoverOneDriveParams`

NewRecoverO365ParamsRecoverOneDriveParamsWithDefaults instantiates a new RecoverO365ParamsRecoverOneDriveParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverO365ParamsRecoverOneDriveParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverO365ParamsRecoverOneDriveParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetObjects

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetObjects() []ObjectOneDriveParam`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetObjectsOk() (*[]ObjectOneDriveParam, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetObjects(v []ObjectOneDriveParam)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverO365ParamsRecoverOneDriveParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverPreservationHoldLibrary

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetRecoverPreservationHoldLibrary() bool`

GetRecoverPreservationHoldLibrary returns the RecoverPreservationHoldLibrary field if non-nil, zero value otherwise.

### GetRecoverPreservationHoldLibraryOk

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetRecoverPreservationHoldLibraryOk() (*bool, bool)`

GetRecoverPreservationHoldLibraryOk returns a tuple with the RecoverPreservationHoldLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverPreservationHoldLibrary

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetRecoverPreservationHoldLibrary(v bool)`

SetRecoverPreservationHoldLibrary sets RecoverPreservationHoldLibrary field to given value.

### HasRecoverPreservationHoldLibrary

`func (o *RecoverO365ParamsRecoverOneDriveParams) HasRecoverPreservationHoldLibrary() bool`

HasRecoverPreservationHoldLibrary returns a boolean if a field has been set.

### SetRecoverPreservationHoldLibraryNil

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetRecoverPreservationHoldLibraryNil(b bool)`

 SetRecoverPreservationHoldLibraryNil sets the value for RecoverPreservationHoldLibrary to be an explicit nil

### UnsetRecoverPreservationHoldLibrary
`func (o *RecoverO365ParamsRecoverOneDriveParams) UnsetRecoverPreservationHoldLibrary()`

UnsetRecoverPreservationHoldLibrary ensures that no value is present for RecoverPreservationHoldLibrary, not even an explicit nil
### GetRecoverUserDefaultDrive

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetRecoverUserDefaultDrive() bool`

GetRecoverUserDefaultDrive returns the RecoverUserDefaultDrive field if non-nil, zero value otherwise.

### GetRecoverUserDefaultDriveOk

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetRecoverUserDefaultDriveOk() (*bool, bool)`

GetRecoverUserDefaultDriveOk returns a tuple with the RecoverUserDefaultDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverUserDefaultDrive

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetRecoverUserDefaultDrive(v bool)`

SetRecoverUserDefaultDrive sets RecoverUserDefaultDrive field to given value.

### HasRecoverUserDefaultDrive

`func (o *RecoverO365ParamsRecoverOneDriveParams) HasRecoverUserDefaultDrive() bool`

HasRecoverUserDefaultDrive returns a boolean if a field has been set.

### SetRecoverUserDefaultDriveNil

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetRecoverUserDefaultDriveNil(b bool)`

 SetRecoverUserDefaultDriveNil sets the value for RecoverUserDefaultDrive to be an explicit nil

### UnsetRecoverUserDefaultDrive
`func (o *RecoverO365ParamsRecoverOneDriveParams) UnsetRecoverUserDefaultDrive()`

UnsetRecoverUserDefaultDrive ensures that no value is present for RecoverUserDefaultDrive, not even an explicit nil
### GetTargetDrive

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetTargetDrive() RecoverOneDriveParamsTargetDrive`

GetTargetDrive returns the TargetDrive field if non-nil, zero value otherwise.

### GetTargetDriveOk

`func (o *RecoverO365ParamsRecoverOneDriveParams) GetTargetDriveOk() (*RecoverOneDriveParamsTargetDrive, bool)`

GetTargetDriveOk returns a tuple with the TargetDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDrive

`func (o *RecoverO365ParamsRecoverOneDriveParams) SetTargetDrive(v RecoverOneDriveParamsTargetDrive)`

SetTargetDrive sets TargetDrive field to given value.

### HasTargetDrive

`func (o *RecoverO365ParamsRecoverOneDriveParams) HasTargetDrive() bool`

HasTargetDrive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


