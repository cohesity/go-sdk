# SfdcMetaInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildObjects** | Pointer to [**[]SfdcDependentObject**](SfdcDependentObject.md) | Specifies the list of child objects. | [optional] 
**ParentObjects** | Pointer to [**[]SfdcDependentObject**](SfdcDependentObject.md) | Specifies the list of parent objects. | [optional] 

## Methods

### NewSfdcMetaInfoResult

`func NewSfdcMetaInfoResult() *SfdcMetaInfoResult`

NewSfdcMetaInfoResult instantiates a new SfdcMetaInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcMetaInfoResultWithDefaults

`func NewSfdcMetaInfoResultWithDefaults() *SfdcMetaInfoResult`

NewSfdcMetaInfoResultWithDefaults instantiates a new SfdcMetaInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildObjects

`func (o *SfdcMetaInfoResult) GetChildObjects() []SfdcDependentObject`

GetChildObjects returns the ChildObjects field if non-nil, zero value otherwise.

### GetChildObjectsOk

`func (o *SfdcMetaInfoResult) GetChildObjectsOk() (*[]SfdcDependentObject, bool)`

GetChildObjectsOk returns a tuple with the ChildObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildObjects

`func (o *SfdcMetaInfoResult) SetChildObjects(v []SfdcDependentObject)`

SetChildObjects sets ChildObjects field to given value.

### HasChildObjects

`func (o *SfdcMetaInfoResult) HasChildObjects() bool`

HasChildObjects returns a boolean if a field has been set.

### SetChildObjectsNil

`func (o *SfdcMetaInfoResult) SetChildObjectsNil(b bool)`

 SetChildObjectsNil sets the value for ChildObjects to be an explicit nil

### UnsetChildObjects
`func (o *SfdcMetaInfoResult) UnsetChildObjects()`

UnsetChildObjects ensures that no value is present for ChildObjects, not even an explicit nil
### GetParentObjects

`func (o *SfdcMetaInfoResult) GetParentObjects() []SfdcDependentObject`

GetParentObjects returns the ParentObjects field if non-nil, zero value otherwise.

### GetParentObjectsOk

`func (o *SfdcMetaInfoResult) GetParentObjectsOk() (*[]SfdcDependentObject, bool)`

GetParentObjectsOk returns a tuple with the ParentObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObjects

`func (o *SfdcMetaInfoResult) SetParentObjects(v []SfdcDependentObject)`

SetParentObjects sets ParentObjects field to given value.

### HasParentObjects

`func (o *SfdcMetaInfoResult) HasParentObjects() bool`

HasParentObjects returns a boolean if a field has been set.

### SetParentObjectsNil

`func (o *SfdcMetaInfoResult) SetParentObjectsNil(b bool)`

 SetParentObjectsNil sets the value for ParentObjects to be an explicit nil

### UnsetParentObjects
`func (o *SfdcMetaInfoResult) UnsetParentObjects()`

UnsetParentObjects ensures that no value is present for ParentObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


