# CapacityByTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPhysicalCapacityBytesTier** | Pointer to **NullableInt64** | maxPhysicalCapacityBytesTier is the maximum physical capacity in bytes of the storage tier. | [optional] 
**StorageTier** | Pointer to **NullableString** | StorageTier is the type of StorageTier. StorageTierType represents the various values for the Storage Tier. &#39;kPCIeSSD&#39; indicates storage tier type of Pci Solid State Drive. &#39;kSATAHDD&#39; indicates storage tier type of SATA Solid State Drive. &#39;kSATAHDD&#39; indicates storage tier type of SATA Hard Disk Drive. &#39;kCLOUD&#39; indicates storage tier type of Cloud. | [optional] 

## Methods

### NewCapacityByTier

`func NewCapacityByTier() *CapacityByTier`

NewCapacityByTier instantiates a new CapacityByTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityByTierWithDefaults

`func NewCapacityByTierWithDefaults() *CapacityByTier`

NewCapacityByTierWithDefaults instantiates a new CapacityByTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPhysicalCapacityBytesTier

`func (o *CapacityByTier) GetMaxPhysicalCapacityBytesTier() int64`

GetMaxPhysicalCapacityBytesTier returns the MaxPhysicalCapacityBytesTier field if non-nil, zero value otherwise.

### GetMaxPhysicalCapacityBytesTierOk

`func (o *CapacityByTier) GetMaxPhysicalCapacityBytesTierOk() (*int64, bool)`

GetMaxPhysicalCapacityBytesTierOk returns a tuple with the MaxPhysicalCapacityBytesTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPhysicalCapacityBytesTier

`func (o *CapacityByTier) SetMaxPhysicalCapacityBytesTier(v int64)`

SetMaxPhysicalCapacityBytesTier sets MaxPhysicalCapacityBytesTier field to given value.

### HasMaxPhysicalCapacityBytesTier

`func (o *CapacityByTier) HasMaxPhysicalCapacityBytesTier() bool`

HasMaxPhysicalCapacityBytesTier returns a boolean if a field has been set.

### SetMaxPhysicalCapacityBytesTierNil

`func (o *CapacityByTier) SetMaxPhysicalCapacityBytesTierNil(b bool)`

 SetMaxPhysicalCapacityBytesTierNil sets the value for MaxPhysicalCapacityBytesTier to be an explicit nil

### UnsetMaxPhysicalCapacityBytesTier
`func (o *CapacityByTier) UnsetMaxPhysicalCapacityBytesTier()`

UnsetMaxPhysicalCapacityBytesTier ensures that no value is present for MaxPhysicalCapacityBytesTier, not even an explicit nil
### GetStorageTier

`func (o *CapacityByTier) GetStorageTier() string`

GetStorageTier returns the StorageTier field if non-nil, zero value otherwise.

### GetStorageTierOk

`func (o *CapacityByTier) GetStorageTierOk() (*string, bool)`

GetStorageTierOk returns a tuple with the StorageTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTier

`func (o *CapacityByTier) SetStorageTier(v string)`

SetStorageTier sets StorageTier field to given value.

### HasStorageTier

`func (o *CapacityByTier) HasStorageTier() bool`

HasStorageTier returns a boolean if a field has been set.

### SetStorageTierNil

`func (o *CapacityByTier) SetStorageTierNil(b bool)`

 SetStorageTierNil sets the value for StorageTier to be an explicit nil

### UnsetStorageTier
`func (o *CapacityByTier) UnsetStorageTier()`

UnsetStorageTier ensures that no value is present for StorageTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


