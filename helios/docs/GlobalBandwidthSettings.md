# GlobalBandwidthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalParams** | Pointer to [**ArchivalBandwidthSettings**](ArchivalBandwidthSettings.md) |  | [optional] 
**TieringParams** | Pointer to [**TieringBandwidthSettings**](TieringBandwidthSettings.md) |  | [optional] 

## Methods

### NewGlobalBandwidthSettings

`func NewGlobalBandwidthSettings() *GlobalBandwidthSettings`

NewGlobalBandwidthSettings instantiates a new GlobalBandwidthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalBandwidthSettingsWithDefaults

`func NewGlobalBandwidthSettingsWithDefaults() *GlobalBandwidthSettings`

NewGlobalBandwidthSettingsWithDefaults instantiates a new GlobalBandwidthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalParams

`func (o *GlobalBandwidthSettings) GetArchivalParams() ArchivalBandwidthSettings`

GetArchivalParams returns the ArchivalParams field if non-nil, zero value otherwise.

### GetArchivalParamsOk

`func (o *GlobalBandwidthSettings) GetArchivalParamsOk() (*ArchivalBandwidthSettings, bool)`

GetArchivalParamsOk returns a tuple with the ArchivalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalParams

`func (o *GlobalBandwidthSettings) SetArchivalParams(v ArchivalBandwidthSettings)`

SetArchivalParams sets ArchivalParams field to given value.

### HasArchivalParams

`func (o *GlobalBandwidthSettings) HasArchivalParams() bool`

HasArchivalParams returns a boolean if a field has been set.

### GetTieringParams

`func (o *GlobalBandwidthSettings) GetTieringParams() TieringBandwidthSettings`

GetTieringParams returns the TieringParams field if non-nil, zero value otherwise.

### GetTieringParamsOk

`func (o *GlobalBandwidthSettings) GetTieringParamsOk() (*TieringBandwidthSettings, bool)`

GetTieringParamsOk returns a tuple with the TieringParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieringParams

`func (o *GlobalBandwidthSettings) SetTieringParams(v TieringBandwidthSettings)`

SetTieringParams sets TieringParams field to given value.

### HasTieringParams

`func (o *GlobalBandwidthSettings) HasTieringParams() bool`

HasTieringParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


