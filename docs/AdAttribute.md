# AdAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdAttributeFlags** | Pointer to **[]string** | Specifies the flags related to the attribute of the AD object. &#39;kEqual&#39; indicates the attribute value of AD object from Snapshot and Production AD are equal. &#39;kNotEqual&#39; indicates the attribute value of AD object from Snapshot and Production AD are not equal. &#39;kNotFound&#39; indicates attribute of the AD object is missing from both Snapshot and Production AD. &#39;kSystem&#39; indicates this is system attribute. This can only be updated by the AD internal component. &#39;kMultiValue&#39; indicates that the attribute is mutli-value attribute. This attribute supports mutli-value merge during attribute restore operation. | [optional] 
**DestinationValue** | Pointer to [**AttributeValue**](AttributeValue.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error message regarding the attribute | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the attribute of the AD object. | [optional] 
**SameValue** | Pointer to [**AttributeValue**](AttributeValue.md) |  | [optional] 
**SourceValue** | Pointer to [**AttributeValue**](AttributeValue.md) |  | [optional] 

## Methods

### NewAdAttribute

`func NewAdAttribute() *AdAttribute`

NewAdAttribute instantiates a new AdAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdAttributeWithDefaults

`func NewAdAttributeWithDefaults() *AdAttribute`

NewAdAttributeWithDefaults instantiates a new AdAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdAttributeFlags

`func (o *AdAttribute) GetAdAttributeFlags() []string`

GetAdAttributeFlags returns the AdAttributeFlags field if non-nil, zero value otherwise.

### GetAdAttributeFlagsOk

`func (o *AdAttribute) GetAdAttributeFlagsOk() (*[]string, bool)`

GetAdAttributeFlagsOk returns a tuple with the AdAttributeFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAttributeFlags

`func (o *AdAttribute) SetAdAttributeFlags(v []string)`

SetAdAttributeFlags sets AdAttributeFlags field to given value.

### HasAdAttributeFlags

`func (o *AdAttribute) HasAdAttributeFlags() bool`

HasAdAttributeFlags returns a boolean if a field has been set.

### SetAdAttributeFlagsNil

`func (o *AdAttribute) SetAdAttributeFlagsNil(b bool)`

 SetAdAttributeFlagsNil sets the value for AdAttributeFlags to be an explicit nil

### UnsetAdAttributeFlags
`func (o *AdAttribute) UnsetAdAttributeFlags()`

UnsetAdAttributeFlags ensures that no value is present for AdAttributeFlags, not even an explicit nil
### GetDestinationValue

`func (o *AdAttribute) GetDestinationValue() AttributeValue`

GetDestinationValue returns the DestinationValue field if non-nil, zero value otherwise.

### GetDestinationValueOk

`func (o *AdAttribute) GetDestinationValueOk() (*AttributeValue, bool)`

GetDestinationValueOk returns a tuple with the DestinationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationValue

`func (o *AdAttribute) SetDestinationValue(v AttributeValue)`

SetDestinationValue sets DestinationValue field to given value.

### HasDestinationValue

`func (o *AdAttribute) HasDestinationValue() bool`

HasDestinationValue returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AdAttribute) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AdAttribute) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AdAttribute) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AdAttribute) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AdAttribute) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AdAttribute) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetName

`func (o *AdAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdAttribute) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdAttribute) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSameValue

`func (o *AdAttribute) GetSameValue() AttributeValue`

GetSameValue returns the SameValue field if non-nil, zero value otherwise.

### GetSameValueOk

`func (o *AdAttribute) GetSameValueOk() (*AttributeValue, bool)`

GetSameValueOk returns a tuple with the SameValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameValue

`func (o *AdAttribute) SetSameValue(v AttributeValue)`

SetSameValue sets SameValue field to given value.

### HasSameValue

`func (o *AdAttribute) HasSameValue() bool`

HasSameValue returns a boolean if a field has been set.

### GetSourceValue

`func (o *AdAttribute) GetSourceValue() AttributeValue`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *AdAttribute) GetSourceValueOk() (*AttributeValue, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *AdAttribute) SetSourceValue(v AttributeValue)`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *AdAttribute) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


