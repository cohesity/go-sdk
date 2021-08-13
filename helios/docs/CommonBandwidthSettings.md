# CommonBandwidthSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Download** | Pointer to [**BandwidthThrottling**](BandwidthThrottling.md) |  | [optional] 
**Upload** | Pointer to [**BandwidthThrottling**](BandwidthThrottling.md) |  | [optional] 

## Methods

### NewCommonBandwidthSettings

`func NewCommonBandwidthSettings() *CommonBandwidthSettings`

NewCommonBandwidthSettings instantiates a new CommonBandwidthSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonBandwidthSettingsWithDefaults

`func NewCommonBandwidthSettingsWithDefaults() *CommonBandwidthSettings`

NewCommonBandwidthSettingsWithDefaults instantiates a new CommonBandwidthSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownload

`func (o *CommonBandwidthSettings) GetDownload() BandwidthThrottling`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *CommonBandwidthSettings) GetDownloadOk() (*BandwidthThrottling, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *CommonBandwidthSettings) SetDownload(v BandwidthThrottling)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *CommonBandwidthSettings) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetUpload

`func (o *CommonBandwidthSettings) GetUpload() BandwidthThrottling`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *CommonBandwidthSettings) GetUploadOk() (*BandwidthThrottling, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *CommonBandwidthSettings) SetUpload(v BandwidthThrottling)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *CommonBandwidthSettings) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


