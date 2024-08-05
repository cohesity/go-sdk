# FreeNodeInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanConnect** | Pointer to **NullableBool** | Deprecated - This field is deprecated, use connectedTo field. | [optional] 
**ChassisModel** | Pointer to **NullableString** | Specifies the model number of the Chassis the Node is installed in. | [optional] 
**ChassisSerial** | Pointer to **NullableString** | Specifies the serial number of the Chassis the Node is installed in. | [optional] 
**ConnectedTo** | Pointer to **NullableBool** | Specifies if this is the node from where this API response was received. | [optional] 
**Hostname** | Pointer to **NullableString** | Specifies the host name of the node. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the node. | [optional] 
**IpmiIp** | Pointer to **NullableString** | Specifies the IPMI IP of the Node. | [optional] 
**Ips** | Pointer to **[]string** | List of discovered ipv4/ipv6 addresses of the node. Ip field returns ips as comma separated single string which is incorrect. | [optional] 
**NodeModel** | Pointer to **NullableString** | Specifies the node model. | [optional] 
**NodeSerial** | Pointer to **NullableString** | Specifies the serial number of the Node. | [optional] 
**NodeUiSlot** | Pointer to **NullableString** | Specifies the position for the UI to display the Node in the Cluster creation page. | [optional] 
**NumSlotsInChassis** | Pointer to **NullableInt32** | Specifies the number of Node slots present in the Chassis where this Node is installed. | [optional] 
**PrimaryIPv4Address** | Pointer to **NullableString** | IPv4 addresses in primary interface&#39;s LAN. | [optional] 
**PrimaryIPv6Address** | Pointer to **NullableString** | IPv6 addresses in primary interface&#39;s LAN. | [optional] 
**ProductModel** | Pointer to **NullableString** | Specifies the product model of the node. | [optional] 
**SlotNumber** | Pointer to **NullableString** | Specifies the number of the slot the Node is installed in. | [optional] 
**SoftwareVersion** | Pointer to **NullableString** | Specifies the version of the software installed on the Node. | [optional] 

## Methods

### NewFreeNodeInformation

`func NewFreeNodeInformation() *FreeNodeInformation`

NewFreeNodeInformation instantiates a new FreeNodeInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeNodeInformationWithDefaults

`func NewFreeNodeInformationWithDefaults() *FreeNodeInformation`

NewFreeNodeInformationWithDefaults instantiates a new FreeNodeInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanConnect

`func (o *FreeNodeInformation) GetCanConnect() bool`

GetCanConnect returns the CanConnect field if non-nil, zero value otherwise.

### GetCanConnectOk

`func (o *FreeNodeInformation) GetCanConnectOk() (*bool, bool)`

GetCanConnectOk returns a tuple with the CanConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanConnect

`func (o *FreeNodeInformation) SetCanConnect(v bool)`

SetCanConnect sets CanConnect field to given value.

### HasCanConnect

`func (o *FreeNodeInformation) HasCanConnect() bool`

HasCanConnect returns a boolean if a field has been set.

### SetCanConnectNil

`func (o *FreeNodeInformation) SetCanConnectNil(b bool)`

 SetCanConnectNil sets the value for CanConnect to be an explicit nil

### UnsetCanConnect
`func (o *FreeNodeInformation) UnsetCanConnect()`

UnsetCanConnect ensures that no value is present for CanConnect, not even an explicit nil
### GetChassisModel

`func (o *FreeNodeInformation) GetChassisModel() string`

GetChassisModel returns the ChassisModel field if non-nil, zero value otherwise.

### GetChassisModelOk

`func (o *FreeNodeInformation) GetChassisModelOk() (*string, bool)`

GetChassisModelOk returns a tuple with the ChassisModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisModel

`func (o *FreeNodeInformation) SetChassisModel(v string)`

SetChassisModel sets ChassisModel field to given value.

### HasChassisModel

`func (o *FreeNodeInformation) HasChassisModel() bool`

HasChassisModel returns a boolean if a field has been set.

### SetChassisModelNil

`func (o *FreeNodeInformation) SetChassisModelNil(b bool)`

 SetChassisModelNil sets the value for ChassisModel to be an explicit nil

### UnsetChassisModel
`func (o *FreeNodeInformation) UnsetChassisModel()`

UnsetChassisModel ensures that no value is present for ChassisModel, not even an explicit nil
### GetChassisSerial

`func (o *FreeNodeInformation) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *FreeNodeInformation) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *FreeNodeInformation) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *FreeNodeInformation) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### SetChassisSerialNil

`func (o *FreeNodeInformation) SetChassisSerialNil(b bool)`

 SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil

### UnsetChassisSerial
`func (o *FreeNodeInformation) UnsetChassisSerial()`

UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
### GetConnectedTo

`func (o *FreeNodeInformation) GetConnectedTo() bool`

GetConnectedTo returns the ConnectedTo field if non-nil, zero value otherwise.

### GetConnectedToOk

`func (o *FreeNodeInformation) GetConnectedToOk() (*bool, bool)`

GetConnectedToOk returns a tuple with the ConnectedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedTo

`func (o *FreeNodeInformation) SetConnectedTo(v bool)`

SetConnectedTo sets ConnectedTo field to given value.

### HasConnectedTo

`func (o *FreeNodeInformation) HasConnectedTo() bool`

HasConnectedTo returns a boolean if a field has been set.

### SetConnectedToNil

`func (o *FreeNodeInformation) SetConnectedToNil(b bool)`

 SetConnectedToNil sets the value for ConnectedTo to be an explicit nil

### UnsetConnectedTo
`func (o *FreeNodeInformation) UnsetConnectedTo()`

UnsetConnectedTo ensures that no value is present for ConnectedTo, not even an explicit nil
### GetHostname

`func (o *FreeNodeInformation) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FreeNodeInformation) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FreeNodeInformation) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *FreeNodeInformation) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *FreeNodeInformation) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *FreeNodeInformation) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetId

`func (o *FreeNodeInformation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FreeNodeInformation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FreeNodeInformation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FreeNodeInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FreeNodeInformation) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FreeNodeInformation) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIpmiIp

`func (o *FreeNodeInformation) GetIpmiIp() string`

GetIpmiIp returns the IpmiIp field if non-nil, zero value otherwise.

### GetIpmiIpOk

`func (o *FreeNodeInformation) GetIpmiIpOk() (*string, bool)`

GetIpmiIpOk returns a tuple with the IpmiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiIp

`func (o *FreeNodeInformation) SetIpmiIp(v string)`

SetIpmiIp sets IpmiIp field to given value.

### HasIpmiIp

`func (o *FreeNodeInformation) HasIpmiIp() bool`

HasIpmiIp returns a boolean if a field has been set.

### SetIpmiIpNil

`func (o *FreeNodeInformation) SetIpmiIpNil(b bool)`

 SetIpmiIpNil sets the value for IpmiIp to be an explicit nil

### UnsetIpmiIp
`func (o *FreeNodeInformation) UnsetIpmiIp()`

UnsetIpmiIp ensures that no value is present for IpmiIp, not even an explicit nil
### GetIps

`func (o *FreeNodeInformation) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *FreeNodeInformation) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *FreeNodeInformation) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *FreeNodeInformation) HasIps() bool`

HasIps returns a boolean if a field has been set.

### SetIpsNil

`func (o *FreeNodeInformation) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *FreeNodeInformation) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil
### GetNodeModel

`func (o *FreeNodeInformation) GetNodeModel() string`

GetNodeModel returns the NodeModel field if non-nil, zero value otherwise.

### GetNodeModelOk

`func (o *FreeNodeInformation) GetNodeModelOk() (*string, bool)`

GetNodeModelOk returns a tuple with the NodeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeModel

`func (o *FreeNodeInformation) SetNodeModel(v string)`

SetNodeModel sets NodeModel field to given value.

### HasNodeModel

`func (o *FreeNodeInformation) HasNodeModel() bool`

HasNodeModel returns a boolean if a field has been set.

### SetNodeModelNil

`func (o *FreeNodeInformation) SetNodeModelNil(b bool)`

 SetNodeModelNil sets the value for NodeModel to be an explicit nil

### UnsetNodeModel
`func (o *FreeNodeInformation) UnsetNodeModel()`

UnsetNodeModel ensures that no value is present for NodeModel, not even an explicit nil
### GetNodeSerial

`func (o *FreeNodeInformation) GetNodeSerial() string`

GetNodeSerial returns the NodeSerial field if non-nil, zero value otherwise.

### GetNodeSerialOk

`func (o *FreeNodeInformation) GetNodeSerialOk() (*string, bool)`

GetNodeSerialOk returns a tuple with the NodeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSerial

`func (o *FreeNodeInformation) SetNodeSerial(v string)`

SetNodeSerial sets NodeSerial field to given value.

### HasNodeSerial

`func (o *FreeNodeInformation) HasNodeSerial() bool`

HasNodeSerial returns a boolean if a field has been set.

### SetNodeSerialNil

`func (o *FreeNodeInformation) SetNodeSerialNil(b bool)`

 SetNodeSerialNil sets the value for NodeSerial to be an explicit nil

### UnsetNodeSerial
`func (o *FreeNodeInformation) UnsetNodeSerial()`

UnsetNodeSerial ensures that no value is present for NodeSerial, not even an explicit nil
### GetNodeUiSlot

`func (o *FreeNodeInformation) GetNodeUiSlot() string`

GetNodeUiSlot returns the NodeUiSlot field if non-nil, zero value otherwise.

### GetNodeUiSlotOk

`func (o *FreeNodeInformation) GetNodeUiSlotOk() (*string, bool)`

GetNodeUiSlotOk returns a tuple with the NodeUiSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUiSlot

`func (o *FreeNodeInformation) SetNodeUiSlot(v string)`

SetNodeUiSlot sets NodeUiSlot field to given value.

### HasNodeUiSlot

`func (o *FreeNodeInformation) HasNodeUiSlot() bool`

HasNodeUiSlot returns a boolean if a field has been set.

### SetNodeUiSlotNil

`func (o *FreeNodeInformation) SetNodeUiSlotNil(b bool)`

 SetNodeUiSlotNil sets the value for NodeUiSlot to be an explicit nil

### UnsetNodeUiSlot
`func (o *FreeNodeInformation) UnsetNodeUiSlot()`

UnsetNodeUiSlot ensures that no value is present for NodeUiSlot, not even an explicit nil
### GetNumSlotsInChassis

`func (o *FreeNodeInformation) GetNumSlotsInChassis() int32`

GetNumSlotsInChassis returns the NumSlotsInChassis field if non-nil, zero value otherwise.

### GetNumSlotsInChassisOk

`func (o *FreeNodeInformation) GetNumSlotsInChassisOk() (*int32, bool)`

GetNumSlotsInChassisOk returns a tuple with the NumSlotsInChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSlotsInChassis

`func (o *FreeNodeInformation) SetNumSlotsInChassis(v int32)`

SetNumSlotsInChassis sets NumSlotsInChassis field to given value.

### HasNumSlotsInChassis

`func (o *FreeNodeInformation) HasNumSlotsInChassis() bool`

HasNumSlotsInChassis returns a boolean if a field has been set.

### SetNumSlotsInChassisNil

`func (o *FreeNodeInformation) SetNumSlotsInChassisNil(b bool)`

 SetNumSlotsInChassisNil sets the value for NumSlotsInChassis to be an explicit nil

### UnsetNumSlotsInChassis
`func (o *FreeNodeInformation) UnsetNumSlotsInChassis()`

UnsetNumSlotsInChassis ensures that no value is present for NumSlotsInChassis, not even an explicit nil
### GetPrimaryIPv4Address

`func (o *FreeNodeInformation) GetPrimaryIPv4Address() string`

GetPrimaryIPv4Address returns the PrimaryIPv4Address field if non-nil, zero value otherwise.

### GetPrimaryIPv4AddressOk

`func (o *FreeNodeInformation) GetPrimaryIPv4AddressOk() (*string, bool)`

GetPrimaryIPv4AddressOk returns a tuple with the PrimaryIPv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIPv4Address

`func (o *FreeNodeInformation) SetPrimaryIPv4Address(v string)`

SetPrimaryIPv4Address sets PrimaryIPv4Address field to given value.

### HasPrimaryIPv4Address

`func (o *FreeNodeInformation) HasPrimaryIPv4Address() bool`

HasPrimaryIPv4Address returns a boolean if a field has been set.

### SetPrimaryIPv4AddressNil

`func (o *FreeNodeInformation) SetPrimaryIPv4AddressNil(b bool)`

 SetPrimaryIPv4AddressNil sets the value for PrimaryIPv4Address to be an explicit nil

### UnsetPrimaryIPv4Address
`func (o *FreeNodeInformation) UnsetPrimaryIPv4Address()`

UnsetPrimaryIPv4Address ensures that no value is present for PrimaryIPv4Address, not even an explicit nil
### GetPrimaryIPv6Address

`func (o *FreeNodeInformation) GetPrimaryIPv6Address() string`

GetPrimaryIPv6Address returns the PrimaryIPv6Address field if non-nil, zero value otherwise.

### GetPrimaryIPv6AddressOk

`func (o *FreeNodeInformation) GetPrimaryIPv6AddressOk() (*string, bool)`

GetPrimaryIPv6AddressOk returns a tuple with the PrimaryIPv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIPv6Address

`func (o *FreeNodeInformation) SetPrimaryIPv6Address(v string)`

SetPrimaryIPv6Address sets PrimaryIPv6Address field to given value.

### HasPrimaryIPv6Address

`func (o *FreeNodeInformation) HasPrimaryIPv6Address() bool`

HasPrimaryIPv6Address returns a boolean if a field has been set.

### SetPrimaryIPv6AddressNil

`func (o *FreeNodeInformation) SetPrimaryIPv6AddressNil(b bool)`

 SetPrimaryIPv6AddressNil sets the value for PrimaryIPv6Address to be an explicit nil

### UnsetPrimaryIPv6Address
`func (o *FreeNodeInformation) UnsetPrimaryIPv6Address()`

UnsetPrimaryIPv6Address ensures that no value is present for PrimaryIPv6Address, not even an explicit nil
### GetProductModel

`func (o *FreeNodeInformation) GetProductModel() string`

GetProductModel returns the ProductModel field if non-nil, zero value otherwise.

### GetProductModelOk

`func (o *FreeNodeInformation) GetProductModelOk() (*string, bool)`

GetProductModelOk returns a tuple with the ProductModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductModel

`func (o *FreeNodeInformation) SetProductModel(v string)`

SetProductModel sets ProductModel field to given value.

### HasProductModel

`func (o *FreeNodeInformation) HasProductModel() bool`

HasProductModel returns a boolean if a field has been set.

### SetProductModelNil

`func (o *FreeNodeInformation) SetProductModelNil(b bool)`

 SetProductModelNil sets the value for ProductModel to be an explicit nil

### UnsetProductModel
`func (o *FreeNodeInformation) UnsetProductModel()`

UnsetProductModel ensures that no value is present for ProductModel, not even an explicit nil
### GetSlotNumber

`func (o *FreeNodeInformation) GetSlotNumber() string`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *FreeNodeInformation) GetSlotNumberOk() (*string, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *FreeNodeInformation) SetSlotNumber(v string)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *FreeNodeInformation) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.

### SetSlotNumberNil

`func (o *FreeNodeInformation) SetSlotNumberNil(b bool)`

 SetSlotNumberNil sets the value for SlotNumber to be an explicit nil

### UnsetSlotNumber
`func (o *FreeNodeInformation) UnsetSlotNumber()`

UnsetSlotNumber ensures that no value is present for SlotNumber, not even an explicit nil
### GetSoftwareVersion

`func (o *FreeNodeInformation) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *FreeNodeInformation) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *FreeNodeInformation) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *FreeNodeInformation) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### SetSoftwareVersionNil

`func (o *FreeNodeInformation) SetSoftwareVersionNil(b bool)`

 SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil

### UnsetSoftwareVersion
`func (o *FreeNodeInformation) UnsetSoftwareVersion()`

UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


