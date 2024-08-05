# UpgradeLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeliosUpgradeStatus** | Pointer to **string** | \&quot;The overall upgrade status\&quot; \&quot;(e.g., Success, InProgress, Failed, Pending).\&quot;  | [optional] 
**HeliosUpgradeVersion** | Pointer to **string** | Helios upgrade version. | [optional] 
**Services** | Pointer to [**[]ServiceUpgradeLog**](ServiceUpgradeLog.md) | List of service upgrade logs. | [optional] 

## Methods

### NewUpgradeLogsResponse

`func NewUpgradeLogsResponse() *UpgradeLogsResponse`

NewUpgradeLogsResponse instantiates a new UpgradeLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeLogsResponseWithDefaults

`func NewUpgradeLogsResponseWithDefaults() *UpgradeLogsResponse`

NewUpgradeLogsResponseWithDefaults instantiates a new UpgradeLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeliosUpgradeStatus

`func (o *UpgradeLogsResponse) GetHeliosUpgradeStatus() string`

GetHeliosUpgradeStatus returns the HeliosUpgradeStatus field if non-nil, zero value otherwise.

### GetHeliosUpgradeStatusOk

`func (o *UpgradeLogsResponse) GetHeliosUpgradeStatusOk() (*string, bool)`

GetHeliosUpgradeStatusOk returns a tuple with the HeliosUpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeliosUpgradeStatus

`func (o *UpgradeLogsResponse) SetHeliosUpgradeStatus(v string)`

SetHeliosUpgradeStatus sets HeliosUpgradeStatus field to given value.

### HasHeliosUpgradeStatus

`func (o *UpgradeLogsResponse) HasHeliosUpgradeStatus() bool`

HasHeliosUpgradeStatus returns a boolean if a field has been set.

### GetHeliosUpgradeVersion

`func (o *UpgradeLogsResponse) GetHeliosUpgradeVersion() string`

GetHeliosUpgradeVersion returns the HeliosUpgradeVersion field if non-nil, zero value otherwise.

### GetHeliosUpgradeVersionOk

`func (o *UpgradeLogsResponse) GetHeliosUpgradeVersionOk() (*string, bool)`

GetHeliosUpgradeVersionOk returns a tuple with the HeliosUpgradeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeliosUpgradeVersion

`func (o *UpgradeLogsResponse) SetHeliosUpgradeVersion(v string)`

SetHeliosUpgradeVersion sets HeliosUpgradeVersion field to given value.

### HasHeliosUpgradeVersion

`func (o *UpgradeLogsResponse) HasHeliosUpgradeVersion() bool`

HasHeliosUpgradeVersion returns a boolean if a field has been set.

### GetServices

`func (o *UpgradeLogsResponse) GetServices() []ServiceUpgradeLog`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *UpgradeLogsResponse) GetServicesOk() (*[]ServiceUpgradeLog, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *UpgradeLogsResponse) SetServices(v []ServiceUpgradeLog)`

SetServices sets Services field to given value.

### HasServices

`func (o *UpgradeLogsResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


