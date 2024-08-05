# CountByTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskCount** | Pointer to **NullableInt64** | DiskCount is the disk number of the storage tier. | [optional] 
**StorageTier** | Pointer to **NullableString** | StorageTier is the type of StorageTier. StorageTierType represents the various values for the Storage Tier. &#39;kPCIeSSD&#39; indicates storage tier type of Pci Solid State Drive. &#39;kSATAHDD&#39; indicates storage tier type of SATA Solid State Drive. &#39;kSATAHDD&#39; indicates storage tier type of SATA Hard Disk Drive. &#39;kCLOUD&#39; indicates storage tier type of Cloud. | [optional] 

## Methods

### NewCountByTier

`func NewCountByTier() *CountByTier`

NewCountByTier instantiates a new CountByTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountByTierWithDefaults

`func NewCountByTierWithDefaults() *CountByTier`

NewCountByTierWithDefaults instantiates a new CountByTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskCount

`func (o *CountByTier) GetDiskCount() int64`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *CountByTier) GetDiskCountOk() (*int64, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *CountByTier) SetDiskCount(v int64)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *CountByTier) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### SetDiskCountNil

`func (o *CountByTier) SetDiskCountNil(b bool)`

 SetDiskCountNil sets the value for DiskCount to be an explicit nil

### UnsetDiskCount
`func (o *CountByTier) UnsetDiskCount()`

UnsetDiskCount ensures that no value is present for DiskCount, not even an explicit nil
### GetStorageTier

`func (o *CountByTier) GetStorageTier() string`

GetStorageTier returns the StorageTier field if non-nil, zero value otherwise.

### GetStorageTierOk

`func (o *CountByTier) GetStorageTierOk() (*string, bool)`

GetStorageTierOk returns a tuple with the StorageTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTier

`func (o *CountByTier) SetStorageTier(v string)`

SetStorageTier sets StorageTier field to given value.

### HasStorageTier

`func (o *CountByTier) HasStorageTier() bool`

HasStorageTier returns a boolean if a field has been set.

### SetStorageTierNil

`func (o *CountByTier) SetStorageTierNil(b bool)`

 SetStorageTierNil sets the value for StorageTier to be an explicit nil

### UnsetStorageTier
`func (o *CountByTier) UnsetStorageTier()`

UnsetStorageTier ensures that no value is present for StorageTier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


