# ClusterNodeOperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]OperationEvents**](OperationEvents.md) | Specifies the list of events that took place during the operation.  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the node. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the IP address of the node. | [optional] 
**Percentage** | Pointer to **int32** | Specifies the approximate completion percentage for the operation. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of the operation. &#x60;Success&#x60; indicates the operation is successful. &#x60;Failed&#x60; indicates the operation failed due to an error. &#x60;InProgress&#x60; indicates the operation is in progress.  | [optional] 

## Methods

### NewClusterNodeOperationStatus

`func NewClusterNodeOperationStatus() *ClusterNodeOperationStatus`

NewClusterNodeOperationStatus instantiates a new ClusterNodeOperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeOperationStatusWithDefaults

`func NewClusterNodeOperationStatusWithDefaults() *ClusterNodeOperationStatus`

NewClusterNodeOperationStatusWithDefaults instantiates a new ClusterNodeOperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ClusterNodeOperationStatus) GetEvents() []OperationEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ClusterNodeOperationStatus) GetEventsOk() (*[]OperationEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ClusterNodeOperationStatus) SetEvents(v []OperationEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ClusterNodeOperationStatus) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *ClusterNodeOperationStatus) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterNodeOperationStatus) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterNodeOperationStatus) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterNodeOperationStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ClusterNodeOperationStatus) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ClusterNodeOperationStatus) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIp

`func (o *ClusterNodeOperationStatus) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClusterNodeOperationStatus) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClusterNodeOperationStatus) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClusterNodeOperationStatus) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ClusterNodeOperationStatus) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ClusterNodeOperationStatus) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetPercentage

`func (o *ClusterNodeOperationStatus) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ClusterNodeOperationStatus) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ClusterNodeOperationStatus) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ClusterNodeOperationStatus) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterNodeOperationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterNodeOperationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterNodeOperationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterNodeOperationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ClusterNodeOperationStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ClusterNodeOperationStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


