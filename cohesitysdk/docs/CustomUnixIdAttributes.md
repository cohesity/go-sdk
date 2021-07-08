# CustomUnixIdAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GidAttrName** | Pointer to **NullableString** | Specifies the custom field name in Active Directory user properties to get the GID. | [optional] 
**UidAttrName** | Pointer to **NullableString** | Specifies the custom field name in Active Directory user properties to get the UID. | [optional] 

## Methods

### NewCustomUnixIdAttributes

`func NewCustomUnixIdAttributes() *CustomUnixIdAttributes`

NewCustomUnixIdAttributes instantiates a new CustomUnixIdAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomUnixIdAttributesWithDefaults

`func NewCustomUnixIdAttributesWithDefaults() *CustomUnixIdAttributes`

NewCustomUnixIdAttributesWithDefaults instantiates a new CustomUnixIdAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGidAttrName

`func (o *CustomUnixIdAttributes) GetGidAttrName() string`

GetGidAttrName returns the GidAttrName field if non-nil, zero value otherwise.

### GetGidAttrNameOk

`func (o *CustomUnixIdAttributes) GetGidAttrNameOk() (*string, bool)`

GetGidAttrNameOk returns a tuple with the GidAttrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidAttrName

`func (o *CustomUnixIdAttributes) SetGidAttrName(v string)`

SetGidAttrName sets GidAttrName field to given value.

### HasGidAttrName

`func (o *CustomUnixIdAttributes) HasGidAttrName() bool`

HasGidAttrName returns a boolean if a field has been set.

### SetGidAttrNameNil

`func (o *CustomUnixIdAttributes) SetGidAttrNameNil(b bool)`

 SetGidAttrNameNil sets the value for GidAttrName to be an explicit nil

### UnsetGidAttrName
`func (o *CustomUnixIdAttributes) UnsetGidAttrName()`

UnsetGidAttrName ensures that no value is present for GidAttrName, not even an explicit nil
### GetUidAttrName

`func (o *CustomUnixIdAttributes) GetUidAttrName() string`

GetUidAttrName returns the UidAttrName field if non-nil, zero value otherwise.

### GetUidAttrNameOk

`func (o *CustomUnixIdAttributes) GetUidAttrNameOk() (*string, bool)`

GetUidAttrNameOk returns a tuple with the UidAttrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidAttrName

`func (o *CustomUnixIdAttributes) SetUidAttrName(v string)`

SetUidAttrName sets UidAttrName field to given value.

### HasUidAttrName

`func (o *CustomUnixIdAttributes) HasUidAttrName() bool`

HasUidAttrName returns a boolean if a field has been set.

### SetUidAttrNameNil

`func (o *CustomUnixIdAttributes) SetUidAttrNameNil(b bool)`

 SetUidAttrNameNil sets the value for UidAttrName to be an explicit nil

### UnsetUidAttrName
`func (o *CustomUnixIdAttributes) UnsetUidAttrName()`

UnsetUidAttrName ensures that no value is present for UidAttrName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


