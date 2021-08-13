# HeliosSourceObjectsStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceName** | Pointer to **NullableString** | Name of the source | [optional] 
**NumAssignedObjects** | Pointer to **NullableInt64** | Number of objects assigned. | [optional] 

## Methods

### NewHeliosSourceObjectsStats

`func NewHeliosSourceObjectsStats() *HeliosSourceObjectsStats`

NewHeliosSourceObjectsStats instantiates a new HeliosSourceObjectsStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosSourceObjectsStatsWithDefaults

`func NewHeliosSourceObjectsStatsWithDefaults() *HeliosSourceObjectsStats`

NewHeliosSourceObjectsStatsWithDefaults instantiates a new HeliosSourceObjectsStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceName

`func (o *HeliosSourceObjectsStats) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HeliosSourceObjectsStats) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HeliosSourceObjectsStats) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *HeliosSourceObjectsStats) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *HeliosSourceObjectsStats) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *HeliosSourceObjectsStats) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetNumAssignedObjects

`func (o *HeliosSourceObjectsStats) GetNumAssignedObjects() int64`

GetNumAssignedObjects returns the NumAssignedObjects field if non-nil, zero value otherwise.

### GetNumAssignedObjectsOk

`func (o *HeliosSourceObjectsStats) GetNumAssignedObjectsOk() (*int64, bool)`

GetNumAssignedObjectsOk returns a tuple with the NumAssignedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssignedObjects

`func (o *HeliosSourceObjectsStats) SetNumAssignedObjects(v int64)`

SetNumAssignedObjects sets NumAssignedObjects field to given value.

### HasNumAssignedObjects

`func (o *HeliosSourceObjectsStats) HasNumAssignedObjects() bool`

HasNumAssignedObjects returns a boolean if a field has been set.

### SetNumAssignedObjectsNil

`func (o *HeliosSourceObjectsStats) SetNumAssignedObjectsNil(b bool)`

 SetNumAssignedObjectsNil sets the value for NumAssignedObjects to be an explicit nil

### UnsetNumAssignedObjects
`func (o *HeliosSourceObjectsStats) UnsetNumAssignedObjects()`

UnsetNumAssignedObjects ensures that no value is present for NumAssignedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


