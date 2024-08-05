# AzureTargetParamsForRecoverAzureEntraId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangePasswordOnSignIn** | Pointer to **NullableBool** | Specifies Whether to force changing password on next sign in. Applies for only user type | [optional] 
**DefaultPassword** | Pointer to **NullableString** | Specifies the default password to be set in case users are recovered. | [optional] 
**IsContainerMemberRecovery** | Pointer to **NullableBool** | Specifies whether recovery should recover all the members of a selected object. Applies for Admin Unit, Group and Dir Roles | [optional] 
**IsRelationOverwrite** | Pointer to **NullableBool** | If true, relationships are restored in overwrite mode (i.e. any relationship created *after* restore point is deleted). For example, if user U1 is member of G1 and G2 in restore point (in selected snapshot), but the same user is member of G1 and G3 in live AAD, membership of G2 is restored and membership of G3 is removed in overwrite-mode while restoring U1. If this field is false, relationship is restored in merge-mode. In merge-mode, membership of G2 is restored but membership of G3 is not removed. | [optional] 
**RecoveryType** | **string** | Specifies the recovery type for the selected azure entra id recoverable object. | 

## Methods

### NewAzureTargetParamsForRecoverAzureEntraId

`func NewAzureTargetParamsForRecoverAzureEntraId(recoveryType string, ) *AzureTargetParamsForRecoverAzureEntraId`

NewAzureTargetParamsForRecoverAzureEntraId instantiates a new AzureTargetParamsForRecoverAzureEntraId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTargetParamsForRecoverAzureEntraIdWithDefaults

`func NewAzureTargetParamsForRecoverAzureEntraIdWithDefaults() *AzureTargetParamsForRecoverAzureEntraId`

NewAzureTargetParamsForRecoverAzureEntraIdWithDefaults instantiates a new AzureTargetParamsForRecoverAzureEntraId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangePasswordOnSignIn

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetChangePasswordOnSignIn() bool`

GetChangePasswordOnSignIn returns the ChangePasswordOnSignIn field if non-nil, zero value otherwise.

### GetChangePasswordOnSignInOk

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetChangePasswordOnSignInOk() (*bool, bool)`

GetChangePasswordOnSignInOk returns a tuple with the ChangePasswordOnSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePasswordOnSignIn

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetChangePasswordOnSignIn(v bool)`

SetChangePasswordOnSignIn sets ChangePasswordOnSignIn field to given value.

### HasChangePasswordOnSignIn

`func (o *AzureTargetParamsForRecoverAzureEntraId) HasChangePasswordOnSignIn() bool`

HasChangePasswordOnSignIn returns a boolean if a field has been set.

### SetChangePasswordOnSignInNil

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetChangePasswordOnSignInNil(b bool)`

 SetChangePasswordOnSignInNil sets the value for ChangePasswordOnSignIn to be an explicit nil

### UnsetChangePasswordOnSignIn
`func (o *AzureTargetParamsForRecoverAzureEntraId) UnsetChangePasswordOnSignIn()`

UnsetChangePasswordOnSignIn ensures that no value is present for ChangePasswordOnSignIn, not even an explicit nil
### GetDefaultPassword

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetDefaultPassword() string`

GetDefaultPassword returns the DefaultPassword field if non-nil, zero value otherwise.

### GetDefaultPasswordOk

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetDefaultPasswordOk() (*string, bool)`

GetDefaultPasswordOk returns a tuple with the DefaultPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPassword

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetDefaultPassword(v string)`

SetDefaultPassword sets DefaultPassword field to given value.

### HasDefaultPassword

`func (o *AzureTargetParamsForRecoverAzureEntraId) HasDefaultPassword() bool`

HasDefaultPassword returns a boolean if a field has been set.

### SetDefaultPasswordNil

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetDefaultPasswordNil(b bool)`

 SetDefaultPasswordNil sets the value for DefaultPassword to be an explicit nil

### UnsetDefaultPassword
`func (o *AzureTargetParamsForRecoverAzureEntraId) UnsetDefaultPassword()`

UnsetDefaultPassword ensures that no value is present for DefaultPassword, not even an explicit nil
### GetIsContainerMemberRecovery

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetIsContainerMemberRecovery() bool`

GetIsContainerMemberRecovery returns the IsContainerMemberRecovery field if non-nil, zero value otherwise.

### GetIsContainerMemberRecoveryOk

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetIsContainerMemberRecoveryOk() (*bool, bool)`

GetIsContainerMemberRecoveryOk returns a tuple with the IsContainerMemberRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainerMemberRecovery

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetIsContainerMemberRecovery(v bool)`

SetIsContainerMemberRecovery sets IsContainerMemberRecovery field to given value.

### HasIsContainerMemberRecovery

`func (o *AzureTargetParamsForRecoverAzureEntraId) HasIsContainerMemberRecovery() bool`

HasIsContainerMemberRecovery returns a boolean if a field has been set.

### SetIsContainerMemberRecoveryNil

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetIsContainerMemberRecoveryNil(b bool)`

 SetIsContainerMemberRecoveryNil sets the value for IsContainerMemberRecovery to be an explicit nil

### UnsetIsContainerMemberRecovery
`func (o *AzureTargetParamsForRecoverAzureEntraId) UnsetIsContainerMemberRecovery()`

UnsetIsContainerMemberRecovery ensures that no value is present for IsContainerMemberRecovery, not even an explicit nil
### GetIsRelationOverwrite

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetIsRelationOverwrite() bool`

GetIsRelationOverwrite returns the IsRelationOverwrite field if non-nil, zero value otherwise.

### GetIsRelationOverwriteOk

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetIsRelationOverwriteOk() (*bool, bool)`

GetIsRelationOverwriteOk returns a tuple with the IsRelationOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRelationOverwrite

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetIsRelationOverwrite(v bool)`

SetIsRelationOverwrite sets IsRelationOverwrite field to given value.

### HasIsRelationOverwrite

`func (o *AzureTargetParamsForRecoverAzureEntraId) HasIsRelationOverwrite() bool`

HasIsRelationOverwrite returns a boolean if a field has been set.

### SetIsRelationOverwriteNil

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetIsRelationOverwriteNil(b bool)`

 SetIsRelationOverwriteNil sets the value for IsRelationOverwrite to be an explicit nil

### UnsetIsRelationOverwrite
`func (o *AzureTargetParamsForRecoverAzureEntraId) UnsetIsRelationOverwrite()`

UnsetIsRelationOverwrite ensures that no value is present for IsRelationOverwrite, not even an explicit nil
### GetRecoveryType

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetRecoveryType() string`

GetRecoveryType returns the RecoveryType field if non-nil, zero value otherwise.

### GetRecoveryTypeOk

`func (o *AzureTargetParamsForRecoverAzureEntraId) GetRecoveryTypeOk() (*string, bool)`

GetRecoveryTypeOk returns a tuple with the RecoveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryType

`func (o *AzureTargetParamsForRecoverAzureEntraId) SetRecoveryType(v string)`

SetRecoveryType sets RecoveryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


