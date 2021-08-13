# DataTieringAnalysisGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the data tiering analysis group. | 
**Source** | [**DataTieringSource**](DataTieringSource.md) |  | 
**Id** | Pointer to **NullableString** | Specifies the ID of the data tiering analysis group. | [optional] 
**LastRun** | Pointer to [**DataTieringAnalysisGroupRun**](DataTieringAnalysisGroupRun.md) |  | [optional] 

## Methods

### NewDataTieringAnalysisGroup

`func NewDataTieringAnalysisGroup(name NullableString, source DataTieringSource, ) *DataTieringAnalysisGroup`

NewDataTieringAnalysisGroup instantiates a new DataTieringAnalysisGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringAnalysisGroupWithDefaults

`func NewDataTieringAnalysisGroupWithDefaults() *DataTieringAnalysisGroup`

NewDataTieringAnalysisGroupWithDefaults instantiates a new DataTieringAnalysisGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataTieringAnalysisGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTieringAnalysisGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTieringAnalysisGroup) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *DataTieringAnalysisGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataTieringAnalysisGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSource

`func (o *DataTieringAnalysisGroup) GetSource() DataTieringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataTieringAnalysisGroup) GetSourceOk() (*DataTieringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataTieringAnalysisGroup) SetSource(v DataTieringSource)`

SetSource sets Source field to given value.


### GetId

`func (o *DataTieringAnalysisGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTieringAnalysisGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTieringAnalysisGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataTieringAnalysisGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataTieringAnalysisGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataTieringAnalysisGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastRun

`func (o *DataTieringAnalysisGroup) GetLastRun() DataTieringAnalysisGroupRun`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *DataTieringAnalysisGroup) GetLastRunOk() (*DataTieringAnalysisGroupRun, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *DataTieringAnalysisGroup) SetLastRun(v DataTieringAnalysisGroupRun)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *DataTieringAnalysisGroup) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


