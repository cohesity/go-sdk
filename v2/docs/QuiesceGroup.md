# QuiesceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuiesceMode** | **NullableString** | Specifies quiesce mode for applying quiesce rules. | 
**QuiesceRules** | [**[]QuiesceRule**](QuiesceRule.md) | Specifies a list of quiesce rules. | 

## Methods

### NewQuiesceGroup

`func NewQuiesceGroup(quiesceMode NullableString, quiesceRules []QuiesceRule, ) *QuiesceGroup`

NewQuiesceGroup instantiates a new QuiesceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuiesceGroupWithDefaults

`func NewQuiesceGroupWithDefaults() *QuiesceGroup`

NewQuiesceGroupWithDefaults instantiates a new QuiesceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuiesceMode

`func (o *QuiesceGroup) GetQuiesceMode() string`

GetQuiesceMode returns the QuiesceMode field if non-nil, zero value otherwise.

### GetQuiesceModeOk

`func (o *QuiesceGroup) GetQuiesceModeOk() (*string, bool)`

GetQuiesceModeOk returns a tuple with the QuiesceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesceMode

`func (o *QuiesceGroup) SetQuiesceMode(v string)`

SetQuiesceMode sets QuiesceMode field to given value.


### SetQuiesceModeNil

`func (o *QuiesceGroup) SetQuiesceModeNil(b bool)`

 SetQuiesceModeNil sets the value for QuiesceMode to be an explicit nil

### UnsetQuiesceMode
`func (o *QuiesceGroup) UnsetQuiesceMode()`

UnsetQuiesceMode ensures that no value is present for QuiesceMode, not even an explicit nil
### GetQuiesceRules

`func (o *QuiesceGroup) GetQuiesceRules() []QuiesceRule`

GetQuiesceRules returns the QuiesceRules field if non-nil, zero value otherwise.

### GetQuiesceRulesOk

`func (o *QuiesceGroup) GetQuiesceRulesOk() (*[]QuiesceRule, bool)`

GetQuiesceRulesOk returns a tuple with the QuiesceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiesceRules

`func (o *QuiesceGroup) SetQuiesceRules(v []QuiesceRule)`

SetQuiesceRules sets QuiesceRules field to given value.


### SetQuiesceRulesNil

`func (o *QuiesceGroup) SetQuiesceRulesNil(b bool)`

 SetQuiesceRulesNil sets the value for QuiesceRules to be an explicit nil

### UnsetQuiesceRules
`func (o *QuiesceGroup) UnsetQuiesceRules()`

UnsetQuiesceRules ensures that no value is present for QuiesceRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


