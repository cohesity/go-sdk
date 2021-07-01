# AttributeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **[]string** | Specifies the flags related to the attribute values of the AD object. &#39;kError&#39; indicates error in conversion of AD Object value to string. The value in the AdAttributValue contains the error message. &#39;kTruncated&#39; indicates the multi valued attribute is truncated when value exceeded &#39;truncate_multivalues&#39; value specified in the request. &#39;kCSV&#39; indicates content in &#39;values&#39; is a comma separated value (CSV) format of a complex object. | [optional] 
**Values** | Pointer to **[]string** | Specifies list of values for the attribute. | [optional] 

## Methods

### NewAttributeValue

`func NewAttributeValue() *AttributeValue`

NewAttributeValue instantiates a new AttributeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeValueWithDefaults

`func NewAttributeValueWithDefaults() *AttributeValue`

NewAttributeValueWithDefaults instantiates a new AttributeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *AttributeValue) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *AttributeValue) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *AttributeValue) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *AttributeValue) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### SetFlagsNil

`func (o *AttributeValue) SetFlagsNil(b bool)`

 SetFlagsNil sets the value for Flags to be an explicit nil

### UnsetFlags
`func (o *AttributeValue) UnsetFlags()`

UnsetFlags ensures that no value is present for Flags, not even an explicit nil
### GetValues

`func (o *AttributeValue) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AttributeValue) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AttributeValue) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AttributeValue) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *AttributeValue) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *AttributeValue) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


