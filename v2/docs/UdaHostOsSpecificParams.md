# UdaHostOsSpecificParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostOsType** | Pointer to **NullableString** | Type of the host operating system. | [optional] 
**Index** | Pointer to [**NullableUdaOSIndexConfigParams**](UdaOSIndexConfigParams.md) |  | [optional] 
**Protection** | Pointer to [**NullableUdaProtectionParams**](UdaProtectionParams.md) |  | [optional] 
**Recovery** | Pointer to [**NullableUdaRecoveryParams**](UdaRecoveryParams.md) |  | [optional] 
**Registration** | Pointer to [**NullableUdaRegistrationParams**](UdaRegistrationParams.md) |  | [optional] 

## Methods

### NewUdaHostOsSpecificParams

`func NewUdaHostOsSpecificParams() *UdaHostOsSpecificParams`

NewUdaHostOsSpecificParams instantiates a new UdaHostOsSpecificParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaHostOsSpecificParamsWithDefaults

`func NewUdaHostOsSpecificParamsWithDefaults() *UdaHostOsSpecificParams`

NewUdaHostOsSpecificParamsWithDefaults instantiates a new UdaHostOsSpecificParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostOsType

`func (o *UdaHostOsSpecificParams) GetHostOsType() string`

GetHostOsType returns the HostOsType field if non-nil, zero value otherwise.

### GetHostOsTypeOk

`func (o *UdaHostOsSpecificParams) GetHostOsTypeOk() (*string, bool)`

GetHostOsTypeOk returns a tuple with the HostOsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOsType

`func (o *UdaHostOsSpecificParams) SetHostOsType(v string)`

SetHostOsType sets HostOsType field to given value.

### HasHostOsType

`func (o *UdaHostOsSpecificParams) HasHostOsType() bool`

HasHostOsType returns a boolean if a field has been set.

### SetHostOsTypeNil

`func (o *UdaHostOsSpecificParams) SetHostOsTypeNil(b bool)`

 SetHostOsTypeNil sets the value for HostOsType to be an explicit nil

### UnsetHostOsType
`func (o *UdaHostOsSpecificParams) UnsetHostOsType()`

UnsetHostOsType ensures that no value is present for HostOsType, not even an explicit nil
### GetIndex

`func (o *UdaHostOsSpecificParams) GetIndex() UdaOSIndexConfigParams`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *UdaHostOsSpecificParams) GetIndexOk() (*UdaOSIndexConfigParams, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *UdaHostOsSpecificParams) SetIndex(v UdaOSIndexConfigParams)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *UdaHostOsSpecificParams) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *UdaHostOsSpecificParams) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *UdaHostOsSpecificParams) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetProtection

`func (o *UdaHostOsSpecificParams) GetProtection() UdaProtectionParams`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *UdaHostOsSpecificParams) GetProtectionOk() (*UdaProtectionParams, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *UdaHostOsSpecificParams) SetProtection(v UdaProtectionParams)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *UdaHostOsSpecificParams) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### SetProtectionNil

`func (o *UdaHostOsSpecificParams) SetProtectionNil(b bool)`

 SetProtectionNil sets the value for Protection to be an explicit nil

### UnsetProtection
`func (o *UdaHostOsSpecificParams) UnsetProtection()`

UnsetProtection ensures that no value is present for Protection, not even an explicit nil
### GetRecovery

`func (o *UdaHostOsSpecificParams) GetRecovery() UdaRecoveryParams`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *UdaHostOsSpecificParams) GetRecoveryOk() (*UdaRecoveryParams, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *UdaHostOsSpecificParams) SetRecovery(v UdaRecoveryParams)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *UdaHostOsSpecificParams) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.

### SetRecoveryNil

`func (o *UdaHostOsSpecificParams) SetRecoveryNil(b bool)`

 SetRecoveryNil sets the value for Recovery to be an explicit nil

### UnsetRecovery
`func (o *UdaHostOsSpecificParams) UnsetRecovery()`

UnsetRecovery ensures that no value is present for Recovery, not even an explicit nil
### GetRegistration

`func (o *UdaHostOsSpecificParams) GetRegistration() UdaRegistrationParams`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *UdaHostOsSpecificParams) GetRegistrationOk() (*UdaRegistrationParams, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *UdaHostOsSpecificParams) SetRegistration(v UdaRegistrationParams)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *UdaHostOsSpecificParams) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### SetRegistrationNil

`func (o *UdaHostOsSpecificParams) SetRegistrationNil(b bool)`

 SetRegistrationNil sets the value for Registration to be an explicit nil

### UnsetRegistration
`func (o *UdaHostOsSpecificParams) UnsetRegistration()`

UnsetRegistration ensures that no value is present for Registration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


