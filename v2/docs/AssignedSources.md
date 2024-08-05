# AssignedSources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceIds** | Pointer to **[]int64** | Specifies a list of source Ids assigned to a principal. | [optional] 
**ViewIds** | Pointer to **[]int64** | Specifies a list of view Ids assigned to a principal. | [optional] 

## Methods

### NewAssignedSources

`func NewAssignedSources() *AssignedSources`

NewAssignedSources instantiates a new AssignedSources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignedSourcesWithDefaults

`func NewAssignedSourcesWithDefaults() *AssignedSources`

NewAssignedSourcesWithDefaults instantiates a new AssignedSources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceIds

`func (o *AssignedSources) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *AssignedSources) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *AssignedSources) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *AssignedSources) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### GetViewIds

`func (o *AssignedSources) GetViewIds() []int64`

GetViewIds returns the ViewIds field if non-nil, zero value otherwise.

### GetViewIdsOk

`func (o *AssignedSources) GetViewIdsOk() (*[]int64, bool)`

GetViewIdsOk returns a tuple with the ViewIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewIds

`func (o *AssignedSources) SetViewIds(v []int64)`

SetViewIds sets ViewIds field to given value.

### HasViewIds

`func (o *AssignedSources) HasViewIds() bool`

HasViewIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


