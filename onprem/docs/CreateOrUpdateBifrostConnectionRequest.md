# CreateOrUpdateBifrostConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **NullableString** | Specifies the id of the tenant which the connection belongs to. | 
**Name** | Pointer to **NullableString** | Specifies the name of the connection. | [optional] 
**Subnet** | [**ConnectionSubnet**](ConnectionSubnet.md) |  | 
**CertificateVersion** | Pointer to **NullableInt64** | Specifies the version of the connection&#39;s certificate. The version is used to revoke/renew connection&#39;s certificates. | [optional] 

## Methods

### NewCreateOrUpdateBifrostConnectionRequest

`func NewCreateOrUpdateBifrostConnectionRequest(tenantId NullableString, subnet ConnectionSubnet, ) *CreateOrUpdateBifrostConnectionRequest`

NewCreateOrUpdateBifrostConnectionRequest instantiates a new CreateOrUpdateBifrostConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateBifrostConnectionRequestWithDefaults

`func NewCreateOrUpdateBifrostConnectionRequestWithDefaults() *CreateOrUpdateBifrostConnectionRequest`

NewCreateOrUpdateBifrostConnectionRequestWithDefaults instantiates a new CreateOrUpdateBifrostConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *CreateOrUpdateBifrostConnectionRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateOrUpdateBifrostConnectionRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateOrUpdateBifrostConnectionRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *CreateOrUpdateBifrostConnectionRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateOrUpdateBifrostConnectionRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetName

`func (o *CreateOrUpdateBifrostConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateBifrostConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateBifrostConnectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrUpdateBifrostConnectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateOrUpdateBifrostConnectionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOrUpdateBifrostConnectionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubnet

`func (o *CreateOrUpdateBifrostConnectionRequest) GetSubnet() ConnectionSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateOrUpdateBifrostConnectionRequest) GetSubnetOk() (*ConnectionSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateOrUpdateBifrostConnectionRequest) SetSubnet(v ConnectionSubnet)`

SetSubnet sets Subnet field to given value.


### GetCertificateVersion

`func (o *CreateOrUpdateBifrostConnectionRequest) GetCertificateVersion() int64`

GetCertificateVersion returns the CertificateVersion field if non-nil, zero value otherwise.

### GetCertificateVersionOk

`func (o *CreateOrUpdateBifrostConnectionRequest) GetCertificateVersionOk() (*int64, bool)`

GetCertificateVersionOk returns a tuple with the CertificateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVersion

`func (o *CreateOrUpdateBifrostConnectionRequest) SetCertificateVersion(v int64)`

SetCertificateVersion sets CertificateVersion field to given value.

### HasCertificateVersion

`func (o *CreateOrUpdateBifrostConnectionRequest) HasCertificateVersion() bool`

HasCertificateVersion returns a boolean if a field has been set.

### SetCertificateVersionNil

`func (o *CreateOrUpdateBifrostConnectionRequest) SetCertificateVersionNil(b bool)`

 SetCertificateVersionNil sets the value for CertificateVersion to be an explicit nil

### UnsetCertificateVersion
`func (o *CreateOrUpdateBifrostConnectionRequest) UnsetCertificateVersion()`

UnsetCertificateVersion ensures that no value is present for CertificateVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


