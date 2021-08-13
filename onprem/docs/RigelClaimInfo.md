# RigelClaimInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** | Specifies the registration status. | [optional] 
**Message** | Pointer to **NullableString** | Specifies possible error message during registration. | [optional] 

## Methods

### NewRigelClaimInfo

`func NewRigelClaimInfo() *RigelClaimInfo`

NewRigelClaimInfo instantiates a new RigelClaimInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRigelClaimInfoWithDefaults

`func NewRigelClaimInfoWithDefaults() *RigelClaimInfo`

NewRigelClaimInfoWithDefaults instantiates a new RigelClaimInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RigelClaimInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RigelClaimInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RigelClaimInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RigelClaimInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RigelClaimInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RigelClaimInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetMessage

`func (o *RigelClaimInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RigelClaimInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RigelClaimInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RigelClaimInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *RigelClaimInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *RigelClaimInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


