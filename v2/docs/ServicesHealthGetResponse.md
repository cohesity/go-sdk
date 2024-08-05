# ServicesHealthGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthStatus** | Pointer to **string** | \&quot;The overall health status of the services (e.g., Healthy, Degraded,\&quot; \&quot; Unhealthy).\&quot;  | [optional] 
**Services** | Pointer to [**[]ServiceHealth**](ServiceHealth.md) | List of services with their respective health statuses. | [optional] 

## Methods

### NewServicesHealthGetResponse

`func NewServicesHealthGetResponse() *ServicesHealthGetResponse`

NewServicesHealthGetResponse instantiates a new ServicesHealthGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesHealthGetResponseWithDefaults

`func NewServicesHealthGetResponseWithDefaults() *ServicesHealthGetResponse`

NewServicesHealthGetResponseWithDefaults instantiates a new ServicesHealthGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthStatus

`func (o *ServicesHealthGetResponse) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *ServicesHealthGetResponse) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *ServicesHealthGetResponse) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *ServicesHealthGetResponse) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetServices

`func (o *ServicesHealthGetResponse) GetServices() []ServiceHealth`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServicesHealthGetResponse) GetServicesOk() (*[]ServiceHealth, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServicesHealthGetResponse) SetServices(v []ServiceHealth)`

SetServices sets Services field to given value.

### HasServices

`func (o *ServicesHealthGetResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


