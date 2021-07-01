# ActivateViewAliasesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to [**[]ViewAlias**](ViewAlias.md) | Aliases created for the view. A view alias allows a directory path inside a view to be mounted using the alias name. | [optional] 

## Methods

### NewActivateViewAliasesResult

`func NewActivateViewAliasesResult() *ActivateViewAliasesResult`

NewActivateViewAliasesResult instantiates a new ActivateViewAliasesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateViewAliasesResultWithDefaults

`func NewActivateViewAliasesResultWithDefaults() *ActivateViewAliasesResult`

NewActivateViewAliasesResultWithDefaults instantiates a new ActivateViewAliasesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *ActivateViewAliasesResult) GetAliases() []ViewAlias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ActivateViewAliasesResult) GetAliasesOk() (*[]ViewAlias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ActivateViewAliasesResult) SetAliases(v []ViewAlias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ActivateViewAliasesResult) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliasesNil

`func (o *ActivateViewAliasesResult) SetAliasesNil(b bool)`

 SetAliasesNil sets the value for Aliases to be an explicit nil

### UnsetAliases
`func (o *ActivateViewAliasesResult) UnsetAliases()`

UnsetAliases ensures that no value is present for Aliases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


