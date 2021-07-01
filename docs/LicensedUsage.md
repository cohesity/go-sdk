# LicensedUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityGiB** | Pointer to **NullableInt64** | Feature usage by the cluster. | [optional] 
**ExpiryTime** | Pointer to **NullableInt64** | Expiry time(epoch) of each feature. There could be multiple expiry time for the given SKU. | [optional] 
**FeatureName** | Pointer to **NullableString** | Name of feature. | [optional] 
**LicenseType** | Pointer to **NullableString** | Type of License | [optional] 
**NumVm** | Pointer to **NullableInt64** | Number of VM spinned. | [optional] 

## Methods

### NewLicensedUsage

`func NewLicensedUsage() *LicensedUsage`

NewLicensedUsage instantiates a new LicensedUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensedUsageWithDefaults

`func NewLicensedUsageWithDefaults() *LicensedUsage`

NewLicensedUsageWithDefaults instantiates a new LicensedUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityGiB

`func (o *LicensedUsage) GetCapacityGiB() int64`

GetCapacityGiB returns the CapacityGiB field if non-nil, zero value otherwise.

### GetCapacityGiBOk

`func (o *LicensedUsage) GetCapacityGiBOk() (*int64, bool)`

GetCapacityGiBOk returns a tuple with the CapacityGiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityGiB

`func (o *LicensedUsage) SetCapacityGiB(v int64)`

SetCapacityGiB sets CapacityGiB field to given value.

### HasCapacityGiB

`func (o *LicensedUsage) HasCapacityGiB() bool`

HasCapacityGiB returns a boolean if a field has been set.

### SetCapacityGiBNil

`func (o *LicensedUsage) SetCapacityGiBNil(b bool)`

 SetCapacityGiBNil sets the value for CapacityGiB to be an explicit nil

### UnsetCapacityGiB
`func (o *LicensedUsage) UnsetCapacityGiB()`

UnsetCapacityGiB ensures that no value is present for CapacityGiB, not even an explicit nil
### GetExpiryTime

`func (o *LicensedUsage) GetExpiryTime() int64`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *LicensedUsage) GetExpiryTimeOk() (*int64, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *LicensedUsage) SetExpiryTime(v int64)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *LicensedUsage) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### SetExpiryTimeNil

`func (o *LicensedUsage) SetExpiryTimeNil(b bool)`

 SetExpiryTimeNil sets the value for ExpiryTime to be an explicit nil

### UnsetExpiryTime
`func (o *LicensedUsage) UnsetExpiryTime()`

UnsetExpiryTime ensures that no value is present for ExpiryTime, not even an explicit nil
### GetFeatureName

`func (o *LicensedUsage) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *LicensedUsage) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *LicensedUsage) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *LicensedUsage) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### SetFeatureNameNil

`func (o *LicensedUsage) SetFeatureNameNil(b bool)`

 SetFeatureNameNil sets the value for FeatureName to be an explicit nil

### UnsetFeatureName
`func (o *LicensedUsage) UnsetFeatureName()`

UnsetFeatureName ensures that no value is present for FeatureName, not even an explicit nil
### GetLicenseType

`func (o *LicensedUsage) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicensedUsage) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicensedUsage) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *LicensedUsage) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### SetLicenseTypeNil

`func (o *LicensedUsage) SetLicenseTypeNil(b bool)`

 SetLicenseTypeNil sets the value for LicenseType to be an explicit nil

### UnsetLicenseType
`func (o *LicensedUsage) UnsetLicenseType()`

UnsetLicenseType ensures that no value is present for LicenseType, not even an explicit nil
### GetNumVm

`func (o *LicensedUsage) GetNumVm() int64`

GetNumVm returns the NumVm field if non-nil, zero value otherwise.

### GetNumVmOk

`func (o *LicensedUsage) GetNumVmOk() (*int64, bool)`

GetNumVmOk returns a tuple with the NumVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVm

`func (o *LicensedUsage) SetNumVm(v int64)`

SetNumVm sets NumVm field to given value.

### HasNumVm

`func (o *LicensedUsage) HasNumVm() bool`

HasNumVm returns a boolean if a field has been set.

### SetNumVmNil

`func (o *LicensedUsage) SetNumVmNil(b bool)`

 SetNumVmNil sets the value for NumVm to be an explicit nil

### UnsetNumVm
`func (o *LicensedUsage) UnsetNumVm()`

UnsetNumVm ensures that no value is present for NumVm, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


