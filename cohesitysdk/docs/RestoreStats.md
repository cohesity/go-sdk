# RestoreStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumClonedObjects** | Pointer to **NullableInt64** | Specifies the count of cloned objects in the given time frame. | [optional] 
**NumRecoveredObjects** | Pointer to **NullableInt64** | Specifies the count of recovered objects in the given time frame. | [optional] 
**StatsByEnvironment** | Pointer to [**[]RestoreEnvStats**](RestoreEnvStats.md) | Specifies the stats of recovery jobs aggregated by the environment type. | [optional] 

## Methods

### NewRestoreStats

`func NewRestoreStats() *RestoreStats`

NewRestoreStats instantiates a new RestoreStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreStatsWithDefaults

`func NewRestoreStatsWithDefaults() *RestoreStats`

NewRestoreStatsWithDefaults instantiates a new RestoreStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumClonedObjects

`func (o *RestoreStats) GetNumClonedObjects() int64`

GetNumClonedObjects returns the NumClonedObjects field if non-nil, zero value otherwise.

### GetNumClonedObjectsOk

`func (o *RestoreStats) GetNumClonedObjectsOk() (*int64, bool)`

GetNumClonedObjectsOk returns a tuple with the NumClonedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumClonedObjects

`func (o *RestoreStats) SetNumClonedObjects(v int64)`

SetNumClonedObjects sets NumClonedObjects field to given value.

### HasNumClonedObjects

`func (o *RestoreStats) HasNumClonedObjects() bool`

HasNumClonedObjects returns a boolean if a field has been set.

### SetNumClonedObjectsNil

`func (o *RestoreStats) SetNumClonedObjectsNil(b bool)`

 SetNumClonedObjectsNil sets the value for NumClonedObjects to be an explicit nil

### UnsetNumClonedObjects
`func (o *RestoreStats) UnsetNumClonedObjects()`

UnsetNumClonedObjects ensures that no value is present for NumClonedObjects, not even an explicit nil
### GetNumRecoveredObjects

`func (o *RestoreStats) GetNumRecoveredObjects() int64`

GetNumRecoveredObjects returns the NumRecoveredObjects field if non-nil, zero value otherwise.

### GetNumRecoveredObjectsOk

`func (o *RestoreStats) GetNumRecoveredObjectsOk() (*int64, bool)`

GetNumRecoveredObjectsOk returns a tuple with the NumRecoveredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRecoveredObjects

`func (o *RestoreStats) SetNumRecoveredObjects(v int64)`

SetNumRecoveredObjects sets NumRecoveredObjects field to given value.

### HasNumRecoveredObjects

`func (o *RestoreStats) HasNumRecoveredObjects() bool`

HasNumRecoveredObjects returns a boolean if a field has been set.

### SetNumRecoveredObjectsNil

`func (o *RestoreStats) SetNumRecoveredObjectsNil(b bool)`

 SetNumRecoveredObjectsNil sets the value for NumRecoveredObjects to be an explicit nil

### UnsetNumRecoveredObjects
`func (o *RestoreStats) UnsetNumRecoveredObjects()`

UnsetNumRecoveredObjects ensures that no value is present for NumRecoveredObjects, not even an explicit nil
### GetStatsByEnvironment

`func (o *RestoreStats) GetStatsByEnvironment() []RestoreEnvStats`

GetStatsByEnvironment returns the StatsByEnvironment field if non-nil, zero value otherwise.

### GetStatsByEnvironmentOk

`func (o *RestoreStats) GetStatsByEnvironmentOk() (*[]RestoreEnvStats, bool)`

GetStatsByEnvironmentOk returns a tuple with the StatsByEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsByEnvironment

`func (o *RestoreStats) SetStatsByEnvironment(v []RestoreEnvStats)`

SetStatsByEnvironment sets StatsByEnvironment field to given value.

### HasStatsByEnvironment

`func (o *RestoreStats) HasStatsByEnvironment() bool`

HasStatsByEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


