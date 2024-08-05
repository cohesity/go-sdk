# ServiceInstallLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Install message. | [optional] 
**ServiceInstallStatus** | Pointer to **string** | \&quot;The install status of services\&quot; \&quot;(e.g., Success, InProgress, Failed, Pending).\&quot;  | [optional] 
**ServiceName** | Pointer to **string** | The name of the service. | [optional] 

## Methods

### NewServiceInstallLog

`func NewServiceInstallLog() *ServiceInstallLog`

NewServiceInstallLog instantiates a new ServiceInstallLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstallLogWithDefaults

`func NewServiceInstallLogWithDefaults() *ServiceInstallLog`

NewServiceInstallLogWithDefaults instantiates a new ServiceInstallLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ServiceInstallLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceInstallLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceInstallLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServiceInstallLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetServiceInstallStatus

`func (o *ServiceInstallLog) GetServiceInstallStatus() string`

GetServiceInstallStatus returns the ServiceInstallStatus field if non-nil, zero value otherwise.

### GetServiceInstallStatusOk

`func (o *ServiceInstallLog) GetServiceInstallStatusOk() (*string, bool)`

GetServiceInstallStatusOk returns a tuple with the ServiceInstallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstallStatus

`func (o *ServiceInstallLog) SetServiceInstallStatus(v string)`

SetServiceInstallStatus sets ServiceInstallStatus field to given value.

### HasServiceInstallStatus

`func (o *ServiceInstallLog) HasServiceInstallStatus() bool`

HasServiceInstallStatus returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceInstallLog) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceInstallLog) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceInstallLog) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceInstallLog) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


