# LegalHold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **string** | The action type for legal hold on a protection run i.e. enable or release | [optional] 

## Methods

### NewLegalHold

`func NewLegalHold() *LegalHold`

NewLegalHold instantiates a new LegalHold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalHoldWithDefaults

`func NewLegalHoldWithDefaults() *LegalHold`

NewLegalHoldWithDefaults instantiates a new LegalHold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *LegalHold) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *LegalHold) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *LegalHold) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *LegalHold) HasActionType() bool`

HasActionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


