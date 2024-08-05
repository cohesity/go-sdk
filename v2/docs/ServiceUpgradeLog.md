# ServiceUpgradeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Messages related to the upgrade. | [optional] 
**ServiceName** | Pointer to **string** | The name of the service. | [optional] 
**ServiceUpgradeStatus** | Pointer to **string** | \&quot;The upgrade status of services\&quot; \&quot;(e.g., Success, InProgress, Failed, Pending).\&quot;  | [optional] 

## Methods

### NewServiceUpgradeLog

`func NewServiceUpgradeLog() *ServiceUpgradeLog`

NewServiceUpgradeLog instantiates a new ServiceUpgradeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceUpgradeLogWithDefaults

`func NewServiceUpgradeLogWithDefaults() *ServiceUpgradeLog`

NewServiceUpgradeLogWithDefaults instantiates a new ServiceUpgradeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ServiceUpgradeLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceUpgradeLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceUpgradeLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServiceUpgradeLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceUpgradeLog) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceUpgradeLog) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceUpgradeLog) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceUpgradeLog) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceUpgradeStatus

`func (o *ServiceUpgradeLog) GetServiceUpgradeStatus() string`

GetServiceUpgradeStatus returns the ServiceUpgradeStatus field if non-nil, zero value otherwise.

### GetServiceUpgradeStatusOk

`func (o *ServiceUpgradeLog) GetServiceUpgradeStatusOk() (*string, bool)`

GetServiceUpgradeStatusOk returns a tuple with the ServiceUpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUpgradeStatus

`func (o *ServiceUpgradeLog) SetServiceUpgradeStatus(v string)`

SetServiceUpgradeStatus sets ServiceUpgradeStatus field to given value.

### HasServiceUpgradeStatus

`func (o *ServiceUpgradeLog) HasServiceUpgradeStatus() bool`

HasServiceUpgradeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


