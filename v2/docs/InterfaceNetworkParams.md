# InterfaceNetworkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondInterfaceParams** | Pointer to [**BondInterfaceNetworkParams**](BondInterfaceNetworkParams.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** | MTU of the network interface. | [optional] 

## Methods

### NewInterfaceNetworkParams

`func NewInterfaceNetworkParams() *InterfaceNetworkParams`

NewInterfaceNetworkParams instantiates a new InterfaceNetworkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceNetworkParamsWithDefaults

`func NewInterfaceNetworkParamsWithDefaults() *InterfaceNetworkParams`

NewInterfaceNetworkParamsWithDefaults instantiates a new InterfaceNetworkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondInterfaceParams

`func (o *InterfaceNetworkParams) GetBondInterfaceParams() BondInterfaceNetworkParams`

GetBondInterfaceParams returns the BondInterfaceParams field if non-nil, zero value otherwise.

### GetBondInterfaceParamsOk

`func (o *InterfaceNetworkParams) GetBondInterfaceParamsOk() (*BondInterfaceNetworkParams, bool)`

GetBondInterfaceParamsOk returns a tuple with the BondInterfaceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondInterfaceParams

`func (o *InterfaceNetworkParams) SetBondInterfaceParams(v BondInterfaceNetworkParams)`

SetBondInterfaceParams sets BondInterfaceParams field to given value.

### HasBondInterfaceParams

`func (o *InterfaceNetworkParams) HasBondInterfaceParams() bool`

HasBondInterfaceParams returns a boolean if a field has been set.

### GetMtu

`func (o *InterfaceNetworkParams) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InterfaceNetworkParams) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InterfaceNetworkParams) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InterfaceNetworkParams) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InterfaceNetworkParams) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InterfaceNetworkParams) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


