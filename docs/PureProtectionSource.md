# PureProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies a unique name of the Protection Source | [optional] 
**StorageArray** | Pointer to [**SanStorageArray**](SanStorageArray.md) |  | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of managed Object in a SAN/Pure Protection Source like a kStorageArray or kVolume. Examples of SAN Objects include &#39;kStorageArray&#39; and &#39;kVolume&#39;. &#39;kStorageArray&#39; indicates that entire SAN storage array is being protected. &#39;kVolume&#39; indicates that volume within the array is being protected. | [optional] 
**Volume** | Pointer to [**SanVolume**](SanVolume.md) |  | [optional] 

## Methods

### NewPureProtectionSource

`func NewPureProtectionSource() *PureProtectionSource`

NewPureProtectionSource instantiates a new PureProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPureProtectionSourceWithDefaults

`func NewPureProtectionSourceWithDefaults() *PureProtectionSource`

NewPureProtectionSourceWithDefaults instantiates a new PureProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PureProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PureProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PureProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PureProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PureProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PureProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStorageArray

`func (o *PureProtectionSource) GetStorageArray() SanStorageArray`

GetStorageArray returns the StorageArray field if non-nil, zero value otherwise.

### GetStorageArrayOk

`func (o *PureProtectionSource) GetStorageArrayOk() (*SanStorageArray, bool)`

GetStorageArrayOk returns a tuple with the StorageArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageArray

`func (o *PureProtectionSource) SetStorageArray(v SanStorageArray)`

SetStorageArray sets StorageArray field to given value.

### HasStorageArray

`func (o *PureProtectionSource) HasStorageArray() bool`

HasStorageArray returns a boolean if a field has been set.

### GetType

`func (o *PureProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PureProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PureProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PureProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PureProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PureProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVolume

`func (o *PureProtectionSource) GetVolume() SanVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *PureProtectionSource) GetVolumeOk() (*SanVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *PureProtectionSource) SetVolume(v SanVolume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *PureProtectionSource) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


