# UpdateInfectedFileParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfectedFileIds** | Pointer to [**[]InfectedFileParam**](InfectedFileParam.md) | Specifies the list of infected file identifiers. | [optional] 
**RemediationState** | Pointer to **NullableString** | Specifies the remediation state of the file. Not setting any value to remediation state will reset the infected file. Remediation State. &#39;kQuarantine&#39; indicates &#39;Quarantine&#39; state of the file. This state blocks the client access. The administrator will have to manually delete, rescan or unquarantine the file. &#39;kUnquarantine&#39; indicates &#39;Unquarantine&#39; state of the file. The administrator has manually moved files from quarantined to the unquarantined state to allow client access. Unquarantined files are not scanned for virus until manually reset. | [optional] 

## Methods

### NewUpdateInfectedFileParams

`func NewUpdateInfectedFileParams() *UpdateInfectedFileParams`

NewUpdateInfectedFileParams instantiates a new UpdateInfectedFileParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfectedFileParamsWithDefaults

`func NewUpdateInfectedFileParamsWithDefaults() *UpdateInfectedFileParams`

NewUpdateInfectedFileParamsWithDefaults instantiates a new UpdateInfectedFileParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfectedFileIds

`func (o *UpdateInfectedFileParams) GetInfectedFileIds() []InfectedFileParam`

GetInfectedFileIds returns the InfectedFileIds field if non-nil, zero value otherwise.

### GetInfectedFileIdsOk

`func (o *UpdateInfectedFileParams) GetInfectedFileIdsOk() (*[]InfectedFileParam, bool)`

GetInfectedFileIdsOk returns a tuple with the InfectedFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfectedFileIds

`func (o *UpdateInfectedFileParams) SetInfectedFileIds(v []InfectedFileParam)`

SetInfectedFileIds sets InfectedFileIds field to given value.

### HasInfectedFileIds

`func (o *UpdateInfectedFileParams) HasInfectedFileIds() bool`

HasInfectedFileIds returns a boolean if a field has been set.

### SetInfectedFileIdsNil

`func (o *UpdateInfectedFileParams) SetInfectedFileIdsNil(b bool)`

 SetInfectedFileIdsNil sets the value for InfectedFileIds to be an explicit nil

### UnsetInfectedFileIds
`func (o *UpdateInfectedFileParams) UnsetInfectedFileIds()`

UnsetInfectedFileIds ensures that no value is present for InfectedFileIds, not even an explicit nil
### GetRemediationState

`func (o *UpdateInfectedFileParams) GetRemediationState() string`

GetRemediationState returns the RemediationState field if non-nil, zero value otherwise.

### GetRemediationStateOk

`func (o *UpdateInfectedFileParams) GetRemediationStateOk() (*string, bool)`

GetRemediationStateOk returns a tuple with the RemediationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationState

`func (o *UpdateInfectedFileParams) SetRemediationState(v string)`

SetRemediationState sets RemediationState field to given value.

### HasRemediationState

`func (o *UpdateInfectedFileParams) HasRemediationState() bool`

HasRemediationState returns a boolean if a field has been set.

### SetRemediationStateNil

`func (o *UpdateInfectedFileParams) SetRemediationStateNil(b bool)`

 SetRemediationStateNil sets the value for RemediationState to be an explicit nil

### UnsetRemediationState
`func (o *UpdateInfectedFileParams) UnsetRemediationState()`

UnsetRemediationState ensures that no value is present for RemediationState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


