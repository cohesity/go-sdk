# ConnectionBandwidthLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Download** | Pointer to [**[]BandwidthLimit**](BandwidthLimit.md) | Specifies the max rate limit at which we download the data. | [optional] 
**Upload** | Pointer to [**[]BandwidthLimit**](BandwidthLimit.md) | Specifies the max rate limit at which we upload the data. | [optional] 
**TenantId** | **NullableString** | The tenant Id corresponding to this request. | 
**Timezone** | Pointer to **NullableString** | Specifies a time zone for the specified time period. The time zone is defined in the following format: &#39;Area/Location&#39;, for example: &#39;America/New_York&#39;. | [optional] 

## Methods

### NewConnectionBandwidthLimits

`func NewConnectionBandwidthLimits(tenantId NullableString, ) *ConnectionBandwidthLimits`

NewConnectionBandwidthLimits instantiates a new ConnectionBandwidthLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionBandwidthLimitsWithDefaults

`func NewConnectionBandwidthLimitsWithDefaults() *ConnectionBandwidthLimits`

NewConnectionBandwidthLimitsWithDefaults instantiates a new ConnectionBandwidthLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownload

`func (o *ConnectionBandwidthLimits) GetDownload() []BandwidthLimit`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ConnectionBandwidthLimits) GetDownloadOk() (*[]BandwidthLimit, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ConnectionBandwidthLimits) SetDownload(v []BandwidthLimit)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ConnectionBandwidthLimits) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### SetDownloadNil

`func (o *ConnectionBandwidthLimits) SetDownloadNil(b bool)`

 SetDownloadNil sets the value for Download to be an explicit nil

### UnsetDownload
`func (o *ConnectionBandwidthLimits) UnsetDownload()`

UnsetDownload ensures that no value is present for Download, not even an explicit nil
### GetUpload

`func (o *ConnectionBandwidthLimits) GetUpload() []BandwidthLimit`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ConnectionBandwidthLimits) GetUploadOk() (*[]BandwidthLimit, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ConnectionBandwidthLimits) SetUpload(v []BandwidthLimit)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ConnectionBandwidthLimits) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### SetUploadNil

`func (o *ConnectionBandwidthLimits) SetUploadNil(b bool)`

 SetUploadNil sets the value for Upload to be an explicit nil

### UnsetUpload
`func (o *ConnectionBandwidthLimits) UnsetUpload()`

UnsetUpload ensures that no value is present for Upload, not even an explicit nil
### GetTenantId

`func (o *ConnectionBandwidthLimits) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ConnectionBandwidthLimits) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ConnectionBandwidthLimits) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *ConnectionBandwidthLimits) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ConnectionBandwidthLimits) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTimezone

`func (o *ConnectionBandwidthLimits) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ConnectionBandwidthLimits) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ConnectionBandwidthLimits) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ConnectionBandwidthLimits) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *ConnectionBandwidthLimits) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *ConnectionBandwidthLimits) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


