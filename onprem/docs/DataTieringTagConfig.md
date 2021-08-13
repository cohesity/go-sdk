# DataTieringTagConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the ID of the data tiering analysis group. | [optional] [readonly] 
**TagsInfo** | [**[]DataTieringTagObject**](DataTieringTagObject.md) | Array of Tag objects. | 

## Methods

### NewDataTieringTagConfig

`func NewDataTieringTagConfig(tagsInfo []DataTieringTagObject, ) *DataTieringTagConfig`

NewDataTieringTagConfig instantiates a new DataTieringTagConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringTagConfigWithDefaults

`func NewDataTieringTagConfigWithDefaults() *DataTieringTagConfig`

NewDataTieringTagConfigWithDefaults instantiates a new DataTieringTagConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataTieringTagConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTieringTagConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTieringTagConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataTieringTagConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataTieringTagConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataTieringTagConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTagsInfo

`func (o *DataTieringTagConfig) GetTagsInfo() []DataTieringTagObject`

GetTagsInfo returns the TagsInfo field if non-nil, zero value otherwise.

### GetTagsInfoOk

`func (o *DataTieringTagConfig) GetTagsInfoOk() (*[]DataTieringTagObject, bool)`

GetTagsInfoOk returns a tuple with the TagsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsInfo

`func (o *DataTieringTagConfig) SetTagsInfo(v []DataTieringTagObject)`

SetTagsInfo sets TagsInfo field to given value.


### SetTagsInfoNil

`func (o *DataTieringTagConfig) SetTagsInfoNil(b bool)`

 SetTagsInfoNil sets the value for TagsInfo to be an explicit nil

### UnsetTagsInfo
`func (o *DataTieringTagConfig) UnsetTagsInfo()`

UnsetTagsInfo ensures that no value is present for TagsInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


