# VMwareObjectId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MorItem** | Pointer to **NullableString** | Specifies the Managed Object Reference Item. | [optional] 
**MorType** | Pointer to **NullableString** | Specifies the Managed Object Reference Type. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies a Universally Unique Identifier (UUID) of a VMware Object. | [optional] 

## Methods

### NewVMwareObjectId

`func NewVMwareObjectId() *VMwareObjectId`

NewVMwareObjectId instantiates a new VMwareObjectId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMwareObjectIdWithDefaults

`func NewVMwareObjectIdWithDefaults() *VMwareObjectId`

NewVMwareObjectIdWithDefaults instantiates a new VMwareObjectId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMorItem

`func (o *VMwareObjectId) GetMorItem() string`

GetMorItem returns the MorItem field if non-nil, zero value otherwise.

### GetMorItemOk

`func (o *VMwareObjectId) GetMorItemOk() (*string, bool)`

GetMorItemOk returns a tuple with the MorItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorItem

`func (o *VMwareObjectId) SetMorItem(v string)`

SetMorItem sets MorItem field to given value.

### HasMorItem

`func (o *VMwareObjectId) HasMorItem() bool`

HasMorItem returns a boolean if a field has been set.

### SetMorItemNil

`func (o *VMwareObjectId) SetMorItemNil(b bool)`

 SetMorItemNil sets the value for MorItem to be an explicit nil

### UnsetMorItem
`func (o *VMwareObjectId) UnsetMorItem()`

UnsetMorItem ensures that no value is present for MorItem, not even an explicit nil
### GetMorType

`func (o *VMwareObjectId) GetMorType() string`

GetMorType returns the MorType field if non-nil, zero value otherwise.

### GetMorTypeOk

`func (o *VMwareObjectId) GetMorTypeOk() (*string, bool)`

GetMorTypeOk returns a tuple with the MorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorType

`func (o *VMwareObjectId) SetMorType(v string)`

SetMorType sets MorType field to given value.

### HasMorType

`func (o *VMwareObjectId) HasMorType() bool`

HasMorType returns a boolean if a field has been set.

### SetMorTypeNil

`func (o *VMwareObjectId) SetMorTypeNil(b bool)`

 SetMorTypeNil sets the value for MorType to be an explicit nil

### UnsetMorType
`func (o *VMwareObjectId) UnsetMorType()`

UnsetMorType ensures that no value is present for MorType, not even an explicit nil
### GetUuid

`func (o *VMwareObjectId) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VMwareObjectId) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VMwareObjectId) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VMwareObjectId) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *VMwareObjectId) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *VMwareObjectId) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


