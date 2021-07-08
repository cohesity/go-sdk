# CompareADObjectsResultADAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttrFlags** | Pointer to **NullableInt32** | Object result flags of type ADAttributeFlags. | [optional] 
**DestValue** | Pointer to [**CompareADObjectsResultADAttributeValue**](CompareADObjectsResultADAttributeValue.md) |  | [optional] 
**LdapName** | Pointer to **NullableString** | LDAP attribute name. | [optional] 
**SameValue** | Pointer to [**CompareADObjectsResultADAttributeValue**](CompareADObjectsResultADAttributeValue.md) |  | [optional] 
**SourceValue** | Pointer to [**CompareADObjectsResultADAttributeValue**](CompareADObjectsResultADAttributeValue.md) |  | [optional] 
**Status** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 

## Methods

### NewCompareADObjectsResultADAttribute

`func NewCompareADObjectsResultADAttribute() *CompareADObjectsResultADAttribute`

NewCompareADObjectsResultADAttribute instantiates a new CompareADObjectsResultADAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompareADObjectsResultADAttributeWithDefaults

`func NewCompareADObjectsResultADAttributeWithDefaults() *CompareADObjectsResultADAttribute`

NewCompareADObjectsResultADAttributeWithDefaults instantiates a new CompareADObjectsResultADAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrFlags

`func (o *CompareADObjectsResultADAttribute) GetAttrFlags() int32`

GetAttrFlags returns the AttrFlags field if non-nil, zero value otherwise.

### GetAttrFlagsOk

`func (o *CompareADObjectsResultADAttribute) GetAttrFlagsOk() (*int32, bool)`

GetAttrFlagsOk returns a tuple with the AttrFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrFlags

`func (o *CompareADObjectsResultADAttribute) SetAttrFlags(v int32)`

SetAttrFlags sets AttrFlags field to given value.

### HasAttrFlags

`func (o *CompareADObjectsResultADAttribute) HasAttrFlags() bool`

HasAttrFlags returns a boolean if a field has been set.

### SetAttrFlagsNil

`func (o *CompareADObjectsResultADAttribute) SetAttrFlagsNil(b bool)`

 SetAttrFlagsNil sets the value for AttrFlags to be an explicit nil

### UnsetAttrFlags
`func (o *CompareADObjectsResultADAttribute) UnsetAttrFlags()`

UnsetAttrFlags ensures that no value is present for AttrFlags, not even an explicit nil
### GetDestValue

`func (o *CompareADObjectsResultADAttribute) GetDestValue() CompareADObjectsResultADAttributeValue`

GetDestValue returns the DestValue field if non-nil, zero value otherwise.

### GetDestValueOk

`func (o *CompareADObjectsResultADAttribute) GetDestValueOk() (*CompareADObjectsResultADAttributeValue, bool)`

GetDestValueOk returns a tuple with the DestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestValue

`func (o *CompareADObjectsResultADAttribute) SetDestValue(v CompareADObjectsResultADAttributeValue)`

SetDestValue sets DestValue field to given value.

### HasDestValue

`func (o *CompareADObjectsResultADAttribute) HasDestValue() bool`

HasDestValue returns a boolean if a field has been set.

### GetLdapName

`func (o *CompareADObjectsResultADAttribute) GetLdapName() string`

GetLdapName returns the LdapName field if non-nil, zero value otherwise.

### GetLdapNameOk

`func (o *CompareADObjectsResultADAttribute) GetLdapNameOk() (*string, bool)`

GetLdapNameOk returns a tuple with the LdapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapName

`func (o *CompareADObjectsResultADAttribute) SetLdapName(v string)`

SetLdapName sets LdapName field to given value.

### HasLdapName

`func (o *CompareADObjectsResultADAttribute) HasLdapName() bool`

HasLdapName returns a boolean if a field has been set.

### SetLdapNameNil

`func (o *CompareADObjectsResultADAttribute) SetLdapNameNil(b bool)`

 SetLdapNameNil sets the value for LdapName to be an explicit nil

### UnsetLdapName
`func (o *CompareADObjectsResultADAttribute) UnsetLdapName()`

UnsetLdapName ensures that no value is present for LdapName, not even an explicit nil
### GetSameValue

`func (o *CompareADObjectsResultADAttribute) GetSameValue() CompareADObjectsResultADAttributeValue`

GetSameValue returns the SameValue field if non-nil, zero value otherwise.

### GetSameValueOk

`func (o *CompareADObjectsResultADAttribute) GetSameValueOk() (*CompareADObjectsResultADAttributeValue, bool)`

GetSameValueOk returns a tuple with the SameValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameValue

`func (o *CompareADObjectsResultADAttribute) SetSameValue(v CompareADObjectsResultADAttributeValue)`

SetSameValue sets SameValue field to given value.

### HasSameValue

`func (o *CompareADObjectsResultADAttribute) HasSameValue() bool`

HasSameValue returns a boolean if a field has been set.

### GetSourceValue

`func (o *CompareADObjectsResultADAttribute) GetSourceValue() CompareADObjectsResultADAttributeValue`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *CompareADObjectsResultADAttribute) GetSourceValueOk() (*CompareADObjectsResultADAttributeValue, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *CompareADObjectsResultADAttribute) SetSourceValue(v CompareADObjectsResultADAttributeValue)`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *CompareADObjectsResultADAttribute) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.

### GetStatus

`func (o *CompareADObjectsResultADAttribute) GetStatus() ErrorProto`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompareADObjectsResultADAttribute) GetStatusOk() (*ErrorProto, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompareADObjectsResultADAttribute) SetStatus(v ErrorProto)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompareADObjectsResultADAttribute) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


