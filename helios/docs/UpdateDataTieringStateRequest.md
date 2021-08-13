# UpdateDataTieringStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the action to be performed on all the specified data tiering groups. &#39;Pause&#39;  specifies to pause. &#39;Resume&#39; specifies to resume. | 
**Ids** | **[]string** | Specifies a list of data tiering groups ids for which the state should change. | 

## Methods

### NewUpdateDataTieringStateRequest

`func NewUpdateDataTieringStateRequest(action NullableString, ids []string, ) *UpdateDataTieringStateRequest`

NewUpdateDataTieringStateRequest instantiates a new UpdateDataTieringStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataTieringStateRequestWithDefaults

`func NewUpdateDataTieringStateRequestWithDefaults() *UpdateDataTieringStateRequest`

NewUpdateDataTieringStateRequestWithDefaults instantiates a new UpdateDataTieringStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateDataTieringStateRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateDataTieringStateRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateDataTieringStateRequest) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *UpdateDataTieringStateRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *UpdateDataTieringStateRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetIds

`func (o *UpdateDataTieringStateRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UpdateDataTieringStateRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UpdateDataTieringStateRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### SetIdsNil

`func (o *UpdateDataTieringStateRequest) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *UpdateDataTieringStateRequest) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


