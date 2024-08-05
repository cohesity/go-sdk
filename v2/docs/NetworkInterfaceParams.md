# NetworkInterfaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeNetworkInterfaces** | Pointer to [**[]NodeNetworkInterfaces**](NodeNetworkInterfaces.md) | List of interfaces on each node. | [optional] 

## Methods

### NewNetworkInterfaceParams

`func NewNetworkInterfaceParams() *NetworkInterfaceParams`

NewNetworkInterfaceParams instantiates a new NetworkInterfaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceParamsWithDefaults

`func NewNetworkInterfaceParamsWithDefaults() *NetworkInterfaceParams`

NewNetworkInterfaceParamsWithDefaults instantiates a new NetworkInterfaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeNetworkInterfaces

`func (o *NetworkInterfaceParams) GetNodeNetworkInterfaces() []NodeNetworkInterfaces`

GetNodeNetworkInterfaces returns the NodeNetworkInterfaces field if non-nil, zero value otherwise.

### GetNodeNetworkInterfacesOk

`func (o *NetworkInterfaceParams) GetNodeNetworkInterfacesOk() (*[]NodeNetworkInterfaces, bool)`

GetNodeNetworkInterfacesOk returns a tuple with the NodeNetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNetworkInterfaces

`func (o *NetworkInterfaceParams) SetNodeNetworkInterfaces(v []NodeNetworkInterfaces)`

SetNodeNetworkInterfaces sets NodeNetworkInterfaces field to given value.

### HasNodeNetworkInterfaces

`func (o *NetworkInterfaceParams) HasNodeNetworkInterfaces() bool`

HasNodeNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


