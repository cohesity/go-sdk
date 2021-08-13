# CommonCreateOrUpdateRigelConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **NullableInt64** | Specifies the Id of the connection which this connector belongs to. | 
**TenantId** | **NullableString** | Specifies the id of the tenant which the connector belongs to. | 
**Name** | **NullableString** | Specifies the name of the connector. | 
**CertificateVersion** | **NullableInt64** | Specifies the version of the connector&#39;s certificate. The version is used to revoke/renew connector&#39;s certificates. | 

## Methods

### NewCommonCreateOrUpdateRigelConnectorRequest

`func NewCommonCreateOrUpdateRigelConnectorRequest(connectionId NullableInt64, tenantId NullableString, name NullableString, certificateVersion NullableInt64, ) *CommonCreateOrUpdateRigelConnectorRequest`

NewCommonCreateOrUpdateRigelConnectorRequest instantiates a new CommonCreateOrUpdateRigelConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonCreateOrUpdateRigelConnectorRequestWithDefaults

`func NewCommonCreateOrUpdateRigelConnectorRequestWithDefaults() *CommonCreateOrUpdateRigelConnectorRequest`

NewCommonCreateOrUpdateRigelConnectorRequestWithDefaults instantiates a new CommonCreateOrUpdateRigelConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *CommonCreateOrUpdateRigelConnectorRequest) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CommonCreateOrUpdateRigelConnectorRequest) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CommonCreateOrUpdateRigelConnectorRequest) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.


### SetConnectionIdNil

`func (o *CommonCreateOrUpdateRigelConnectorRequest) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *CommonCreateOrUpdateRigelConnectorRequest) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetTenantId

`func (o *CommonCreateOrUpdateRigelConnectorRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CommonCreateOrUpdateRigelConnectorRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CommonCreateOrUpdateRigelConnectorRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *CommonCreateOrUpdateRigelConnectorRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CommonCreateOrUpdateRigelConnectorRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetName

`func (o *CommonCreateOrUpdateRigelConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonCreateOrUpdateRigelConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonCreateOrUpdateRigelConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonCreateOrUpdateRigelConnectorRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonCreateOrUpdateRigelConnectorRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCertificateVersion

`func (o *CommonCreateOrUpdateRigelConnectorRequest) GetCertificateVersion() int64`

GetCertificateVersion returns the CertificateVersion field if non-nil, zero value otherwise.

### GetCertificateVersionOk

`func (o *CommonCreateOrUpdateRigelConnectorRequest) GetCertificateVersionOk() (*int64, bool)`

GetCertificateVersionOk returns a tuple with the CertificateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVersion

`func (o *CommonCreateOrUpdateRigelConnectorRequest) SetCertificateVersion(v int64)`

SetCertificateVersion sets CertificateVersion field to given value.


### SetCertificateVersionNil

`func (o *CommonCreateOrUpdateRigelConnectorRequest) SetCertificateVersionNil(b bool)`

 SetCertificateVersionNil sets the value for CertificateVersion to be an explicit nil

### UnsetCertificateVersion
`func (o *CommonCreateOrUpdateRigelConnectorRequest) UnsetCertificateVersion()`

UnsetCertificateVersion ensures that no value is present for CertificateVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


