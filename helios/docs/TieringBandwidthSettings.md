# TieringBandwidthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Download** | Pointer to [**BandwidthThrottling**](BandwidthThrottling.md) |  | [optional] 
**Upload** | Pointer to [**BandwidthThrottling**](BandwidthThrottling.md) |  | [optional] 

## Methods

### NewTieringBandwidthSettings

`func NewTieringBandwidthSettings() *TieringBandwidthSettings`

NewTieringBandwidthSettings instantiates a new TieringBandwidthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieringBandwidthSettingsWithDefaults

`func NewTieringBandwidthSettingsWithDefaults() *TieringBandwidthSettings`

NewTieringBandwidthSettingsWithDefaults instantiates a new TieringBandwidthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownload

`func (o *TieringBandwidthSettings) GetDownload() BandwidthThrottling`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *TieringBandwidthSettings) GetDownloadOk() (*BandwidthThrottling, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *TieringBandwidthSettings) SetDownload(v BandwidthThrottling)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *TieringBandwidthSettings) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetUpload

`func (o *TieringBandwidthSettings) GetUpload() BandwidthThrottling`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *TieringBandwidthSettings) GetUploadOk() (*BandwidthThrottling, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *TieringBandwidthSettings) SetUpload(v BandwidthThrottling)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *TieringBandwidthSettings) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


