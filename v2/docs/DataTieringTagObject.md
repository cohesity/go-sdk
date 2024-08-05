# DataTieringTagObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]DataTieringTag**](DataTieringTag.md) | Array of tag label and value. | [optional] 
**Type** | Pointer to **NullableString** | Specifies type of tag. | [optional] 

## Methods

### NewDataTieringTagObject

`func NewDataTieringTagObject() *DataTieringTagObject`

NewDataTieringTagObject instantiates a new DataTieringTagObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringTagObjectWithDefaults

`func NewDataTieringTagObjectWithDefaults() *DataTieringTagObject`

NewDataTieringTagObjectWithDefaults instantiates a new DataTieringTagObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *DataTieringTagObject) GetTags() []DataTieringTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DataTieringTagObject) GetTagsOk() (*[]DataTieringTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DataTieringTagObject) SetTags(v []DataTieringTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DataTieringTagObject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *DataTieringTagObject) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *DataTieringTagObject) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetType

`func (o *DataTieringTagObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataTieringTagObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataTieringTagObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataTieringTagObject) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DataTieringTagObject) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DataTieringTagObject) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


