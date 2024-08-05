# ServiceHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthStatus** | Pointer to **string** | \&quot;The health status of the service (e.g., Healthy, Degraded,\&quot; \&quot; Unhealthy).\&quot;  | [optional] 
**ServiceName** | Pointer to **string** | The name of the service. | [optional] 

## Methods

### NewServiceHealth

`func NewServiceHealth() *ServiceHealth`

NewServiceHealth instantiates a new ServiceHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHealthWithDefaults

`func NewServiceHealthWithDefaults() *ServiceHealth`

NewServiceHealthWithDefaults instantiates a new ServiceHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthStatus

`func (o *ServiceHealth) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *ServiceHealth) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *ServiceHealth) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *ServiceHealth) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceHealth) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceHealth) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceHealth) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceHealth) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


