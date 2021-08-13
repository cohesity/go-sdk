# McmTenantSourceRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRegistrationId** | Pointer to **NullableString** | Specifies the id assigned to the entry. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the ID of the source on the cluster. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment of the source. | [optional] 
**ObjectType** | Pointer to **NullableString** | Specifies the type of the object. | [optional] 
**IsRootSource** | Pointer to **NullableBool** | Specifies whether or not this object is a root level source (registered directly rather than the child of a registered source). | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the tenants which have access to this object. | [optional] 
**AccountId** | Pointer to **NullableString** | Specifies the account ID to which this source belongs. | [optional] 
**DmaasTenantId** | Pointer to **NullableString** | Specifies the DMaaS tenant ID to which this source belongs. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the ID of the cluster to which the source is registered. | [optional] 
**ObjectHash** | Pointer to **NullableString** | Specifies the object hash of the source. | [optional] 
**UserDetails** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedAtMsecs** | Pointer to **NullableInt64** | Specifies the timestamp at which this entry was created in msecs. | [optional] 
**UpdatedAtMsecs** | Pointer to **NullableInt64** | Specifies the timestamp at which this entry was updated in msecs. | [optional] 
**SourceRegistration** | Pointer to [**SourceRegistration**](SourceRegistration.md) |  | [optional] 
**EnableAppParams** | Pointer to [**[]ObjectActionRequest**](ObjectActionRequest.md) | Specifies the list of enable application params which have been performed on this source. | [optional] 
**Credentials** | Pointer to [**SourceRegistrationCredentials**](SourceRegistrationCredentials.md) |  | [optional] 

## Methods

### NewMcmTenantSourceRegistration

`func NewMcmTenantSourceRegistration() *McmTenantSourceRegistration`

NewMcmTenantSourceRegistration instantiates a new McmTenantSourceRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmTenantSourceRegistrationWithDefaults

`func NewMcmTenantSourceRegistrationWithDefaults() *McmTenantSourceRegistration`

NewMcmTenantSourceRegistrationWithDefaults instantiates a new McmTenantSourceRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceRegistrationId

`func (o *McmTenantSourceRegistration) GetSourceRegistrationId() string`

GetSourceRegistrationId returns the SourceRegistrationId field if non-nil, zero value otherwise.

### GetSourceRegistrationIdOk

`func (o *McmTenantSourceRegistration) GetSourceRegistrationIdOk() (*string, bool)`

GetSourceRegistrationIdOk returns a tuple with the SourceRegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegistrationId

`func (o *McmTenantSourceRegistration) SetSourceRegistrationId(v string)`

SetSourceRegistrationId sets SourceRegistrationId field to given value.

### HasSourceRegistrationId

`func (o *McmTenantSourceRegistration) HasSourceRegistrationId() bool`

HasSourceRegistrationId returns a boolean if a field has been set.

### SetSourceRegistrationIdNil

`func (o *McmTenantSourceRegistration) SetSourceRegistrationIdNil(b bool)`

 SetSourceRegistrationIdNil sets the value for SourceRegistrationId to be an explicit nil

### UnsetSourceRegistrationId
`func (o *McmTenantSourceRegistration) UnsetSourceRegistrationId()`

UnsetSourceRegistrationId ensures that no value is present for SourceRegistrationId, not even an explicit nil
### GetSourceId

`func (o *McmTenantSourceRegistration) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *McmTenantSourceRegistration) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *McmTenantSourceRegistration) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *McmTenantSourceRegistration) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *McmTenantSourceRegistration) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *McmTenantSourceRegistration) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetEnvironment

`func (o *McmTenantSourceRegistration) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *McmTenantSourceRegistration) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *McmTenantSourceRegistration) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *McmTenantSourceRegistration) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *McmTenantSourceRegistration) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *McmTenantSourceRegistration) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetObjectType

`func (o *McmTenantSourceRegistration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *McmTenantSourceRegistration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *McmTenantSourceRegistration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *McmTenantSourceRegistration) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *McmTenantSourceRegistration) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *McmTenantSourceRegistration) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetIsRootSource

`func (o *McmTenantSourceRegistration) GetIsRootSource() bool`

GetIsRootSource returns the IsRootSource field if non-nil, zero value otherwise.

### GetIsRootSourceOk

`func (o *McmTenantSourceRegistration) GetIsRootSourceOk() (*bool, bool)`

GetIsRootSourceOk returns a tuple with the IsRootSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRootSource

`func (o *McmTenantSourceRegistration) SetIsRootSource(v bool)`

SetIsRootSource sets IsRootSource field to given value.

### HasIsRootSource

`func (o *McmTenantSourceRegistration) HasIsRootSource() bool`

HasIsRootSource returns a boolean if a field has been set.

### SetIsRootSourceNil

`func (o *McmTenantSourceRegistration) SetIsRootSourceNil(b bool)`

 SetIsRootSourceNil sets the value for IsRootSource to be an explicit nil

### UnsetIsRootSource
`func (o *McmTenantSourceRegistration) UnsetIsRootSource()`

UnsetIsRootSource ensures that no value is present for IsRootSource, not even an explicit nil
### GetTenantIds

`func (o *McmTenantSourceRegistration) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *McmTenantSourceRegistration) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *McmTenantSourceRegistration) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *McmTenantSourceRegistration) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetAccountId

`func (o *McmTenantSourceRegistration) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *McmTenantSourceRegistration) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *McmTenantSourceRegistration) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *McmTenantSourceRegistration) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *McmTenantSourceRegistration) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *McmTenantSourceRegistration) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetDmaasTenantId

`func (o *McmTenantSourceRegistration) GetDmaasTenantId() string`

GetDmaasTenantId returns the DmaasTenantId field if non-nil, zero value otherwise.

### GetDmaasTenantIdOk

`func (o *McmTenantSourceRegistration) GetDmaasTenantIdOk() (*string, bool)`

GetDmaasTenantIdOk returns a tuple with the DmaasTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmaasTenantId

`func (o *McmTenantSourceRegistration) SetDmaasTenantId(v string)`

SetDmaasTenantId sets DmaasTenantId field to given value.

### HasDmaasTenantId

`func (o *McmTenantSourceRegistration) HasDmaasTenantId() bool`

HasDmaasTenantId returns a boolean if a field has been set.

### SetDmaasTenantIdNil

`func (o *McmTenantSourceRegistration) SetDmaasTenantIdNil(b bool)`

 SetDmaasTenantIdNil sets the value for DmaasTenantId to be an explicit nil

### UnsetDmaasTenantId
`func (o *McmTenantSourceRegistration) UnsetDmaasTenantId()`

UnsetDmaasTenantId ensures that no value is present for DmaasTenantId, not even an explicit nil
### GetClusterId

`func (o *McmTenantSourceRegistration) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmTenantSourceRegistration) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmTenantSourceRegistration) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmTenantSourceRegistration) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmTenantSourceRegistration) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmTenantSourceRegistration) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetObjectHash

`func (o *McmTenantSourceRegistration) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *McmTenantSourceRegistration) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *McmTenantSourceRegistration) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *McmTenantSourceRegistration) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *McmTenantSourceRegistration) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *McmTenantSourceRegistration) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetUserDetails

`func (o *McmTenantSourceRegistration) GetUserDetails() User`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *McmTenantSourceRegistration) GetUserDetailsOk() (*User, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *McmTenantSourceRegistration) SetUserDetails(v User)`

SetUserDetails sets UserDetails field to given value.

### HasUserDetails

`func (o *McmTenantSourceRegistration) HasUserDetails() bool`

HasUserDetails returns a boolean if a field has been set.

### GetCreatedAtMsecs

`func (o *McmTenantSourceRegistration) GetCreatedAtMsecs() int64`

GetCreatedAtMsecs returns the CreatedAtMsecs field if non-nil, zero value otherwise.

### GetCreatedAtMsecsOk

`func (o *McmTenantSourceRegistration) GetCreatedAtMsecsOk() (*int64, bool)`

GetCreatedAtMsecsOk returns a tuple with the CreatedAtMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtMsecs

`func (o *McmTenantSourceRegistration) SetCreatedAtMsecs(v int64)`

SetCreatedAtMsecs sets CreatedAtMsecs field to given value.

### HasCreatedAtMsecs

`func (o *McmTenantSourceRegistration) HasCreatedAtMsecs() bool`

HasCreatedAtMsecs returns a boolean if a field has been set.

### SetCreatedAtMsecsNil

`func (o *McmTenantSourceRegistration) SetCreatedAtMsecsNil(b bool)`

 SetCreatedAtMsecsNil sets the value for CreatedAtMsecs to be an explicit nil

### UnsetCreatedAtMsecs
`func (o *McmTenantSourceRegistration) UnsetCreatedAtMsecs()`

UnsetCreatedAtMsecs ensures that no value is present for CreatedAtMsecs, not even an explicit nil
### GetUpdatedAtMsecs

`func (o *McmTenantSourceRegistration) GetUpdatedAtMsecs() int64`

GetUpdatedAtMsecs returns the UpdatedAtMsecs field if non-nil, zero value otherwise.

### GetUpdatedAtMsecsOk

`func (o *McmTenantSourceRegistration) GetUpdatedAtMsecsOk() (*int64, bool)`

GetUpdatedAtMsecsOk returns a tuple with the UpdatedAtMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtMsecs

`func (o *McmTenantSourceRegistration) SetUpdatedAtMsecs(v int64)`

SetUpdatedAtMsecs sets UpdatedAtMsecs field to given value.

### HasUpdatedAtMsecs

`func (o *McmTenantSourceRegistration) HasUpdatedAtMsecs() bool`

HasUpdatedAtMsecs returns a boolean if a field has been set.

### SetUpdatedAtMsecsNil

`func (o *McmTenantSourceRegistration) SetUpdatedAtMsecsNil(b bool)`

 SetUpdatedAtMsecsNil sets the value for UpdatedAtMsecs to be an explicit nil

### UnsetUpdatedAtMsecs
`func (o *McmTenantSourceRegistration) UnsetUpdatedAtMsecs()`

UnsetUpdatedAtMsecs ensures that no value is present for UpdatedAtMsecs, not even an explicit nil
### GetSourceRegistration

`func (o *McmTenantSourceRegistration) GetSourceRegistration() SourceRegistration`

GetSourceRegistration returns the SourceRegistration field if non-nil, zero value otherwise.

### GetSourceRegistrationOk

`func (o *McmTenantSourceRegistration) GetSourceRegistrationOk() (*SourceRegistration, bool)`

GetSourceRegistrationOk returns a tuple with the SourceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegistration

`func (o *McmTenantSourceRegistration) SetSourceRegistration(v SourceRegistration)`

SetSourceRegistration sets SourceRegistration field to given value.

### HasSourceRegistration

`func (o *McmTenantSourceRegistration) HasSourceRegistration() bool`

HasSourceRegistration returns a boolean if a field has been set.

### GetEnableAppParams

`func (o *McmTenantSourceRegistration) GetEnableAppParams() []ObjectActionRequest`

GetEnableAppParams returns the EnableAppParams field if non-nil, zero value otherwise.

### GetEnableAppParamsOk

`func (o *McmTenantSourceRegistration) GetEnableAppParamsOk() (*[]ObjectActionRequest, bool)`

GetEnableAppParamsOk returns a tuple with the EnableAppParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAppParams

`func (o *McmTenantSourceRegistration) SetEnableAppParams(v []ObjectActionRequest)`

SetEnableAppParams sets EnableAppParams field to given value.

### HasEnableAppParams

`func (o *McmTenantSourceRegistration) HasEnableAppParams() bool`

HasEnableAppParams returns a boolean if a field has been set.

### GetCredentials

`func (o *McmTenantSourceRegistration) GetCredentials() SourceRegistrationCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *McmTenantSourceRegistration) GetCredentialsOk() (*SourceRegistrationCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *McmTenantSourceRegistration) SetCredentials(v SourceRegistrationCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *McmTenantSourceRegistration) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


