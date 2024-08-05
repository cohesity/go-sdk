# InterfaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the interface. | 
**NetworkParams** | Pointer to [**InterfaceNetworkParams**](InterfaceNetworkParams.md) |  | [optional] 

## Methods

### NewInterfaceParams

`func NewInterfaceParams(name string, ) *InterfaceParams`

NewInterfaceParams instantiates a new InterfaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceParamsWithDefaults

`func NewInterfaceParamsWithDefaults() *InterfaceParams`

NewInterfaceParamsWithDefaults instantiates a new InterfaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InterfaceParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceParams) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkParams

`func (o *InterfaceParams) GetNetworkParams() InterfaceNetworkParams`

GetNetworkParams returns the NetworkParams field if non-nil, zero value otherwise.

### GetNetworkParamsOk

`func (o *InterfaceParams) GetNetworkParamsOk() (*InterfaceNetworkParams, bool)`

GetNetworkParamsOk returns a tuple with the NetworkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkParams

`func (o *InterfaceParams) SetNetworkParams(v InterfaceNetworkParams)`

SetNetworkParams sets NetworkParams field to given value.

### HasNetworkParams

`func (o *InterfaceParams) HasNetworkParams() bool`

HasNetworkParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


