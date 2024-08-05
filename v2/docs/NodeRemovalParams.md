# NodeRemovalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancel** | **NullableBool** | If true, cancels node removal that is already in progress. | 
**IsOffline** | Pointer to **NullableBool** | Specifies whether node being removed is offline. | [optional] [default to false]
**IsValidateOnly** | Pointer to **NullableBool** | Specifies whether request is for pre-check validations only | [optional] [default to false]

## Methods

### NewNodeRemovalParams

`func NewNodeRemovalParams(cancel NullableBool, ) *NodeRemovalParams`

NewNodeRemovalParams instantiates a new NodeRemovalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeRemovalParamsWithDefaults

`func NewNodeRemovalParamsWithDefaults() *NodeRemovalParams`

NewNodeRemovalParamsWithDefaults instantiates a new NodeRemovalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancel

`func (o *NodeRemovalParams) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *NodeRemovalParams) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *NodeRemovalParams) SetCancel(v bool)`

SetCancel sets Cancel field to given value.


### SetCancelNil

`func (o *NodeRemovalParams) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *NodeRemovalParams) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil
### GetIsOffline

`func (o *NodeRemovalParams) GetIsOffline() bool`

GetIsOffline returns the IsOffline field if non-nil, zero value otherwise.

### GetIsOfflineOk

`func (o *NodeRemovalParams) GetIsOfflineOk() (*bool, bool)`

GetIsOfflineOk returns a tuple with the IsOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOffline

`func (o *NodeRemovalParams) SetIsOffline(v bool)`

SetIsOffline sets IsOffline field to given value.

### HasIsOffline

`func (o *NodeRemovalParams) HasIsOffline() bool`

HasIsOffline returns a boolean if a field has been set.

### SetIsOfflineNil

`func (o *NodeRemovalParams) SetIsOfflineNil(b bool)`

 SetIsOfflineNil sets the value for IsOffline to be an explicit nil

### UnsetIsOffline
`func (o *NodeRemovalParams) UnsetIsOffline()`

UnsetIsOffline ensures that no value is present for IsOffline, not even an explicit nil
### GetIsValidateOnly

`func (o *NodeRemovalParams) GetIsValidateOnly() bool`

GetIsValidateOnly returns the IsValidateOnly field if non-nil, zero value otherwise.

### GetIsValidateOnlyOk

`func (o *NodeRemovalParams) GetIsValidateOnlyOk() (*bool, bool)`

GetIsValidateOnlyOk returns a tuple with the IsValidateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidateOnly

`func (o *NodeRemovalParams) SetIsValidateOnly(v bool)`

SetIsValidateOnly sets IsValidateOnly field to given value.

### HasIsValidateOnly

`func (o *NodeRemovalParams) HasIsValidateOnly() bool`

HasIsValidateOnly returns a boolean if a field has been set.

### SetIsValidateOnlyNil

`func (o *NodeRemovalParams) SetIsValidateOnlyNil(b bool)`

 SetIsValidateOnlyNil sets the value for IsValidateOnly to be an explicit nil

### UnsetIsValidateOnly
`func (o *NodeRemovalParams) UnsetIsValidateOnly()`

UnsetIsValidateOnly ensures that no value is present for IsValidateOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


