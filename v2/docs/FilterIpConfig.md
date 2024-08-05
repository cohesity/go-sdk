# FilterIpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedIpAddresses** | Pointer to **[]string** | Specifies the IP addresses that should be used exclusively during recovery. Cannot be set if deniedIpAddresses is set. | [optional] 
**DeniedIpAddresses** | Pointer to **[]string** | Specifies the IP addresses that should not be used during recovery recovery. Cannot be set if allowedIpAddresses is set. | [optional] 

## Methods

### NewFilterIpConfig

`func NewFilterIpConfig() *FilterIpConfig`

NewFilterIpConfig instantiates a new FilterIpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterIpConfigWithDefaults

`func NewFilterIpConfigWithDefaults() *FilterIpConfig`

NewFilterIpConfigWithDefaults instantiates a new FilterIpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedIpAddresses

`func (o *FilterIpConfig) GetAllowedIpAddresses() []string`

GetAllowedIpAddresses returns the AllowedIpAddresses field if non-nil, zero value otherwise.

### GetAllowedIpAddressesOk

`func (o *FilterIpConfig) GetAllowedIpAddressesOk() (*[]string, bool)`

GetAllowedIpAddressesOk returns a tuple with the AllowedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpAddresses

`func (o *FilterIpConfig) SetAllowedIpAddresses(v []string)`

SetAllowedIpAddresses sets AllowedIpAddresses field to given value.

### HasAllowedIpAddresses

`func (o *FilterIpConfig) HasAllowedIpAddresses() bool`

HasAllowedIpAddresses returns a boolean if a field has been set.

### SetAllowedIpAddressesNil

`func (o *FilterIpConfig) SetAllowedIpAddressesNil(b bool)`

 SetAllowedIpAddressesNil sets the value for AllowedIpAddresses to be an explicit nil

### UnsetAllowedIpAddresses
`func (o *FilterIpConfig) UnsetAllowedIpAddresses()`

UnsetAllowedIpAddresses ensures that no value is present for AllowedIpAddresses, not even an explicit nil
### GetDeniedIpAddresses

`func (o *FilterIpConfig) GetDeniedIpAddresses() []string`

GetDeniedIpAddresses returns the DeniedIpAddresses field if non-nil, zero value otherwise.

### GetDeniedIpAddressesOk

`func (o *FilterIpConfig) GetDeniedIpAddressesOk() (*[]string, bool)`

GetDeniedIpAddressesOk returns a tuple with the DeniedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedIpAddresses

`func (o *FilterIpConfig) SetDeniedIpAddresses(v []string)`

SetDeniedIpAddresses sets DeniedIpAddresses field to given value.

### HasDeniedIpAddresses

`func (o *FilterIpConfig) HasDeniedIpAddresses() bool`

HasDeniedIpAddresses returns a boolean if a field has been set.

### SetDeniedIpAddressesNil

`func (o *FilterIpConfig) SetDeniedIpAddressesNil(b bool)`

 SetDeniedIpAddressesNil sets the value for DeniedIpAddresses to be an explicit nil

### UnsetDeniedIpAddresses
`func (o *FilterIpConfig) UnsetDeniedIpAddresses()`

UnsetDeniedIpAddresses ensures that no value is present for DeniedIpAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


