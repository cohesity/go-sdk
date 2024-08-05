# McmRigelClaimResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **NullableInt64** | Connection id for rigel instance. | [optional] 
**DataplaneEndpoint** | Pointer to **NullableString** | Endpoint for associated data plane. | [optional] 
**HeliosCertificate** | Pointer to **NullableString** | Specifies the Helios certificate that can be used to authenticate api calls made from Helios to Rigel. | [optional] 
**RegionId** | Pointer to **NullableString** | Specifies the region id of the Rigel cluster. | [optional] 
**RigelCaChain** | Pointer to **NullableString** | Specifies the CA chain that is used to sign the Rigel certificate. | [optional] 
**RigelCertificate** | Pointer to **NullableString** | Specifies the Rigel certificate. | [optional] 
**RigelGuid** | Pointer to **NullableInt64** | Unique id for rigel instance. | [optional] 
**RigelPrivateKey** | Pointer to **NullableString** | Specifies the Rigel private key. | [optional] 
**RigelType** | Pointer to **NullableString** | Specifies the Rigel type that is being claimed. | [optional] 
**RigelUseCase** | Pointer to **NullableString** | Specifies the Rigel use case. | [optional] 
**TenantCaChain** | Pointer to **[]string** | Specifies the Tenant CA chain. | [optional] 
**TenantId** | Pointer to **NullableString** | Tenant id associated with the claimed rigel. | [optional] 

## Methods

### NewMcmRigelClaimResponseParams

`func NewMcmRigelClaimResponseParams() *McmRigelClaimResponseParams`

NewMcmRigelClaimResponseParams instantiates a new McmRigelClaimResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmRigelClaimResponseParamsWithDefaults

`func NewMcmRigelClaimResponseParamsWithDefaults() *McmRigelClaimResponseParams`

NewMcmRigelClaimResponseParamsWithDefaults instantiates a new McmRigelClaimResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *McmRigelClaimResponseParams) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *McmRigelClaimResponseParams) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *McmRigelClaimResponseParams) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *McmRigelClaimResponseParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *McmRigelClaimResponseParams) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *McmRigelClaimResponseParams) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetDataplaneEndpoint

`func (o *McmRigelClaimResponseParams) GetDataplaneEndpoint() string`

GetDataplaneEndpoint returns the DataplaneEndpoint field if non-nil, zero value otherwise.

### GetDataplaneEndpointOk

`func (o *McmRigelClaimResponseParams) GetDataplaneEndpointOk() (*string, bool)`

GetDataplaneEndpointOk returns a tuple with the DataplaneEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataplaneEndpoint

`func (o *McmRigelClaimResponseParams) SetDataplaneEndpoint(v string)`

SetDataplaneEndpoint sets DataplaneEndpoint field to given value.

### HasDataplaneEndpoint

`func (o *McmRigelClaimResponseParams) HasDataplaneEndpoint() bool`

HasDataplaneEndpoint returns a boolean if a field has been set.

### SetDataplaneEndpointNil

`func (o *McmRigelClaimResponseParams) SetDataplaneEndpointNil(b bool)`

 SetDataplaneEndpointNil sets the value for DataplaneEndpoint to be an explicit nil

### UnsetDataplaneEndpoint
`func (o *McmRigelClaimResponseParams) UnsetDataplaneEndpoint()`

UnsetDataplaneEndpoint ensures that no value is present for DataplaneEndpoint, not even an explicit nil
### GetHeliosCertificate

`func (o *McmRigelClaimResponseParams) GetHeliosCertificate() string`

GetHeliosCertificate returns the HeliosCertificate field if non-nil, zero value otherwise.

### GetHeliosCertificateOk

`func (o *McmRigelClaimResponseParams) GetHeliosCertificateOk() (*string, bool)`

GetHeliosCertificateOk returns a tuple with the HeliosCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeliosCertificate

`func (o *McmRigelClaimResponseParams) SetHeliosCertificate(v string)`

SetHeliosCertificate sets HeliosCertificate field to given value.

### HasHeliosCertificate

`func (o *McmRigelClaimResponseParams) HasHeliosCertificate() bool`

HasHeliosCertificate returns a boolean if a field has been set.

### SetHeliosCertificateNil

`func (o *McmRigelClaimResponseParams) SetHeliosCertificateNil(b bool)`

 SetHeliosCertificateNil sets the value for HeliosCertificate to be an explicit nil

### UnsetHeliosCertificate
`func (o *McmRigelClaimResponseParams) UnsetHeliosCertificate()`

UnsetHeliosCertificate ensures that no value is present for HeliosCertificate, not even an explicit nil
### GetRegionId

`func (o *McmRigelClaimResponseParams) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *McmRigelClaimResponseParams) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *McmRigelClaimResponseParams) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *McmRigelClaimResponseParams) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### SetRegionIdNil

`func (o *McmRigelClaimResponseParams) SetRegionIdNil(b bool)`

 SetRegionIdNil sets the value for RegionId to be an explicit nil

### UnsetRegionId
`func (o *McmRigelClaimResponseParams) UnsetRegionId()`

UnsetRegionId ensures that no value is present for RegionId, not even an explicit nil
### GetRigelCaChain

`func (o *McmRigelClaimResponseParams) GetRigelCaChain() string`

GetRigelCaChain returns the RigelCaChain field if non-nil, zero value otherwise.

### GetRigelCaChainOk

`func (o *McmRigelClaimResponseParams) GetRigelCaChainOk() (*string, bool)`

GetRigelCaChainOk returns a tuple with the RigelCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelCaChain

`func (o *McmRigelClaimResponseParams) SetRigelCaChain(v string)`

SetRigelCaChain sets RigelCaChain field to given value.

### HasRigelCaChain

`func (o *McmRigelClaimResponseParams) HasRigelCaChain() bool`

HasRigelCaChain returns a boolean if a field has been set.

### SetRigelCaChainNil

`func (o *McmRigelClaimResponseParams) SetRigelCaChainNil(b bool)`

 SetRigelCaChainNil sets the value for RigelCaChain to be an explicit nil

### UnsetRigelCaChain
`func (o *McmRigelClaimResponseParams) UnsetRigelCaChain()`

UnsetRigelCaChain ensures that no value is present for RigelCaChain, not even an explicit nil
### GetRigelCertificate

`func (o *McmRigelClaimResponseParams) GetRigelCertificate() string`

GetRigelCertificate returns the RigelCertificate field if non-nil, zero value otherwise.

### GetRigelCertificateOk

`func (o *McmRigelClaimResponseParams) GetRigelCertificateOk() (*string, bool)`

GetRigelCertificateOk returns a tuple with the RigelCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelCertificate

`func (o *McmRigelClaimResponseParams) SetRigelCertificate(v string)`

SetRigelCertificate sets RigelCertificate field to given value.

### HasRigelCertificate

`func (o *McmRigelClaimResponseParams) HasRigelCertificate() bool`

HasRigelCertificate returns a boolean if a field has been set.

### SetRigelCertificateNil

`func (o *McmRigelClaimResponseParams) SetRigelCertificateNil(b bool)`

 SetRigelCertificateNil sets the value for RigelCertificate to be an explicit nil

### UnsetRigelCertificate
`func (o *McmRigelClaimResponseParams) UnsetRigelCertificate()`

UnsetRigelCertificate ensures that no value is present for RigelCertificate, not even an explicit nil
### GetRigelGuid

`func (o *McmRigelClaimResponseParams) GetRigelGuid() int64`

GetRigelGuid returns the RigelGuid field if non-nil, zero value otherwise.

### GetRigelGuidOk

`func (o *McmRigelClaimResponseParams) GetRigelGuidOk() (*int64, bool)`

GetRigelGuidOk returns a tuple with the RigelGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelGuid

`func (o *McmRigelClaimResponseParams) SetRigelGuid(v int64)`

SetRigelGuid sets RigelGuid field to given value.

### HasRigelGuid

`func (o *McmRigelClaimResponseParams) HasRigelGuid() bool`

HasRigelGuid returns a boolean if a field has been set.

### SetRigelGuidNil

`func (o *McmRigelClaimResponseParams) SetRigelGuidNil(b bool)`

 SetRigelGuidNil sets the value for RigelGuid to be an explicit nil

### UnsetRigelGuid
`func (o *McmRigelClaimResponseParams) UnsetRigelGuid()`

UnsetRigelGuid ensures that no value is present for RigelGuid, not even an explicit nil
### GetRigelPrivateKey

`func (o *McmRigelClaimResponseParams) GetRigelPrivateKey() string`

GetRigelPrivateKey returns the RigelPrivateKey field if non-nil, zero value otherwise.

### GetRigelPrivateKeyOk

`func (o *McmRigelClaimResponseParams) GetRigelPrivateKeyOk() (*string, bool)`

GetRigelPrivateKeyOk returns a tuple with the RigelPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelPrivateKey

`func (o *McmRigelClaimResponseParams) SetRigelPrivateKey(v string)`

SetRigelPrivateKey sets RigelPrivateKey field to given value.

### HasRigelPrivateKey

`func (o *McmRigelClaimResponseParams) HasRigelPrivateKey() bool`

HasRigelPrivateKey returns a boolean if a field has been set.

### SetRigelPrivateKeyNil

`func (o *McmRigelClaimResponseParams) SetRigelPrivateKeyNil(b bool)`

 SetRigelPrivateKeyNil sets the value for RigelPrivateKey to be an explicit nil

### UnsetRigelPrivateKey
`func (o *McmRigelClaimResponseParams) UnsetRigelPrivateKey()`

UnsetRigelPrivateKey ensures that no value is present for RigelPrivateKey, not even an explicit nil
### GetRigelType

`func (o *McmRigelClaimResponseParams) GetRigelType() string`

GetRigelType returns the RigelType field if non-nil, zero value otherwise.

### GetRigelTypeOk

`func (o *McmRigelClaimResponseParams) GetRigelTypeOk() (*string, bool)`

GetRigelTypeOk returns a tuple with the RigelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelType

`func (o *McmRigelClaimResponseParams) SetRigelType(v string)`

SetRigelType sets RigelType field to given value.

### HasRigelType

`func (o *McmRigelClaimResponseParams) HasRigelType() bool`

HasRigelType returns a boolean if a field has been set.

### SetRigelTypeNil

`func (o *McmRigelClaimResponseParams) SetRigelTypeNil(b bool)`

 SetRigelTypeNil sets the value for RigelType to be an explicit nil

### UnsetRigelType
`func (o *McmRigelClaimResponseParams) UnsetRigelType()`

UnsetRigelType ensures that no value is present for RigelType, not even an explicit nil
### GetRigelUseCase

`func (o *McmRigelClaimResponseParams) GetRigelUseCase() string`

GetRigelUseCase returns the RigelUseCase field if non-nil, zero value otherwise.

### GetRigelUseCaseOk

`func (o *McmRigelClaimResponseParams) GetRigelUseCaseOk() (*string, bool)`

GetRigelUseCaseOk returns a tuple with the RigelUseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelUseCase

`func (o *McmRigelClaimResponseParams) SetRigelUseCase(v string)`

SetRigelUseCase sets RigelUseCase field to given value.

### HasRigelUseCase

`func (o *McmRigelClaimResponseParams) HasRigelUseCase() bool`

HasRigelUseCase returns a boolean if a field has been set.

### SetRigelUseCaseNil

`func (o *McmRigelClaimResponseParams) SetRigelUseCaseNil(b bool)`

 SetRigelUseCaseNil sets the value for RigelUseCase to be an explicit nil

### UnsetRigelUseCase
`func (o *McmRigelClaimResponseParams) UnsetRigelUseCase()`

UnsetRigelUseCase ensures that no value is present for RigelUseCase, not even an explicit nil
### GetTenantCaChain

`func (o *McmRigelClaimResponseParams) GetTenantCaChain() []*string`

GetTenantCaChain returns the TenantCaChain field if non-nil, zero value otherwise.

### GetTenantCaChainOk

`func (o *McmRigelClaimResponseParams) GetTenantCaChainOk() (*[]*string, bool)`

GetTenantCaChainOk returns a tuple with the TenantCaChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCaChain

`func (o *McmRigelClaimResponseParams) SetTenantCaChain(v []*string)`

SetTenantCaChain sets TenantCaChain field to given value.

### HasTenantCaChain

`func (o *McmRigelClaimResponseParams) HasTenantCaChain() bool`

HasTenantCaChain returns a boolean if a field has been set.

### SetTenantCaChainNil

`func (o *McmRigelClaimResponseParams) SetTenantCaChainNil(b bool)`

 SetTenantCaChainNil sets the value for TenantCaChain to be an explicit nil

### UnsetTenantCaChain
`func (o *McmRigelClaimResponseParams) UnsetTenantCaChain()`

UnsetTenantCaChain ensures that no value is present for TenantCaChain, not even an explicit nil
### GetTenantId

`func (o *McmRigelClaimResponseParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *McmRigelClaimResponseParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *McmRigelClaimResponseParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *McmRigelClaimResponseParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *McmRigelClaimResponseParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *McmRigelClaimResponseParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


