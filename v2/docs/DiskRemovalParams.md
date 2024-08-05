# DiskRemovalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancel** | **NullableBool** | If true, cancels disk removal which is already in progress. | 
**IsValidateOnly** | Pointer to **NullableBool** | Specifies whether request is for pre-check validations only | [optional] [default to false]

## Methods

### NewDiskRemovalParams

`func NewDiskRemovalParams(cancel NullableBool, ) *DiskRemovalParams`

NewDiskRemovalParams instantiates a new DiskRemovalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskRemovalParamsWithDefaults

`func NewDiskRemovalParamsWithDefaults() *DiskRemovalParams`

NewDiskRemovalParamsWithDefaults instantiates a new DiskRemovalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancel

`func (o *DiskRemovalParams) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *DiskRemovalParams) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *DiskRemovalParams) SetCancel(v bool)`

SetCancel sets Cancel field to given value.


### SetCancelNil

`func (o *DiskRemovalParams) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *DiskRemovalParams) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil
### GetIsValidateOnly

`func (o *DiskRemovalParams) GetIsValidateOnly() bool`

GetIsValidateOnly returns the IsValidateOnly field if non-nil, zero value otherwise.

### GetIsValidateOnlyOk

`func (o *DiskRemovalParams) GetIsValidateOnlyOk() (*bool, bool)`

GetIsValidateOnlyOk returns a tuple with the IsValidateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidateOnly

`func (o *DiskRemovalParams) SetIsValidateOnly(v bool)`

SetIsValidateOnly sets IsValidateOnly field to given value.

### HasIsValidateOnly

`func (o *DiskRemovalParams) HasIsValidateOnly() bool`

HasIsValidateOnly returns a boolean if a field has been set.

### SetIsValidateOnlyNil

`func (o *DiskRemovalParams) SetIsValidateOnlyNil(b bool)`

 SetIsValidateOnlyNil sets the value for IsValidateOnly to be an explicit nil

### UnsetIsValidateOnly
`func (o *DiskRemovalParams) UnsetIsValidateOnly()`

UnsetIsValidateOnly ensures that no value is present for IsValidateOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


