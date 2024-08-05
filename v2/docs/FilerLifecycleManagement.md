# FilerLifecycleManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]FilerLifecycleRule**](FilerLifecycleRule.md) | Specifies Lifecycle configuration rules for a filer view. A maximum of 100 rules can be specified. | [optional] 

## Methods

### NewFilerLifecycleManagement

`func NewFilerLifecycleManagement() *FilerLifecycleManagement`

NewFilerLifecycleManagement instantiates a new FilerLifecycleManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilerLifecycleManagementWithDefaults

`func NewFilerLifecycleManagementWithDefaults() *FilerLifecycleManagement`

NewFilerLifecycleManagementWithDefaults instantiates a new FilerLifecycleManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *FilerLifecycleManagement) GetRules() []FilerLifecycleRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FilerLifecycleManagement) GetRulesOk() (*[]FilerLifecycleRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FilerLifecycleManagement) SetRules(v []FilerLifecycleRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FilerLifecycleManagement) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *FilerLifecycleManagement) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *FilerLifecycleManagement) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


