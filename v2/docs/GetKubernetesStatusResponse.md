# GetKubernetesStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentKubernetesVersion** | Pointer to **string** | specifies the current Kubernetes version | [optional] 
**OverallK8SHealthStatus** | Pointer to **string** | Kubernetes Infra Overall Health Status | [optional] 

## Methods

### NewGetKubernetesStatusResponse

`func NewGetKubernetesStatusResponse() *GetKubernetesStatusResponse`

NewGetKubernetesStatusResponse instantiates a new GetKubernetesStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesStatusResponseWithDefaults

`func NewGetKubernetesStatusResponseWithDefaults() *GetKubernetesStatusResponse`

NewGetKubernetesStatusResponseWithDefaults instantiates a new GetKubernetesStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentKubernetesVersion

`func (o *GetKubernetesStatusResponse) GetCurrentKubernetesVersion() string`

GetCurrentKubernetesVersion returns the CurrentKubernetesVersion field if non-nil, zero value otherwise.

### GetCurrentKubernetesVersionOk

`func (o *GetKubernetesStatusResponse) GetCurrentKubernetesVersionOk() (*string, bool)`

GetCurrentKubernetesVersionOk returns a tuple with the CurrentKubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKubernetesVersion

`func (o *GetKubernetesStatusResponse) SetCurrentKubernetesVersion(v string)`

SetCurrentKubernetesVersion sets CurrentKubernetesVersion field to given value.

### HasCurrentKubernetesVersion

`func (o *GetKubernetesStatusResponse) HasCurrentKubernetesVersion() bool`

HasCurrentKubernetesVersion returns a boolean if a field has been set.

### GetOverallK8SHealthStatus

`func (o *GetKubernetesStatusResponse) GetOverallK8SHealthStatus() string`

GetOverallK8SHealthStatus returns the OverallK8SHealthStatus field if non-nil, zero value otherwise.

### GetOverallK8SHealthStatusOk

`func (o *GetKubernetesStatusResponse) GetOverallK8SHealthStatusOk() (*string, bool)`

GetOverallK8SHealthStatusOk returns a tuple with the OverallK8SHealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallK8SHealthStatus

`func (o *GetKubernetesStatusResponse) SetOverallK8SHealthStatus(v string)`

SetOverallK8SHealthStatus sets OverallK8SHealthStatus field to given value.

### HasOverallK8SHealthStatus

`func (o *GetKubernetesStatusResponse) HasOverallK8SHealthStatus() bool`

HasOverallK8SHealthStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


