# SourceAttributeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterAttribute** | Pointer to **NullableString** | Specifies the filter attribute for the source. | [optional] 
**AttributeValues** | Pointer to **[]string** | Specifies the list of attribute values for above filter. | [optional] 

## Methods

### NewSourceAttributeFilter

`func NewSourceAttributeFilter() *SourceAttributeFilter`

NewSourceAttributeFilter instantiates a new SourceAttributeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAttributeFilterWithDefaults

`func NewSourceAttributeFilterWithDefaults() *SourceAttributeFilter`

NewSourceAttributeFilterWithDefaults instantiates a new SourceAttributeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterAttribute

`func (o *SourceAttributeFilter) GetFilterAttribute() string`

GetFilterAttribute returns the FilterAttribute field if non-nil, zero value otherwise.

### GetFilterAttributeOk

`func (o *SourceAttributeFilter) GetFilterAttributeOk() (*string, bool)`

GetFilterAttributeOk returns a tuple with the FilterAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterAttribute

`func (o *SourceAttributeFilter) SetFilterAttribute(v string)`

SetFilterAttribute sets FilterAttribute field to given value.

### HasFilterAttribute

`func (o *SourceAttributeFilter) HasFilterAttribute() bool`

HasFilterAttribute returns a boolean if a field has been set.

### SetFilterAttributeNil

`func (o *SourceAttributeFilter) SetFilterAttributeNil(b bool)`

 SetFilterAttributeNil sets the value for FilterAttribute to be an explicit nil

### UnsetFilterAttribute
`func (o *SourceAttributeFilter) UnsetFilterAttribute()`

UnsetFilterAttribute ensures that no value is present for FilterAttribute, not even an explicit nil
### GetAttributeValues

`func (o *SourceAttributeFilter) GetAttributeValues() []string`

GetAttributeValues returns the AttributeValues field if non-nil, zero value otherwise.

### GetAttributeValuesOk

`func (o *SourceAttributeFilter) GetAttributeValuesOk() (*[]string, bool)`

GetAttributeValuesOk returns a tuple with the AttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValues

`func (o *SourceAttributeFilter) SetAttributeValues(v []string)`

SetAttributeValues sets AttributeValues field to given value.

### HasAttributeValues

`func (o *SourceAttributeFilter) HasAttributeValues() bool`

HasAttributeValues returns a boolean if a field has been set.

### SetAttributeValuesNil

`func (o *SourceAttributeFilter) SetAttributeValuesNil(b bool)`

 SetAttributeValuesNil sets the value for AttributeValues to be an explicit nil

### UnsetAttributeValues
`func (o *SourceAttributeFilter) UnsetAttributeValues()`

UnsetAttributeValues ensures that no value is present for AttributeValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


