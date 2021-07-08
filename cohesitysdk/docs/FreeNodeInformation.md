# FreeNodeInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisSerial** | Pointer to **NullableString** | Specifies the serial number of the Chassis the Node is installed in. | [optional] 
**ConnectedTo** | Pointer to **NullableBool** | Specifies whether or not this is the Node that is sending the response. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the node. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the IP address of the Node. | [optional] 
**IpmiIp** | Pointer to **NullableString** | Specifies the IPMI IP of the Node. | [optional] 
**NodeSerial** | Pointer to **NullableString** | Specifies the serial number of the Node. | [optional] 
**NodeUiSlot** | Pointer to **NullableString** | Specifies the postition for the UI to display the Node in the Cluster creation page. | [optional] 
**NumSlotsInChassis** | Pointer to **NullableInt32** | Specifies the number of Node slots present in the Chassis where this Node is installed. | [optional] 
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
### GetIp

`func (o *FreeNodeInformation) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FreeNodeInformation) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FreeNodeInformation) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FreeNodeInformation) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *FreeNodeInformation) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *FreeNodeInformation) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
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


