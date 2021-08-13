# DataTieringAnalysisGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the ID of the data tiering analysis group. | [optional] 
**LastRun** | Pointer to [**DataTieringAnalysisGroupRun**](DataTieringAnalysisGroupRun.md) |  | [optional] 

## Methods

### NewDataTieringAnalysisGroupAllOf

`func NewDataTieringAnalysisGroupAllOf() *DataTieringAnalysisGroupAllOf`

NewDataTieringAnalysisGroupAllOf instantiates a new DataTieringAnalysisGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringAnalysisGroupAllOfWithDefaults

`func NewDataTieringAnalysisGroupAllOfWithDefaults() *DataTieringAnalysisGroupAllOf`

NewDataTieringAnalysisGroupAllOfWithDefaults instantiates a new DataTieringAnalysisGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataTieringAnalysisGroupAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTieringAnalysisGroupAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTieringAnalysisGroupAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataTieringAnalysisGroupAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataTieringAnalysisGroupAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataTieringAnalysisGroupAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastRun

`func (o *DataTieringAnalysisGroupAllOf) GetLastRun() DataTieringAnalysisGroupRun`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *DataTieringAnalysisGroupAllOf) GetLastRunOk() (*DataTieringAnalysisGroupRun, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *DataTieringAnalysisGroupAllOf) SetLastRun(v DataTieringAnalysisGroupRun)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *DataTieringAnalysisGroupAllOf) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


