# McmTenantProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SfAccountId** | Pointer to **NullableString** | Specifies the Salesforce account ID of this account. | [optional] 
**UserSid** | Pointer to **NullableString** | Specifies the sid of the logged in user. | [optional] 
**Clusters** | Pointer to [**[]McmClusterIdentifier**](McmClusterIdentifier.md) | Specifies the list of clusters owned by this account id. | [optional] 
**TenantProfiles** | Pointer to [**[]McmTenantProfile**](McmTenantProfile.md) | Specifies the list of tenant profiles associated to this account if any. | [optional] 

## Methods

### NewMcmTenantProfiles

`func NewMcmTenantProfiles() *McmTenantProfiles`

NewMcmTenantProfiles instantiates a new McmTenantProfiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmTenantProfilesWithDefaults

`func NewMcmTenantProfilesWithDefaults() *McmTenantProfiles`

NewMcmTenantProfilesWithDefaults instantiates a new McmTenantProfiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSfAccountId

`func (o *McmTenantProfiles) GetSfAccountId() string`

GetSfAccountId returns the SfAccountId field if non-nil, zero value otherwise.

### GetSfAccountIdOk

`func (o *McmTenantProfiles) GetSfAccountIdOk() (*string, bool)`

GetSfAccountIdOk returns a tuple with the SfAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfAccountId

`func (o *McmTenantProfiles) SetSfAccountId(v string)`

SetSfAccountId sets SfAccountId field to given value.

### HasSfAccountId

`func (o *McmTenantProfiles) HasSfAccountId() bool`

HasSfAccountId returns a boolean if a field has been set.

### SetSfAccountIdNil

`func (o *McmTenantProfiles) SetSfAccountIdNil(b bool)`

 SetSfAccountIdNil sets the value for SfAccountId to be an explicit nil

### UnsetSfAccountId
`func (o *McmTenantProfiles) UnsetSfAccountId()`

UnsetSfAccountId ensures that no value is present for SfAccountId, not even an explicit nil
### GetUserSid

`func (o *McmTenantProfiles) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *McmTenantProfiles) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *McmTenantProfiles) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *McmTenantProfiles) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### SetUserSidNil

`func (o *McmTenantProfiles) SetUserSidNil(b bool)`

 SetUserSidNil sets the value for UserSid to be an explicit nil

### UnsetUserSid
`func (o *McmTenantProfiles) UnsetUserSid()`

UnsetUserSid ensures that no value is present for UserSid, not even an explicit nil
### GetClusters

`func (o *McmTenantProfiles) GetClusters() []McmClusterIdentifier`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *McmTenantProfiles) GetClustersOk() (*[]McmClusterIdentifier, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *McmTenantProfiles) SetClusters(v []McmClusterIdentifier)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *McmTenantProfiles) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### SetClustersNil

`func (o *McmTenantProfiles) SetClustersNil(b bool)`

 SetClustersNil sets the value for Clusters to be an explicit nil

### UnsetClusters
`func (o *McmTenantProfiles) UnsetClusters()`

UnsetClusters ensures that no value is present for Clusters, not even an explicit nil
### GetTenantProfiles

`func (o *McmTenantProfiles) GetTenantProfiles() []McmTenantProfile`

GetTenantProfiles returns the TenantProfiles field if non-nil, zero value otherwise.

### GetTenantProfilesOk

`func (o *McmTenantProfiles) GetTenantProfilesOk() (*[]McmTenantProfile, bool)`

GetTenantProfilesOk returns a tuple with the TenantProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantProfiles

`func (o *McmTenantProfiles) SetTenantProfiles(v []McmTenantProfile)`

SetTenantProfiles sets TenantProfiles field to given value.

### HasTenantProfiles

`func (o *McmTenantProfiles) HasTenantProfiles() bool`

HasTenantProfiles returns a boolean if a field has been set.

### SetTenantProfilesNil

`func (o *McmTenantProfiles) SetTenantProfilesNil(b bool)`

 SetTenantProfilesNil sets the value for TenantProfiles to be an explicit nil

### UnsetTenantProfiles
`func (o *McmTenantProfiles) UnsetTenantProfiles()`

UnsetTenantProfiles ensures that no value is present for TenantProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


