# DataTieringAnalysisGroupRuns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsResponseTruncated** | Pointer to **NullableBool** | Indicates whether the result is truncated due to hitting maximum size limit | [optional] 
**Runs** | Pointer to [**[]DataTieringAnalysisGroupRun**](DataTieringAnalysisGroupRun.md) | Specifies the data tiering analysis group runs. | [optional] 

## Methods

### NewDataTieringAnalysisGroupRuns

`func NewDataTieringAnalysisGroupRuns() *DataTieringAnalysisGroupRuns`

NewDataTieringAnalysisGroupRuns instantiates a new DataTieringAnalysisGroupRuns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringAnalysisGroupRunsWithDefaults

`func NewDataTieringAnalysisGroupRunsWithDefaults() *DataTieringAnalysisGroupRuns`

NewDataTieringAnalysisGroupRunsWithDefaults instantiates a new DataTieringAnalysisGroupRuns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsResponseTruncated

`func (o *DataTieringAnalysisGroupRuns) GetIsResponseTruncated() bool`

GetIsResponseTruncated returns the IsResponseTruncated field if non-nil, zero value otherwise.

### GetIsResponseTruncatedOk

`func (o *DataTieringAnalysisGroupRuns) GetIsResponseTruncatedOk() (*bool, bool)`

GetIsResponseTruncatedOk returns a tuple with the IsResponseTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResponseTruncated

`func (o *DataTieringAnalysisGroupRuns) SetIsResponseTruncated(v bool)`

SetIsResponseTruncated sets IsResponseTruncated field to given value.

### HasIsResponseTruncated

`func (o *DataTieringAnalysisGroupRuns) HasIsResponseTruncated() bool`

HasIsResponseTruncated returns a boolean if a field has been set.

### SetIsResponseTruncatedNil

`func (o *DataTieringAnalysisGroupRuns) SetIsResponseTruncatedNil(b bool)`

 SetIsResponseTruncatedNil sets the value for IsResponseTruncated to be an explicit nil

### UnsetIsResponseTruncated
`func (o *DataTieringAnalysisGroupRuns) UnsetIsResponseTruncated()`

UnsetIsResponseTruncated ensures that no value is present for IsResponseTruncated, not even an explicit nil
### GetRuns

`func (o *DataTieringAnalysisGroupRuns) GetRuns() []DataTieringAnalysisGroupRun`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *DataTieringAnalysisGroupRuns) GetRunsOk() (*[]DataTieringAnalysisGroupRun, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *DataTieringAnalysisGroupRuns) SetRuns(v []DataTieringAnalysisGroupRun)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *DataTieringAnalysisGroupRuns) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### SetRunsNil

`func (o *DataTieringAnalysisGroupRuns) SetRunsNil(b bool)`

 SetRunsNil sets the value for Runs to be an explicit nil

### UnsetRuns
`func (o *DataTieringAnalysisGroupRuns) UnsetRuns()`

UnsetRuns ensures that no value is present for Runs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


