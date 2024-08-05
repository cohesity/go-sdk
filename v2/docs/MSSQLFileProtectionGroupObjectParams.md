# MSSQLFileProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object being protected. If this is a non leaf level object, then the object will be auto-protected unless leaf objects are specified for exclusion. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object being protected. | [optional] [readonly] 
**SourceType** | Pointer to **NullableString** | Specifies the type of source being protected. | [optional] [readonly] 

## Methods

### NewMSSQLFileProtectionGroupObjectParams

`func NewMSSQLFileProtectionGroupObjectParams(id NullableInt64, ) *MSSQLFileProtectionGroupObjectParams`

NewMSSQLFileProtectionGroupObjectParams instantiates a new MSSQLFileProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLFileProtectionGroupObjectParamsWithDefaults

`func NewMSSQLFileProtectionGroupObjectParamsWithDefaults() *MSSQLFileProtectionGroupObjectParams`

NewMSSQLFileProtectionGroupObjectParamsWithDefaults instantiates a new MSSQLFileProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MSSQLFileProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MSSQLFileProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MSSQLFileProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *MSSQLFileProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MSSQLFileProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *MSSQLFileProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MSSQLFileProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MSSQLFileProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MSSQLFileProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MSSQLFileProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MSSQLFileProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceType

`func (o *MSSQLFileProtectionGroupObjectParams) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MSSQLFileProtectionGroupObjectParams) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MSSQLFileProtectionGroupObjectParams) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *MSSQLFileProtectionGroupObjectParams) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *MSSQLFileProtectionGroupObjectParams) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *MSSQLFileProtectionGroupObjectParams) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


