# RecoverSiteParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering the doc libs of a site, if one or more of doc libs failed to recover. Default value is false. | [optional] 
**Objects** | [**[]ObjectSiteParam**](ObjectSiteParam.md) | Specifies a list of site params associated with the objects to recover. | 
**RecoverPreservationHoldLibrary** | Pointer to **NullableBool** | Specifies whether to recover Preservation Hold Library associated with the Sites selected for restore. Default value is false. | [optional] 
**TargetDomainObjectId** | Pointer to [**NullableRecoverSiteParamsTargetDomainObjectId**](RecoverSiteParamsTargetDomainObjectId.md) |  | [optional] 
**TargetSite** | Pointer to [**RecoverSiteParamsTargetSite**](RecoverSiteParamsTargetSite.md) |  | [optional] 

## Methods

### NewRecoverSiteParams

`func NewRecoverSiteParams(objects []ObjectSiteParam, ) *RecoverSiteParams`

NewRecoverSiteParams instantiates a new RecoverSiteParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverSiteParamsWithDefaults

`func NewRecoverSiteParamsWithDefaults() *RecoverSiteParams`

NewRecoverSiteParamsWithDefaults instantiates a new RecoverSiteParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverSiteParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverSiteParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverSiteParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverSiteParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverSiteParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverSiteParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetObjects

`func (o *RecoverSiteParams) GetObjects() []ObjectSiteParam`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *RecoverSiteParams) GetObjectsOk() (*[]ObjectSiteParam, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *RecoverSiteParams) SetObjects(v []ObjectSiteParam)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *RecoverSiteParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *RecoverSiteParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetRecoverPreservationHoldLibrary

`func (o *RecoverSiteParams) GetRecoverPreservationHoldLibrary() bool`

GetRecoverPreservationHoldLibrary returns the RecoverPreservationHoldLibrary field if non-nil, zero value otherwise.

### GetRecoverPreservationHoldLibraryOk

`func (o *RecoverSiteParams) GetRecoverPreservationHoldLibraryOk() (*bool, bool)`

GetRecoverPreservationHoldLibraryOk returns a tuple with the RecoverPreservationHoldLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverPreservationHoldLibrary

`func (o *RecoverSiteParams) SetRecoverPreservationHoldLibrary(v bool)`

SetRecoverPreservationHoldLibrary sets RecoverPreservationHoldLibrary field to given value.

### HasRecoverPreservationHoldLibrary

`func (o *RecoverSiteParams) HasRecoverPreservationHoldLibrary() bool`

HasRecoverPreservationHoldLibrary returns a boolean if a field has been set.

### SetRecoverPreservationHoldLibraryNil

`func (o *RecoverSiteParams) SetRecoverPreservationHoldLibraryNil(b bool)`

 SetRecoverPreservationHoldLibraryNil sets the value for RecoverPreservationHoldLibrary to be an explicit nil

### UnsetRecoverPreservationHoldLibrary
`func (o *RecoverSiteParams) UnsetRecoverPreservationHoldLibrary()`

UnsetRecoverPreservationHoldLibrary ensures that no value is present for RecoverPreservationHoldLibrary, not even an explicit nil
### GetTargetDomainObjectId

`func (o *RecoverSiteParams) GetTargetDomainObjectId() RecoverSiteParamsTargetDomainObjectId`

GetTargetDomainObjectId returns the TargetDomainObjectId field if non-nil, zero value otherwise.

### GetTargetDomainObjectIdOk

`func (o *RecoverSiteParams) GetTargetDomainObjectIdOk() (*RecoverSiteParamsTargetDomainObjectId, bool)`

GetTargetDomainObjectIdOk returns a tuple with the TargetDomainObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDomainObjectId

`func (o *RecoverSiteParams) SetTargetDomainObjectId(v RecoverSiteParamsTargetDomainObjectId)`

SetTargetDomainObjectId sets TargetDomainObjectId field to given value.

### HasTargetDomainObjectId

`func (o *RecoverSiteParams) HasTargetDomainObjectId() bool`

HasTargetDomainObjectId returns a boolean if a field has been set.

### SetTargetDomainObjectIdNil

`func (o *RecoverSiteParams) SetTargetDomainObjectIdNil(b bool)`

 SetTargetDomainObjectIdNil sets the value for TargetDomainObjectId to be an explicit nil

### UnsetTargetDomainObjectId
`func (o *RecoverSiteParams) UnsetTargetDomainObjectId()`

UnsetTargetDomainObjectId ensures that no value is present for TargetDomainObjectId, not even an explicit nil
### GetTargetSite

`func (o *RecoverSiteParams) GetTargetSite() RecoverSiteParamsTargetSite`

GetTargetSite returns the TargetSite field if non-nil, zero value otherwise.

### GetTargetSiteOk

`func (o *RecoverSiteParams) GetTargetSiteOk() (*RecoverSiteParamsTargetSite, bool)`

GetTargetSiteOk returns a tuple with the TargetSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSite

`func (o *RecoverSiteParams) SetTargetSite(v RecoverSiteParamsTargetSite)`

SetTargetSite sets TargetSite field to given value.

### HasTargetSite

`func (o *RecoverSiteParams) HasTargetSite() bool`

HasTargetSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


