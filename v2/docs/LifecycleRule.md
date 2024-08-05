# LifecycleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortIncompleteMultipartUploadAction** | Pointer to [**LifecycleRuleAbortIncompleteMultipartUploadAction**](LifecycleRuleAbortIncompleteMultipartUploadAction.md) |  | [optional] 
**Expiration** | Pointer to [**LifecycleRuleExpiration**](LifecycleRuleExpiration.md) |  | [optional] 
**Filter** | Pointer to [**LifecycleRuleFilter**](LifecycleRuleFilter.md) |  | [optional] 
**Id** | **NullableString** | Specifies the Unique identifier for the rule. The value cannot be longer than 255 characters. | 
**NonCurrentVersionExpirationAction** | Pointer to [**LifecycleRuleNonCurrentVersionExpirationAction**](LifecycleRuleNonCurrentVersionExpirationAction.md) |  | [optional] 
**Prefix** | Pointer to **NullableString** | Specifies the prefix used to identify objects that a lifecycle rule applies to. | [optional] 
**Status** | **NullableBool** | Specifies if the rule is currently being applied. | 

## Methods

### NewLifecycleRule

`func NewLifecycleRule(id NullableString, status NullableBool, ) *LifecycleRule`

NewLifecycleRule instantiates a new LifecycleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleWithDefaults

`func NewLifecycleRuleWithDefaults() *LifecycleRule`

NewLifecycleRuleWithDefaults instantiates a new LifecycleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortIncompleteMultipartUploadAction

`func (o *LifecycleRule) GetAbortIncompleteMultipartUploadAction() LifecycleRuleAbortIncompleteMultipartUploadAction`

GetAbortIncompleteMultipartUploadAction returns the AbortIncompleteMultipartUploadAction field if non-nil, zero value otherwise.

### GetAbortIncompleteMultipartUploadActionOk

`func (o *LifecycleRule) GetAbortIncompleteMultipartUploadActionOk() (*LifecycleRuleAbortIncompleteMultipartUploadAction, bool)`

GetAbortIncompleteMultipartUploadActionOk returns a tuple with the AbortIncompleteMultipartUploadAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortIncompleteMultipartUploadAction

`func (o *LifecycleRule) SetAbortIncompleteMultipartUploadAction(v LifecycleRuleAbortIncompleteMultipartUploadAction)`

SetAbortIncompleteMultipartUploadAction sets AbortIncompleteMultipartUploadAction field to given value.

### HasAbortIncompleteMultipartUploadAction

`func (o *LifecycleRule) HasAbortIncompleteMultipartUploadAction() bool`

HasAbortIncompleteMultipartUploadAction returns a boolean if a field has been set.

### GetExpiration

`func (o *LifecycleRule) GetExpiration() LifecycleRuleExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *LifecycleRule) GetExpirationOk() (*LifecycleRuleExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *LifecycleRule) SetExpiration(v LifecycleRuleExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *LifecycleRule) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetFilter

`func (o *LifecycleRule) GetFilter() LifecycleRuleFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LifecycleRule) GetFilterOk() (*LifecycleRuleFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LifecycleRule) SetFilter(v LifecycleRuleFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LifecycleRule) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *LifecycleRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LifecycleRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LifecycleRule) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *LifecycleRule) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LifecycleRule) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNonCurrentVersionExpirationAction

`func (o *LifecycleRule) GetNonCurrentVersionExpirationAction() LifecycleRuleNonCurrentVersionExpirationAction`

GetNonCurrentVersionExpirationAction returns the NonCurrentVersionExpirationAction field if non-nil, zero value otherwise.

### GetNonCurrentVersionExpirationActionOk

`func (o *LifecycleRule) GetNonCurrentVersionExpirationActionOk() (*LifecycleRuleNonCurrentVersionExpirationAction, bool)`

GetNonCurrentVersionExpirationActionOk returns a tuple with the NonCurrentVersionExpirationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCurrentVersionExpirationAction

`func (o *LifecycleRule) SetNonCurrentVersionExpirationAction(v LifecycleRuleNonCurrentVersionExpirationAction)`

SetNonCurrentVersionExpirationAction sets NonCurrentVersionExpirationAction field to given value.

### HasNonCurrentVersionExpirationAction

`func (o *LifecycleRule) HasNonCurrentVersionExpirationAction() bool`

HasNonCurrentVersionExpirationAction returns a boolean if a field has been set.

### GetPrefix

`func (o *LifecycleRule) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *LifecycleRule) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *LifecycleRule) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *LifecycleRule) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *LifecycleRule) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *LifecycleRule) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetStatus

`func (o *LifecycleRule) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LifecycleRule) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LifecycleRule) SetStatus(v bool)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *LifecycleRule) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *LifecycleRule) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


