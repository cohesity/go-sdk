# ArchivalTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the archival target. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of the archival target. | [optional] 
**VaultId** | Pointer to **NullableInt64** | The id of the archival vault. | [optional] 

## Methods

### NewArchivalTarget

`func NewArchivalTarget() *ArchivalTarget`

NewArchivalTarget instantiates a new ArchivalTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalTargetWithDefaults

`func NewArchivalTargetWithDefaults() *ArchivalTarget`

NewArchivalTargetWithDefaults instantiates a new ArchivalTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArchivalTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArchivalTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArchivalTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArchivalTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ArchivalTarget) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ArchivalTarget) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *ArchivalTarget) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArchivalTarget) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArchivalTarget) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ArchivalTarget) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ArchivalTarget) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ArchivalTarget) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVaultId

`func (o *ArchivalTarget) GetVaultId() int64`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *ArchivalTarget) GetVaultIdOk() (*int64, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *ArchivalTarget) SetVaultId(v int64)`

SetVaultId sets VaultId field to given value.

### HasVaultId

`func (o *ArchivalTarget) HasVaultId() bool`

HasVaultId returns a boolean if a field has been set.

### SetVaultIdNil

`func (o *ArchivalTarget) SetVaultIdNil(b bool)`

 SetVaultIdNil sets the value for VaultId to be an explicit nil

### UnsetVaultId
`func (o *ArchivalTarget) UnsetVaultId()`

UnsetVaultId ensures that no value is present for VaultId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


