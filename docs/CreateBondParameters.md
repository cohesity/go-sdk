# CreateBondParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BondingMode** | Pointer to **NullableString** | Specifies the bonding mode to use for this bond. If not specified, this value will default to &#39;kActiveBackup&#39;. &#39;kActiveBackup&#39; indicates active backup bonding mode. &#39;k802_3ad&#39; indicates 802.3ad bonding mode. | [optional] 
**Name** | **NullableString** | Specifies a unique name to identify the bond being created. | 
**Slaves** | **[]string** | Specifies the names of the slaves of this bond. | 

## Methods

### NewCreateBondParameters

`func NewCreateBondParameters(name NullableString, slaves []string, ) *CreateBondParameters`

NewCreateBondParameters instantiates a new CreateBondParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBondParametersWithDefaults

`func NewCreateBondParametersWithDefaults() *CreateBondParameters`

NewCreateBondParametersWithDefaults instantiates a new CreateBondParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBondingMode

`func (o *CreateBondParameters) GetBondingMode() string`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *CreateBondParameters) GetBondingModeOk() (*string, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *CreateBondParameters) SetBondingMode(v string)`

SetBondingMode sets BondingMode field to given value.

### HasBondingMode

`func (o *CreateBondParameters) HasBondingMode() bool`

HasBondingMode returns a boolean if a field has been set.

### SetBondingModeNil

`func (o *CreateBondParameters) SetBondingModeNil(b bool)`

 SetBondingModeNil sets the value for BondingMode to be an explicit nil

### UnsetBondingMode
`func (o *CreateBondParameters) UnsetBondingMode()`

UnsetBondingMode ensures that no value is present for BondingMode, not even an explicit nil
### GetName

`func (o *CreateBondParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBondParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBondParameters) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateBondParameters) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateBondParameters) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSlaves

`func (o *CreateBondParameters) GetSlaves() []string`

GetSlaves returns the Slaves field if non-nil, zero value otherwise.

### GetSlavesOk

`func (o *CreateBondParameters) GetSlavesOk() (*[]string, bool)`

GetSlavesOk returns a tuple with the Slaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaves

`func (o *CreateBondParameters) SetSlaves(v []string)`

SetSlaves sets Slaves field to given value.


### SetSlavesNil

`func (o *CreateBondParameters) SetSlavesNil(b bool)`

 SetSlavesNil sets the value for Slaves to be an explicit nil

### UnsetSlaves
`func (o *CreateBondParameters) UnsetSlaves()`

UnsetSlaves ensures that no value is present for Slaves, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


