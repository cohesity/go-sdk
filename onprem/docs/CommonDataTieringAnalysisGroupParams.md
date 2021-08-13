# CommonDataTieringAnalysisGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the data tiering analysis group. | 
**Source** | [**DataTieringSource**](DataTieringSource.md) |  | 

## Methods

### NewCommonDataTieringAnalysisGroupParams

`func NewCommonDataTieringAnalysisGroupParams(name NullableString, source DataTieringSource, ) *CommonDataTieringAnalysisGroupParams`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


