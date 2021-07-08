# RunNowParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseIds** | Pointer to **[]int64** | Specifies the ids of the DB&#39;s to perform run now on. | [optional] 
**PhysicalParams** | Pointer to [**RunNowPhysicalParameters**](RunNowPhysicalParameters.md) |  | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the source id of the Databases to perform the Run Now operation on. | [optional] 

## Methods

### NewRunNowParameters

`func NewRunNowParameters() *RunNowParameters`

NewRunNowParameters instantiates a new RunNowParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunNowParametersWithDefaults

`func NewRunNowParametersWithDefaults() *RunNowParameters`

NewRunNowParametersWithDefaults instantiates a new RunNowParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseIds

`func (o *RunNowParameters) GetDatabaseIds() []int64`

GetDatabaseIds returns the DatabaseIds field if non-nil, zero value otherwise.

### GetDatabaseIdsOk

`func (o *RunNowParameters) GetDatabaseIdsOk() (*[]int64, bool)`

GetDatabaseIdsOk returns a tuple with the DatabaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseIds

`func (o *RunNowParameters) SetDatabaseIds(v []int64)`

SetDatabaseIds sets DatabaseIds field to given value.

### HasDatabaseIds

`func (o *RunNowParameters) HasDatabaseIds() bool`

HasDatabaseIds returns a boolean if a field has been set.

### SetDatabaseIdsNil

`func (o *RunNowParameters) SetDatabaseIdsNil(b bool)`

 SetDatabaseIdsNil sets the value for DatabaseIds to be an explicit nil

### UnsetDatabaseIds
`func (o *RunNowParameters) UnsetDatabaseIds()`

UnsetDatabaseIds ensures that no value is present for DatabaseIds, not even an explicit nil
### GetPhysicalParams

`func (o *RunNowParameters) GetPhysicalParams() RunNowPhysicalParameters`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *RunNowParameters) GetPhysicalParamsOk() (*RunNowPhysicalParameters, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *RunNowParameters) SetPhysicalParams(v RunNowPhysicalParameters)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *RunNowParameters) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.

### GetSourceId

`func (o *RunNowParameters) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *RunNowParameters) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *RunNowParameters) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *RunNowParameters) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *RunNowParameters) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *RunNowParameters) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


