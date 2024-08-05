# IsilonRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackUpSMBVolumes** | Pointer to **NullableBool** | Specifies whether or not to back up SMB Volumes. | [optional] 
**Credentials** | [**Credentials**](Credentials.md) |  | 
**Endpoint** | **NullableString** | Specifies the IP Address Endpoint for the Isilon Source. | 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**SmbCredentials** | Pointer to [**SmbMountCredentials**](SmbMountCredentials.md) |  | [optional] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 

## Methods

### NewIsilonRegistrationParams

`func NewIsilonRegistrationParams(credentials Credentials, endpoint NullableString, ) *IsilonRegistrationParams`

NewIsilonRegistrationParams instantiates a new IsilonRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonRegistrationParamsWithDefaults

`func NewIsilonRegistrationParamsWithDefaults() *IsilonRegistrationParams`

NewIsilonRegistrationParamsWithDefaults instantiates a new IsilonRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackUpSMBVolumes

`func (o *IsilonRegistrationParams) GetBackUpSMBVolumes() bool`

GetBackUpSMBVolumes returns the BackUpSMBVolumes field if non-nil, zero value otherwise.

### GetBackUpSMBVolumesOk

`func (o *IsilonRegistrationParams) GetBackUpSMBVolumesOk() (*bool, bool)`

GetBackUpSMBVolumesOk returns a tuple with the BackUpSMBVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpSMBVolumes

`func (o *IsilonRegistrationParams) SetBackUpSMBVolumes(v bool)`

SetBackUpSMBVolumes sets BackUpSMBVolumes field to given value.

### HasBackUpSMBVolumes

`func (o *IsilonRegistrationParams) HasBackUpSMBVolumes() bool`

HasBackUpSMBVolumes returns a boolean if a field has been set.

### SetBackUpSMBVolumesNil

`func (o *IsilonRegistrationParams) SetBackUpSMBVolumesNil(b bool)`

 SetBackUpSMBVolumesNil sets the value for BackUpSMBVolumes to be an explicit nil

### UnsetBackUpSMBVolumes
`func (o *IsilonRegistrationParams) UnsetBackUpSMBVolumes()`

UnsetBackUpSMBVolumes ensures that no value is present for BackUpSMBVolumes, not even an explicit nil
### GetCredentials

`func (o *IsilonRegistrationParams) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *IsilonRegistrationParams) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *IsilonRegistrationParams) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.


### GetEndpoint

`func (o *IsilonRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IsilonRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IsilonRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *IsilonRegistrationParams) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *IsilonRegistrationParams) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetFilterIpConfig

`func (o *IsilonRegistrationParams) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *IsilonRegistrationParams) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *IsilonRegistrationParams) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *IsilonRegistrationParams) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetSmbCredentials

`func (o *IsilonRegistrationParams) GetSmbCredentials() SmbMountCredentials`

GetSmbCredentials returns the SmbCredentials field if non-nil, zero value otherwise.

### GetSmbCredentialsOk

`func (o *IsilonRegistrationParams) GetSmbCredentialsOk() (*SmbMountCredentials, bool)`

GetSmbCredentialsOk returns a tuple with the SmbCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbCredentials

`func (o *IsilonRegistrationParams) SetSmbCredentials(v SmbMountCredentials)`

SetSmbCredentials sets SmbCredentials field to given value.

### HasSmbCredentials

`func (o *IsilonRegistrationParams) HasSmbCredentials() bool`

HasSmbCredentials returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *IsilonRegistrationParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *IsilonRegistrationParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *IsilonRegistrationParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *IsilonRegistrationParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


