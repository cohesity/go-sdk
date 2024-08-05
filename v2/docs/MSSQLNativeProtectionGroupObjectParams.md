# MSSQLNativeProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object being protected. If this is a non leaf level object, then the object will be auto-protected unless leaf objects are specified for exclusion. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object being protected. | [optional] [readonly] 
**SourceType** | Pointer to **NullableString** | Specifies the type of source being protected. | [optional] [readonly] 

## Methods

### NewMSSQLNativeProtectionGroupObjectParams

`func NewMSSQLNativeProtectionGroupObjectParams(id NullableInt64, ) *MSSQLNativeProtectionGroupObjectParams`

NewMSSQLNativeProtectionGroupObjectParams instantiates a new MSSQLNativeProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLNativeProtectionGroupObjectParamsWithDefaults

`func NewMSSQLNativeProtectionGroupObjectParamsWithDefaults() *MSSQLNativeProtectionGroupObjectParams`

NewMSSQLNativeProtectionGroupObjectParamsWithDefaults instantiates a new MSSQLNativeProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MSSQLNativeProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MSSQLNativeProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MSSQLNativeProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *MSSQLNativeProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MSSQLNativeProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *MSSQLNativeProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MSSQLNativeProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MSSQLNativeProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MSSQLNativeProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MSSQLNativeProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MSSQLNativeProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceType

`func (o *MSSQLNativeProtectionGroupObjectParams) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MSSQLNativeProtectionGroupObjectParams) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MSSQLNativeProtectionGroupObjectParams) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *MSSQLNativeProtectionGroupObjectParams) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *MSSQLNativeProtectionGroupObjectParams) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *MSSQLNativeProtectionGroupObjectParams) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


