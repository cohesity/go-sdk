# CommonNoSqlRecoveryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedConfigs** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Specifies the advanced configuration for a recovery job. | [optional] 
**BandwidthMBPS** | Pointer to **NullableInt64** | Specifies the maximum network bandwidth that each concurrent IO Stream can use for exchanging data with the cluster. | [optional] 
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams that will be created to exchange data with the cluster. | [optional] 
**Overwrite** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**RecoverTo** | Pointer to **NullableInt64** | Specifies the &#39;Source Registration ID&#39; of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**Warnings** | Pointer to **[]string** | This field will hold the warnings in cases where the job status is SucceededWithWarnings. | [optional] [readonly] 

## Methods

### NewCommonNoSqlRecoveryOptions

`func NewCommonNoSqlRecoveryOptions() *CommonNoSqlRecoveryOptions`

NewCommonNoSqlRecoveryOptions instantiates a new CommonNoSqlRecoveryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonNoSqlRecoveryOptionsWithDefaults

`func NewCommonNoSqlRecoveryOptionsWithDefaults() *CommonNoSqlRecoveryOptions`

NewCommonNoSqlRecoveryOptionsWithDefaults instantiates a new CommonNoSqlRecoveryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedConfigs

`func (o *CommonNoSqlRecoveryOptions) GetAdvancedConfigs() []KeyValuePair`

GetAdvancedConfigs returns the AdvancedConfigs field if non-nil, zero value otherwise.

### GetAdvancedConfigsOk

`func (o *CommonNoSqlRecoveryOptions) GetAdvancedConfigsOk() (*[]KeyValuePair, bool)`

GetAdvancedConfigsOk returns a tuple with the AdvancedConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedConfigs

`func (o *CommonNoSqlRecoveryOptions) SetAdvancedConfigs(v []KeyValuePair)`

SetAdvancedConfigs sets AdvancedConfigs field to given value.

### HasAdvancedConfigs

`func (o *CommonNoSqlRecoveryOptions) HasAdvancedConfigs() bool`

HasAdvancedConfigs returns a boolean if a field has been set.

### SetAdvancedConfigsNil

`func (o *CommonNoSqlRecoveryOptions) SetAdvancedConfigsNil(b bool)`

 SetAdvancedConfigsNil sets the value for AdvancedConfigs to be an explicit nil

### UnsetAdvancedConfigs
`func (o *CommonNoSqlRecoveryOptions) UnsetAdvancedConfigs()`

UnsetAdvancedConfigs ensures that no value is present for AdvancedConfigs, not even an explicit nil
### GetBandwidthMBPS

`func (o *CommonNoSqlRecoveryOptions) GetBandwidthMBPS() int64`

GetBandwidthMBPS returns the BandwidthMBPS field if non-nil, zero value otherwise.

### GetBandwidthMBPSOk

`func (o *CommonNoSqlRecoveryOptions) GetBandwidthMBPSOk() (*int64, bool)`

GetBandwidthMBPSOk returns a tuple with the BandwidthMBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMBPS

`func (o *CommonNoSqlRecoveryOptions) SetBandwidthMBPS(v int64)`

SetBandwidthMBPS sets BandwidthMBPS field to given value.

### HasBandwidthMBPS

`func (o *CommonNoSqlRecoveryOptions) HasBandwidthMBPS() bool`

HasBandwidthMBPS returns a boolean if a field has been set.

### SetBandwidthMBPSNil

`func (o *CommonNoSqlRecoveryOptions) SetBandwidthMBPSNil(b bool)`

 SetBandwidthMBPSNil sets the value for BandwidthMBPS to be an explicit nil

### UnsetBandwidthMBPS
`func (o *CommonNoSqlRecoveryOptions) UnsetBandwidthMBPS()`

UnsetBandwidthMBPS ensures that no value is present for BandwidthMBPS, not even an explicit nil
### GetConcurrency

`func (o *CommonNoSqlRecoveryOptions) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *CommonNoSqlRecoveryOptions) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *CommonNoSqlRecoveryOptions) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *CommonNoSqlRecoveryOptions) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *CommonNoSqlRecoveryOptions) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *CommonNoSqlRecoveryOptions) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetOverwrite

`func (o *CommonNoSqlRecoveryOptions) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *CommonNoSqlRecoveryOptions) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *CommonNoSqlRecoveryOptions) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *CommonNoSqlRecoveryOptions) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *CommonNoSqlRecoveryOptions) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *CommonNoSqlRecoveryOptions) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetRecoverTo

`func (o *CommonNoSqlRecoveryOptions) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *CommonNoSqlRecoveryOptions) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *CommonNoSqlRecoveryOptions) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *CommonNoSqlRecoveryOptions) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *CommonNoSqlRecoveryOptions) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *CommonNoSqlRecoveryOptions) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetWarnings

`func (o *CommonNoSqlRecoveryOptions) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CommonNoSqlRecoveryOptions) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CommonNoSqlRecoveryOptions) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *CommonNoSqlRecoveryOptions) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *CommonNoSqlRecoveryOptions) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *CommonNoSqlRecoveryOptions) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


