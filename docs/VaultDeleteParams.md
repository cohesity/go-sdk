# VaultDeleteParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceDelete** | Pointer to **NullableBool** | Specifies whether to force delete the vault. If the flag is set to true, the RemovalState of the vault is changed to &#39;kMarkedForRemoval&#39; and Eventually vault is removed from cluster config and archived metadata from scribe is removed without necessarily deleting the associated archived data. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies an id that identifies the Vault. | [optional] 
**IncludeMarkedForRemoval** | Pointer to **NullableBool** | Specifies if Vaults that are marked for removal should be returned. | [optional] 
**Retry** | Pointer to **NullableBool** | Specifies whether to retry a request after failure. | [optional] 

## Methods

### NewVaultDeleteParams

`func NewVaultDeleteParams() *VaultDeleteParams`

NewVaultDeleteParams instantiates a new VaultDeleteParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVaultDeleteParamsWithDefaults

`func NewVaultDeleteParamsWithDefaults() *VaultDeleteParams`

NewVaultDeleteParamsWithDefaults instantiates a new VaultDeleteParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceDelete

`func (o *VaultDeleteParams) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *VaultDeleteParams) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *VaultDeleteParams) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.

### HasForceDelete

`func (o *VaultDeleteParams) HasForceDelete() bool`

HasForceDelete returns a boolean if a field has been set.

### SetForceDeleteNil

`func (o *VaultDeleteParams) SetForceDeleteNil(b bool)`

 SetForceDeleteNil sets the value for ForceDelete to be an explicit nil

### UnsetForceDelete
`func (o *VaultDeleteParams) UnsetForceDelete()`

UnsetForceDelete ensures that no value is present for ForceDelete, not even an explicit nil
### GetId

`func (o *VaultDeleteParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VaultDeleteParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VaultDeleteParams) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VaultDeleteParams) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VaultDeleteParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VaultDeleteParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncludeMarkedForRemoval

`func (o *VaultDeleteParams) GetIncludeMarkedForRemoval() bool`

GetIncludeMarkedForRemoval returns the IncludeMarkedForRemoval field if non-nil, zero value otherwise.

### GetIncludeMarkedForRemovalOk

`func (o *VaultDeleteParams) GetIncludeMarkedForRemovalOk() (*bool, bool)`

GetIncludeMarkedForRemovalOk returns a tuple with the IncludeMarkedForRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMarkedForRemoval

`func (o *VaultDeleteParams) SetIncludeMarkedForRemoval(v bool)`

SetIncludeMarkedForRemoval sets IncludeMarkedForRemoval field to given value.

### HasIncludeMarkedForRemoval

`func (o *VaultDeleteParams) HasIncludeMarkedForRemoval() bool`

HasIncludeMarkedForRemoval returns a boolean if a field has been set.

### SetIncludeMarkedForRemovalNil

`func (o *VaultDeleteParams) SetIncludeMarkedForRemovalNil(b bool)`

 SetIncludeMarkedForRemovalNil sets the value for IncludeMarkedForRemoval to be an explicit nil

### UnsetIncludeMarkedForRemoval
`func (o *VaultDeleteParams) UnsetIncludeMarkedForRemoval()`

UnsetIncludeMarkedForRemoval ensures that no value is present for IncludeMarkedForRemoval, not even an explicit nil
### GetRetry

`func (o *VaultDeleteParams) GetRetry() bool`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *VaultDeleteParams) GetRetryOk() (*bool, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *VaultDeleteParams) SetRetry(v bool)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *VaultDeleteParams) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### SetRetryNil

`func (o *VaultDeleteParams) SetRetryNil(b bool)`

 SetRetryNil sets the value for Retry to be an explicit nil

### UnsetRetry
`func (o *VaultDeleteParams) UnsetRetry()`

UnsetRetry ensures that no value is present for Retry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


