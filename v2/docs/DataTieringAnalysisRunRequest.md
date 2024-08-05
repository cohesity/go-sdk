# DataTieringAnalysisRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shares** | Pointer to [**[]DataTieringAnalysisShareInfo**](DataTieringAnalysisShareInfo.md) | Specifies the list of shares to analyse. | [optional] 

## Methods

### NewDataTieringAnalysisRunRequest

`func NewDataTieringAnalysisRunRequest() *DataTieringAnalysisRunRequest`

NewDataTieringAnalysisRunRequest instantiates a new DataTieringAnalysisRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringAnalysisRunRequestWithDefaults

`func NewDataTieringAnalysisRunRequestWithDefaults() *DataTieringAnalysisRunRequest`

NewDataTieringAnalysisRunRequestWithDefaults instantiates a new DataTieringAnalysisRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShares

`func (o *DataTieringAnalysisRunRequest) GetShares() []DataTieringAnalysisShareInfo`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *DataTieringAnalysisRunRequest) GetSharesOk() (*[]DataTieringAnalysisShareInfo, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *DataTieringAnalysisRunRequest) SetShares(v []DataTieringAnalysisShareInfo)`

SetShares sets Shares field to given value.

### HasShares

`func (o *DataTieringAnalysisRunRequest) HasShares() bool`

HasShares returns a boolean if a field has been set.

### SetSharesNil

`func (o *DataTieringAnalysisRunRequest) SetSharesNil(b bool)`

 SetSharesNil sets the value for Shares to be an explicit nil

### UnsetShares
`func (o *DataTieringAnalysisRunRequest) UnsetShares()`

UnsetShares ensures that no value is present for Shares, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


