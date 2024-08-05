# RigelConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateVersion** | Pointer to **NullableInt64** | Specifies the version of the connector&#39;s certificate. The version is used to revoke/renew connector&#39;s certificates. | [optional] 
**CertificateVersionList** | Pointer to **[]int64** | Specifies the list of accepted version of the connector&#39;s certificate. The version is used to revoke/renew connector&#39;s certificates. | [optional] 
**ConnectionId** | Pointer to **NullableInt64** | Specifies the Id of the connection which this connector belongs to. | [optional] 
**ConnectionStatus** | Pointer to [**ConnectorConnectionInfo**](ConnectorConnectionInfo.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the connector. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the connector. | [optional] 

## Methods

### NewRigelConnector

`func NewRigelConnector() *RigelConnector`

NewRigelConnector instantiates a new RigelConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRigelConnectorWithDefaults

`func NewRigelConnectorWithDefaults() *RigelConnector`

NewRigelConnectorWithDefaults instantiates a new RigelConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateVersion

`func (o *RigelConnector) GetCertificateVersion() int64`

GetCertificateVersion returns the CertificateVersion field if non-nil, zero value otherwise.

### GetCertificateVersionOk

`func (o *RigelConnector) GetCertificateVersionOk() (*int64, bool)`

GetCertificateVersionOk returns a tuple with the CertificateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVersion

`func (o *RigelConnector) SetCertificateVersion(v int64)`

SetCertificateVersion sets CertificateVersion field to given value.

### HasCertificateVersion

`func (o *RigelConnector) HasCertificateVersion() bool`

HasCertificateVersion returns a boolean if a field has been set.

### SetCertificateVersionNil

`func (o *RigelConnector) SetCertificateVersionNil(b bool)`

 SetCertificateVersionNil sets the value for CertificateVersion to be an explicit nil

### UnsetCertificateVersion
`func (o *RigelConnector) UnsetCertificateVersion()`

UnsetCertificateVersion ensures that no value is present for CertificateVersion, not even an explicit nil
### GetCertificateVersionList

`func (o *RigelConnector) GetCertificateVersionList() []int64`

GetCertificateVersionList returns the CertificateVersionList field if non-nil, zero value otherwise.

### GetCertificateVersionListOk

`func (o *RigelConnector) GetCertificateVersionListOk() (*[]int64, bool)`

GetCertificateVersionListOk returns a tuple with the CertificateVersionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVersionList

`func (o *RigelConnector) SetCertificateVersionList(v []int64)`

SetCertificateVersionList sets CertificateVersionList field to given value.

### HasCertificateVersionList

`func (o *RigelConnector) HasCertificateVersionList() bool`

HasCertificateVersionList returns a boolean if a field has been set.

### SetCertificateVersionListNil

`func (o *RigelConnector) SetCertificateVersionListNil(b bool)`

 SetCertificateVersionListNil sets the value for CertificateVersionList to be an explicit nil

### UnsetCertificateVersionList
`func (o *RigelConnector) UnsetCertificateVersionList()`

UnsetCertificateVersionList ensures that no value is present for CertificateVersionList, not even an explicit nil
### GetConnectionId

`func (o *RigelConnector) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *RigelConnector) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *RigelConnector) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *RigelConnector) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *RigelConnector) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *RigelConnector) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnectionStatus

`func (o *RigelConnector) GetConnectionStatus() ConnectorConnectionInfo`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *RigelConnector) GetConnectionStatusOk() (*ConnectorConnectionInfo, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *RigelConnector) SetConnectionStatus(v ConnectorConnectionInfo)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *RigelConnector) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetId

`func (o *RigelConnector) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RigelConnector) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RigelConnector) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RigelConnector) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RigelConnector) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RigelConnector) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *RigelConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RigelConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RigelConnector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RigelConnector) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RigelConnector) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RigelConnector) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


