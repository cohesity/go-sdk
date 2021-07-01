# IcapConnectionStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedConnectionStatus** | Pointer to **[]string** | Specifies the failed connection status of Icap servers. | [optional] 
**SucceededConnectionStatus** | Pointer to **[]string** | Specifies the success connection status of Icap servers. | [optional] 

## Methods

### NewIcapConnectionStatusResponse

`func NewIcapConnectionStatusResponse() *IcapConnectionStatusResponse`

NewIcapConnectionStatusResponse instantiates a new IcapConnectionStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIcapConnectionStatusResponseWithDefaults

`func NewIcapConnectionStatusResponseWithDefaults() *IcapConnectionStatusResponse`

NewIcapConnectionStatusResponseWithDefaults instantiates a new IcapConnectionStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedConnectionStatus

`func (o *IcapConnectionStatusResponse) GetFailedConnectionStatus() []string`

GetFailedConnectionStatus returns the FailedConnectionStatus field if non-nil, zero value otherwise.

### GetFailedConnectionStatusOk

`func (o *IcapConnectionStatusResponse) GetFailedConnectionStatusOk() (*[]string, bool)`

GetFailedConnectionStatusOk returns a tuple with the FailedConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedConnectionStatus

`func (o *IcapConnectionStatusResponse) SetFailedConnectionStatus(v []string)`

SetFailedConnectionStatus sets FailedConnectionStatus field to given value.

### HasFailedConnectionStatus

`func (o *IcapConnectionStatusResponse) HasFailedConnectionStatus() bool`

HasFailedConnectionStatus returns a boolean if a field has been set.

### SetFailedConnectionStatusNil

`func (o *IcapConnectionStatusResponse) SetFailedConnectionStatusNil(b bool)`

 SetFailedConnectionStatusNil sets the value for FailedConnectionStatus to be an explicit nil

### UnsetFailedConnectionStatus
`func (o *IcapConnectionStatusResponse) UnsetFailedConnectionStatus()`

UnsetFailedConnectionStatus ensures that no value is present for FailedConnectionStatus, not even an explicit nil
### GetSucceededConnectionStatus

`func (o *IcapConnectionStatusResponse) GetSucceededConnectionStatus() []string`

GetSucceededConnectionStatus returns the SucceededConnectionStatus field if non-nil, zero value otherwise.

### GetSucceededConnectionStatusOk

`func (o *IcapConnectionStatusResponse) GetSucceededConnectionStatusOk() (*[]string, bool)`

GetSucceededConnectionStatusOk returns a tuple with the SucceededConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededConnectionStatus

`func (o *IcapConnectionStatusResponse) SetSucceededConnectionStatus(v []string)`

SetSucceededConnectionStatus sets SucceededConnectionStatus field to given value.

### HasSucceededConnectionStatus

`func (o *IcapConnectionStatusResponse) HasSucceededConnectionStatus() bool`

HasSucceededConnectionStatus returns a boolean if a field has been set.

### SetSucceededConnectionStatusNil

`func (o *IcapConnectionStatusResponse) SetSucceededConnectionStatusNil(b bool)`

 SetSucceededConnectionStatusNil sets the value for SucceededConnectionStatus to be an explicit nil

### UnsetSucceededConnectionStatus
`func (o *IcapConnectionStatusResponse) UnsetSucceededConnectionStatus()`

UnsetSucceededConnectionStatus ensures that no value is present for SucceededConnectionStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


