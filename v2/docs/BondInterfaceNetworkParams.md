# BondInterfaceNetworkParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondingMode** | Pointer to **NullableString** | Bonding mode of the interface. | [optional] 
**LacpRate** | Pointer to **NullableString** | Rate option to use for link partner to transmit LACPDU packets in 802.3ad mode. | [optional] 
**XmitHashPolicy** | Pointer to **NullableString** | Transmit hash policy to use for selection in 802.3ad mode. | [optional] 

## Methods

### NewBondInterfaceNetworkParams

`func NewBondInterfaceNetworkParams() *BondInterfaceNetworkParams`

NewBondInterfaceNetworkParams instantiates a new BondInterfaceNetworkParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondInterfaceNetworkParamsWithDefaults

`func NewBondInterfaceNetworkParamsWithDefaults() *BondInterfaceNetworkParams`

NewBondInterfaceNetworkParamsWithDefaults instantiates a new BondInterfaceNetworkParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondingMode

`func (o *BondInterfaceNetworkParams) GetBondingMode() string`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *BondInterfaceNetworkParams) GetBondingModeOk() (*string, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *BondInterfaceNetworkParams) SetBondingMode(v string)`

SetBondingMode sets BondingMode field to given value.

### HasBondingMode

`func (o *BondInterfaceNetworkParams) HasBondingMode() bool`

HasBondingMode returns a boolean if a field has been set.

### SetBondingModeNil

`func (o *BondInterfaceNetworkParams) SetBondingModeNil(b bool)`

 SetBondingModeNil sets the value for BondingMode to be an explicit nil

### UnsetBondingMode
`func (o *BondInterfaceNetworkParams) UnsetBondingMode()`

UnsetBondingMode ensures that no value is present for BondingMode, not even an explicit nil
### GetLacpRate

`func (o *BondInterfaceNetworkParams) GetLacpRate() string`

GetLacpRate returns the LacpRate field if non-nil, zero value otherwise.

### GetLacpRateOk

`func (o *BondInterfaceNetworkParams) GetLacpRateOk() (*string, bool)`

GetLacpRateOk returns a tuple with the LacpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacpRate

`func (o *BondInterfaceNetworkParams) SetLacpRate(v string)`

SetLacpRate sets LacpRate field to given value.

### HasLacpRate

`func (o *BondInterfaceNetworkParams) HasLacpRate() bool`

HasLacpRate returns a boolean if a field has been set.

### SetLacpRateNil

`func (o *BondInterfaceNetworkParams) SetLacpRateNil(b bool)`

 SetLacpRateNil sets the value for LacpRate to be an explicit nil

### UnsetLacpRate
`func (o *BondInterfaceNetworkParams) UnsetLacpRate()`

UnsetLacpRate ensures that no value is present for LacpRate, not even an explicit nil
### GetXmitHashPolicy

`func (o *BondInterfaceNetworkParams) GetXmitHashPolicy() string`

GetXmitHashPolicy returns the XmitHashPolicy field if non-nil, zero value otherwise.

### GetXmitHashPolicyOk

`func (o *BondInterfaceNetworkParams) GetXmitHashPolicyOk() (*string, bool)`

GetXmitHashPolicyOk returns a tuple with the XmitHashPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmitHashPolicy

`func (o *BondInterfaceNetworkParams) SetXmitHashPolicy(v string)`

SetXmitHashPolicy sets XmitHashPolicy field to given value.

### HasXmitHashPolicy

`func (o *BondInterfaceNetworkParams) HasXmitHashPolicy() bool`

HasXmitHashPolicy returns a boolean if a field has been set.

### SetXmitHashPolicyNil

`func (o *BondInterfaceNetworkParams) SetXmitHashPolicyNil(b bool)`

 SetXmitHashPolicyNil sets the value for XmitHashPolicy to be an explicit nil

### UnsetXmitHashPolicy
`func (o *BondInterfaceNetworkParams) UnsetXmitHashPolicy()`

UnsetXmitHashPolicy ensures that no value is present for XmitHashPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


