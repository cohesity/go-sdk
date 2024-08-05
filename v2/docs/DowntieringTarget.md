# DowntieringTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DowntieredDataLocations** | Pointer to [**[]DowntieredDataLocation**](DowntieredDataLocation.md) | Specifies a list of mapping between sources and its corresponding viewNames and mountPaths, where the sources were downtiered. | [optional] [readonly] 
**MountPathPrefix** | Pointer to **NullableString** | Specifies the mount path prefix inside the view. | [optional] 
**StorageDomainId** | **NullableInt64** | Specifies the Storage Domain ID where the view will be kept. | 
**ViewNamePrefix** | **NullableString** | Specifies the view name prefix. | 

## Methods

### NewDowntieringTarget

`func NewDowntieringTarget(storageDomainId NullableInt64, viewNamePrefix NullableString, ) *DowntieringTarget`

NewDowntieringTarget instantiates a new DowntieringTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntieringTargetWithDefaults

`func NewDowntieringTargetWithDefaults() *DowntieringTarget`

NewDowntieringTargetWithDefaults instantiates a new DowntieringTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDowntieredDataLocations

`func (o *DowntieringTarget) GetDowntieredDataLocations() []DowntieredDataLocation`

GetDowntieredDataLocations returns the DowntieredDataLocations field if non-nil, zero value otherwise.

### GetDowntieredDataLocationsOk

`func (o *DowntieringTarget) GetDowntieredDataLocationsOk() (*[]DowntieredDataLocation, bool)`

GetDowntieredDataLocationsOk returns a tuple with the DowntieredDataLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntieredDataLocations

`func (o *DowntieringTarget) SetDowntieredDataLocations(v []DowntieredDataLocation)`

SetDowntieredDataLocations sets DowntieredDataLocations field to given value.

### HasDowntieredDataLocations

`func (o *DowntieringTarget) HasDowntieredDataLocations() bool`

HasDowntieredDataLocations returns a boolean if a field has been set.

### SetDowntieredDataLocationsNil

`func (o *DowntieringTarget) SetDowntieredDataLocationsNil(b bool)`

 SetDowntieredDataLocationsNil sets the value for DowntieredDataLocations to be an explicit nil

### UnsetDowntieredDataLocations
`func (o *DowntieringTarget) UnsetDowntieredDataLocations()`

UnsetDowntieredDataLocations ensures that no value is present for DowntieredDataLocations, not even an explicit nil
### GetMountPathPrefix

`func (o *DowntieringTarget) GetMountPathPrefix() string`

GetMountPathPrefix returns the MountPathPrefix field if non-nil, zero value otherwise.

### GetMountPathPrefixOk

`func (o *DowntieringTarget) GetMountPathPrefixOk() (*string, bool)`

GetMountPathPrefixOk returns a tuple with the MountPathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPathPrefix

`func (o *DowntieringTarget) SetMountPathPrefix(v string)`

SetMountPathPrefix sets MountPathPrefix field to given value.

### HasMountPathPrefix

`func (o *DowntieringTarget) HasMountPathPrefix() bool`

HasMountPathPrefix returns a boolean if a field has been set.

### SetMountPathPrefixNil

`func (o *DowntieringTarget) SetMountPathPrefixNil(b bool)`

 SetMountPathPrefixNil sets the value for MountPathPrefix to be an explicit nil

### UnsetMountPathPrefix
`func (o *DowntieringTarget) UnsetMountPathPrefix()`

UnsetMountPathPrefix ensures that no value is present for MountPathPrefix, not even an explicit nil
### GetStorageDomainId

`func (o *DowntieringTarget) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *DowntieringTarget) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *DowntieringTarget) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.


### SetStorageDomainIdNil

`func (o *DowntieringTarget) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *DowntieringTarget) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetViewNamePrefix

`func (o *DowntieringTarget) GetViewNamePrefix() string`

GetViewNamePrefix returns the ViewNamePrefix field if non-nil, zero value otherwise.

### GetViewNamePrefixOk

`func (o *DowntieringTarget) GetViewNamePrefixOk() (*string, bool)`

GetViewNamePrefixOk returns a tuple with the ViewNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNamePrefix

`func (o *DowntieringTarget) SetViewNamePrefix(v string)`

SetViewNamePrefix sets ViewNamePrefix field to given value.


### SetViewNamePrefixNil

`func (o *DowntieringTarget) SetViewNamePrefixNil(b bool)`

 SetViewNamePrefixNil sets the value for ViewNamePrefix to be an explicit nil

### UnsetViewNamePrefix
`func (o *DowntieringTarget) UnsetViewNamePrefix()`

UnsetViewNamePrefix ensures that no value is present for ViewNamePrefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


