# InterfaceGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the interface group. | 
**NetworkParams** | Pointer to [**InterfaceGroupNetworkParams**](InterfaceGroupNetworkParams.md) |  | [optional] 
**NodeInterfaceParams** | [**[]NodeInterfaceParams**](NodeInterfaceParams.md) | Node and interface parameters. | 
**Type** | **string** | Type of the interface group. | 

## Methods

### NewInterfaceGroupParams

`func NewInterfaceGroupParams(name string, nodeInterfaceParams []NodeInterfaceParams, type_ string, ) *InterfaceGroupParams`

NewInterfaceGroupParams instantiates a new InterfaceGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceGroupParamsWithDefaults

`func NewInterfaceGroupParamsWithDefaults() *InterfaceGroupParams`

NewInterfaceGroupParamsWithDefaults instantiates a new InterfaceGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InterfaceGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceGroupParams) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkParams

`func (o *InterfaceGroupParams) GetNetworkParams() InterfaceGroupNetworkParams`

GetNetworkParams returns the NetworkParams field if non-nil, zero value otherwise.

### GetNetworkParamsOk

`func (o *InterfaceGroupParams) GetNetworkParamsOk() (*InterfaceGroupNetworkParams, bool)`

GetNetworkParamsOk returns a tuple with the NetworkParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkParams

`func (o *InterfaceGroupParams) SetNetworkParams(v InterfaceGroupNetworkParams)`

SetNetworkParams sets NetworkParams field to given value.

### HasNetworkParams

`func (o *InterfaceGroupParams) HasNetworkParams() bool`

HasNetworkParams returns a boolean if a field has been set.

### GetNodeInterfaceParams

`func (o *InterfaceGroupParams) GetNodeInterfaceParams() []NodeInterfaceParams`

GetNodeInterfaceParams returns the NodeInterfaceParams field if non-nil, zero value otherwise.

### GetNodeInterfaceParamsOk

`func (o *InterfaceGroupParams) GetNodeInterfaceParamsOk() (*[]NodeInterfaceParams, bool)`

GetNodeInterfaceParamsOk returns a tuple with the NodeInterfaceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInterfaceParams

`func (o *InterfaceGroupParams) SetNodeInterfaceParams(v []NodeInterfaceParams)`

SetNodeInterfaceParams sets NodeInterfaceParams field to given value.


### GetType

`func (o *InterfaceGroupParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceGroupParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceGroupParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


