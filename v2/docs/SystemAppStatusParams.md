# SystemAppStatusParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | Pointer to **int32** | Specifies the number of available replicas. | [optional] 
**ConfiguredReplicas** | Pointer to **int32** | Specifies the number of configured replicas. | [optional] 
**Name** | Pointer to **string** | Specifies the name of the system app. | [optional] 
**ReadyReplicas** | Pointer to **int32** | Specifies the number of ready replicas. | [optional] 
**ServiceEndpoint** | Pointer to **string** | Specifies the service endpoint. | [optional] 

## Methods

### NewSystemAppStatusParams

`func NewSystemAppStatusParams() *SystemAppStatusParams`

NewSystemAppStatusParams instantiates a new SystemAppStatusParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemAppStatusParamsWithDefaults

`func NewSystemAppStatusParamsWithDefaults() *SystemAppStatusParams`

NewSystemAppStatusParamsWithDefaults instantiates a new SystemAppStatusParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableReplicas

`func (o *SystemAppStatusParams) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *SystemAppStatusParams) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *SystemAppStatusParams) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *SystemAppStatusParams) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetConfiguredReplicas

`func (o *SystemAppStatusParams) GetConfiguredReplicas() int32`

GetConfiguredReplicas returns the ConfiguredReplicas field if non-nil, zero value otherwise.

### GetConfiguredReplicasOk

`func (o *SystemAppStatusParams) GetConfiguredReplicasOk() (*int32, bool)`

GetConfiguredReplicasOk returns a tuple with the ConfiguredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredReplicas

`func (o *SystemAppStatusParams) SetConfiguredReplicas(v int32)`

SetConfiguredReplicas sets ConfiguredReplicas field to given value.

### HasConfiguredReplicas

`func (o *SystemAppStatusParams) HasConfiguredReplicas() bool`

HasConfiguredReplicas returns a boolean if a field has been set.

### GetName

`func (o *SystemAppStatusParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemAppStatusParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemAppStatusParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemAppStatusParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *SystemAppStatusParams) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *SystemAppStatusParams) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *SystemAppStatusParams) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *SystemAppStatusParams) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetServiceEndpoint

`func (o *SystemAppStatusParams) GetServiceEndpoint() string`

GetServiceEndpoint returns the ServiceEndpoint field if non-nil, zero value otherwise.

### GetServiceEndpointOk

`func (o *SystemAppStatusParams) GetServiceEndpointOk() (*string, bool)`

GetServiceEndpointOk returns a tuple with the ServiceEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoint

`func (o *SystemAppStatusParams) SetServiceEndpoint(v string)`

SetServiceEndpoint sets ServiceEndpoint field to given value.

### HasServiceEndpoint

`func (o *SystemAppStatusParams) HasServiceEndpoint() bool`

HasServiceEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


