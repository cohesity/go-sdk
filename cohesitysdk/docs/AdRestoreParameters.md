# AdRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdOptions** | Pointer to [**AdRestoreOptions**](AdRestoreOptions.md) |  | [optional] 
**Credentials** | Pointer to [**Credentials**](Credentials.md) |  | [optional] 
**MountAndRestore** | Pointer to **NullableBool** | Specifies the option to mount the AD snapshot database and restore the AD objects in a single restore task. AdOptions must be set if this is set to true. | [optional] 
**Port** | Pointer to **NullableInt32** | Specifies the port on which the AD domain controller&#39;s NTDS database will be mounted. | [optional] 

## Methods

### NewAdRestoreParameters

`func NewAdRestoreParameters() *AdRestoreParameters`

NewAdRestoreParameters instantiates a new AdRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdRestoreParametersWithDefaults

`func NewAdRestoreParametersWithDefaults() *AdRestoreParameters`

NewAdRestoreParametersWithDefaults instantiates a new AdRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdOptions

`func (o *AdRestoreParameters) GetAdOptions() AdRestoreOptions`

GetAdOptions returns the AdOptions field if non-nil, zero value otherwise.

### GetAdOptionsOk

`func (o *AdRestoreParameters) GetAdOptionsOk() (*AdRestoreOptions, bool)`

GetAdOptionsOk returns a tuple with the AdOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdOptions

`func (o *AdRestoreParameters) SetAdOptions(v AdRestoreOptions)`

SetAdOptions sets AdOptions field to given value.

### HasAdOptions

`func (o *AdRestoreParameters) HasAdOptions() bool`

HasAdOptions returns a boolean if a field has been set.

### GetCredentials

`func (o *AdRestoreParameters) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AdRestoreParameters) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AdRestoreParameters) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AdRestoreParameters) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetMountAndRestore

`func (o *AdRestoreParameters) GetMountAndRestore() bool`

GetMountAndRestore returns the MountAndRestore field if non-nil, zero value otherwise.

### GetMountAndRestoreOk

`func (o *AdRestoreParameters) GetMountAndRestoreOk() (*bool, bool)`

GetMountAndRestoreOk returns a tuple with the MountAndRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountAndRestore

`func (o *AdRestoreParameters) SetMountAndRestore(v bool)`

SetMountAndRestore sets MountAndRestore field to given value.

### HasMountAndRestore

`func (o *AdRestoreParameters) HasMountAndRestore() bool`

HasMountAndRestore returns a boolean if a field has been set.

### SetMountAndRestoreNil

`func (o *AdRestoreParameters) SetMountAndRestoreNil(b bool)`

 SetMountAndRestoreNil sets the value for MountAndRestore to be an explicit nil

### UnsetMountAndRestore
`func (o *AdRestoreParameters) UnsetMountAndRestore()`

UnsetMountAndRestore ensures that no value is present for MountAndRestore, not even an explicit nil
### GetPort

`func (o *AdRestoreParameters) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AdRestoreParameters) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AdRestoreParameters) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AdRestoreParameters) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *AdRestoreParameters) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *AdRestoreParameters) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


