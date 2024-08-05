# UpdateBifrostConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateVersion** | Pointer to **NullableInt64** | Specifies the version of the connection&#39;s certificate. The version is used to revoke/renew connection&#39;s certificates. | [optional] 
**Connectors** | Pointer to **[]string** | Specifies the ids of the connectors in this connection | [optional] 
**Id** | **NullableInt64** | Specifies the id of the connection. | 
**Name** | **NullableString** | Specifies the name of the connection. | 
**NetworkConnectionInfo** | Pointer to [**NetworkConnectionInfo**](NetworkConnectionInfo.md) |  | [optional] 

## Methods

### NewUpdateBifrostConnectionRequest

`func NewUpdateBifrostConnectionRequest(id NullableInt64, name NullableString, ) *UpdateBifrostConnectionRequest`

NewUpdateBifrostConnectionRequest instantiates a new UpdateBifrostConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBifrostConnectionRequestWithDefaults

`func NewUpdateBifrostConnectionRequestWithDefaults() *UpdateBifrostConnectionRequest`

NewUpdateBifrostConnectionRequestWithDefaults instantiates a new UpdateBifrostConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateVersion

`func (o *UpdateBifrostConnectionRequest) GetCertificateVersion() int64`

GetCertificateVersion returns the CertificateVersion field if non-nil, zero value otherwise.

### GetCertificateVersionOk

`func (o *UpdateBifrostConnectionRequest) GetCertificateVersionOk() (*int64, bool)`

GetCertificateVersionOk returns a tuple with the CertificateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVersion

`func (o *UpdateBifrostConnectionRequest) SetCertificateVersion(v int64)`

SetCertificateVersion sets CertificateVersion field to given value.

### HasCertificateVersion

`func (o *UpdateBifrostConnectionRequest) HasCertificateVersion() bool`

HasCertificateVersion returns a boolean if a field has been set.

### SetCertificateVersionNil

`func (o *UpdateBifrostConnectionRequest) SetCertificateVersionNil(b bool)`

 SetCertificateVersionNil sets the value for CertificateVersion to be an explicit nil

### UnsetCertificateVersion
`func (o *UpdateBifrostConnectionRequest) UnsetCertificateVersion()`

UnsetCertificateVersion ensures that no value is present for CertificateVersion, not even an explicit nil
### GetConnectors

`func (o *UpdateBifrostConnectionRequest) GetConnectors() []string`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *UpdateBifrostConnectionRequest) GetConnectorsOk() (*[]string, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *UpdateBifrostConnectionRequest) SetConnectors(v []string)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *UpdateBifrostConnectionRequest) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetId

`func (o *UpdateBifrostConnectionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateBifrostConnectionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateBifrostConnectionRequest) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *UpdateBifrostConnectionRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdateBifrostConnectionRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *UpdateBifrostConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBifrostConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBifrostConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateBifrostConnectionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateBifrostConnectionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetworkConnectionInfo

`func (o *UpdateBifrostConnectionRequest) GetNetworkConnectionInfo() NetworkConnectionInfo`

GetNetworkConnectionInfo returns the NetworkConnectionInfo field if non-nil, zero value otherwise.

### GetNetworkConnectionInfoOk

`func (o *UpdateBifrostConnectionRequest) GetNetworkConnectionInfoOk() (*NetworkConnectionInfo, bool)`

GetNetworkConnectionInfoOk returns a tuple with the NetworkConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConnectionInfo

`func (o *UpdateBifrostConnectionRequest) SetNetworkConnectionInfo(v NetworkConnectionInfo)`

SetNetworkConnectionInfo sets NetworkConnectionInfo field to given value.

### HasNetworkConnectionInfo

`func (o *UpdateBifrostConnectionRequest) HasNetworkConnectionInfo() bool`

HasNetworkConnectionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


