# HardwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisModel** | Pointer to **NullableString** | Chassis model. | [optional] 
**ChassisSerial** | Pointer to **NullableString** | Chassis serial number programmed by manufacturer. | [optional] 
**ChassisType** | Pointer to **NullableString** | Chassis serial number programmed by manufacturer. | [optional] 
**CohesityChassisSerial** | Pointer to **NullableString** | Chassis serial number programmed by cohesity software. | [optional] 
**CohesityNodeSerial** | Pointer to **NullableString** | Node serial number programmed by cohesity software. | [optional] 
**HbaModel** | Pointer to **NullableString** | Specifies the HBA model type for the given node. | [optional] 
**IpmiLanChannel** | Pointer to **NullableString** | Specifies the channel through which the IPMI interface communicates on the network. | [optional] 
**MaxSlots** | Pointer to **NullableString** | Maximum number of slots. | [optional] 
**NodeModel** | Pointer to **NullableString** | Node model. | [optional] 
**NodeSerial** | Pointer to **NullableString** | Node serial number programmed by manufacturer. | [optional] 
**ProductModel** | Pointer to **NullableString** | Specifies the product model for the given node. | [optional] 
**ProductModelType** | Pointer to **NullableString** | Specifies the type of the product model for the given node. | [optional] 
**SlotNumber** | Pointer to **NullableString** | Slot number of the node in the chassis. | [optional] 

## Methods

### NewHardwareInfo

`func NewHardwareInfo() *HardwareInfo`

NewHardwareInfo instantiates a new HardwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHardwareInfoWithDefaults

`func NewHardwareInfoWithDefaults() *HardwareInfo`

NewHardwareInfoWithDefaults instantiates a new HardwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisModel

`func (o *HardwareInfo) GetChassisModel() string`

GetChassisModel returns the ChassisModel field if non-nil, zero value otherwise.

### GetChassisModelOk

`func (o *HardwareInfo) GetChassisModelOk() (*string, bool)`

GetChassisModelOk returns a tuple with the ChassisModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisModel

`func (o *HardwareInfo) SetChassisModel(v string)`

SetChassisModel sets ChassisModel field to given value.

### HasChassisModel

`func (o *HardwareInfo) HasChassisModel() bool`

HasChassisModel returns a boolean if a field has been set.

### SetChassisModelNil

`func (o *HardwareInfo) SetChassisModelNil(b bool)`

 SetChassisModelNil sets the value for ChassisModel to be an explicit nil

### UnsetChassisModel
`func (o *HardwareInfo) UnsetChassisModel()`

UnsetChassisModel ensures that no value is present for ChassisModel, not even an explicit nil
### GetChassisSerial

`func (o *HardwareInfo) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *HardwareInfo) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *HardwareInfo) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *HardwareInfo) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### SetChassisSerialNil

`func (o *HardwareInfo) SetChassisSerialNil(b bool)`

 SetChassisSerialNil sets the value for ChassisSerial to be an explicit nil

### UnsetChassisSerial
`func (o *HardwareInfo) UnsetChassisSerial()`

UnsetChassisSerial ensures that no value is present for ChassisSerial, not even an explicit nil
### GetChassisType

`func (o *HardwareInfo) GetChassisType() string`

GetChassisType returns the ChassisType field if non-nil, zero value otherwise.

### GetChassisTypeOk

`func (o *HardwareInfo) GetChassisTypeOk() (*string, bool)`

GetChassisTypeOk returns a tuple with the ChassisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisType

`func (o *HardwareInfo) SetChassisType(v string)`

SetChassisType sets ChassisType field to given value.

### HasChassisType

`func (o *HardwareInfo) HasChassisType() bool`

HasChassisType returns a boolean if a field has been set.

### SetChassisTypeNil

`func (o *HardwareInfo) SetChassisTypeNil(b bool)`

 SetChassisTypeNil sets the value for ChassisType to be an explicit nil

### UnsetChassisType
`func (o *HardwareInfo) UnsetChassisType()`

UnsetChassisType ensures that no value is present for ChassisType, not even an explicit nil
### GetCohesityChassisSerial

`func (o *HardwareInfo) GetCohesityChassisSerial() string`

GetCohesityChassisSerial returns the CohesityChassisSerial field if non-nil, zero value otherwise.

### GetCohesityChassisSerialOk

`func (o *HardwareInfo) GetCohesityChassisSerialOk() (*string, bool)`

GetCohesityChassisSerialOk returns a tuple with the CohesityChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesityChassisSerial

`func (o *HardwareInfo) SetCohesityChassisSerial(v string)`

SetCohesityChassisSerial sets CohesityChassisSerial field to given value.

### HasCohesityChassisSerial

`func (o *HardwareInfo) HasCohesityChassisSerial() bool`

HasCohesityChassisSerial returns a boolean if a field has been set.

### SetCohesityChassisSerialNil

`func (o *HardwareInfo) SetCohesityChassisSerialNil(b bool)`

 SetCohesityChassisSerialNil sets the value for CohesityChassisSerial to be an explicit nil

### UnsetCohesityChassisSerial
`func (o *HardwareInfo) UnsetCohesityChassisSerial()`

UnsetCohesityChassisSerial ensures that no value is present for CohesityChassisSerial, not even an explicit nil
### GetCohesityNodeSerial

`func (o *HardwareInfo) GetCohesityNodeSerial() string`

GetCohesityNodeSerial returns the CohesityNodeSerial field if non-nil, zero value otherwise.

### GetCohesityNodeSerialOk

`func (o *HardwareInfo) GetCohesityNodeSerialOk() (*string, bool)`

GetCohesityNodeSerialOk returns a tuple with the CohesityNodeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesityNodeSerial

`func (o *HardwareInfo) SetCohesityNodeSerial(v string)`

SetCohesityNodeSerial sets CohesityNodeSerial field to given value.

### HasCohesityNodeSerial

`func (o *HardwareInfo) HasCohesityNodeSerial() bool`

HasCohesityNodeSerial returns a boolean if a field has been set.

### SetCohesityNodeSerialNil

`func (o *HardwareInfo) SetCohesityNodeSerialNil(b bool)`

 SetCohesityNodeSerialNil sets the value for CohesityNodeSerial to be an explicit nil

### UnsetCohesityNodeSerial
`func (o *HardwareInfo) UnsetCohesityNodeSerial()`

UnsetCohesityNodeSerial ensures that no value is present for CohesityNodeSerial, not even an explicit nil
### GetHbaModel

`func (o *HardwareInfo) GetHbaModel() string`

GetHbaModel returns the HbaModel field if non-nil, zero value otherwise.

### GetHbaModelOk

`func (o *HardwareInfo) GetHbaModelOk() (*string, bool)`

GetHbaModelOk returns a tuple with the HbaModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaModel

`func (o *HardwareInfo) SetHbaModel(v string)`

SetHbaModel sets HbaModel field to given value.

### HasHbaModel

`func (o *HardwareInfo) HasHbaModel() bool`

HasHbaModel returns a boolean if a field has been set.

### SetHbaModelNil

`func (o *HardwareInfo) SetHbaModelNil(b bool)`

 SetHbaModelNil sets the value for HbaModel to be an explicit nil

### UnsetHbaModel
`func (o *HardwareInfo) UnsetHbaModel()`

UnsetHbaModel ensures that no value is present for HbaModel, not even an explicit nil
### GetIpmiLanChannel

`func (o *HardwareInfo) GetIpmiLanChannel() string`

GetIpmiLanChannel returns the IpmiLanChannel field if non-nil, zero value otherwise.

### GetIpmiLanChannelOk

`func (o *HardwareInfo) GetIpmiLanChannelOk() (*string, bool)`

GetIpmiLanChannelOk returns a tuple with the IpmiLanChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiLanChannel

`func (o *HardwareInfo) SetIpmiLanChannel(v string)`

SetIpmiLanChannel sets IpmiLanChannel field to given value.

### HasIpmiLanChannel

`func (o *HardwareInfo) HasIpmiLanChannel() bool`

HasIpmiLanChannel returns a boolean if a field has been set.

### SetIpmiLanChannelNil

`func (o *HardwareInfo) SetIpmiLanChannelNil(b bool)`

 SetIpmiLanChannelNil sets the value for IpmiLanChannel to be an explicit nil

### UnsetIpmiLanChannel
`func (o *HardwareInfo) UnsetIpmiLanChannel()`

UnsetIpmiLanChannel ensures that no value is present for IpmiLanChannel, not even an explicit nil
### GetMaxSlots

`func (o *HardwareInfo) GetMaxSlots() string`

GetMaxSlots returns the MaxSlots field if non-nil, zero value otherwise.

### GetMaxSlotsOk

`func (o *HardwareInfo) GetMaxSlotsOk() (*string, bool)`

GetMaxSlotsOk returns a tuple with the MaxSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSlots

`func (o *HardwareInfo) SetMaxSlots(v string)`

SetMaxSlots sets MaxSlots field to given value.

### HasMaxSlots

`func (o *HardwareInfo) HasMaxSlots() bool`

HasMaxSlots returns a boolean if a field has been set.

### SetMaxSlotsNil

`func (o *HardwareInfo) SetMaxSlotsNil(b bool)`

 SetMaxSlotsNil sets the value for MaxSlots to be an explicit nil

### UnsetMaxSlots
`func (o *HardwareInfo) UnsetMaxSlots()`

UnsetMaxSlots ensures that no value is present for MaxSlots, not even an explicit nil
### GetNodeModel

`func (o *HardwareInfo) GetNodeModel() string`

GetNodeModel returns the NodeModel field if non-nil, zero value otherwise.

### GetNodeModelOk

`func (o *HardwareInfo) GetNodeModelOk() (*string, bool)`

GetNodeModelOk returns a tuple with the NodeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeModel

`func (o *HardwareInfo) SetNodeModel(v string)`

SetNodeModel sets NodeModel field to given value.

### HasNodeModel

`func (o *HardwareInfo) HasNodeModel() bool`

HasNodeModel returns a boolean if a field has been set.

### SetNodeModelNil

`func (o *HardwareInfo) SetNodeModelNil(b bool)`

 SetNodeModelNil sets the value for NodeModel to be an explicit nil

### UnsetNodeModel
`func (o *HardwareInfo) UnsetNodeModel()`

UnsetNodeModel ensures that no value is present for NodeModel, not even an explicit nil
### GetNodeSerial

`func (o *HardwareInfo) GetNodeSerial() string`

GetNodeSerial returns the NodeSerial field if non-nil, zero value otherwise.

### GetNodeSerialOk

`func (o *HardwareInfo) GetNodeSerialOk() (*string, bool)`

GetNodeSerialOk returns a tuple with the NodeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSerial

`func (o *HardwareInfo) SetNodeSerial(v string)`

SetNodeSerial sets NodeSerial field to given value.

### HasNodeSerial

`func (o *HardwareInfo) HasNodeSerial() bool`

HasNodeSerial returns a boolean if a field has been set.

### SetNodeSerialNil

`func (o *HardwareInfo) SetNodeSerialNil(b bool)`

 SetNodeSerialNil sets the value for NodeSerial to be an explicit nil

### UnsetNodeSerial
`func (o *HardwareInfo) UnsetNodeSerial()`

UnsetNodeSerial ensures that no value is present for NodeSerial, not even an explicit nil
### GetProductModel

`func (o *HardwareInfo) GetProductModel() string`

GetProductModel returns the ProductModel field if non-nil, zero value otherwise.

### GetProductModelOk

`func (o *HardwareInfo) GetProductModelOk() (*string, bool)`

GetProductModelOk returns a tuple with the ProductModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductModel

`func (o *HardwareInfo) SetProductModel(v string)`

SetProductModel sets ProductModel field to given value.

### HasProductModel

`func (o *HardwareInfo) HasProductModel() bool`

HasProductModel returns a boolean if a field has been set.

### SetProductModelNil

`func (o *HardwareInfo) SetProductModelNil(b bool)`

 SetProductModelNil sets the value for ProductModel to be an explicit nil

### UnsetProductModel
`func (o *HardwareInfo) UnsetProductModel()`

UnsetProductModel ensures that no value is present for ProductModel, not even an explicit nil
### GetProductModelType

`func (o *HardwareInfo) GetProductModelType() string`

GetProductModelType returns the ProductModelType field if non-nil, zero value otherwise.

### GetProductModelTypeOk

`func (o *HardwareInfo) GetProductModelTypeOk() (*string, bool)`

GetProductModelTypeOk returns a tuple with the ProductModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductModelType

`func (o *HardwareInfo) SetProductModelType(v string)`

SetProductModelType sets ProductModelType field to given value.

### HasProductModelType

`func (o *HardwareInfo) HasProductModelType() bool`

HasProductModelType returns a boolean if a field has been set.

### SetProductModelTypeNil

`func (o *HardwareInfo) SetProductModelTypeNil(b bool)`

 SetProductModelTypeNil sets the value for ProductModelType to be an explicit nil

### UnsetProductModelType
`func (o *HardwareInfo) UnsetProductModelType()`

UnsetProductModelType ensures that no value is present for ProductModelType, not even an explicit nil
### GetSlotNumber

`func (o *HardwareInfo) GetSlotNumber() string`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *HardwareInfo) GetSlotNumberOk() (*string, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *HardwareInfo) SetSlotNumber(v string)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *HardwareInfo) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.

### SetSlotNumberNil

`func (o *HardwareInfo) SetSlotNumberNil(b bool)`

 SetSlotNumberNil sets the value for SlotNumber to be an explicit nil

### UnsetSlotNumber
`func (o *HardwareInfo) UnsetSlotNumber()`

UnsetSlotNumber ensures that no value is present for SlotNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


