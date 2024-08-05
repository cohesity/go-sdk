# HeliosTierLevelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsTiering** | Pointer to [**HeliosAWSTiers**](HeliosAWSTiers.md) |  | [optional] 
**AzureTiering** | Pointer to [**HeliosAzureTiers**](HeliosAzureTiers.md) |  | [optional] 
**CloudPlatform** | **NullableString** | Specifies the cloud platform to enable tiering. | 
**GoogleTiering** | Pointer to [**HeliosGoogleTiers**](HeliosGoogleTiers.md) |  | [optional] 
**OracleTiering** | Pointer to [**HeliosOracleTiers**](HeliosOracleTiers.md) |  | [optional] 

## Methods

### NewHeliosTierLevelSettings

`func NewHeliosTierLevelSettings(cloudPlatform NullableString, ) *HeliosTierLevelSettings`

NewHeliosTierLevelSettings instantiates a new HeliosTierLevelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTierLevelSettingsWithDefaults

`func NewHeliosTierLevelSettingsWithDefaults() *HeliosTierLevelSettings`

NewHeliosTierLevelSettingsWithDefaults instantiates a new HeliosTierLevelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsTiering

`func (o *HeliosTierLevelSettings) GetAwsTiering() HeliosAWSTiers`

GetAwsTiering returns the AwsTiering field if non-nil, zero value otherwise.

### GetAwsTieringOk

`func (o *HeliosTierLevelSettings) GetAwsTieringOk() (*HeliosAWSTiers, bool)`

GetAwsTieringOk returns a tuple with the AwsTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTiering

`func (o *HeliosTierLevelSettings) SetAwsTiering(v HeliosAWSTiers)`

SetAwsTiering sets AwsTiering field to given value.

### HasAwsTiering

`func (o *HeliosTierLevelSettings) HasAwsTiering() bool`

HasAwsTiering returns a boolean if a field has been set.

### GetAzureTiering

`func (o *HeliosTierLevelSettings) GetAzureTiering() HeliosAzureTiers`

GetAzureTiering returns the AzureTiering field if non-nil, zero value otherwise.

### GetAzureTieringOk

`func (o *HeliosTierLevelSettings) GetAzureTieringOk() (*HeliosAzureTiers, bool)`

GetAzureTieringOk returns a tuple with the AzureTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTiering

`func (o *HeliosTierLevelSettings) SetAzureTiering(v HeliosAzureTiers)`

SetAzureTiering sets AzureTiering field to given value.

### HasAzureTiering

`func (o *HeliosTierLevelSettings) HasAzureTiering() bool`

HasAzureTiering returns a boolean if a field has been set.

### GetCloudPlatform

`func (o *HeliosTierLevelSettings) GetCloudPlatform() string`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *HeliosTierLevelSettings) GetCloudPlatformOk() (*string, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *HeliosTierLevelSettings) SetCloudPlatform(v string)`

SetCloudPlatform sets CloudPlatform field to given value.


### SetCloudPlatformNil

`func (o *HeliosTierLevelSettings) SetCloudPlatformNil(b bool)`

 SetCloudPlatformNil sets the value for CloudPlatform to be an explicit nil

### UnsetCloudPlatform
`func (o *HeliosTierLevelSettings) UnsetCloudPlatform()`

UnsetCloudPlatform ensures that no value is present for CloudPlatform, not even an explicit nil
### GetGoogleTiering

`func (o *HeliosTierLevelSettings) GetGoogleTiering() HeliosGoogleTiers`

GetGoogleTiering returns the GoogleTiering field if non-nil, zero value otherwise.

### GetGoogleTieringOk

`func (o *HeliosTierLevelSettings) GetGoogleTieringOk() (*HeliosGoogleTiers, bool)`

GetGoogleTieringOk returns a tuple with the GoogleTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTiering

`func (o *HeliosTierLevelSettings) SetGoogleTiering(v HeliosGoogleTiers)`

SetGoogleTiering sets GoogleTiering field to given value.

### HasGoogleTiering

`func (o *HeliosTierLevelSettings) HasGoogleTiering() bool`

HasGoogleTiering returns a boolean if a field has been set.

### GetOracleTiering

`func (o *HeliosTierLevelSettings) GetOracleTiering() HeliosOracleTiers`

GetOracleTiering returns the OracleTiering field if non-nil, zero value otherwise.

### GetOracleTieringOk

`func (o *HeliosTierLevelSettings) GetOracleTieringOk() (*HeliosOracleTiers, bool)`

GetOracleTieringOk returns a tuple with the OracleTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTiering

`func (o *HeliosTierLevelSettings) SetOracleTiering(v HeliosOracleTiers)`

SetOracleTiering sets OracleTiering field to given value.

### HasOracleTiering

`func (o *HeliosTierLevelSettings) HasOracleTiering() bool`

HasOracleTiering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


