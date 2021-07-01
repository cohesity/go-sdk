# UpgradePhysicalAgentsMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies the status message returned after initiating an upgrade request. Status of each agent upgrade can be obtained by listing Physical Servers using the GET /public/protectionSources operation. | [optional] 

## Methods

### NewUpgradePhysicalAgentsMessage

`func NewUpgradePhysicalAgentsMessage() *UpgradePhysicalAgentsMessage`

NewUpgradePhysicalAgentsMessage instantiates a new UpgradePhysicalAgentsMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradePhysicalAgentsMessageWithDefaults

`func NewUpgradePhysicalAgentsMessageWithDefaults() *UpgradePhysicalAgentsMessage`

NewUpgradePhysicalAgentsMessageWithDefaults instantiates a new UpgradePhysicalAgentsMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpgradePhysicalAgentsMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpgradePhysicalAgentsMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpgradePhysicalAgentsMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpgradePhysicalAgentsMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *UpgradePhysicalAgentsMessage) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UpgradePhysicalAgentsMessage) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


