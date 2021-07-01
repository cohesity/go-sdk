# ExternalNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AxonNetworkName** | Pointer to **NullableString** | Name of axon network corresponding to this external network. Used for internal purposes only. | [optional] 
**VlanId** | Pointer to **NullableInt32** | VLAN id of the network. | [optional] 
**VlanNetworkName** | Pointer to **NullableString** | Display name for the UI to select the external network. | [optional] 

## Methods

### NewExternalNetworkInfo

`func NewExternalNetworkInfo() *ExternalNetworkInfo`

NewExternalNetworkInfo instantiates a new ExternalNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalNetworkInfoWithDefaults

`func NewExternalNetworkInfoWithDefaults() *ExternalNetworkInfo`

NewExternalNetworkInfoWithDefaults instantiates a new ExternalNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxonNetworkName

`func (o *ExternalNetworkInfo) GetAxonNetworkName() string`

GetAxonNetworkName returns the AxonNetworkName field if non-nil, zero value otherwise.

### GetAxonNetworkNameOk

`func (o *ExternalNetworkInfo) GetAxonNetworkNameOk() (*string, bool)`

GetAxonNetworkNameOk returns a tuple with the AxonNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxonNetworkName

`func (o *ExternalNetworkInfo) SetAxonNetworkName(v string)`

SetAxonNetworkName sets AxonNetworkName field to given value.

### HasAxonNetworkName

`func (o *ExternalNetworkInfo) HasAxonNetworkName() bool`

HasAxonNetworkName returns a boolean if a field has been set.

### SetAxonNetworkNameNil

`func (o *ExternalNetworkInfo) SetAxonNetworkNameNil(b bool)`

 SetAxonNetworkNameNil sets the value for AxonNetworkName to be an explicit nil

### UnsetAxonNetworkName
`func (o *ExternalNetworkInfo) UnsetAxonNetworkName()`

UnsetAxonNetworkName ensures that no value is present for AxonNetworkName, not even an explicit nil
### GetVlanId

`func (o *ExternalNetworkInfo) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ExternalNetworkInfo) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ExternalNetworkInfo) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ExternalNetworkInfo) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *ExternalNetworkInfo) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *ExternalNetworkInfo) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetVlanNetworkName

`func (o *ExternalNetworkInfo) GetVlanNetworkName() string`

GetVlanNetworkName returns the VlanNetworkName field if non-nil, zero value otherwise.

### GetVlanNetworkNameOk

`func (o *ExternalNetworkInfo) GetVlanNetworkNameOk() (*string, bool)`

GetVlanNetworkNameOk returns a tuple with the VlanNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNetworkName

`func (o *ExternalNetworkInfo) SetVlanNetworkName(v string)`

SetVlanNetworkName sets VlanNetworkName field to given value.

### HasVlanNetworkName

`func (o *ExternalNetworkInfo) HasVlanNetworkName() bool`

HasVlanNetworkName returns a boolean if a field has been set.

### SetVlanNetworkNameNil

`func (o *ExternalNetworkInfo) SetVlanNetworkNameNil(b bool)`

 SetVlanNetworkNameNil sets the value for VlanNetworkName to be an explicit nil

### UnsetVlanNetworkName
`func (o *ExternalNetworkInfo) UnsetVlanNetworkName()`

UnsetVlanNetworkName ensures that no value is present for VlanNetworkName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


