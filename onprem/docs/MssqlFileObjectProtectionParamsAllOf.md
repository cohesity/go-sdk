# MssqlFileObjectProtectionParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]MssqlFileObjectProtection**](MssqlFileObjectProtection.md) | Specifies the list of objects to be protected. | 

## Methods

### NewMssqlFileObjectProtectionParamsAllOf

`func NewMssqlFileObjectProtectionParamsAllOf(objects []MssqlFileObjectProtection, ) *MssqlFileObjectProtectionParamsAllOf`

NewMssqlFileObjectProtectionParamsAllOf instantiates a new MssqlFileObjectProtectionParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMssqlFileObjectProtectionParamsAllOfWithDefaults

`func NewMssqlFileObjectProtectionParamsAllOfWithDefaults() *MssqlFileObjectProtectionParamsAllOf`

NewMssqlFileObjectProtectionParamsAllOfWithDefaults instantiates a new MssqlFileObjectProtectionParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *MssqlFileObjectProtectionParamsAllOf) GetObjects() []MssqlFileObjectProtection`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MssqlFileObjectProtectionParamsAllOf) GetObjectsOk() (*[]MssqlFileObjectProtection, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MssqlFileObjectProtectionParamsAllOf) SetObjects(v []MssqlFileObjectProtection)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *MssqlFileObjectProtectionParamsAllOf) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *MssqlFileObjectProtectionParamsAllOf) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


