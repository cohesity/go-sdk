# VMFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterString** | Pointer to **NullableString** | Specifies the filter string using wildcard supported strings or regular expressions. | [optional] 
**IsRegularExpression** | Pointer to **NullableBool** | Specifies whether the provided filter string is a regular expression or not. This needs to be explicitly set to true if user is trying to filter by regular expressions. Not providing this value in case of regular expression can result in unintended results. The default value is assumed to be false. | [optional] [default to false]
**CaseSensitive** | Pointer to **NullableBool** | Specifies whether the provided filter string is case sensitive or not. This needs to be explicitly set to true if user is trying to filter by case sensitive expressions. The default value is assumed to be false. | [optional] [default to false]

## Methods

### NewVMFilter

`func NewVMFilter() *VMFilter`

NewVMFilter instantiates a new VMFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMFilterWithDefaults

`func NewVMFilterWithDefaults() *VMFilter`

NewVMFilterWithDefaults instantiates a new VMFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterString

`func (o *VMFilter) GetFilterString() string`

GetFilterString returns the FilterString field if non-nil, zero value otherwise.

### GetFilterStringOk

`func (o *VMFilter) GetFilterStringOk() (*string, bool)`

GetFilterStringOk returns a tuple with the FilterString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterString

`func (o *VMFilter) SetFilterString(v string)`

SetFilterString sets FilterString field to given value.

### HasFilterString

`func (o *VMFilter) HasFilterString() bool`

HasFilterString returns a boolean if a field has been set.

### SetFilterStringNil

`func (o *VMFilter) SetFilterStringNil(b bool)`

 SetFilterStringNil sets the value for FilterString to be an explicit nil

### UnsetFilterString
`func (o *VMFilter) UnsetFilterString()`

UnsetFilterString ensures that no value is present for FilterString, not even an explicit nil
### GetIsRegularExpression

`func (o *VMFilter) GetIsRegularExpression() bool`

GetIsRegularExpression returns the IsRegularExpression field if non-nil, zero value otherwise.

### GetIsRegularExpressionOk

`func (o *VMFilter) GetIsRegularExpressionOk() (*bool, bool)`

GetIsRegularExpressionOk returns a tuple with the IsRegularExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegularExpression

`func (o *VMFilter) SetIsRegularExpression(v bool)`

SetIsRegularExpression sets IsRegularExpression field to given value.

### HasIsRegularExpression

`func (o *VMFilter) HasIsRegularExpression() bool`

HasIsRegularExpression returns a boolean if a field has been set.

### SetIsRegularExpressionNil

`func (o *VMFilter) SetIsRegularExpressionNil(b bool)`

 SetIsRegularExpressionNil sets the value for IsRegularExpression to be an explicit nil

### UnsetIsRegularExpression
`func (o *VMFilter) UnsetIsRegularExpression()`

UnsetIsRegularExpression ensures that no value is present for IsRegularExpression, not even an explicit nil
### GetCaseSensitive

`func (o *VMFilter) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *VMFilter) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *VMFilter) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *VMFilter) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### SetCaseSensitiveNil

`func (o *VMFilter) SetCaseSensitiveNil(b bool)`

 SetCaseSensitiveNil sets the value for CaseSensitive to be an explicit nil

### UnsetCaseSensitive
`func (o *VMFilter) UnsetCaseSensitive()`

UnsetCaseSensitive ensures that no value is present for CaseSensitive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


