# FlashBladeProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSystem** | Pointer to [**FlashBladeFileSystem**](FlashBladeFileSystem.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies a unique name of the Protection Source. | [optional] 
**StorageArray** | Pointer to [**FlashBladeStorageArray**](FlashBladeStorageArray.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of managed object in a Pure Storage FlashBlade like &#39;kStorageArray&#39; or &#39;kFileSystem&#39;. &#39;kStorageArray&#39; indicates a top level Pure Storage FlashBlade array. &#39;kFileSystem&#39; indicates a Pure Storage FlashBlade file system within the array. | [optional] 

## Methods

### NewFlashBladeProtectionSource

`func NewFlashBladeProtectionSource() *FlashBladeProtectionSource`

NewFlashBladeProtectionSource instantiates a new FlashBladeProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashBladeProtectionSourceWithDefaults

`func NewFlashBladeProtectionSourceWithDefaults() *FlashBladeProtectionSource`

NewFlashBladeProtectionSourceWithDefaults instantiates a new FlashBladeProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSystem

`func (o *FlashBladeProtectionSource) GetFileSystem() FlashBladeFileSystem`

GetFileSystem returns the FileSystem field if non-nil, zero value otherwise.

### GetFileSystemOk

`func (o *FlashBladeProtectionSource) GetFileSystemOk() (*FlashBladeFileSystem, bool)`

GetFileSystemOk returns a tuple with the FileSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystem

`func (o *FlashBladeProtectionSource) SetFileSystem(v FlashBladeFileSystem)`

SetFileSystem sets FileSystem field to given value.

### HasFileSystem

`func (o *FlashBladeProtectionSource) HasFileSystem() bool`

HasFileSystem returns a boolean if a field has been set.

### GetName

`func (o *FlashBladeProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlashBladeProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlashBladeProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlashBladeProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FlashBladeProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FlashBladeProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStorageArray

`func (o *FlashBladeProtectionSource) GetStorageArray() FlashBladeStorageArray`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *FlashBladeProtectionSource) GetStorageArrayOk() (*FlashBladeStorageArray, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *FlashBladeProtectionSource) SetStorageArray(v FlashBladeStorageArray)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *FlashBladeProtectionSource) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.

### GetType

`func (o *FlashBladeProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlashBladeProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlashBladeProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlashBladeProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FlashBladeProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FlashBladeProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


