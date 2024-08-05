# ArchivalTargetTierInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsTiering** | Pointer to [**AWSTiers**](AWSTiers.md) |  | [optional] 
**AzureTiering** | Pointer to [**AzureTiers**](AzureTiers.md) |  | [optional] 
**CloudPlatform** | **NullableString** | Specifies the cloud platform to enable tiering. | 
**GoogleTiering** | Pointer to [**GoogleTiers**](GoogleTiers.md) |  | [optional] 
**OracleTiering** | Pointer to [**OracleTiers**](OracleTiers.md) |  | [optional] 
**CurrentTierType** | Pointer to **NullableString** | Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run. | [optional] 

## Methods

### NewArchivalTargetTierInfo

`func NewArchivalTargetTierInfo(cloudPlatform NullableString, ) *ArchivalTargetTierInfo`

NewArchivalTargetTierInfo instantiates a new ArchivalTargetTierInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalTargetTierInfoWithDefaults

`func NewArchivalTargetTierInfoWithDefaults() *ArchivalTargetTierInfo`

NewArchivalTargetTierInfoWithDefaults instantiates a new ArchivalTargetTierInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsTiering

`func (o *ArchivalTargetTierInfo) GetAwsTiering() AWSTiers`

GetAwsTiering returns the AwsTiering field if non-nil, zero value otherwise.

### GetAwsTieringOk

`func (o *ArchivalTargetTierInfo) GetAwsTieringOk() (*AWSTiers, bool)`

GetAwsTieringOk returns a tuple with the AwsTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTiering

`func (o *ArchivalTargetTierInfo) SetAwsTiering(v AWSTiers)`

SetAwsTiering sets AwsTiering field to given value.

### HasAwsTiering

`func (o *ArchivalTargetTierInfo) HasAwsTiering() bool`

HasAwsTiering returns a boolean if a field has been set.

### GetAzureTiering

`func (o *ArchivalTargetTierInfo) GetAzureTiering() AzureTiers`

GetAzureTiering returns the AzureTiering field if non-nil, zero value otherwise.

### GetAzureTieringOk

`func (o *ArchivalTargetTierInfo) GetAzureTieringOk() (*AzureTiers, bool)`

GetAzureTieringOk returns a tuple with the AzureTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTiering

`func (o *ArchivalTargetTierInfo) SetAzureTiering(v AzureTiers)`

SetAzureTiering sets AzureTiering field to given value.

### HasAzureTiering

`func (o *ArchivalTargetTierInfo) HasAzureTiering() bool`

HasAzureTiering returns a boolean if a field has been set.

### GetCloudPlatform

`func (o *ArchivalTargetTierInfo) GetCloudPlatform() string`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *ArchivalTargetTierInfo) GetCloudPlatformOk() (*string, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *ArchivalTargetTierInfo) SetCloudPlatform(v string)`

SetCloudPlatform sets CloudPlatform field to given value.


### SetCloudPlatformNil

`func (o *ArchivalTargetTierInfo) SetCloudPlatformNil(b bool)`

 SetCloudPlatformNil sets the value for CloudPlatform to be an explicit nil

### UnsetCloudPlatform
`func (o *ArchivalTargetTierInfo) UnsetCloudPlatform()`

UnsetCloudPlatform ensures that no value is present for CloudPlatform, not even an explicit nil
### GetGoogleTiering

`func (o *ArchivalTargetTierInfo) GetGoogleTiering() GoogleTiers`

GetGoogleTiering returns the GoogleTiering field if non-nil, zero value otherwise.

### GetGoogleTieringOk

`func (o *ArchivalTargetTierInfo) GetGoogleTieringOk() (*GoogleTiers, bool)`

GetGoogleTieringOk returns a tuple with the GoogleTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleTiering

`func (o *ArchivalTargetTierInfo) SetGoogleTiering(v GoogleTiers)`

SetGoogleTiering sets GoogleTiering field to given value.

### HasGoogleTiering

`func (o *ArchivalTargetTierInfo) HasGoogleTiering() bool`

HasGoogleTiering returns a boolean if a field has been set.

### GetOracleTiering

`func (o *ArchivalTargetTierInfo) GetOracleTiering() OracleTiers`

GetOracleTiering returns the OracleTiering field if non-nil, zero value otherwise.

### GetOracleTieringOk

`func (o *ArchivalTargetTierInfo) GetOracleTieringOk() (*OracleTiers, bool)`

GetOracleTieringOk returns a tuple with the OracleTiering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTiering

`func (o *ArchivalTargetTierInfo) SetOracleTiering(v OracleTiers)`

SetOracleTiering sets OracleTiering field to given value.

### HasOracleTiering

`func (o *ArchivalTargetTierInfo) HasOracleTiering() bool`

HasOracleTiering returns a boolean if a field has been set.

### GetCurrentTierType

`func (o *ArchivalTargetTierInfo) GetCurrentTierType() string`

GetCurrentTierType returns the CurrentTierType field if non-nil, zero value otherwise.

### GetCurrentTierTypeOk

`func (o *ArchivalTargetTierInfo) GetCurrentTierTypeOk() (*string, bool)`

GetCurrentTierTypeOk returns a tuple with the CurrentTierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTierType

`func (o *ArchivalTargetTierInfo) SetCurrentTierType(v string)`

SetCurrentTierType sets CurrentTierType field to given value.

### HasCurrentTierType

`func (o *ArchivalTargetTierInfo) HasCurrentTierType() bool`

HasCurrentTierType returns a boolean if a field has been set.

### SetCurrentTierTypeNil

`func (o *ArchivalTargetTierInfo) SetCurrentTierTypeNil(b bool)`

 SetCurrentTierTypeNil sets the value for CurrentTierType to be an explicit nil

### UnsetCurrentTierType
`func (o *ArchivalTargetTierInfo) UnsetCurrentTierType()`

UnsetCurrentTierType ensures that no value is present for CurrentTierType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


