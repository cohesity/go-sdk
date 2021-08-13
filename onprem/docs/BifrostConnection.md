# BifrostConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the id of the connection. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the connection. | [optional] 
**Subnet** | Pointer to [**ConnectionSubnet**](ConnectionSubnet.md) |  | [optional] 
**CertificateVersion** | Pointer to **NullableInt64** | Specifies the version of the connection&#39;s certificate. The version is used to revoke/renew connection&#39;s certificates. | [optional] 
**NetworkConnectionInfo** | Pointer to [**NetworkConnectionInfo**](NetworkConnectionInfo.md) |  | [optional] 

## Methods

### NewBifrostConnection

`func NewBifrostConnection() *BifrostConnection`

NewBifrostConnection instantiates a new BifrostConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBifrostConnectionWithDefaults

`func NewBifrostConnectionWithDefaults() *BifrostConnection`

NewBifrostConnectionWithDefaults instantiates a new BifrostConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BifrostConnection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BifrostConnection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BifrostConnection) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BifrostConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BifrostConnection) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BifrostConnection) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BifrostConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BifrostConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BifrostConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BifrostConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BifrostConnection) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BifrostConnection) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubnet

`func (o *BifrostConnection) GetSubnet() ConnectionSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *BifrostConnection) GetSubnetOk() (*ConnectionSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *BifrostConnection) SetSubnet(v ConnectionSubnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *BifrostConnection) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetCertificateVersion

`func (o *BifrostConnection) GetCertificateVersion() int64`

GetCertificateVersion returns the CertificateVersion field if non-nil, zero value otherwise.

### GetCertificateVersionOk

`func (o *BifrostConnection) GetCertificateVersionOk() (*int64, bool)`

GetCertificateVersionOk returns a tuple with the CertificateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateVersion

`func (o *BifrostConnection) SetCertificateVersion(v int64)`

SetCertificateVersion sets CertificateVersion field to given value.

### HasCertificateVersion

`func (o *BifrostConnection) HasCertificateVersion() bool`

HasCertificateVersion returns a boolean if a field has been set.

### SetCertificateVersionNil

`func (o *BifrostConnection) SetCertificateVersionNil(b bool)`

 SetCertificateVersionNil sets the value for CertificateVersion to be an explicit nil

### UnsetCertificateVersion
`func (o *BifrostConnection) UnsetCertificateVersion()`

UnsetCertificateVersion ensures that no value is present for CertificateVersion, not even an explicit nil
### GetNetworkConnectionInfo

`func (o *BifrostConnection) GetNetworkConnectionInfo() NetworkConnectionInfo`

GetNetworkConnectionInfo returns the NetworkConnectionInfo field if non-nil, zero value otherwise.

### GetNetworkConnectionInfoOk

`func (o *BifrostConnection) GetNetworkConnectionInfoOk() (*NetworkConnectionInfo, bool)`

GetNetworkConnectionInfoOk returns a tuple with the NetworkConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConnectionInfo

`func (o *BifrostConnection) SetNetworkConnectionInfo(v NetworkConnectionInfo)`

SetNetworkConnectionInfo sets NetworkConnectionInfo field to given value.

### HasNetworkConnectionInfo

`func (o *BifrostConnection) HasNetworkConnectionInfo() bool`

HasNetworkConnectionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


