# AAGInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the AAG name. | [optional] 
**ObjectId** | Pointer to **NullableInt64** | Specifies the AAG object Id. | [optional] 

## Methods

### NewAAGInfo

`func NewAAGInfo() *AAGInfo`

NewAAGInfo instantiates a new AAGInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAAGInfoWithDefaults

`func NewAAGInfoWithDefaults() *AAGInfo`

NewAAGInfoWithDefaults instantiates a new AAGInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AAGInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AAGInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AAGInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AAGInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AAGInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AAGInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjectId

`func (o *AAGInfo) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AAGInfo) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AAGInfo) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *AAGInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *AAGInfo) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *AAGInfo) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


