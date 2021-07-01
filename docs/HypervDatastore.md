# HypervDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **NullableInt32** | Specifies the capacity of the datastore in bytes. | [optional] 
**FreeSpace** | Pointer to **NullableInt32** | Specifies the available space on the datastore in bytes. | [optional] 
**MountPoints** | Pointer to **[]string** | Specifies the available mount points on the datastore. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the datastore object like kFileShare or kVolume. overrideDescription: true Specifies the type of a HyperV datastore object. &#39;kFileShare&#39; indicates SMB file share datastore. &#39;kVolume&#39; indicates a volume which can a LUN. | [optional] 

## Methods

### NewHypervDatastore

`func NewHypervDatastore() *HypervDatastore`

NewHypervDatastore instantiates a new HypervDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHypervDatastoreWithDefaults

`func NewHypervDatastoreWithDefaults() *HypervDatastore`

NewHypervDatastoreWithDefaults instantiates a new HypervDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *HypervDatastore) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *HypervDatastore) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *HypervDatastore) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *HypervDatastore) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### SetCapacityNil

`func (o *HypervDatastore) SetCapacityNil(b bool)`

 SetCapacityNil sets the value for Capacity to be an explicit nil

### UnsetCapacity
`func (o *HypervDatastore) UnsetCapacity()`

UnsetCapacity ensures that no value is present for Capacity, not even an explicit nil
### GetFreeSpace

`func (o *HypervDatastore) GetFreeSpace() int32`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *HypervDatastore) GetFreeSpaceOk() (*int32, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *HypervDatastore) SetFreeSpace(v int32)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *HypervDatastore) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### SetFreeSpaceNil

`func (o *HypervDatastore) SetFreeSpaceNil(b bool)`

 SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil

### UnsetFreeSpace
`func (o *HypervDatastore) UnsetFreeSpace()`

UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
### GetMountPoints

`func (o *HypervDatastore) GetMountPoints() []string`

GetMountPoints returns the MountPoints field if non-nil, zero value otherwise.

### GetMountPointsOk

`func (o *HypervDatastore) GetMountPointsOk() (*[]string, bool)`

GetMountPointsOk returns a tuple with the MountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoints

`func (o *HypervDatastore) SetMountPoints(v []string)`

SetMountPoints sets MountPoints field to given value.

### HasMountPoints

`func (o *HypervDatastore) HasMountPoints() bool`

HasMountPoints returns a boolean if a field has been set.

### SetMountPointsNil

`func (o *HypervDatastore) SetMountPointsNil(b bool)`

 SetMountPointsNil sets the value for MountPoints to be an explicit nil

### UnsetMountPoints
`func (o *HypervDatastore) UnsetMountPoints()`

UnsetMountPoints ensures that no value is present for MountPoints, not even an explicit nil
### GetType

`func (o *HypervDatastore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HypervDatastore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HypervDatastore) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HypervDatastore) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HypervDatastore) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HypervDatastore) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


