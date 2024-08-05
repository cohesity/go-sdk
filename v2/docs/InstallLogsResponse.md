# InstallLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeliosInstallStatus** | Pointer to **string** | \&quot;The overall install status \&quot; \&quot;(e.g., Success, InProgress, Failed, Pending).\&quot;  | [optional] 
**HeliosInstallVersion** | Pointer to **string** | Helios install version. | [optional] 
**Services** | Pointer to [**[]ServiceInstallLog**](ServiceInstallLog.md) | List of service install logs. | [optional] 

## Methods

### NewInstallLogsResponse

`func NewInstallLogsResponse() *InstallLogsResponse`

NewInstallLogsResponse instantiates a new InstallLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallLogsResponseWithDefaults

`func NewInstallLogsResponseWithDefaults() *InstallLogsResponse`

NewInstallLogsResponseWithDefaults instantiates a new InstallLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeliosInstallStatus

`func (o *InstallLogsResponse) GetHeliosInstallStatus() string`

GetHeliosInstallStatus returns the HeliosInstallStatus field if non-nil, zero value otherwise.

### GetHeliosInstallStatusOk

`func (o *InstallLogsResponse) GetHeliosInstallStatusOk() (*string, bool)`

GetHeliosInstallStatusOk returns a tuple with the HeliosInstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeliosInstallStatus

`func (o *InstallLogsResponse) SetHeliosInstallStatus(v string)`

SetHeliosInstallStatus sets HeliosInstallStatus field to given value.

### HasHeliosInstallStatus

`func (o *InstallLogsResponse) HasHeliosInstallStatus() bool`

HasHeliosInstallStatus returns a boolean if a field has been set.

### GetHeliosInstallVersion

`func (o *InstallLogsResponse) GetHeliosInstallVersion() string`

GetHeliosInstallVersion returns the HeliosInstallVersion field if non-nil, zero value otherwise.

### GetHeliosInstallVersionOk

`func (o *InstallLogsResponse) GetHeliosInstallVersionOk() (*string, bool)`

GetHeliosInstallVersionOk returns a tuple with the HeliosInstallVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeliosInstallVersion

`func (o *InstallLogsResponse) SetHeliosInstallVersion(v string)`

SetHeliosInstallVersion sets HeliosInstallVersion field to given value.

### HasHeliosInstallVersion

`func (o *InstallLogsResponse) HasHeliosInstallVersion() bool`

HasHeliosInstallVersion returns a boolean if a field has been set.

### GetServices

`func (o *InstallLogsResponse) GetServices() []ServiceInstallLog`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InstallLogsResponse) GetServicesOk() (*[]ServiceInstallLog, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InstallLogsResponse) SetServices(v []ServiceInstallLog)`

SetServices sets Services field to given value.

### HasServices

`func (o *InstallLogsResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


