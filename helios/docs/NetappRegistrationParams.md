# NetappRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **NullableString** | Specifies the Netapp source type. Can be either kCluster or kVServer (SVM). | 
**Endpoint** | **NullableString** | Specifies the Hostname or IP Address Endpoint for the Netapp Source. | 
**Credentials** | [**Credentials**](Credentials.md) |  | 
**BackUpSMBVolumes** | Pointer to **NullableBool** | Specifies whether or not to back up SMB Volumes. | [optional] 
**SmbCredentials** | Pointer to [**SmbMountCredentials**](SmbMountCredentials.md) |  | [optional] 
**FilterIpConfig** | Pointer to [**FilterIpConfig**](FilterIpConfig.md) |  | [optional] 
**ThrottlingConfig** | Pointer to [**NasThrottlingConfig**](NasThrottlingConfig.md) |  | [optional] 

## Methods

### NewNetappRegistrationParams

`func NewNetappRegistrationParams(sourceType NullableString, endpoint NullableString, credentials Credentials, ) *NetappRegistrationParams`

NewNetappRegistrationParams instantiates a new NetappRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappRegistrationParamsWithDefaults

`func NewNetappRegistrationParamsWithDefaults() *NetappRegistrationParams`

NewNetappRegistrationParamsWithDefaults instantiates a new NetappRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *NetappRegistrationParams) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *NetappRegistrationParams) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *NetappRegistrationParams) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### SetSourceTypeNil

`func (o *NetappRegistrationParams) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *NetappRegistrationParams) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetEndpoint

`func (o *NetappRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *NetappRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *NetappRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *NetappRegistrationParams) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *NetappRegistrationParams) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetCredentials

`func (o *NetappRegistrationParams) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *NetappRegistrationParams) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *NetappRegistrationParams) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.


### GetBackUpSMBVolumes

`func (o *NetappRegistrationParams) GetBackUpSMBVolumes() bool`

GetBackUpSMBVolumes returns the BackUpSMBVolumes field if non-nil, zero value otherwise.

### GetBackUpSMBVolumesOk

`func (o *NetappRegistrationParams) GetBackUpSMBVolumesOk() (*bool, bool)`

GetBackUpSMBVolumesOk returns a tuple with the BackUpSMBVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackUpSMBVolumes

`func (o *NetappRegistrationParams) SetBackUpSMBVolumes(v bool)`

SetBackUpSMBVolumes sets BackUpSMBVolumes field to given value.

### HasBackUpSMBVolumes

`func (o *NetappRegistrationParams) HasBackUpSMBVolumes() bool`

HasBackUpSMBVolumes returns a boolean if a field has been set.

### SetBackUpSMBVolumesNil

`func (o *NetappRegistrationParams) SetBackUpSMBVolumesNil(b bool)`

 SetBackUpSMBVolumesNil sets the value for BackUpSMBVolumes to be an explicit nil

### UnsetBackUpSMBVolumes
`func (o *NetappRegistrationParams) UnsetBackUpSMBVolumes()`

UnsetBackUpSMBVolumes ensures that no value is present for BackUpSMBVolumes, not even an explicit nil
### GetSmbCredentials

`func (o *NetappRegistrationParams) GetSmbCredentials() SmbMountCredentials`

GetSmbCredentials returns the SmbCredentials field if non-nil, zero value otherwise.

### GetSmbCredentialsOk

`func (o *NetappRegistrationParams) GetSmbCredentialsOk() (*SmbMountCredentials, bool)`

GetSmbCredentialsOk returns a tuple with the SmbCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbCredentials

`func (o *NetappRegistrationParams) SetSmbCredentials(v SmbMountCredentials)`

SetSmbCredentials sets SmbCredentials field to given value.

### HasSmbCredentials

`func (o *NetappRegistrationParams) HasSmbCredentials() bool`

HasSmbCredentials returns a boolean if a field has been set.

### GetFilterIpConfig

`func (o *NetappRegistrationParams) GetFilterIpConfig() FilterIpConfig`

GetFilterIpConfig returns the FilterIpConfig field if non-nil, zero value otherwise.

### GetFilterIpConfigOk

`func (o *NetappRegistrationParams) GetFilterIpConfigOk() (*FilterIpConfig, bool)`

GetFilterIpConfigOk returns a tuple with the FilterIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterIpConfig

`func (o *NetappRegistrationParams) SetFilterIpConfig(v FilterIpConfig)`

SetFilterIpConfig sets FilterIpConfig field to given value.

### HasFilterIpConfig

`func (o *NetappRegistrationParams) HasFilterIpConfig() bool`

HasFilterIpConfig returns a boolean if a field has been set.

### GetThrottlingConfig

`func (o *NetappRegistrationParams) GetThrottlingConfig() NasThrottlingConfig`

GetThrottlingConfig returns the ThrottlingConfig field if non-nil, zero value otherwise.

### GetThrottlingConfigOk

`func (o *NetappRegistrationParams) GetThrottlingConfigOk() (*NasThrottlingConfig, bool)`

GetThrottlingConfigOk returns a tuple with the ThrottlingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingConfig

`func (o *NetappRegistrationParams) SetThrottlingConfig(v NasThrottlingConfig)`

SetThrottlingConfig sets ThrottlingConfig field to given value.

### HasThrottlingConfig

`func (o *NetappRegistrationParams) HasThrottlingConfig() bool`

HasThrottlingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


