# Management

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthStatus** | Pointer to **string** | Kubernetes Management Cluster Overall Health Status | [optional] 

## Methods

### NewManagement

`func NewManagement() *Management`

NewManagement instantiates a new Management object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementWithDefaults

`func NewManagementWithDefaults() *Management`

NewManagementWithDefaults instantiates a new Management object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthStatus

`func (o *Management) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *Management) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *Management) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *Management) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


