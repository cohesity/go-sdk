# ServiceActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | Pointer to **string** | Name of the service. | [optional] 
**ServiceStatus** | Pointer to **string** | Images and Status of ondemand service. | [optional] 

## Methods

### NewServiceActionResponse

`func NewServiceActionResponse() *ServiceActionResponse`

NewServiceActionResponse instantiates a new ServiceActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceActionResponseWithDefaults

`func NewServiceActionResponseWithDefaults() *ServiceActionResponse`

NewServiceActionResponseWithDefaults instantiates a new ServiceActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *ServiceActionResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceActionResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceActionResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceActionResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceStatus

`func (o *ServiceActionResponse) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServiceActionResponse) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServiceActionResponse) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *ServiceActionResponse) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


