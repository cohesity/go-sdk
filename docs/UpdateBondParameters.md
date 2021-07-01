# UpdateBondParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondingMode** | **NullableString** | Specifies the new bonding mode. &#39;kActiveBackup&#39; indicates active backup bonding mode. &#39;k802_3ad&#39; indicates 802.3ad bonding mode. | 
**LacpRate** | Pointer to **NullableString** | Specifies the LACP rate. If not specified, This value will default to 0 (slow). | [optional] 
**Name** | **NullableString** | Specifies the name of the bond being updated. | 
**XmitHashPolicy** | Pointer to **NullableString** | Specifies the xmit hash policy. If not specified, This value will default to 1 (layer3+4). | [optional] 

## Methods

### NewUpdateBondParameters

`func NewUpdateBondParameters(bondingMode NullableString, name NullableString, ) *UpdateBondParameters`

NewUpdateBondParameters instantiates a new UpdateBondParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBondParametersWithDefaults

`func NewUpdateBondParametersWithDefaults() *UpdateBondParameters`

NewUpdateBondParametersWithDefaults instantiates a new UpdateBondParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondingMode

`func (o *UpdateBondParameters) GetBondingMode() string`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *UpdateBondParameters) GetBondingModeOk() (*string, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *UpdateBondParameters) SetBondingMode(v string)`

SetBondingMode sets BondingMode field to given value.


### SetBondingModeNil

`func (o *UpdateBondParameters) SetBondingModeNil(b bool)`

 SetBondingModeNil sets the value for BondingMode to be an explicit nil

### UnsetBondingMode
`func (o *UpdateBondParameters) UnsetBondingMode()`

UnsetBondingMode ensures that no value is present for BondingMode, not even an explicit nil
### GetLacpRate

`func (o *UpdateBondParameters) GetLacpRate() string`

GetLacpRate returns the LacpRate field if non-nil, zero value otherwise.

### GetLacpRateOk

`func (o *UpdateBondParameters) GetLacpRateOk() (*string, bool)`

GetLacpRateOk returns a tuple with the LacpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacpRate

`func (o *UpdateBondParameters) SetLacpRate(v string)`

SetLacpRate sets LacpRate field to given value.

### HasLacpRate

`func (o *UpdateBondParameters) HasLacpRate() bool`

HasLacpRate returns a boolean if a field has been set.

### SetLacpRateNil

`func (o *UpdateBondParameters) SetLacpRateNil(b bool)`

 SetLacpRateNil sets the value for LacpRate to be an explicit nil

### UnsetLacpRate
`func (o *UpdateBondParameters) UnsetLacpRate()`

UnsetLacpRate ensures that no value is present for LacpRate, not even an explicit nil
### GetName

`func (o *UpdateBondParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBondParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBondParameters) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateBondParameters) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateBondParameters) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetXmitHashPolicy

`func (o *UpdateBondParameters) GetXmitHashPolicy() string`

GetXmitHashPolicy returns the XmitHashPolicy field if non-nil, zero value otherwise.

### GetXmitHashPolicyOk

`func (o *UpdateBondParameters) GetXmitHashPolicyOk() (*string, bool)`

GetXmitHashPolicyOk returns a tuple with the XmitHashPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmitHashPolicy

`func (o *UpdateBondParameters) SetXmitHashPolicy(v string)`

SetXmitHashPolicy sets XmitHashPolicy field to given value.

### HasXmitHashPolicy

`func (o *UpdateBondParameters) HasXmitHashPolicy() bool`

HasXmitHashPolicy returns a boolean if a field has been set.

### SetXmitHashPolicyNil

`func (o *UpdateBondParameters) SetXmitHashPolicyNil(b bool)`

 SetXmitHashPolicyNil sets the value for XmitHashPolicy to be an explicit nil

### UnsetXmitHashPolicy
`func (o *UpdateBondParameters) UnsetXmitHashPolicy()`

UnsetXmitHashPolicy ensures that no value is present for XmitHashPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


