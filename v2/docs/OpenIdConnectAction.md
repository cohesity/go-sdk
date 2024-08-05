# OpenIdConnectAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies an action to perform on an Open ID Connect Identity Provider. The following actions are currently supported: 1. &#39;RefreshPublicKeys&#39;: Refreshes the public keys currently stored on the cluster for the user sending the request. In order to do this, the public key URL specified in the current users Open ID configuration will be polled for a new public key. | 

## Methods

### NewOpenIdConnectAction

`func NewOpenIdConnectAction(action NullableString, ) *OpenIdConnectAction`

NewOpenIdConnectAction instantiates a new OpenIdConnectAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectActionWithDefaults

`func NewOpenIdConnectActionWithDefaults() *OpenIdConnectAction`

NewOpenIdConnectActionWithDefaults instantiates a new OpenIdConnectAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OpenIdConnectAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OpenIdConnectAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OpenIdConnectAction) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *OpenIdConnectAction) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *OpenIdConnectAction) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


