# GetConnectionBandwidthResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectorGroupId** | Pointer to **NullableInt64** | Specifies the connector group id of connector groups. | [optional] 
**Download** | Pointer to [**[]BandwidthLimit**](BandwidthLimit.md) | Specifies the max rate limit at which we download the data. | [optional] 
**TenantId** | **NullableString** | The tenant Id corresponding to this request. | 
**Timezone** | Pointer to **NullableString** | Specifies a time zone for the specified time period. The time zone is defined in the following format: &#39;Area/Location&#39;, for example: &#39;America/New_York&#39;. | [optional] 
**Upload** | Pointer to [**[]BandwidthLimit**](BandwidthLimit.md) | Specifies the max rate limit at which we upload the data. | [optional] 
**ConnectionId** | Pointer to **NullableInt64** | Connection Id for which bandwidth settings are to be returned | [optional] 
**ConnectorGroups** | Pointer to [**[]ConnectorGroup**](ConnectorGroup.md) | Specifies the list of connector groups. | [optional] 

## Methods

### NewGetConnectionBandwidthResponseBody

`func NewGetConnectionBandwidthResponseBody(tenantId NullableString, ) *GetConnectionBandwidthResponseBody`

NewGetConnectionBandwidthResponseBody instantiates a new GetConnectionBandwidthResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionBandwidthResponseBodyWithDefaults

`func NewGetConnectionBandwidthResponseBodyWithDefaults() *GetConnectionBandwidthResponseBody`

NewGetConnectionBandwidthResponseBodyWithDefaults instantiates a new GetConnectionBandwidthResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectorGroupId

`func (o *GetConnectionBandwidthResponseBody) GetConnectorGroupId() int64`

GetConnectorGroupId returns the ConnectorGroupId field if non-nil, zero value otherwise.

### GetConnectorGroupIdOk

`func (o *GetConnectionBandwidthResponseBody) GetConnectorGroupIdOk() (*int64, bool)`

GetConnectorGroupIdOk returns a tuple with the ConnectorGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGroupId

`func (o *GetConnectionBandwidthResponseBody) SetConnectorGroupId(v int64)`

SetConnectorGroupId sets ConnectorGroupId field to given value.

### HasConnectorGroupId

`func (o *GetConnectionBandwidthResponseBody) HasConnectorGroupId() bool`

HasConnectorGroupId returns a boolean if a field has been set.

### SetConnectorGroupIdNil

`func (o *GetConnectionBandwidthResponseBody) SetConnectorGroupIdNil(b bool)`

 SetConnectorGroupIdNil sets the value for ConnectorGroupId to be an explicit nil

### UnsetConnectorGroupId
`func (o *GetConnectionBandwidthResponseBody) UnsetConnectorGroupId()`

UnsetConnectorGroupId ensures that no value is present for ConnectorGroupId, not even an explicit nil
### GetDownload

`func (o *GetConnectionBandwidthResponseBody) GetDownload() []BandwidthLimit`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *GetConnectionBandwidthResponseBody) GetDownloadOk() (*[]BandwidthLimit, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *GetConnectionBandwidthResponseBody) SetDownload(v []BandwidthLimit)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *GetConnectionBandwidthResponseBody) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### SetDownloadNil

`func (o *GetConnectionBandwidthResponseBody) SetDownloadNil(b bool)`

 SetDownloadNil sets the value for Download to be an explicit nil

### UnsetDownload
`func (o *GetConnectionBandwidthResponseBody) UnsetDownload()`

UnsetDownload ensures that no value is present for Download, not even an explicit nil
### GetTenantId

`func (o *GetConnectionBandwidthResponseBody) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetConnectionBandwidthResponseBody) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetConnectionBandwidthResponseBody) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *GetConnectionBandwidthResponseBody) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *GetConnectionBandwidthResponseBody) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTimezone

`func (o *GetConnectionBandwidthResponseBody) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetConnectionBandwidthResponseBody) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetConnectionBandwidthResponseBody) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetConnectionBandwidthResponseBody) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *GetConnectionBandwidthResponseBody) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *GetConnectionBandwidthResponseBody) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetUpload

`func (o *GetConnectionBandwidthResponseBody) GetUpload() []BandwidthLimit`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *GetConnectionBandwidthResponseBody) GetUploadOk() (*[]BandwidthLimit, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *GetConnectionBandwidthResponseBody) SetUpload(v []BandwidthLimit)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *GetConnectionBandwidthResponseBody) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### SetUploadNil

`func (o *GetConnectionBandwidthResponseBody) SetUploadNil(b bool)`

 SetUploadNil sets the value for Upload to be an explicit nil

### UnsetUpload
`func (o *GetConnectionBandwidthResponseBody) UnsetUpload()`

UnsetUpload ensures that no value is present for Upload, not even an explicit nil
### GetConnectionId

`func (o *GetConnectionBandwidthResponseBody) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GetConnectionBandwidthResponseBody) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GetConnectionBandwidthResponseBody) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *GetConnectionBandwidthResponseBody) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *GetConnectionBandwidthResponseBody) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *GetConnectionBandwidthResponseBody) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetConnectorGroups

`func (o *GetConnectionBandwidthResponseBody) GetConnectorGroups() []ConnectorGroup`

GetConnectorGroups returns the ConnectorGroups field if non-nil, zero value otherwise.

### GetConnectorGroupsOk

`func (o *GetConnectionBandwidthResponseBody) GetConnectorGroupsOk() (*[]ConnectorGroup, bool)`

GetConnectorGroupsOk returns a tuple with the ConnectorGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGroups

`func (o *GetConnectionBandwidthResponseBody) SetConnectorGroups(v []ConnectorGroup)`

SetConnectorGroups sets ConnectorGroups field to given value.

### HasConnectorGroups

`func (o *GetConnectionBandwidthResponseBody) HasConnectorGroups() bool`

HasConnectorGroups returns a boolean if a field has been set.

### SetConnectorGroupsNil

`func (o *GetConnectionBandwidthResponseBody) SetConnectorGroupsNil(b bool)`

 SetConnectorGroupsNil sets the value for ConnectorGroups to be an explicit nil

### UnsetConnectorGroups
`func (o *GetConnectionBandwidthResponseBody) UnsetConnectorGroups()`

UnsetConnectorGroups ensures that no value is present for ConnectorGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


