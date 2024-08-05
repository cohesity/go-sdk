# NasThrottlingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullBackupConfig** | Pointer to [**NasFullThrottlingConfig**](NasFullThrottlingConfig.md) |  | [optional] 
**IncrementalBackupConfig** | Pointer to [**NasIncrementalThrottlingConfig**](NasIncrementalThrottlingConfig.md) |  | [optional] 

## Methods

### NewNasThrottlingConfig

`func NewNasThrottlingConfig() *NasThrottlingConfig`

NewNasThrottlingConfig instantiates a new NasThrottlingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasThrottlingConfigWithDefaults

`func NewNasThrottlingConfigWithDefaults() *NasThrottlingConfig`

NewNasThrottlingConfigWithDefaults instantiates a new NasThrottlingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullBackupConfig

`func (o *NasThrottlingConfig) GetFullBackupConfig() NasFullThrottlingConfig`

GetFullBackupConfig returns the FullBackupConfig field if non-nil, zero value otherwise.

### GetFullBackupConfigOk

`func (o *NasThrottlingConfig) GetFullBackupConfigOk() (*NasFullThrottlingConfig, bool)`

GetFullBackupConfigOk returns a tuple with the FullBackupConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupConfig

`func (o *NasThrottlingConfig) SetFullBackupConfig(v NasFullThrottlingConfig)`

SetFullBackupConfig sets FullBackupConfig field to given value.

### HasFullBackupConfig

`func (o *NasThrottlingConfig) HasFullBackupConfig() bool`

HasFullBackupConfig returns a boolean if a field has been set.

### GetIncrementalBackupConfig

`func (o *NasThrottlingConfig) GetIncrementalBackupConfig() NasIncrementalThrottlingConfig`

GetIncrementalBackupConfig returns the IncrementalBackupConfig field if non-nil, zero value otherwise.

### GetIncrementalBackupConfigOk

`func (o *NasThrottlingConfig) GetIncrementalBackupConfigOk() (*NasIncrementalThrottlingConfig, bool)`

GetIncrementalBackupConfigOk returns a tuple with the IncrementalBackupConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupConfig

`func (o *NasThrottlingConfig) SetIncrementalBackupConfig(v NasIncrementalThrottlingConfig)`

SetIncrementalBackupConfig sets IncrementalBackupConfig field to given value.

### HasIncrementalBackupConfig

`func (o *NasThrottlingConfig) HasIncrementalBackupConfig() bool`

HasIncrementalBackupConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


