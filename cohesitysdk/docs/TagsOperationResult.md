# TagsOperationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocErrors** | Pointer to [**[]DocError**](DocError.md) | DocErrors are document errors incurred in yoda service while tagging. | [optional] 

## Methods

### NewTagsOperationResult

`func NewTagsOperationResult() *TagsOperationResult`

NewTagsOperationResult instantiates a new TagsOperationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsOperationResultWithDefaults

`func NewTagsOperationResultWithDefaults() *TagsOperationResult`

NewTagsOperationResultWithDefaults instantiates a new TagsOperationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocErrors

`func (o *TagsOperationResult) GetDocErrors() []DocError`

GetDocErrors returns the DocErrors field if non-nil, zero value otherwise.

### GetDocErrorsOk

`func (o *TagsOperationResult) GetDocErrorsOk() (*[]DocError, bool)`

GetDocErrorsOk returns a tuple with the DocErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocErrors

`func (o *TagsOperationResult) SetDocErrors(v []DocError)`

SetDocErrors sets DocErrors field to given value.

### HasDocErrors

`func (o *TagsOperationResult) HasDocErrors() bool`

HasDocErrors returns a boolean if a field has been set.

### SetDocErrorsNil

`func (o *TagsOperationResult) SetDocErrorsNil(b bool)`

 SetDocErrorsNil sets the value for DocErrors to be an explicit nil

### UnsetDocErrors
`func (o *TagsOperationResult) UnsetDocErrors()`

UnsetDocErrors ensures that no value is present for DocErrors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


