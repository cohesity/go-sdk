# SfdcObjectFieldExclusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeFieldNames** | Pointer to **[]string** | Specifies the list of Sfdc API names of the fields to be excluded in this object. | [optional] 
**ObjectId** | Pointer to **int64** | Specifies the id of the object in which some fields are to be excluded. This should be a leaf level object. | [optional] 

## Methods

### NewSfdcObjectFieldExclusion

`func NewSfdcObjectFieldExclusion() *SfdcObjectFieldExclusion`

NewSfdcObjectFieldExclusion instantiates a new SfdcObjectFieldExclusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSfdcObjectFieldExclusionWithDefaults

`func NewSfdcObjectFieldExclusionWithDefaults() *SfdcObjectFieldExclusion`

NewSfdcObjectFieldExclusionWithDefaults instantiates a new SfdcObjectFieldExclusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeFieldNames

`func (o *SfdcObjectFieldExclusion) GetExcludeFieldNames() []string`

GetExcludeFieldNames returns the ExcludeFieldNames field if non-nil, zero value otherwise.

### GetExcludeFieldNamesOk

`func (o *SfdcObjectFieldExclusion) GetExcludeFieldNamesOk() (*[]string, bool)`

GetExcludeFieldNamesOk returns a tuple with the ExcludeFieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFieldNames

`func (o *SfdcObjectFieldExclusion) SetExcludeFieldNames(v []string)`

SetExcludeFieldNames sets ExcludeFieldNames field to given value.

### HasExcludeFieldNames

`func (o *SfdcObjectFieldExclusion) HasExcludeFieldNames() bool`

HasExcludeFieldNames returns a boolean if a field has been set.

### GetObjectId

`func (o *SfdcObjectFieldExclusion) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *SfdcObjectFieldExclusion) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *SfdcObjectFieldExclusion) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *SfdcObjectFieldExclusion) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


