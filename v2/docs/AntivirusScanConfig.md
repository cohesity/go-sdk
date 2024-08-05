# AntivirusScanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockAccessOnScanFailure** | Pointer to **NullableBool** | Specifies whether block access to the file when antivirus scan fails. | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Specifies whether the antivirus service is enabled or not. | [optional] 
**MaximumScanFileSize** | Pointer to **NullableInt64** | Specifies maximum file size that will be sent to antivirus server for scanning. if greater than zero, the file size that exceeds this size would be skipped from virus scan. | [optional] 
**PrefixScanFilter** | Pointer to [**AntivirusScanConfigPrefixScanFilter**](AntivirusScanConfigPrefixScanFilter.md) |  | [optional] 
**S3TaggingFilter** | Pointer to [**AntivirusScanConfigS3TaggingFilter**](AntivirusScanConfigS3TaggingFilter.md) |  | [optional] 
**ScanFilter** | Pointer to [**AntivirusScanConfigScanFilter**](AntivirusScanConfigScanFilter.md) |  | [optional] 
**ScanOnAccess** | Pointer to **NullableBool** | Specifies whether to scan a SMB file or S3 object before it is opened/GET. | [optional] 
**ScanOnClose** | Pointer to **NullableBool** | Specifies whether to scan a SMB file when it is closed after modify. | [optional] 
**ScanOnPut** | Pointer to **NullableBool** | Specifies whether to scan a S3 object after it is PUT. | [optional] 
**ScanTimeoutUsecs** | **NullableInt32** | Specifies the maximum amount of time that a scan can take before timing out. | 

## Methods

### NewAntivirusScanConfig

`func NewAntivirusScanConfig(scanTimeoutUsecs NullableInt32, ) *AntivirusScanConfig`

NewAntivirusScanConfig instantiates a new AntivirusScanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusScanConfigWithDefaults

`func NewAntivirusScanConfigWithDefaults() *AntivirusScanConfig`

NewAntivirusScanConfigWithDefaults instantiates a new AntivirusScanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockAccessOnScanFailure

`func (o *AntivirusScanConfig) GetBlockAccessOnScanFailure() bool`

GetBlockAccessOnScanFailure returns the BlockAccessOnScanFailure field if non-nil, zero value otherwise.

### GetBlockAccessOnScanFailureOk

`func (o *AntivirusScanConfig) GetBlockAccessOnScanFailureOk() (*bool, bool)`

GetBlockAccessOnScanFailureOk returns a tuple with the BlockAccessOnScanFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAccessOnScanFailure

`func (o *AntivirusScanConfig) SetBlockAccessOnScanFailure(v bool)`

SetBlockAccessOnScanFailure sets BlockAccessOnScanFailure field to given value.

### HasBlockAccessOnScanFailure

`func (o *AntivirusScanConfig) HasBlockAccessOnScanFailure() bool`

HasBlockAccessOnScanFailure returns a boolean if a field has been set.

### SetBlockAccessOnScanFailureNil

`func (o *AntivirusScanConfig) SetBlockAccessOnScanFailureNil(b bool)`

 SetBlockAccessOnScanFailureNil sets the value for BlockAccessOnScanFailure to be an explicit nil

### UnsetBlockAccessOnScanFailure
`func (o *AntivirusScanConfig) UnsetBlockAccessOnScanFailure()`

UnsetBlockAccessOnScanFailure ensures that no value is present for BlockAccessOnScanFailure, not even an explicit nil
### GetIsEnabled

`func (o *AntivirusScanConfig) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AntivirusScanConfig) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AntivirusScanConfig) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AntivirusScanConfig) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *AntivirusScanConfig) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *AntivirusScanConfig) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil
### GetMaximumScanFileSize

`func (o *AntivirusScanConfig) GetMaximumScanFileSize() int64`

GetMaximumScanFileSize returns the MaximumScanFileSize field if non-nil, zero value otherwise.

### GetMaximumScanFileSizeOk

`func (o *AntivirusScanConfig) GetMaximumScanFileSizeOk() (*int64, bool)`

GetMaximumScanFileSizeOk returns a tuple with the MaximumScanFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumScanFileSize

`func (o *AntivirusScanConfig) SetMaximumScanFileSize(v int64)`

SetMaximumScanFileSize sets MaximumScanFileSize field to given value.

### HasMaximumScanFileSize

`func (o *AntivirusScanConfig) HasMaximumScanFileSize() bool`

HasMaximumScanFileSize returns a boolean if a field has been set.

### SetMaximumScanFileSizeNil

`func (o *AntivirusScanConfig) SetMaximumScanFileSizeNil(b bool)`

 SetMaximumScanFileSizeNil sets the value for MaximumScanFileSize to be an explicit nil

### UnsetMaximumScanFileSize
`func (o *AntivirusScanConfig) UnsetMaximumScanFileSize()`

UnsetMaximumScanFileSize ensures that no value is present for MaximumScanFileSize, not even an explicit nil
### GetPrefixScanFilter

`func (o *AntivirusScanConfig) GetPrefixScanFilter() AntivirusScanConfigPrefixScanFilter`

GetPrefixScanFilter returns the PrefixScanFilter field if non-nil, zero value otherwise.

### GetPrefixScanFilterOk

`func (o *AntivirusScanConfig) GetPrefixScanFilterOk() (*AntivirusScanConfigPrefixScanFilter, bool)`

GetPrefixScanFilterOk returns a tuple with the PrefixScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixScanFilter

`func (o *AntivirusScanConfig) SetPrefixScanFilter(v AntivirusScanConfigPrefixScanFilter)`

SetPrefixScanFilter sets PrefixScanFilter field to given value.

### HasPrefixScanFilter

`func (o *AntivirusScanConfig) HasPrefixScanFilter() bool`

HasPrefixScanFilter returns a boolean if a field has been set.

### GetS3TaggingFilter

`func (o *AntivirusScanConfig) GetS3TaggingFilter() AntivirusScanConfigS3TaggingFilter`

GetS3TaggingFilter returns the S3TaggingFilter field if non-nil, zero value otherwise.

### GetS3TaggingFilterOk

`func (o *AntivirusScanConfig) GetS3TaggingFilterOk() (*AntivirusScanConfigS3TaggingFilter, bool)`

GetS3TaggingFilterOk returns a tuple with the S3TaggingFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3TaggingFilter

`func (o *AntivirusScanConfig) SetS3TaggingFilter(v AntivirusScanConfigS3TaggingFilter)`

SetS3TaggingFilter sets S3TaggingFilter field to given value.

### HasS3TaggingFilter

`func (o *AntivirusScanConfig) HasS3TaggingFilter() bool`

HasS3TaggingFilter returns a boolean if a field has been set.

### GetScanFilter

`func (o *AntivirusScanConfig) GetScanFilter() AntivirusScanConfigScanFilter`

GetScanFilter returns the ScanFilter field if non-nil, zero value otherwise.

### GetScanFilterOk

`func (o *AntivirusScanConfig) GetScanFilterOk() (*AntivirusScanConfigScanFilter, bool)`

GetScanFilterOk returns a tuple with the ScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFilter

`func (o *AntivirusScanConfig) SetScanFilter(v AntivirusScanConfigScanFilter)`

SetScanFilter sets ScanFilter field to given value.

### HasScanFilter

`func (o *AntivirusScanConfig) HasScanFilter() bool`

HasScanFilter returns a boolean if a field has been set.

### GetScanOnAccess

`func (o *AntivirusScanConfig) GetScanOnAccess() bool`

GetScanOnAccess returns the ScanOnAccess field if non-nil, zero value otherwise.

### GetScanOnAccessOk

`func (o *AntivirusScanConfig) GetScanOnAccessOk() (*bool, bool)`

GetScanOnAccessOk returns a tuple with the ScanOnAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanOnAccess

`func (o *AntivirusScanConfig) SetScanOnAccess(v bool)`

SetScanOnAccess sets ScanOnAccess field to given value.

### HasScanOnAccess

`func (o *AntivirusScanConfig) HasScanOnAccess() bool`

HasScanOnAccess returns a boolean if a field has been set.

### SetScanOnAccessNil

`func (o *AntivirusScanConfig) SetScanOnAccessNil(b bool)`

 SetScanOnAccessNil sets the value for ScanOnAccess to be an explicit nil

### UnsetScanOnAccess
`func (o *AntivirusScanConfig) UnsetScanOnAccess()`

UnsetScanOnAccess ensures that no value is present for ScanOnAccess, not even an explicit nil
### GetScanOnClose

`func (o *AntivirusScanConfig) GetScanOnClose() bool`

GetScanOnClose returns the ScanOnClose field if non-nil, zero value otherwise.

### GetScanOnCloseOk

`func (o *AntivirusScanConfig) GetScanOnCloseOk() (*bool, bool)`

GetScanOnCloseOk returns a tuple with the ScanOnClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanOnClose

`func (o *AntivirusScanConfig) SetScanOnClose(v bool)`

SetScanOnClose sets ScanOnClose field to given value.

### HasScanOnClose

`func (o *AntivirusScanConfig) HasScanOnClose() bool`

HasScanOnClose returns a boolean if a field has been set.

### SetScanOnCloseNil

`func (o *AntivirusScanConfig) SetScanOnCloseNil(b bool)`

 SetScanOnCloseNil sets the value for ScanOnClose to be an explicit nil

### UnsetScanOnClose
`func (o *AntivirusScanConfig) UnsetScanOnClose()`

UnsetScanOnClose ensures that no value is present for ScanOnClose, not even an explicit nil
### GetScanOnPut

`func (o *AntivirusScanConfig) GetScanOnPut() bool`

GetScanOnPut returns the ScanOnPut field if non-nil, zero value otherwise.

### GetScanOnPutOk

`func (o *AntivirusScanConfig) GetScanOnPutOk() (*bool, bool)`

GetScanOnPutOk returns a tuple with the ScanOnPut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanOnPut

`func (o *AntivirusScanConfig) SetScanOnPut(v bool)`

SetScanOnPut sets ScanOnPut field to given value.

### HasScanOnPut

`func (o *AntivirusScanConfig) HasScanOnPut() bool`

HasScanOnPut returns a boolean if a field has been set.

### SetScanOnPutNil

`func (o *AntivirusScanConfig) SetScanOnPutNil(b bool)`

 SetScanOnPutNil sets the value for ScanOnPut to be an explicit nil

### UnsetScanOnPut
`func (o *AntivirusScanConfig) UnsetScanOnPut()`

UnsetScanOnPut ensures that no value is present for ScanOnPut, not even an explicit nil
### GetScanTimeoutUsecs

`func (o *AntivirusScanConfig) GetScanTimeoutUsecs() int32`

GetScanTimeoutUsecs returns the ScanTimeoutUsecs field if non-nil, zero value otherwise.

### GetScanTimeoutUsecsOk

`func (o *AntivirusScanConfig) GetScanTimeoutUsecsOk() (*int32, bool)`

GetScanTimeoutUsecsOk returns a tuple with the ScanTimeoutUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTimeoutUsecs

`func (o *AntivirusScanConfig) SetScanTimeoutUsecs(v int32)`

SetScanTimeoutUsecs sets ScanTimeoutUsecs field to given value.


### SetScanTimeoutUsecsNil

`func (o *AntivirusScanConfig) SetScanTimeoutUsecsNil(b bool)`

 SetScanTimeoutUsecsNil sets the value for ScanTimeoutUsecs to be an explicit nil

### UnsetScanTimeoutUsecs
`func (o *AntivirusScanConfig) UnsetScanTimeoutUsecs()`

UnsetScanTimeoutUsecs ensures that no value is present for ScanTimeoutUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


