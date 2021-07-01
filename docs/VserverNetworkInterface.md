# VserverNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataProtocols** | Pointer to **[]string** | Array of Data Protocols.  Specifies the set of data protocols supported by this interface. &#39;kNfs&#39; indicates NFS connections. &#39;kCifs&#39; indicates SMB (CIFS) connections. &#39;kIscsi&#39; indicates iSCSI connections. &#39;kFc&#39; indicates Fiber Channel connections. &#39;kFcache&#39; indicates Flex Cache connections. &#39;kHttp&#39; indicates HTTP connections. &#39;kNdmp&#39; indicates NDMP connections. &#39;kManagement&#39; indicates non-data connections used for management purposes. &#39;kNvme&#39; indicates NVMe connections. | [optional] 
**IpAddress** | Pointer to **NullableString** | Specifies the IP address of this interface. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of this interface. | [optional] 

## Methods

### NewVserverNetworkInterface

`func NewVserverNetworkInterface() *VserverNetworkInterface`

NewVserverNetworkInterface instantiates a new VserverNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVserverNetworkInterfaceWithDefaults

`func NewVserverNetworkInterfaceWithDefaults() *VserverNetworkInterface`

NewVserverNetworkInterfaceWithDefaults instantiates a new VserverNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataProtocols

`func (o *VserverNetworkInterface) GetDataProtocols() []string`

GetDataProtocols returns the DataProtocols field if non-nil, zero value otherwise.

### GetDataProtocolsOk

`func (o *VserverNetworkInterface) GetDataProtocolsOk() (*[]string, bool)`

GetDataProtocolsOk returns a tuple with the DataProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtocols

`func (o *VserverNetworkInterface) SetDataProtocols(v []string)`

SetDataProtocols sets DataProtocols field to given value.

### HasDataProtocols

`func (o *VserverNetworkInterface) HasDataProtocols() bool`

HasDataProtocols returns a boolean if a field has been set.

### SetDataProtocolsNil

`func (o *VserverNetworkInterface) SetDataProtocolsNil(b bool)`

 SetDataProtocolsNil sets the value for DataProtocols to be an explicit nil

### UnsetDataProtocols
`func (o *VserverNetworkInterface) UnsetDataProtocols()`

UnsetDataProtocols ensures that no value is present for DataProtocols, not even an explicit nil
### GetIpAddress

`func (o *VserverNetworkInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VserverNetworkInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VserverNetworkInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VserverNetworkInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VserverNetworkInterface) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VserverNetworkInterface) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetName

`func (o *VserverNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VserverNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VserverNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VserverNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VserverNetworkInterface) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VserverNetworkInterface) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


