# WorkerEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** | Specifies the end point of the worker which is combination of both endpoint and port number. | [optional] 

## Methods

### NewWorkerEndpoint

`func NewWorkerEndpoint() *WorkerEndpoint`

NewWorkerEndpoint instantiates a new WorkerEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkerEndpointWithDefaults

`func NewWorkerEndpointWithDefaults() *WorkerEndpoint`

NewWorkerEndpointWithDefaults instantiates a new WorkerEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *WorkerEndpoint) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WorkerEndpoint) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WorkerEndpoint) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *WorkerEndpoint) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


