# CommonDataTieringAnalysisGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the data tiering analysis group. | 
**Schedule** | Pointer to [**DataTieringSchedule**](DataTieringSchedule.md) |  | [optional] 
**Source** | Pointer to [**DataTieringSource**](DataTieringSource.md) |  | [optional] 

## Methods

### NewCommonDataTieringAnalysisGroupParams

`func NewCommonDataTieringAnalysisGroupParams(name NullableString, ) *CommonDataTieringAnalysisGroupParams`

NewCommonDataTieringAnalysisGroupParams instantiates a new CommonDataTieringAnalysisGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonDataTieringAnalysisGroupParamsWithDefaults

`func NewCommonDataTieringAnalysisGroupParamsWithDefaults() *CommonDataTieringAnalysisGroupParams`

NewCommonDataTieringAnalysisGroupParamsWithDefaults instantiates a new CommonDataTieringAnalysisGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonDataTieringAnalysisGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonDataTieringAnalysisGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonDataTieringAnalysisGroupParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CommonDataTieringAnalysisGroupParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonDataTieringAnalysisGroupParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSchedule

`func (o *CommonDataTieringAnalysisGroupParams) GetSchedule() DataTieringSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CommonDataTieringAnalysisGroupParams) GetScheduleOk() (*DataTieringSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CommonDataTieringAnalysisGroupParams) SetSchedule(v DataTieringSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CommonDataTieringAnalysisGroupParams) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSource

`func (o *CommonDataTieringAnalysisGroupParams) GetSource() DataTieringSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CommonDataTieringAnalysisGroupParams) GetSourceOk() (*DataTieringSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CommonDataTieringAnalysisGroupParams) SetSource(v DataTieringSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CommonDataTieringAnalysisGroupParams) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


