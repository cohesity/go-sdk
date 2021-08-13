# AntivirusScanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockAccessOnScanFailure** | Pointer to **NullableBool** | Specifies whether block access to the file when antivirus scan fails. | [optional] 
**IsEnabled** | Pointer to **NullableBool** | Specifies whether the antivirus service is enabled or not. | [optional] 
**MaximumScanFileSize** | Pointer to **NullableInt64** | Specifies maximum file size that will be sent to antivirus server for scanning. if greater than zero, the file size that exceeds this size would be skipped from virus scan. | [optional] 
**ScanFilter** | Pointer to [**FileExtensionFilter**](FileExtensionFilter.md) | Files extension that meets these filter criteria will be sent to antivirus server for the scan. | [optional] 
**ScanOnAccess** | Pointer to **NullableBool** | Specifies whether to scan a file when it is opened. | [optional] 
**ScanOnClose** | Pointer to **NullableBool** | Specifies whether to scan a file when it is closed after modify. | [optional] 
**ScanTimeoutUsecs** | Pointer to **NullableInt32** | Specifies the maximum amount of time that a scan can take before timing out. | [optional] 

## Methods

### NewAntivirusScanConfig

`func NewAntivirusScanConfig() *AntivirusScanConfig`

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
### GetScanFilter

`func (o *AntivirusScanConfig) GetScanFilter() FileExtensionFilter`

GetScanFilter returns the ScanFilter field if non-nil, zero value otherwise.

### GetScanFilterOk

`func (o *AntivirusScanConfig) GetScanFilterOk() (*FileExtensionFilter, bool)`

GetScanFilterOk returns a tuple with the ScanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFilter

`func (o *AntivirusScanConfig) SetScanFilter(v FileExtensionFilter)`

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

### HasScanTimeoutUsecs

`func (o *AntivirusScanConfig) HasScanTimeoutUsecs() bool`

HasScanTimeoutUsecs returns a boolean if a field has been set.

### SetScanTimeoutUsecsNil

`func (o *AntivirusScanConfig) SetScanTimeoutUsecsNil(b bool)`

 SetScanTimeoutUsecsNil sets the value for ScanTimeoutUsecs to be an explicit nil

### UnsetScanTimeoutUsecs
`func (o *AntivirusScanConfig) UnsetScanTimeoutUsecs()`

UnsetScanTimeoutUsecs ensures that no value is present for ScanTimeoutUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


