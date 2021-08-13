# AddRemoteDiskResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteDisks** | Pointer to [**[]CreateRemoteDiskStatus**](CreateRemoteDiskStatus.md) | Specifies a list of remote disk creating status. | [optional] 

## Methods

### NewAddRemoteDiskResponseBody

`func NewAddRemoteDiskResponseBody() *AddRemoteDiskResponseBody`

NewAddRemoteDiskResponseBody instantiates a new AddRemoteDiskResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRemoteDiskResponseBodyWithDefaults

`func NewAddRemoteDiskResponseBodyWithDefaults() *AddRemoteDiskResponseBody`

NewAddRemoteDiskResponseBodyWithDefaults instantiates a new AddRemoteDiskResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteDisks

`func (o *AddRemoteDiskResponseBody) GetRemoteDisks() []CreateRemoteDiskStatus`

GetRemoteDisks returns the RemoteDisks field if non-nil, zero value otherwise.

### GetRemoteDisksOk

`func (o *AddRemoteDiskResponseBody) GetRemoteDisksOk() (*[]CreateRemoteDiskStatus, bool)`

GetRemoteDisksOk returns a tuple with the RemoteDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDisks

`func (o *AddRemoteDiskResponseBody) SetRemoteDisks(v []CreateRemoteDiskStatus)`

SetRemoteDisks sets RemoteDisks field to given value.

### HasRemoteDisks

`func (o *AddRemoteDiskResponseBody) HasRemoteDisks() bool`

HasRemoteDisks returns a boolean if a field has been set.

### SetRemoteDisksNil

`func (o *AddRemoteDiskResponseBody) SetRemoteDisksNil(b bool)`

 SetRemoteDisksNil sets the value for RemoteDisks to be an explicit nil

### UnsetRemoteDisks
`func (o *AddRemoteDiskResponseBody) UnsetRemoteDisks()`

UnsetRemoteDisks ensures that no value is present for RemoteDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


