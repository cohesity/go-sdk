# NodeRemovalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancel** | **NullableBool** | If true, cancels node removal that is already in progress. | 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


