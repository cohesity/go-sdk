# ADRestoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectInfo** | Pointer to [**CompareADObjectsResultADObject**](CompareADObjectsResultADObject.md) |  | [optional] 
**Status** | Pointer to [**ADObjectRestoreStatus**](ADObjectRestoreStatus.md) |  | [optional] 

## Methods

### NewADRestoreStatus

`func NewADRestoreStatus() *ADRestoreStatus`

NewADRestoreStatus instantiates a new ADRestoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADRestoreStatusWithDefaults

`func NewADRestoreStatusWithDefaults() *ADRestoreStatus`

NewADRestoreStatusWithDefaults instantiates a new ADRestoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectInfo

`func (o *ADRestoreStatus) GetObjectInfo() CompareADObjectsResultADObject`

GetObjectInfo returns the ObjectInfo field if non-nil, zero value otherwise.

### GetObjectInfoOk

`func (o *ADRestoreStatus) GetObjectInfoOk() (*CompareADObjectsResultADObject, bool)`

GetObjectInfoOk returns a tuple with the ObjectInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInfo

`func (o *ADRestoreStatus) SetObjectInfo(v CompareADObjectsResultADObject)`

SetObjectInfo sets ObjectInfo field to given value.

### HasObjectInfo

`func (o *ADRestoreStatus) HasObjectInfo() bool`

HasObjectInfo returns a boolean if a field has been set.

### GetStatus

`func (o *ADRestoreStatus) GetStatus() ADObjectRestoreStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ADRestoreStatus) GetStatusOk() (*ADObjectRestoreStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ADRestoreStatus) SetStatus(v ADObjectRestoreStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ADRestoreStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


