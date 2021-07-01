# CapacityByTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageTier** | Pointer to **NullableString** | StorageTier is the type of StorageTier. StorageTierType represents the various values for the Storage Tier. &#39;kPCIeSSD&#39; indicates storage tier type of Pci Solid State Drive. &#39;kSATAHDD&#39; indicates storage tier type of SATA Solid State Drive. &#39;kSATAHDD&#39; indicates storage tier type of SATA Hard Disk Drive. &#39;kCLOUD&#39; indicates storage tier type of Cloud. | [optional] 
**TierMaxPhysicalCapacityBytes** | Pointer to **NullableInt64** | TierMaxPhysicalCapacityBytes is the maximum physical capacity in bytes of the storage tier. | [optional] 

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
### GetTierMaxPhysicalCapacityBytes

`func (o *CapacityByTier) GetTierMaxPhysicalCapacityBytes() int64`

GetTierMaxPhysicalCapacityBytes returns the TierMaxPhysicalCapacityBytes field if non-nil, zero value otherwise.

### GetTierMaxPhysicalCapacityBytesOk

`func (o *CapacityByTier) GetTierMaxPhysicalCapacityBytesOk() (*int64, bool)`

GetTierMaxPhysicalCapacityBytesOk returns a tuple with the TierMaxPhysicalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierMaxPhysicalCapacityBytes

`func (o *CapacityByTier) SetTierMaxPhysicalCapacityBytes(v int64)`

SetTierMaxPhysicalCapacityBytes sets TierMaxPhysicalCapacityBytes field to given value.

### HasTierMaxPhysicalCapacityBytes

`func (o *CapacityByTier) HasTierMaxPhysicalCapacityBytes() bool`

HasTierMaxPhysicalCapacityBytes returns a boolean if a field has been set.

### SetTierMaxPhysicalCapacityBytesNil

`func (o *CapacityByTier) SetTierMaxPhysicalCapacityBytesNil(b bool)`

 SetTierMaxPhysicalCapacityBytesNil sets the value for TierMaxPhysicalCapacityBytes to be an explicit nil

### UnsetTierMaxPhysicalCapacityBytes
`func (o *CapacityByTier) UnsetTierMaxPhysicalCapacityBytes()`

UnsetTierMaxPhysicalCapacityBytes ensures that no value is present for TierMaxPhysicalCapacityBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


