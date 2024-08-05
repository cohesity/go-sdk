# SfdcProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]SfdcProtectionGroupObjectParams**](SfdcProtectionGroupObjectParams.md) | Specifies the list of objects to be protected. | 

## Methods

### NewSfdcProtectionGroupParams

`func NewSfdcProtectionGroupParams(objects []SfdcProtectionGroupObjectParams, ) *SfdcProtectionGroupParams`

NewSfdcProtectionGroupParams instantiates a new SfdcProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcProtectionGroupParamsWithDefaults

`func NewSfdcProtectionGroupParamsWithDefaults() *SfdcProtectionGroupParams`

NewSfdcProtectionGroupParamsWithDefaults instantiates a new SfdcProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *SfdcProtectionGroupParams) GetObjects() []SfdcProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *SfdcProtectionGroupParams) GetObjectsOk() (*[]SfdcProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *SfdcProtectionGroupParams) SetObjects(v []SfdcProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *SfdcProtectionGroupParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *SfdcProtectionGroupParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


