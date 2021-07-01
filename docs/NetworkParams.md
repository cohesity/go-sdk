# NetworkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondingOpts** | Pointer to [**BondingOpts**](BondingOpts.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNetworkParams

`func NewNetworkParams() *NetworkParams`

NewNetworkParams instantiates a new NetworkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkParamsWithDefaults

`func NewNetworkParamsWithDefaults() *NetworkParams`

NewNetworkParamsWithDefaults instantiates a new NetworkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondingOpts

`func (o *NetworkParams) GetBondingOpts() BondingOpts`

GetBondingOpts returns the BondingOpts field if non-nil, zero value otherwise.

### GetBondingOptsOk

`func (o *NetworkParams) GetBondingOptsOk() (*BondingOpts, bool)`

GetBondingOptsOk returns a tuple with the BondingOpts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingOpts

`func (o *NetworkParams) SetBondingOpts(v BondingOpts)`

SetBondingOpts sets BondingOpts field to given value.

### HasBondingOpts

`func (o *NetworkParams) HasBondingOpts() bool`

HasBondingOpts returns a boolean if a field has been set.

### GetMtu

`func (o *NetworkParams) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkParams) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkParams) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkParams) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *NetworkParams) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *NetworkParams) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


