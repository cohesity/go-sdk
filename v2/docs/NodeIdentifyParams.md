# NodeIdentifyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identify** | **NullableBool** | Turn on/off node led light if set to true/false respectively. | 

## Methods

### NewNodeIdentifyParams

`func NewNodeIdentifyParams(identify NullableBool, ) *NodeIdentifyParams`

NewNodeIdentifyParams instantiates a new NodeIdentifyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeIdentifyParamsWithDefaults

`func NewNodeIdentifyParamsWithDefaults() *NodeIdentifyParams`

NewNodeIdentifyParamsWithDefaults instantiates a new NodeIdentifyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentify

`func (o *NodeIdentifyParams) GetIdentify() bool`

GetIdentify returns the Identify field if non-nil, zero value otherwise.

### GetIdentifyOk

`func (o *NodeIdentifyParams) GetIdentifyOk() (*bool, bool)`

GetIdentifyOk returns a tuple with the Identify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentify

`func (o *NodeIdentifyParams) SetIdentify(v bool)`

SetIdentify sets Identify field to given value.


### SetIdentifyNil

`func (o *NodeIdentifyParams) SetIdentifyNil(b bool)`

 SetIdentifyNil sets the value for Identify to be an explicit nil

### UnsetIdentify
`func (o *NodeIdentifyParams) UnsetIdentify()`

UnsetIdentify ensures that no value is present for Identify, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


