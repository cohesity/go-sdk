# ServicesStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadedImages** | Pointer to **[]string** | Images loaded by the service. | [optional] 
**ServiceName** | Pointer to **string** | Name of the service. | [optional] 
**ServiceStatus** | Pointer to **string** | Images and Status of ondemand service. | [optional] 
**SupportedImages** | Pointer to **[]string** | Images supported by the service. | [optional] 

## Methods

### NewServicesStatusResponse

`func NewServicesStatusResponse() *ServicesStatusResponse`

NewServicesStatusResponse instantiates a new ServicesStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesStatusResponseWithDefaults

`func NewServicesStatusResponseWithDefaults() *ServicesStatusResponse`

NewServicesStatusResponseWithDefaults instantiates a new ServicesStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadedImages

`func (o *ServicesStatusResponse) GetLoadedImages() []string`

GetLoadedImages returns the LoadedImages field if non-nil, zero value otherwise.

### GetLoadedImagesOk

`func (o *ServicesStatusResponse) GetLoadedImagesOk() (*[]string, bool)`

GetLoadedImagesOk returns a tuple with the LoadedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedImages

`func (o *ServicesStatusResponse) SetLoadedImages(v []string)`

SetLoadedImages sets LoadedImages field to given value.

### HasLoadedImages

`func (o *ServicesStatusResponse) HasLoadedImages() bool`

HasLoadedImages returns a boolean if a field has been set.

### GetServiceName

`func (o *ServicesStatusResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServicesStatusResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServicesStatusResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServicesStatusResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceStatus

`func (o *ServicesStatusResponse) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServicesStatusResponse) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServicesStatusResponse) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *ServicesStatusResponse) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetSupportedImages

`func (o *ServicesStatusResponse) GetSupportedImages() []string`

GetSupportedImages returns the SupportedImages field if non-nil, zero value otherwise.

### GetSupportedImagesOk

`func (o *ServicesStatusResponse) GetSupportedImagesOk() (*[]string, bool)`

GetSupportedImagesOk returns a tuple with the SupportedImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedImages

`func (o *ServicesStatusResponse) SetSupportedImages(v []string)`

SetSupportedImages sets SupportedImages field to given value.

### HasSupportedImages

`func (o *ServicesStatusResponse) HasSupportedImages() bool`

HasSupportedImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


