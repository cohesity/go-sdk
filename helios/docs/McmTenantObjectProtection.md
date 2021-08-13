# McmTenantObjectProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectProtectionId** | Pointer to **NullableString** | Specifies the id assigned to the entry. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the tenants which have access to this object. | [optional] 
**AccountId** | Pointer to **NullableString** | Specifies the account ID to which this source belongs. | [optional] 
**DmaasTenantId** | Pointer to **NullableString** | Specifies the DMaaS tenant ID to which this source belongs. | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the ID of the cluster to which the source is registered. | [optional] 
**ObjectHash** | Pointer to **NullableString** | Specifies the object hash of the source. | [optional] 
**UserDetails** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedAtMsecs** | Pointer to **NullableInt64** | Specifies the timestamp at which this entry was created in msecs. | [optional] 
**UpdatedAtMsecs** | Pointer to **NullableInt64** | Specifies the timestamp at which this entry was updated in msecs. | [optional] 
**ObjectProtection** | Pointer to [**ProtectedObjectInfo**](ProtectedObjectInfo.md) |  | [optional] 

## Methods

### NewMcmTenantObjectProtection

`func NewMcmTenantObjectProtection() *McmTenantObjectProtection`

NewMcmTenantObjectProtection instantiates a new McmTenantObjectProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmTenantObjectProtectionWithDefaults

`func NewMcmTenantObjectProtectionWithDefaults() *McmTenantObjectProtection`

NewMcmTenantObjectProtectionWithDefaults instantiates a new McmTenantObjectProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectProtectionId

`func (o *McmTenantObjectProtection) GetObjectProtectionId() string`

GetObjectProtectionId returns the ObjectProtectionId field if non-nil, zero value otherwise.

### GetObjectProtectionIdOk

`func (o *McmTenantObjectProtection) GetObjectProtectionIdOk() (*string, bool)`

GetObjectProtectionIdOk returns a tuple with the ObjectProtectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtectionId

`func (o *McmTenantObjectProtection) SetObjectProtectionId(v string)`

SetObjectProtectionId sets ObjectProtectionId field to given value.

### HasObjectProtectionId

`func (o *McmTenantObjectProtection) HasObjectProtectionId() bool`

HasObjectProtectionId returns a boolean if a field has been set.

### SetObjectProtectionIdNil

`func (o *McmTenantObjectProtection) SetObjectProtectionIdNil(b bool)`

 SetObjectProtectionIdNil sets the value for ObjectProtectionId to be an explicit nil

### UnsetObjectProtectionId
`func (o *McmTenantObjectProtection) UnsetObjectProtectionId()`

UnsetObjectProtectionId ensures that no value is present for ObjectProtectionId, not even an explicit nil
### GetTenantIds

`func (o *McmTenantObjectProtection) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *McmTenantObjectProtection) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *McmTenantObjectProtection) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *McmTenantObjectProtection) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetAccountId

`func (o *McmTenantObjectProtection) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *McmTenantObjectProtection) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *McmTenantObjectProtection) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *McmTenantObjectProtection) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *McmTenantObjectProtection) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *McmTenantObjectProtection) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetDmaasTenantId

`func (o *McmTenantObjectProtection) GetDmaasTenantId() string`

GetDmaasTenantId returns the DmaasTenantId field if non-nil, zero value otherwise.

### GetDmaasTenantIdOk

`func (o *McmTenantObjectProtection) GetDmaasTenantIdOk() (*string, bool)`

GetDmaasTenantIdOk returns a tuple with the DmaasTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmaasTenantId

`func (o *McmTenantObjectProtection) SetDmaasTenantId(v string)`

SetDmaasTenantId sets DmaasTenantId field to given value.

### HasDmaasTenantId

`func (o *McmTenantObjectProtection) HasDmaasTenantId() bool`

HasDmaasTenantId returns a boolean if a field has been set.

### SetDmaasTenantIdNil

`func (o *McmTenantObjectProtection) SetDmaasTenantIdNil(b bool)`

 SetDmaasTenantIdNil sets the value for DmaasTenantId to be an explicit nil

### UnsetDmaasTenantId
`func (o *McmTenantObjectProtection) UnsetDmaasTenantId()`

UnsetDmaasTenantId ensures that no value is present for DmaasTenantId, not even an explicit nil
### GetClusterId

`func (o *McmTenantObjectProtection) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *McmTenantObjectProtection) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *McmTenantObjectProtection) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *McmTenantObjectProtection) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *McmTenantObjectProtection) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *McmTenantObjectProtection) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetObjectHash

`func (o *McmTenantObjectProtection) GetObjectHash() string`

GetObjectHash returns the ObjectHash field if non-nil, zero value otherwise.

### GetObjectHashOk

`func (o *McmTenantObjectProtection) GetObjectHashOk() (*string, bool)`

GetObjectHashOk returns a tuple with the ObjectHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectHash

`func (o *McmTenantObjectProtection) SetObjectHash(v string)`

SetObjectHash sets ObjectHash field to given value.

### HasObjectHash

`func (o *McmTenantObjectProtection) HasObjectHash() bool`

HasObjectHash returns a boolean if a field has been set.

### SetObjectHashNil

`func (o *McmTenantObjectProtection) SetObjectHashNil(b bool)`

 SetObjectHashNil sets the value for ObjectHash to be an explicit nil

### UnsetObjectHash
`func (o *McmTenantObjectProtection) UnsetObjectHash()`

UnsetObjectHash ensures that no value is present for ObjectHash, not even an explicit nil
### GetUserDetails

`func (o *McmTenantObjectProtection) GetUserDetails() User`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *McmTenantObjectProtection) GetUserDetailsOk() (*User, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *McmTenantObjectProtection) SetUserDetails(v User)`

SetUserDetails sets UserDetails field to given value.

### HasUserDetails

`func (o *McmTenantObjectProtection) HasUserDetails() bool`

HasUserDetails returns a boolean if a field has been set.

### GetCreatedAtMsecs

`func (o *McmTenantObjectProtection) GetCreatedAtMsecs() int64`

GetCreatedAtMsecs returns the CreatedAtMsecs field if non-nil, zero value otherwise.

### GetCreatedAtMsecsOk

`func (o *McmTenantObjectProtection) GetCreatedAtMsecsOk() (*int64, bool)`

GetCreatedAtMsecsOk returns a tuple with the CreatedAtMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtMsecs

`func (o *McmTenantObjectProtection) SetCreatedAtMsecs(v int64)`

SetCreatedAtMsecs sets CreatedAtMsecs field to given value.

### HasCreatedAtMsecs

`func (o *McmTenantObjectProtection) HasCreatedAtMsecs() bool`

HasCreatedAtMsecs returns a boolean if a field has been set.

### SetCreatedAtMsecsNil

`func (o *McmTenantObjectProtection) SetCreatedAtMsecsNil(b bool)`

 SetCreatedAtMsecsNil sets the value for CreatedAtMsecs to be an explicit nil

### UnsetCreatedAtMsecs
`func (o *McmTenantObjectProtection) UnsetCreatedAtMsecs()`

UnsetCreatedAtMsecs ensures that no value is present for CreatedAtMsecs, not even an explicit nil
### GetUpdatedAtMsecs

`func (o *McmTenantObjectProtection) GetUpdatedAtMsecs() int64`

GetUpdatedAtMsecs returns the UpdatedAtMsecs field if non-nil, zero value otherwise.

### GetUpdatedAtMsecsOk

`func (o *McmTenantObjectProtection) GetUpdatedAtMsecsOk() (*int64, bool)`

GetUpdatedAtMsecsOk returns a tuple with the UpdatedAtMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtMsecs

`func (o *McmTenantObjectProtection) SetUpdatedAtMsecs(v int64)`

SetUpdatedAtMsecs sets UpdatedAtMsecs field to given value.

### HasUpdatedAtMsecs

`func (o *McmTenantObjectProtection) HasUpdatedAtMsecs() bool`

HasUpdatedAtMsecs returns a boolean if a field has been set.

### SetUpdatedAtMsecsNil

`func (o *McmTenantObjectProtection) SetUpdatedAtMsecsNil(b bool)`

 SetUpdatedAtMsecsNil sets the value for UpdatedAtMsecs to be an explicit nil

### UnsetUpdatedAtMsecs
`func (o *McmTenantObjectProtection) UnsetUpdatedAtMsecs()`

UnsetUpdatedAtMsecs ensures that no value is present for UpdatedAtMsecs, not even an explicit nil
### GetObjectProtection

`func (o *McmTenantObjectProtection) GetObjectProtection() ProtectedObjectInfo`

GetObjectProtection returns the ObjectProtection field if non-nil, zero value otherwise.

### GetObjectProtectionOk

`func (o *McmTenantObjectProtection) GetObjectProtectionOk() (*ProtectedObjectInfo, bool)`

GetObjectProtectionOk returns a tuple with the ObjectProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectProtection

`func (o *McmTenantObjectProtection) SetObjectProtection(v ProtectedObjectInfo)`

SetObjectProtection sets ObjectProtection field to given value.

### HasObjectProtection

`func (o *McmTenantObjectProtection) HasObjectProtection() bool`

HasObjectProtection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


