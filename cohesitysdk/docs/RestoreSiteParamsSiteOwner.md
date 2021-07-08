# RestoreSiteParamsSiteOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveVec** | Pointer to [**[]RestoreSiteParamsSiteOwnerDrive**](RestoreSiteParamsSiteOwnerDrive.md) | The list of drives that are being restored. | [optional] 
**Object** | Pointer to [**RestoreObject**](RestoreObject.md) |  | [optional] 

## Methods

### NewRestoreSiteParamsSiteOwner

`func NewRestoreSiteParamsSiteOwner() *RestoreSiteParamsSiteOwner`

NewRestoreSiteParamsSiteOwner instantiates a new RestoreSiteParamsSiteOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSiteParamsSiteOwnerWithDefaults

`func NewRestoreSiteParamsSiteOwnerWithDefaults() *RestoreSiteParamsSiteOwner`

NewRestoreSiteParamsSiteOwnerWithDefaults instantiates a new RestoreSiteParamsSiteOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveVec

`func (o *RestoreSiteParamsSiteOwner) GetDriveVec() []RestoreSiteParamsSiteOwnerDrive`

GetDriveVec returns the DriveVec field if non-nil, zero value otherwise.

### GetDriveVecOk

`func (o *RestoreSiteParamsSiteOwner) GetDriveVecOk() (*[]RestoreSiteParamsSiteOwnerDrive, bool)`

GetDriveVecOk returns a tuple with the DriveVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveVec

`func (o *RestoreSiteParamsSiteOwner) SetDriveVec(v []RestoreSiteParamsSiteOwnerDrive)`

SetDriveVec sets DriveVec field to given value.

### HasDriveVec

`func (o *RestoreSiteParamsSiteOwner) HasDriveVec() bool`

HasDriveVec returns a boolean if a field has been set.

### SetDriveVecNil

`func (o *RestoreSiteParamsSiteOwner) SetDriveVecNil(b bool)`

 SetDriveVecNil sets the value for DriveVec to be an explicit nil

### UnsetDriveVec
`func (o *RestoreSiteParamsSiteOwner) UnsetDriveVec()`

UnsetDriveVec ensures that no value is present for DriveVec, not even an explicit nil
### GetObject

`func (o *RestoreSiteParamsSiteOwner) GetObject() RestoreObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RestoreSiteParamsSiteOwner) GetObjectOk() (*RestoreObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RestoreSiteParamsSiteOwner) SetObject(v RestoreObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *RestoreSiteParamsSiteOwner) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


