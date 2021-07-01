# EulaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseKey** | **NullableString** | Specifies the license key. | 
**SignedByUser** | Pointer to **NullableString** | Specifies the login account name for the Cohesity user who accepted the End User License Agreement. | [optional] [readonly] 
**SignedTime** | Pointer to **NullableInt64** | Specifies the time that the End User License Agreement was accepted. | [optional] [readonly] 
**SignedVersion** | **NullableInt64** | Specifies the version of the End User License Agreement that was accepted. | 

## Methods

### NewEulaConfig

`func NewEulaConfig(licenseKey NullableString, signedVersion NullableInt64, ) *EulaConfig`

NewEulaConfig instantiates a new EulaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEulaConfigWithDefaults

`func NewEulaConfigWithDefaults() *EulaConfig`

NewEulaConfigWithDefaults instantiates a new EulaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseKey

`func (o *EulaConfig) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *EulaConfig) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *EulaConfig) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.


### SetLicenseKeyNil

`func (o *EulaConfig) SetLicenseKeyNil(b bool)`

 SetLicenseKeyNil sets the value for LicenseKey to be an explicit nil

### UnsetLicenseKey
`func (o *EulaConfig) UnsetLicenseKey()`

UnsetLicenseKey ensures that no value is present for LicenseKey, not even an explicit nil
### GetSignedByUser

`func (o *EulaConfig) GetSignedByUser() string`

GetSignedByUser returns the SignedByUser field if non-nil, zero value otherwise.

### GetSignedByUserOk

`func (o *EulaConfig) GetSignedByUserOk() (*string, bool)`

GetSignedByUserOk returns a tuple with the SignedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedByUser

`func (o *EulaConfig) SetSignedByUser(v string)`

SetSignedByUser sets SignedByUser field to given value.

### HasSignedByUser

`func (o *EulaConfig) HasSignedByUser() bool`

HasSignedByUser returns a boolean if a field has been set.

### SetSignedByUserNil

`func (o *EulaConfig) SetSignedByUserNil(b bool)`

 SetSignedByUserNil sets the value for SignedByUser to be an explicit nil

### UnsetSignedByUser
`func (o *EulaConfig) UnsetSignedByUser()`

UnsetSignedByUser ensures that no value is present for SignedByUser, not even an explicit nil
### GetSignedTime

`func (o *EulaConfig) GetSignedTime() int64`

GetSignedTime returns the SignedTime field if non-nil, zero value otherwise.

### GetSignedTimeOk

`func (o *EulaConfig) GetSignedTimeOk() (*int64, bool)`

GetSignedTimeOk returns a tuple with the SignedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedTime

`func (o *EulaConfig) SetSignedTime(v int64)`

SetSignedTime sets SignedTime field to given value.

### HasSignedTime

`func (o *EulaConfig) HasSignedTime() bool`

HasSignedTime returns a boolean if a field has been set.

### SetSignedTimeNil

`func (o *EulaConfig) SetSignedTimeNil(b bool)`

 SetSignedTimeNil sets the value for SignedTime to be an explicit nil

### UnsetSignedTime
`func (o *EulaConfig) UnsetSignedTime()`

UnsetSignedTime ensures that no value is present for SignedTime, not even an explicit nil
### GetSignedVersion

`func (o *EulaConfig) GetSignedVersion() int64`

GetSignedVersion returns the SignedVersion field if non-nil, zero value otherwise.

### GetSignedVersionOk

`func (o *EulaConfig) GetSignedVersionOk() (*int64, bool)`

GetSignedVersionOk returns a tuple with the SignedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedVersion

`func (o *EulaConfig) SetSignedVersion(v int64)`

SetSignedVersion sets SignedVersion field to given value.


### SetSignedVersionNil

`func (o *EulaConfig) SetSignedVersionNil(b bool)`

 SetSignedVersionNil sets the value for SignedVersion to be an explicit nil

### UnsetSignedVersion
`func (o *EulaConfig) UnsetSignedVersion()`

UnsetSignedVersion ensures that no value is present for SignedVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


