# DataTieringObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 
**AnalysisInfo** | Pointer to [**DataTieringObjectAnalysisInfo**](DataTieringObjectAnalysisInfo.md) |  | [optional] 

## Methods

### NewDataTieringObjectInfo

`func NewDataTieringObjectInfo(id NullableInt64, ) *DataTieringObjectInfo`

NewDataTieringObjectInfo instantiates a new DataTieringObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringObjectInfoWithDefaults

`func NewDataTieringObjectInfoWithDefaults() *DataTieringObjectInfo`

NewDataTieringObjectInfoWithDefaults instantiates a new DataTieringObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataTieringObjectInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTieringObjectInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTieringObjectInfo) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *DataTieringObjectInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataTieringObjectInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DataTieringObjectInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTieringObjectInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTieringObjectInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataTieringObjectInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DataTieringObjectInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataTieringObjectInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAnalysisInfo

`func (o *DataTieringObjectInfo) GetAnalysisInfo() DataTieringObjectAnalysisInfo`

GetAnalysisInfo returns the AnalysisInfo field if non-nil, zero value otherwise.

### GetAnalysisInfoOk

`func (o *DataTieringObjectInfo) GetAnalysisInfoOk() (*DataTieringObjectAnalysisInfo, bool)`

GetAnalysisInfoOk returns a tuple with the AnalysisInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisInfo

`func (o *DataTieringObjectInfo) SetAnalysisInfo(v DataTieringObjectAnalysisInfo)`

SetAnalysisInfo sets AnalysisInfo field to given value.

### HasAnalysisInfo

`func (o *DataTieringObjectInfo) HasAnalysisInfo() bool`

HasAnalysisInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


