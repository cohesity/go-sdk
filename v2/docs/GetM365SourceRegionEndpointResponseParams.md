# GetM365SourceRegionEndpointResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceAuthEndpoint** | Pointer to **NullableString** | Specifies the device authorization endpoint to be used for Microsoft graph calls. | [optional] 
**GraphEndpoint** | Pointer to **NullableString** | Specifies the Microsoft graph host url to be used for graph calls. | [optional] 
**Region** | Pointer to **NullableString** | Specifies the scope of the region. For eg NA for North America or AS for Australia. For Azure Gov cloud it can be USG or USGov. | [optional] 
**SubRegion** | Pointer to **NullableString** | Specifies the scope of the sub region. | [optional] 
**TenantRegion** | Pointer to **NullableString** | Specifies the tenant region for the given domain. This can be either Default(Commercial), GCC, GCC High or DoD. This is different from the Geo location which is represented by the region parameter. | [optional] 
**TokenEndpoint** | Pointer to **NullableString** | Specifies the token endpoint of the Microsoft365 source. | [optional] 

## Methods

### NewGetM365SourceRegionEndpointResponseParams

`func NewGetM365SourceRegionEndpointResponseParams() *GetM365SourceRegionEndpointResponseParams`

NewGetM365SourceRegionEndpointResponseParams instantiates a new GetM365SourceRegionEndpointResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetM365SourceRegionEndpointResponseParamsWithDefaults

`func NewGetM365SourceRegionEndpointResponseParamsWithDefaults() *GetM365SourceRegionEndpointResponseParams`

NewGetM365SourceRegionEndpointResponseParamsWithDefaults instantiates a new GetM365SourceRegionEndpointResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceAuthEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) GetDeviceAuthEndpoint() string`

GetDeviceAuthEndpoint returns the DeviceAuthEndpoint field if non-nil, zero value otherwise.

### GetDeviceAuthEndpointOk

`func (o *GetM365SourceRegionEndpointResponseParams) GetDeviceAuthEndpointOk() (*string, bool)`

GetDeviceAuthEndpointOk returns a tuple with the DeviceAuthEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) SetDeviceAuthEndpoint(v string)`

SetDeviceAuthEndpoint sets DeviceAuthEndpoint field to given value.

### HasDeviceAuthEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) HasDeviceAuthEndpoint() bool`

HasDeviceAuthEndpoint returns a boolean if a field has been set.

### SetDeviceAuthEndpointNil

`func (o *GetM365SourceRegionEndpointResponseParams) SetDeviceAuthEndpointNil(b bool)`

 SetDeviceAuthEndpointNil sets the value for DeviceAuthEndpoint to be an explicit nil

### UnsetDeviceAuthEndpoint
`func (o *GetM365SourceRegionEndpointResponseParams) UnsetDeviceAuthEndpoint()`

UnsetDeviceAuthEndpoint ensures that no value is present for DeviceAuthEndpoint, not even an explicit nil
### GetGraphEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) GetGraphEndpoint() string`

GetGraphEndpoint returns the GraphEndpoint field if non-nil, zero value otherwise.

### GetGraphEndpointOk

`func (o *GetM365SourceRegionEndpointResponseParams) GetGraphEndpointOk() (*string, bool)`

GetGraphEndpointOk returns a tuple with the GraphEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) SetGraphEndpoint(v string)`

SetGraphEndpoint sets GraphEndpoint field to given value.

### HasGraphEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) HasGraphEndpoint() bool`

HasGraphEndpoint returns a boolean if a field has been set.

### SetGraphEndpointNil

`func (o *GetM365SourceRegionEndpointResponseParams) SetGraphEndpointNil(b bool)`

 SetGraphEndpointNil sets the value for GraphEndpoint to be an explicit nil

### UnsetGraphEndpoint
`func (o *GetM365SourceRegionEndpointResponseParams) UnsetGraphEndpoint()`

UnsetGraphEndpoint ensures that no value is present for GraphEndpoint, not even an explicit nil
### GetRegion

`func (o *GetM365SourceRegionEndpointResponseParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetM365SourceRegionEndpointResponseParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetM365SourceRegionEndpointResponseParams) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetM365SourceRegionEndpointResponseParams) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *GetM365SourceRegionEndpointResponseParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *GetM365SourceRegionEndpointResponseParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSubRegion

`func (o *GetM365SourceRegionEndpointResponseParams) GetSubRegion() string`

GetSubRegion returns the SubRegion field if non-nil, zero value otherwise.

### GetSubRegionOk

`func (o *GetM365SourceRegionEndpointResponseParams) GetSubRegionOk() (*string, bool)`

GetSubRegionOk returns a tuple with the SubRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRegion

`func (o *GetM365SourceRegionEndpointResponseParams) SetSubRegion(v string)`

SetSubRegion sets SubRegion field to given value.

### HasSubRegion

`func (o *GetM365SourceRegionEndpointResponseParams) HasSubRegion() bool`

HasSubRegion returns a boolean if a field has been set.

### SetSubRegionNil

`func (o *GetM365SourceRegionEndpointResponseParams) SetSubRegionNil(b bool)`

 SetSubRegionNil sets the value for SubRegion to be an explicit nil

### UnsetSubRegion
`func (o *GetM365SourceRegionEndpointResponseParams) UnsetSubRegion()`

UnsetSubRegion ensures that no value is present for SubRegion, not even an explicit nil
### GetTenantRegion

`func (o *GetM365SourceRegionEndpointResponseParams) GetTenantRegion() string`

GetTenantRegion returns the TenantRegion field if non-nil, zero value otherwise.

### GetTenantRegionOk

`func (o *GetM365SourceRegionEndpointResponseParams) GetTenantRegionOk() (*string, bool)`

GetTenantRegionOk returns a tuple with the TenantRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantRegion

`func (o *GetM365SourceRegionEndpointResponseParams) SetTenantRegion(v string)`

SetTenantRegion sets TenantRegion field to given value.

### HasTenantRegion

`func (o *GetM365SourceRegionEndpointResponseParams) HasTenantRegion() bool`

HasTenantRegion returns a boolean if a field has been set.

### SetTenantRegionNil

`func (o *GetM365SourceRegionEndpointResponseParams) SetTenantRegionNil(b bool)`

 SetTenantRegionNil sets the value for TenantRegion to be an explicit nil

### UnsetTenantRegion
`func (o *GetM365SourceRegionEndpointResponseParams) UnsetTenantRegion()`

UnsetTenantRegion ensures that no value is present for TenantRegion, not even an explicit nil
### GetTokenEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *GetM365SourceRegionEndpointResponseParams) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *GetM365SourceRegionEndpointResponseParams) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### SetTokenEndpointNil

`func (o *GetM365SourceRegionEndpointResponseParams) SetTokenEndpointNil(b bool)`

 SetTokenEndpointNil sets the value for TokenEndpoint to be an explicit nil

### UnsetTokenEndpoint
`func (o *GetM365SourceRegionEndpointResponseParams) UnsetTokenEndpoint()`

UnsetTokenEndpoint ensures that no value is present for TokenEndpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


