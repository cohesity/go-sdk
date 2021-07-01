# AttributeRestoreInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **[]string** | Specifes the error messages corresponding to restore of the attribute. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the attribute of the AD object. | [optional] 

## Methods

### NewAttributeRestoreInformation

`func NewAttributeRestoreInformation() *AttributeRestoreInformation`

NewAttributeRestoreInformation instantiates a new AttributeRestoreInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeRestoreInformationWithDefaults

`func NewAttributeRestoreInformationWithDefaults() *AttributeRestoreInformation`

NewAttributeRestoreInformationWithDefaults instantiates a new AttributeRestoreInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *AttributeRestoreInformation) GetErrorMessage() []string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AttributeRestoreInformation) GetErrorMessageOk() (*[]string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AttributeRestoreInformation) SetErrorMessage(v []string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AttributeRestoreInformation) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AttributeRestoreInformation) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AttributeRestoreInformation) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetName

`func (o *AttributeRestoreInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeRestoreInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeRestoreInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttributeRestoreInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AttributeRestoreInformation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AttributeRestoreInformation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


