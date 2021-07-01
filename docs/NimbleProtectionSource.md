# NimbleProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies a unique name of the Protection Source | [optional] 
**StorageArray** | Pointer to [**SanStorageArray**](SanStorageArray.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of managed Object in a SAN/Nimble Protection Source like a kStorageArray or kVolume. Examples of SAN Objects include &#39;kStorageArray&#39; and &#39;kVolume&#39;. &#39;kStorageArray&#39; indicates that entire SAN storage array is being protected. &#39;kVolume&#39; indicates that volume within the array is being protected. | [optional] 
**Volume** | Pointer to [**SanVolume**](SanVolume.md) |  | [optional] 

## Methods

### NewNimbleProtectionSource

`func NewNimbleProtectionSource() *NimbleProtectionSource`

NewNimbleProtectionSource instantiates a new NimbleProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNimbleProtectionSourceWithDefaults

`func NewNimbleProtectionSourceWithDefaults() *NimbleProtectionSource`

NewNimbleProtectionSourceWithDefaults instantiates a new NimbleProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NimbleProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NimbleProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NimbleProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NimbleProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NimbleProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NimbleProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStorageArray

`func (o *NimbleProtectionSource) GetStorageArray() SanStorageArray`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *NimbleProtectionSource) GetStorageArrayOk() (*SanStorageArray, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *NimbleProtectionSource) SetStorageArray(v SanStorageArray)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *NimbleProtectionSource) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.

### GetType

`func (o *NimbleProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NimbleProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NimbleProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NimbleProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NimbleProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NimbleProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVolume

`func (o *NimbleProtectionSource) GetVolume() SanVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *NimbleProtectionSource) GetVolumeOk() (*SanVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *NimbleProtectionSource) SetVolume(v SanVolume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *NimbleProtectionSource) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


