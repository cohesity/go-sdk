# TenantProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstituentId** | Pointer to **NullableInt64** | Specifies the constituent id of the proxy. | [optional] 
**IpAddress** | Pointer to **NullableString** | Specifies the ip address of the proxy. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantProxy

`func NewTenantProxy() *TenantProxy`

NewTenantProxy instantiates a new TenantProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantProxyWithDefaults

`func NewTenantProxyWithDefaults() *TenantProxy`

NewTenantProxyWithDefaults instantiates a new TenantProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstituentId

`func (o *TenantProxy) GetConstituentId() int64`

GetConstituentId returns the ConstituentId field if non-nil, zero value otherwise.

### GetConstituentIdOk

`func (o *TenantProxy) GetConstituentIdOk() (*int64, bool)`

GetConstituentIdOk returns a tuple with the ConstituentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituentId

`func (o *TenantProxy) SetConstituentId(v int64)`

SetConstituentId sets ConstituentId field to given value.

### HasConstituentId

`func (o *TenantProxy) HasConstituentId() bool`

HasConstituentId returns a boolean if a field has been set.

### SetConstituentIdNil

`func (o *TenantProxy) SetConstituentIdNil(b bool)`

 SetConstituentIdNil sets the value for ConstituentId to be an explicit nil

### UnsetConstituentId
`func (o *TenantProxy) UnsetConstituentId()`

UnsetConstituentId ensures that no value is present for ConstituentId, not even an explicit nil
### GetIpAddress

`func (o *TenantProxy) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *TenantProxy) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *TenantProxy) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *TenantProxy) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *TenantProxy) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *TenantProxy) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetTenantId

`func (o *TenantProxy) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantProxy) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantProxy) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantProxy) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantProxy) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantProxy) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


