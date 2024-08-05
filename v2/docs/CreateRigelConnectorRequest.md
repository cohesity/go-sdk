# CreateRigelConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateVersion** | Pointer to **NullableInt64** | Specifies the version of the connector&#39;s certificate. The version is used to revoke/renew connector&#39;s certificates. | [optional] 
**CertificateVersionList** | Pointer to **[]int64** | Specifies the list of accepted version of the connector&#39;s certificate. The version is used to revoke/renew connector&#39;s certificates. | [optional] 
**ConnectionId** | **NullableInt64** | Specifies the Id of the connection which this connector belongs to. | 
**ConnectorCloudMetadata** | Pointer to [**ConnectorCloudMetadata**](ConnectorCloudMetadata.md) |  | [optional] 
**Name** | **NullableString** | Specifies the name of the connector. | 
**TenantId** | **NullableString** | Specifies the id of the tenant which the connector belongs to. | 
**ConnectorId** | **NullableInt64** | Specifies the Id of the connector. | 

## Methods

### NewCreateRigelConnectorRequest

`func NewCreateRigelConnectorRequest(connectionId NullableInt64, name NullableString, tenantId NullableString, connectorId NullableInt64, ) *CreateRigelConnectorRequest`

NewCreateRigelConnectorRequest instantiates a new CreateRigelConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRigelConnectorRequestWithDefaults

`func NewCreateRigelConnectorRequestWithDefaults() *CreateRigelConnectorRequest`

NewCreateRigelConnectorRequestWithDefaults instantiates a new CreateRigelConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateVersion

`func (o *CreateRigelConnectorRequest) GetCertificateVersion() int64`

GetCertificateVersion returns the CertificateVersion field if non-nil, zero value otherwise.

### GetCertificateVersionOk

`func (o *CreateRigelConnectorRequest) GetCertificateVersionOk() (*int64, bool)`

GetCertificateVersionOk returns a tuple with the CertificateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVersion

`func (o *CreateRigelConnectorRequest) SetCertificateVersion(v int64)`

SetCertificateVersion sets CertificateVersion field to given value.

### HasCertificateVersion

`func (o *CreateRigelConnectorRequest) HasCertificateVersion() bool`

HasCertificateVersion returns a boolean if a field has been set.

### SetCertificateVersionNil

`func (o *CreateRigelConnectorRequest) SetCertificateVersionNil(b bool)`

 SetCertificateVersionNil sets the value for CertificateVersion to be an explicit nil

### UnsetCertificateVersion
`func (o *CreateRigelConnectorRequest) UnsetCertificateVersion()`

UnsetCertificateVersion ensures that no value is present for CertificateVersion, not even an explicit nil
### GetCertificateVersionList

`func (o *CreateRigelConnectorRequest) GetCertificateVersionList() []int64`

GetCertificateVersionList returns the CertificateVersionList field if non-nil, zero value otherwise.

### GetCertificateVersionListOk

`func (o *CreateRigelConnectorRequest) GetCertificateVersionListOk() (*[]int64, bool)`

GetCertificateVersionListOk returns a tuple with the CertificateVersionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVersionList

`func (o *CreateRigelConnectorRequest) SetCertificateVersionList(v []int64)`

SetCertificateVersionList sets CertificateVersionList field to given value.

### HasCertificateVersionList

`func (o *CreateRigelConnectorRequest) HasCertificateVersionList() bool`

HasCertificateVersionList returns a boolean if a field has been set.

### SetCertificateVersionListNil

`func (o *CreateRigelConnectorRequest) SetCertificateVersionListNil(b bool)`

 SetCertificateVersionListNil sets the value for CertificateVersionList to be an explicit nil

### UnsetCertificateVersionList
`func (o *CreateRigelConnectorRequest) UnsetCertificateVersionList()`

UnsetCertificateVersionList ensures that no value is present for CertificateVersionList, not even an explicit nil
### GetConnectionId

`func (o *CreateRigelConnectorRequest) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CreateRigelConnectorRequest) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CreateRigelConnectorRequest) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.


### SetConnectionIdNil

`func (o *CreateRigelConnectorRequest) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *CreateRigelConnectorRequest) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnectorCloudMetadata

`func (o *CreateRigelConnectorRequest) GetConnectorCloudMetadata() ConnectorCloudMetadata`

GetConnectorCloudMetadata returns the ConnectorCloudMetadata field if non-nil, zero value otherwise.

### GetConnectorCloudMetadataOk

`func (o *CreateRigelConnectorRequest) GetConnectorCloudMetadataOk() (*ConnectorCloudMetadata, bool)`

GetConnectorCloudMetadataOk returns a tuple with the ConnectorCloudMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorCloudMetadata

`func (o *CreateRigelConnectorRequest) SetConnectorCloudMetadata(v ConnectorCloudMetadata)`

SetConnectorCloudMetadata sets ConnectorCloudMetadata field to given value.

### HasConnectorCloudMetadata

`func (o *CreateRigelConnectorRequest) HasConnectorCloudMetadata() bool`

HasConnectorCloudMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateRigelConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRigelConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRigelConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateRigelConnectorRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateRigelConnectorRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantId

`func (o *CreateRigelConnectorRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CreateRigelConnectorRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CreateRigelConnectorRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *CreateRigelConnectorRequest) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CreateRigelConnectorRequest) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetConnectorId

`func (o *CreateRigelConnectorRequest) GetConnectorId() int64`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *CreateRigelConnectorRequest) GetConnectorIdOk() (*int64, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *CreateRigelConnectorRequest) SetConnectorId(v int64)`

SetConnectorId sets ConnectorId field to given value.


### SetConnectorIdNil

`func (o *CreateRigelConnectorRequest) SetConnectorIdNil(b bool)`

 SetConnectorIdNil sets the value for ConnectorId to be an explicit nil

### UnsetConnectorId
`func (o *CreateRigelConnectorRequest) UnsetConnectorId()`

UnsetConnectorId ensures that no value is present for ConnectorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


