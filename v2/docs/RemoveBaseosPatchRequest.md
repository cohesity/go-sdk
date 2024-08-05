# RemoveBaseosPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceRemove** | Pointer to **NullableBool** | If patch files should be removed even for inprogress patch | [optional] 
**PatchName** | **string** | Name of the hotfix with security patch | 

## Methods

### NewRemoveBaseosPatchRequest

`func NewRemoveBaseosPatchRequest(patchName string, ) *RemoveBaseosPatchRequest`

NewRemoveBaseosPatchRequest instantiates a new RemoveBaseosPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveBaseosPatchRequestWithDefaults

`func NewRemoveBaseosPatchRequestWithDefaults() *RemoveBaseosPatchRequest`

NewRemoveBaseosPatchRequestWithDefaults instantiates a new RemoveBaseosPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceRemove

`func (o *RemoveBaseosPatchRequest) GetForceRemove() bool`

GetForceRemove returns the ForceRemove field if non-nil, zero value otherwise.

### GetForceRemoveOk

`func (o *RemoveBaseosPatchRequest) GetForceRemoveOk() (*bool, bool)`

GetForceRemoveOk returns a tuple with the ForceRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRemove

`func (o *RemoveBaseosPatchRequest) SetForceRemove(v bool)`

SetForceRemove sets ForceRemove field to given value.

### HasForceRemove

`func (o *RemoveBaseosPatchRequest) HasForceRemove() bool`

HasForceRemove returns a boolean if a field has been set.

### SetForceRemoveNil

`func (o *RemoveBaseosPatchRequest) SetForceRemoveNil(b bool)`

 SetForceRemoveNil sets the value for ForceRemove to be an explicit nil

### UnsetForceRemove
`func (o *RemoveBaseosPatchRequest) UnsetForceRemove()`

UnsetForceRemove ensures that no value is present for ForceRemove, not even an explicit nil
### GetPatchName

`func (o *RemoveBaseosPatchRequest) GetPatchName() string`

GetPatchName returns the PatchName field if non-nil, zero value otherwise.

### GetPatchNameOk

`func (o *RemoveBaseosPatchRequest) GetPatchNameOk() (*string, bool)`

GetPatchNameOk returns a tuple with the PatchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchName

`func (o *RemoveBaseosPatchRequest) SetPatchName(v string)`

SetPatchName sets PatchName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


