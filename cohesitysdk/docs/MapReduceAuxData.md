# MapReduceAuxData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatternVec** | Pointer to [**[]Pattern**](Pattern.md) | Pattern auxiliary data for a MapReduce. | [optional] 

## Methods

### NewMapReduceAuxData

`func NewMapReduceAuxData() *MapReduceAuxData`

NewMapReduceAuxData instantiates a new MapReduceAuxData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapReduceAuxDataWithDefaults

`func NewMapReduceAuxDataWithDefaults() *MapReduceAuxData`

NewMapReduceAuxDataWithDefaults instantiates a new MapReduceAuxData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatternVec

`func (o *MapReduceAuxData) GetPatternVec() []Pattern`

GetPatternVec returns the PatternVec field if non-nil, zero value otherwise.

### GetPatternVecOk

`func (o *MapReduceAuxData) GetPatternVecOk() (*[]Pattern, bool)`

GetPatternVecOk returns a tuple with the PatternVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternVec

`func (o *MapReduceAuxData) SetPatternVec(v []Pattern)`

SetPatternVec sets PatternVec field to given value.

### HasPatternVec

`func (o *MapReduceAuxData) HasPatternVec() bool`

HasPatternVec returns a boolean if a field has been set.

### SetPatternVecNil

`func (o *MapReduceAuxData) SetPatternVecNil(b bool)`

 SetPatternVecNil sets the value for PatternVec to be an explicit nil

### UnsetPatternVec
`func (o *MapReduceAuxData) UnsetPatternVec()`

UnsetPatternVec ensures that no value is present for PatternVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


