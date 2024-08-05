# ApplyPatchesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **NullableBool** | Specifies all the available patches should be applied. | [optional] 
**Service** | Pointer to **NullableString** | Specifies the name of the service whose patch should be applied. | [optional] 

## Methods

### NewApplyPatchesRequest

`func NewApplyPatchesRequest() *ApplyPatchesRequest`

NewApplyPatchesRequest instantiates a new ApplyPatchesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyPatchesRequestWithDefaults

`func NewApplyPatchesRequestWithDefaults() *ApplyPatchesRequest`

NewApplyPatchesRequestWithDefaults instantiates a new ApplyPatchesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *ApplyPatchesRequest) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ApplyPatchesRequest) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ApplyPatchesRequest) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ApplyPatchesRequest) HasAll() bool`

HasAll returns a boolean if a field has been set.

### SetAllNil

`func (o *ApplyPatchesRequest) SetAllNil(b bool)`

 SetAllNil sets the value for All to be an explicit nil

### UnsetAll
`func (o *ApplyPatchesRequest) UnsetAll()`

UnsetAll ensures that no value is present for All, not even an explicit nil
### GetService

`func (o *ApplyPatchesRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ApplyPatchesRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ApplyPatchesRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ApplyPatchesRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *ApplyPatchesRequest) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *ApplyPatchesRequest) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


