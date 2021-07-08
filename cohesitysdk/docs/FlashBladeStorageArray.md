# FlashBladeStorageArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityBytes** | Pointer to **NullableInt64** | Specifies the total capacity in bytes of the Pure Storage FlashBlade Array. | [optional] 
**Id** | Pointer to **NullableString** | Specifies a unique id of a Pure Storage FlashBlade Array. The id is unique across Cohesity Clusters. | [optional] 
**Networks** | Pointer to [**[]FlashBladeNetworkInterface**](FlashBladeNetworkInterface.md) | Specifies the network interfaces of the Pure Storage FlashBlade Array. | [optional] 
**PhysicalUsedBytes** | Pointer to **NullableInt64** | Specifies the space used for physical data in bytes. | [optional] 
**Revision** | Pointer to **NullableString** | Specifies the revision of the Pure Storage FlashBlade software. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the software version running on the Pure Storage FlashBlade Array. | [optional] 

## Methods

### NewFlashBladeStorageArray

`func NewFlashBladeStorageArray() *FlashBladeStorageArray`

NewFlashBladeStorageArray instantiates a new FlashBladeStorageArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashBladeStorageArrayWithDefaults

`func NewFlashBladeStorageArrayWithDefaults() *FlashBladeStorageArray`

NewFlashBladeStorageArrayWithDefaults instantiates a new FlashBladeStorageArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityBytes

`func (o *FlashBladeStorageArray) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *FlashBladeStorageArray) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *FlashBladeStorageArray) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.

### HasCapacityBytes

`func (o *FlashBladeStorageArray) HasCapacityBytes() bool`

HasCapacityBytes returns a boolean if a field has been set.

### SetCapacityBytesNil

`func (o *FlashBladeStorageArray) SetCapacityBytesNil(b bool)`

 SetCapacityBytesNil sets the value for CapacityBytes to be an explicit nil

### UnsetCapacityBytes
`func (o *FlashBladeStorageArray) UnsetCapacityBytes()`

UnsetCapacityBytes ensures that no value is present for CapacityBytes, not even an explicit nil
### GetId

`func (o *FlashBladeStorageArray) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlashBladeStorageArray) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlashBladeStorageArray) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlashBladeStorageArray) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FlashBladeStorageArray) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FlashBladeStorageArray) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNetworks

`func (o *FlashBladeStorageArray) GetNetworks() []FlashBladeNetworkInterface`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *FlashBladeStorageArray) GetNetworksOk() (*[]FlashBladeNetworkInterface, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *FlashBladeStorageArray) SetNetworks(v []FlashBladeNetworkInterface)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *FlashBladeStorageArray) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### SetNetworksNil

`func (o *FlashBladeStorageArray) SetNetworksNil(b bool)`

 SetNetworksNil sets the value for Networks to be an explicit nil

### UnsetNetworks
`func (o *FlashBladeStorageArray) UnsetNetworks()`

UnsetNetworks ensures that no value is present for Networks, not even an explicit nil
### GetPhysicalUsedBytes

`func (o *FlashBladeStorageArray) GetPhysicalUsedBytes() int64`

GetPhysicalUsedBytes returns the PhysicalUsedBytes field if non-nil, zero value otherwise.

### GetPhysicalUsedBytesOk

`func (o *FlashBladeStorageArray) GetPhysicalUsedBytesOk() (*int64, bool)`

GetPhysicalUsedBytesOk returns a tuple with the PhysicalUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalUsedBytes

`func (o *FlashBladeStorageArray) SetPhysicalUsedBytes(v int64)`

SetPhysicalUsedBytes sets PhysicalUsedBytes field to given value.

### HasPhysicalUsedBytes

`func (o *FlashBladeStorageArray) HasPhysicalUsedBytes() bool`

HasPhysicalUsedBytes returns a boolean if a field has been set.

### SetPhysicalUsedBytesNil

`func (o *FlashBladeStorageArray) SetPhysicalUsedBytesNil(b bool)`

 SetPhysicalUsedBytesNil sets the value for PhysicalUsedBytes to be an explicit nil

### UnsetPhysicalUsedBytes
`func (o *FlashBladeStorageArray) UnsetPhysicalUsedBytes()`

UnsetPhysicalUsedBytes ensures that no value is present for PhysicalUsedBytes, not even an explicit nil
### GetRevision

`func (o *FlashBladeStorageArray) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *FlashBladeStorageArray) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *FlashBladeStorageArray) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *FlashBladeStorageArray) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *FlashBladeStorageArray) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *FlashBladeStorageArray) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetVersion

`func (o *FlashBladeStorageArray) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FlashBladeStorageArray) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FlashBladeStorageArray) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FlashBladeStorageArray) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *FlashBladeStorageArray) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *FlashBladeStorageArray) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


