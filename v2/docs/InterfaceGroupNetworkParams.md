# InterfaceGroupNetworkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondInterfaceParams** | Pointer to [**BondInterfaceNetworkParams**](BondInterfaceNetworkParams.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** | MTU of the network interface group. | [optional] 

## Methods

### NewInterfaceGroupNetworkParams

`func NewInterfaceGroupNetworkParams() *InterfaceGroupNetworkParams`

NewInterfaceGroupNetworkParams instantiates a new InterfaceGroupNetworkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceGroupNetworkParamsWithDefaults

`func NewInterfaceGroupNetworkParamsWithDefaults() *InterfaceGroupNetworkParams`

NewInterfaceGroupNetworkParamsWithDefaults instantiates a new InterfaceGroupNetworkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondInterfaceParams

`func (o *InterfaceGroupNetworkParams) GetBondInterfaceParams() BondInterfaceNetworkParams`

GetBondInterfaceParams returns the BondInterfaceParams field if non-nil, zero value otherwise.

### GetBondInterfaceParamsOk

`func (o *InterfaceGroupNetworkParams) GetBondInterfaceParamsOk() (*BondInterfaceNetworkParams, bool)`

GetBondInterfaceParamsOk returns a tuple with the BondInterfaceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondInterfaceParams

`func (o *InterfaceGroupNetworkParams) SetBondInterfaceParams(v BondInterfaceNetworkParams)`

SetBondInterfaceParams sets BondInterfaceParams field to given value.

### HasBondInterfaceParams

`func (o *InterfaceGroupNetworkParams) HasBondInterfaceParams() bool`

HasBondInterfaceParams returns a boolean if a field has been set.

### GetMtu

`func (o *InterfaceGroupNetworkParams) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InterfaceGroupNetworkParams) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InterfaceGroupNetworkParams) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InterfaceGroupNetworkParams) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InterfaceGroupNetworkParams) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InterfaceGroupNetworkParams) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


