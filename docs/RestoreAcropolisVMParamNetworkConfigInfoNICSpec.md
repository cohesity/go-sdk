# RestoreAcropolisVMParamNetworkConfigInfoNICSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **NullableString** | IP address to assign to the NIC. | [optional] 
**NetworkUuid** | Pointer to **NullableString** | The UUID of the network to which the NIC is to be attached. | [optional] 

## Methods

### NewRestoreAcropolisVMParamNetworkConfigInfoNICSpec

`func NewRestoreAcropolisVMParamNetworkConfigInfoNICSpec() *RestoreAcropolisVMParamNetworkConfigInfoNICSpec`

NewRestoreAcropolisVMParamNetworkConfigInfoNICSpec instantiates a new RestoreAcropolisVMParamNetworkConfigInfoNICSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAcropolisVMParamNetworkConfigInfoNICSpecWithDefaults

`func NewRestoreAcropolisVMParamNetworkConfigInfoNICSpecWithDefaults() *RestoreAcropolisVMParamNetworkConfigInfoNICSpec`

NewRestoreAcropolisVMParamNetworkConfigInfoNICSpecWithDefaults instantiates a new RestoreAcropolisVMParamNetworkConfigInfoNICSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetNetworkUuid

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) GetNetworkUuid() string`

GetNetworkUuid returns the NetworkUuid field if non-nil, zero value otherwise.

### GetNetworkUuidOk

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) GetNetworkUuidOk() (*string, bool)`

GetNetworkUuidOk returns a tuple with the NetworkUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUuid

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) SetNetworkUuid(v string)`

SetNetworkUuid sets NetworkUuid field to given value.

### HasNetworkUuid

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) HasNetworkUuid() bool`

HasNetworkUuid returns a boolean if a field has been set.

### SetNetworkUuidNil

`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) SetNetworkUuidNil(b bool)`

 SetNetworkUuidNil sets the value for NetworkUuid to be an explicit nil

### UnsetNetworkUuid
`func (o *RestoreAcropolisVMParamNetworkConfigInfoNICSpec) UnsetNetworkUuid()`

UnsetNetworkUuid ensures that no value is present for NetworkUuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


