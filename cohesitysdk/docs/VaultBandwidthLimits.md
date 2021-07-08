# VaultBandwidthLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Download** | Pointer to [**BandwidthLimit**](BandwidthLimit.md) |  | [optional] 
**Upload** | Pointer to [**BandwidthLimit**](BandwidthLimit.md) |  | [optional] 

## Methods

### NewVaultBandwidthLimits

`func NewVaultBandwidthLimits() *VaultBandwidthLimits`

NewVaultBandwidthLimits instantiates a new VaultBandwidthLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultBandwidthLimitsWithDefaults

`func NewVaultBandwidthLimitsWithDefaults() *VaultBandwidthLimits`

NewVaultBandwidthLimitsWithDefaults instantiates a new VaultBandwidthLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownload

`func (o *VaultBandwidthLimits) GetDownload() BandwidthLimit`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *VaultBandwidthLimits) GetDownloadOk() (*BandwidthLimit, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *VaultBandwidthLimits) SetDownload(v BandwidthLimit)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *VaultBandwidthLimits) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetUpload

`func (o *VaultBandwidthLimits) GetUpload() BandwidthLimit`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *VaultBandwidthLimits) GetUploadOk() (*BandwidthLimit, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *VaultBandwidthLimits) SetUpload(v BandwidthLimit)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *VaultBandwidthLimits) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


