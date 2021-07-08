# NetappVserverInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataProtocols** | Pointer to **[]string** | Array of Data Protocols.  Specifies the set of data protocols supported by this Vserver. The kManagement protocol is not supported for this case. &#39;kNfs&#39; indicates NFS connections. &#39;kCifs&#39; indicates SMB (CIFS) connections. &#39;kIscsi&#39; indicates iSCSI connections. &#39;kFc&#39; indicates Fiber Channel connections. &#39;kFcache&#39; indicates Flex Cache connections. &#39;kHttp&#39; indicates HTTP connections. &#39;kNdmp&#39; indicates NDMP connections. &#39;kManagement&#39; indicates non-data connections used for management purposes. &#39;kNvme&#39; indicates NVMe connections. | [optional] 
**Interfaces** | Pointer to [**[]VserverNetworkInterface**](VserverNetworkInterface.md) | Array of Interfaces.  Specifies information about all interfaces on this Vserver. | [optional] 
**RootCifsShare** | Pointer to [**CifsShareInfo**](CifsShareInfo.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of this Vserver. Specifies the type of the NetApp Vserver. &#39;kData&#39; indicates the Vserver is used for data backup and restore. &#39;kAdmin&#39; indicates the Vserver is used for cluster-wide management. &#39;kSystem&#39; indicates the Vserver is used for cluster-scoped communications in an IPspace. &#39;kNode&#39; indicates the Vserver is used as the physical controller. &#39;kUnknown&#39; indicates the Vserver is used for an unknown purpose. | [optional] 

## Methods

### NewNetappVserverInfo

`func NewNetappVserverInfo() *NetappVserverInfo`

NewNetappVserverInfo instantiates a new NetappVserverInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappVserverInfoWithDefaults

`func NewNetappVserverInfoWithDefaults() *NetappVserverInfo`

NewNetappVserverInfoWithDefaults instantiates a new NetappVserverInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataProtocols

`func (o *NetappVserverInfo) GetDataProtocols() []string`

GetDataProtocols returns the DataProtocols field if non-nil, zero value otherwise.

### GetDataProtocolsOk

`func (o *NetappVserverInfo) GetDataProtocolsOk() (*[]string, bool)`

GetDataProtocolsOk returns a tuple with the DataProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtocols

`func (o *NetappVserverInfo) SetDataProtocols(v []string)`

SetDataProtocols sets DataProtocols field to given value.

### HasDataProtocols

`func (o *NetappVserverInfo) HasDataProtocols() bool`

HasDataProtocols returns a boolean if a field has been set.

### SetDataProtocolsNil

`func (o *NetappVserverInfo) SetDataProtocolsNil(b bool)`

 SetDataProtocolsNil sets the value for DataProtocols to be an explicit nil

### UnsetDataProtocols
`func (o *NetappVserverInfo) UnsetDataProtocols()`

UnsetDataProtocols ensures that no value is present for DataProtocols, not even an explicit nil
### GetInterfaces

`func (o *NetappVserverInfo) GetInterfaces() []VserverNetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NetappVserverInfo) GetInterfacesOk() (*[]VserverNetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NetappVserverInfo) SetInterfaces(v []VserverNetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *NetappVserverInfo) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### SetInterfacesNil

`func (o *NetappVserverInfo) SetInterfacesNil(b bool)`

 SetInterfacesNil sets the value for Interfaces to be an explicit nil

### UnsetInterfaces
`func (o *NetappVserverInfo) UnsetInterfaces()`

UnsetInterfaces ensures that no value is present for Interfaces, not even an explicit nil
### GetRootCifsShare

`func (o *NetappVserverInfo) GetRootCifsShare() CifsShareInfo`

GetRootCifsShare returns the RootCifsShare field if non-nil, zero value otherwise.

### GetRootCifsShareOk

`func (o *NetappVserverInfo) GetRootCifsShareOk() (*CifsShareInfo, bool)`

GetRootCifsShareOk returns a tuple with the RootCifsShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCifsShare

`func (o *NetappVserverInfo) SetRootCifsShare(v CifsShareInfo)`

SetRootCifsShare sets RootCifsShare field to given value.

### HasRootCifsShare

`func (o *NetappVserverInfo) HasRootCifsShare() bool`

HasRootCifsShare returns a boolean if a field has been set.

### GetType

`func (o *NetappVserverInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetappVserverInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetappVserverInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetappVserverInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NetappVserverInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NetappVserverInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


