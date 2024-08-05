# UptieringTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DowntieredDataLocations** | Pointer to [**[]DowntieredDataLocation**](DowntieredDataLocation.md) | Specifies a list of mapping between sources and its corresponding viewNames and mountPaths, where the sources were downtiered. | [optional] 
**StorageDomainId** | **NullableInt64** | Specifies the Storage Domain ID where the view will be kept. | 

## Methods

### NewUptieringTarget

`func NewUptieringTarget(storageDomainId NullableInt64, ) *UptieringTarget`

NewUptieringTarget instantiates a new UptieringTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUptieringTargetWithDefaults

`func NewUptieringTargetWithDefaults() *UptieringTarget`

NewUptieringTargetWithDefaults instantiates a new UptieringTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDowntieredDataLocations

`func (o *UptieringTarget) GetDowntieredDataLocations() []DowntieredDataLocation`

GetDowntieredDataLocations returns the DowntieredDataLocations field if non-nil, zero value otherwise.

### GetDowntieredDataLocationsOk

`func (o *UptieringTarget) GetDowntieredDataLocationsOk() (*[]DowntieredDataLocation, bool)`

GetDowntieredDataLocationsOk returns a tuple with the DowntieredDataLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowntieredDataLocations

`func (o *UptieringTarget) SetDowntieredDataLocations(v []DowntieredDataLocation)`

SetDowntieredDataLocations sets DowntieredDataLocations field to given value.

### HasDowntieredDataLocations

`func (o *UptieringTarget) HasDowntieredDataLocations() bool`

HasDowntieredDataLocations returns a boolean if a field has been set.

### SetDowntieredDataLocationsNil

`func (o *UptieringTarget) SetDowntieredDataLocationsNil(b bool)`

 SetDowntieredDataLocationsNil sets the value for DowntieredDataLocations to be an explicit nil

### UnsetDowntieredDataLocations
`func (o *UptieringTarget) UnsetDowntieredDataLocations()`

UnsetDowntieredDataLocations ensures that no value is present for DowntieredDataLocations, not even an explicit nil
### GetStorageDomainId

`func (o *UptieringTarget) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *UptieringTarget) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *UptieringTarget) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.


### SetStorageDomainIdNil

`func (o *UptieringTarget) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *UptieringTarget) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


