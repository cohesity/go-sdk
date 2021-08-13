# RevertPatchesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | Specifies the name of the service whose patch should be reverted. | 
**PatchLevel** | Pointer to **NullableInt64** | Specifies optional patch level upto which the patches should be reverted. If given, it should be between 1 and the current patch level inclusive. | [optional] 

## Methods

### NewRevertPatchesRequest

`func NewRevertPatchesRequest(service string, ) *RevertPatchesRequest`

NewRevertPatchesRequest instantiates a new RevertPatchesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevertPatchesRequestWithDefaults

`func NewRevertPatchesRequestWithDefaults() *RevertPatchesRequest`

NewRevertPatchesRequestWithDefaults instantiates a new RevertPatchesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *RevertPatchesRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *RevertPatchesRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *RevertPatchesRequest) SetService(v string)`

SetService sets Service field to given value.


### GetPatchLevel

`func (o *RevertPatchesRequest) GetPatchLevel() int64`

GetPatchLevel returns the PatchLevel field if non-nil, zero value otherwise.

### GetPatchLevelOk

`func (o *RevertPatchesRequest) GetPatchLevelOk() (*int64, bool)`

GetPatchLevelOk returns a tuple with the PatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchLevel

`func (o *RevertPatchesRequest) SetPatchLevel(v int64)`

SetPatchLevel sets PatchLevel field to given value.

### HasPatchLevel

`func (o *RevertPatchesRequest) HasPatchLevel() bool`

HasPatchLevel returns a boolean if a field has been set.

### SetPatchLevelNil

`func (o *RevertPatchesRequest) SetPatchLevelNil(b bool)`

 SetPatchLevelNil sets the value for PatchLevel to be an explicit nil

### UnsetPatchLevel
`func (o *RevertPatchesRequest) UnsetPatchLevel()`

UnsetPatchLevel ensures that no value is present for PatchLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


