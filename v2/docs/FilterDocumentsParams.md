# FilterDocumentsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentFilterType** | **NullableString** | Specifies the filter type for Documents to be restored. | 
**FilterExpression** | Pointer to **NullableString** | A filter expression to match Documents content to be restored. | [optional] 
**IdRegex** | Pointer to **NullableString** | A regular expression to match Documents ID&#39;s to be restored. | [optional] 

## Methods

### NewFilterDocumentsParams

`func NewFilterDocumentsParams(documentFilterType NullableString, ) *FilterDocumentsParams`

NewFilterDocumentsParams instantiates a new FilterDocumentsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDocumentsParamsWithDefaults

`func NewFilterDocumentsParamsWithDefaults() *FilterDocumentsParams`

NewFilterDocumentsParamsWithDefaults instantiates a new FilterDocumentsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentFilterType

`func (o *FilterDocumentsParams) GetDocumentFilterType() string`

GetDocumentFilterType returns the DocumentFilterType field if non-nil, zero value otherwise.

### GetDocumentFilterTypeOk

`func (o *FilterDocumentsParams) GetDocumentFilterTypeOk() (*string, bool)`

GetDocumentFilterTypeOk returns a tuple with the DocumentFilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFilterType

`func (o *FilterDocumentsParams) SetDocumentFilterType(v string)`

SetDocumentFilterType sets DocumentFilterType field to given value.


### SetDocumentFilterTypeNil

`func (o *FilterDocumentsParams) SetDocumentFilterTypeNil(b bool)`

 SetDocumentFilterTypeNil sets the value for DocumentFilterType to be an explicit nil

### UnsetDocumentFilterType
`func (o *FilterDocumentsParams) UnsetDocumentFilterType()`

UnsetDocumentFilterType ensures that no value is present for DocumentFilterType, not even an explicit nil
### GetFilterExpression

`func (o *FilterDocumentsParams) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *FilterDocumentsParams) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *FilterDocumentsParams) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *FilterDocumentsParams) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### SetFilterExpressionNil

`func (o *FilterDocumentsParams) SetFilterExpressionNil(b bool)`

 SetFilterExpressionNil sets the value for FilterExpression to be an explicit nil

### UnsetFilterExpression
`func (o *FilterDocumentsParams) UnsetFilterExpression()`

UnsetFilterExpression ensures that no value is present for FilterExpression, not even an explicit nil
### GetIdRegex

`func (o *FilterDocumentsParams) GetIdRegex() string`

GetIdRegex returns the IdRegex field if non-nil, zero value otherwise.

### GetIdRegexOk

`func (o *FilterDocumentsParams) GetIdRegexOk() (*string, bool)`

GetIdRegexOk returns a tuple with the IdRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdRegex

`func (o *FilterDocumentsParams) SetIdRegex(v string)`

SetIdRegex sets IdRegex field to given value.

### HasIdRegex

`func (o *FilterDocumentsParams) HasIdRegex() bool`

HasIdRegex returns a boolean if a field has been set.

### SetIdRegexNil

`func (o *FilterDocumentsParams) SetIdRegexNil(b bool)`

 SetIdRegexNil sets the value for IdRegex to be an explicit nil

### UnsetIdRegex
`func (o *FilterDocumentsParams) UnsetIdRegex()`

UnsetIdRegex ensures that no value is present for IdRegex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


