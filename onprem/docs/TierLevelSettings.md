# TierLevelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudPlatform** | **NullableString** | Specifies the cloud platform to enable tiering. | 
**AwsTiering** | Pointer to [**AWSTiers**](AWSTiers.md) |  | [optional] 
**AzureTiering** | Pointer to [**AzureTiers**](AzureTiers.md) |  | [optional] 
**GoogleTiering** | Pointer to [**GoogleTiers**](GoogleTiers.md) |  | [optional] 
**OracleTiering** | Pointer to [**OracleTiers**](OracleTiers.md) |  | [optional] 

## Methods

### NewTierLevelSettings

`func NewTierLevelSettings(cloudPlatform NullableString, ) *TierLevelSettings`

NewTierLevelSettings instantiates a new TierLevelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierLevelSettingsWithDefaults

`func NewTierLevelSettingsWithDefaults() *TierLevelSettings`

NewTierLevelSettingsWithDefaults instantiates a new TierLevelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudPlatform

`func (o *TierLevelSettings) GetCloudPlatform() string`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *TierLevelSettings) GetCloudPlatformOk() (*string, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *TierLevelSettings) SetCloudPlatform(v string)`

SetCloudPlatform sets CloudPlatform field to given value.


### SetCloudPlatformNil

`func (o *TierLevelSettings) SetCloudPlatformNil(b bool)`

 SetCloudPlatformNil sets the value for CloudPlatform to be an explicit nil

### UnsetCloudPlatform
`func (o *TierLevelSettings) UnsetCloudPlatform()`

UnsetCloudPlatform ensures that no value is present for CloudPlatform, not even an explicit nil
### GetAwsTiering

`func (o *TierLevelSettings) GetAwsTiering() AWSTiers`

GetAwsTiering returns the AwsTiering field if non-nil, zero value otherwise.

### GetAwsTieringOk

`func (o *TierLevelSettings) GetAwsTieringOk() (*AWSTiers, bool)`

GetAwsTieringOk returns a tuple with the AwsTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTiering

`func (o *TierLevelSettings) SetAwsTiering(v AWSTiers)`

SetAwsTiering sets AwsTiering field to given value.

### HasAwsTiering

`func (o *TierLevelSettings) HasAwsTiering() bool`

HasAwsTiering returns a boolean if a field has been set.

### GetAzureTiering

`func (o *TierLevelSettings) GetAzureTiering() AzureTiers`

GetAzureTiering returns the AzureTiering field if non-nil, zero value otherwise.

### GetAzureTieringOk

`func (o *TierLevelSettings) GetAzureTieringOk() (*AzureTiers, bool)`

GetAzureTieringOk returns a tuple with the AzureTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTiering

`func (o *TierLevelSettings) SetAzureTiering(v AzureTiers)`

SetAzureTiering sets AzureTiering field to given value.

### HasAzureTiering

`func (o *TierLevelSettings) HasAzureTiering() bool`

HasAzureTiering returns a boolean if a field has been set.

### GetGoogleTiering

`func (o *TierLevelSettings) GetGoogleTiering() GoogleTiers`

GetGoogleTiering returns the GoogleTiering field if non-nil, zero value otherwise.

### GetGoogleTieringOk

`func (o *TierLevelSettings) GetGoogleTieringOk() (*GoogleTiers, bool)`

GetGoogleTieringOk returns a tuple with the GoogleTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTiering

`func (o *TierLevelSettings) SetGoogleTiering(v GoogleTiers)`

SetGoogleTiering sets GoogleTiering field to given value.

### HasGoogleTiering

`func (o *TierLevelSettings) HasGoogleTiering() bool`

HasGoogleTiering returns a boolean if a field has been set.

### GetOracleTiering

`func (o *TierLevelSettings) GetOracleTiering() OracleTiers`

GetOracleTiering returns the OracleTiering field if non-nil, zero value otherwise.

### GetOracleTieringOk

`func (o *TierLevelSettings) GetOracleTieringOk() (*OracleTiers, bool)`

GetOracleTieringOk returns a tuple with the OracleTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTiering

`func (o *TierLevelSettings) SetOracleTiering(v OracleTiers)`

SetOracleTiering sets OracleTiering field to given value.

### HasOracleTiering

`func (o *TierLevelSettings) HasOracleTiering() bool`

HasOracleTiering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


