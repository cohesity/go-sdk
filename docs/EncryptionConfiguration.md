# EncryptionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableEncryption** | Pointer to **NullableBool** | Specifies whether or not to enable encryption. If encryption is enabled, all data on the Cluster will be encrypted. This can only be enabled at Cluster creation time and cannot be disabled later. | [optional] 
**EnableFipsMode** | Pointer to **NullableBool** | Specifies whether or not to enable FIPS mode. EnableEncryption must be set to true in order to enable FIPS. | [optional] 
**RotationPeriod** | Pointer to **NullableInt32** | Specifies the rotation period for encryption keys in days. | [optional] 

## Methods

### NewEncryptionConfiguration

`func NewEncryptionConfiguration() *EncryptionConfiguration`

NewEncryptionConfiguration instantiates a new EncryptionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionConfigurationWithDefaults

`func NewEncryptionConfigurationWithDefaults() *EncryptionConfiguration`

NewEncryptionConfigurationWithDefaults instantiates a new EncryptionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableEncryption

`func (o *EncryptionConfiguration) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *EncryptionConfiguration) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *EncryptionConfiguration) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.

### HasEnableEncryption

`func (o *EncryptionConfiguration) HasEnableEncryption() bool`

HasEnableEncryption returns a boolean if a field has been set.

### SetEnableEncryptionNil

`func (o *EncryptionConfiguration) SetEnableEncryptionNil(b bool)`

 SetEnableEncryptionNil sets the value for EnableEncryption to be an explicit nil

### UnsetEnableEncryption
`func (o *EncryptionConfiguration) UnsetEnableEncryption()`

UnsetEnableEncryption ensures that no value is present for EnableEncryption, not even an explicit nil
### GetEnableFipsMode

`func (o *EncryptionConfiguration) GetEnableFipsMode() bool`

GetEnableFipsMode returns the EnableFipsMode field if non-nil, zero value otherwise.

### GetEnableFipsModeOk

`func (o *EncryptionConfiguration) GetEnableFipsModeOk() (*bool, bool)`

GetEnableFipsModeOk returns a tuple with the EnableFipsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFipsMode

`func (o *EncryptionConfiguration) SetEnableFipsMode(v bool)`

SetEnableFipsMode sets EnableFipsMode field to given value.

### HasEnableFipsMode

`func (o *EncryptionConfiguration) HasEnableFipsMode() bool`

HasEnableFipsMode returns a boolean if a field has been set.

### SetEnableFipsModeNil

`func (o *EncryptionConfiguration) SetEnableFipsModeNil(b bool)`

 SetEnableFipsModeNil sets the value for EnableFipsMode to be an explicit nil

### UnsetEnableFipsMode
`func (o *EncryptionConfiguration) UnsetEnableFipsMode()`

UnsetEnableFipsMode ensures that no value is present for EnableFipsMode, not even an explicit nil
### GetRotationPeriod

`func (o *EncryptionConfiguration) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *EncryptionConfiguration) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *EncryptionConfiguration) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *EncryptionConfiguration) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### SetRotationPeriodNil

`func (o *EncryptionConfiguration) SetRotationPeriodNil(b bool)`

 SetRotationPeriodNil sets the value for RotationPeriod to be an explicit nil

### UnsetRotationPeriod
`func (o *EncryptionConfiguration) UnsetRotationPeriod()`

UnsetRotationPeriod ensures that no value is present for RotationPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


