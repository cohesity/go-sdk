# IdpUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **[]string** | Specifies the Idp groups that the user is part of. As the user may not be registered on the cluster, we may have to capture the idp group membership. This way, if a group is created on the cluster later, users will instantly have access to tenantIds from that group as well. | [optional] 
**IdpId** | Pointer to **NullableInt64** | Specifies the unique Id assigned by the Cluster for the IdP. | [optional] 
**IssuerId** | Pointer to **NullableString** | Specifies the unique identifier assigned by the vendor for this Cluster. | [optional] 
**UserId** | Pointer to **NullableString** | Specifies the unique identifier assigned by the vendor for the user. | [optional] 
**Vendor** | Pointer to **NullableString** | Specifies the vendor providing the IdP service. | [optional] 

## Methods

### NewIdpUserInfo

`func NewIdpUserInfo() *IdpUserInfo`

NewIdpUserInfo instantiates a new IdpUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpUserInfoWithDefaults

`func NewIdpUserInfoWithDefaults() *IdpUserInfo`

NewIdpUserInfoWithDefaults instantiates a new IdpUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *IdpUserInfo) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IdpUserInfo) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IdpUserInfo) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IdpUserInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *IdpUserInfo) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *IdpUserInfo) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetIdpId

`func (o *IdpUserInfo) GetIdpId() int64`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *IdpUserInfo) GetIdpIdOk() (*int64, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *IdpUserInfo) SetIdpId(v int64)`

SetIdpId sets IdpId field to given value.

### HasIdpId

`func (o *IdpUserInfo) HasIdpId() bool`

HasIdpId returns a boolean if a field has been set.

### SetIdpIdNil

`func (o *IdpUserInfo) SetIdpIdNil(b bool)`

 SetIdpIdNil sets the value for IdpId to be an explicit nil

### UnsetIdpId
`func (o *IdpUserInfo) UnsetIdpId()`

UnsetIdpId ensures that no value is present for IdpId, not even an explicit nil
### GetIssuerId

`func (o *IdpUserInfo) GetIssuerId() string`

GetIssuerId returns the IssuerId field if non-nil, zero value otherwise.

### GetIssuerIdOk

`func (o *IdpUserInfo) GetIssuerIdOk() (*string, bool)`

GetIssuerIdOk returns a tuple with the IssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerId

`func (o *IdpUserInfo) SetIssuerId(v string)`

SetIssuerId sets IssuerId field to given value.

### HasIssuerId

`func (o *IdpUserInfo) HasIssuerId() bool`

HasIssuerId returns a boolean if a field has been set.

### SetIssuerIdNil

`func (o *IdpUserInfo) SetIssuerIdNil(b bool)`

 SetIssuerIdNil sets the value for IssuerId to be an explicit nil

### UnsetIssuerId
`func (o *IdpUserInfo) UnsetIssuerId()`

UnsetIssuerId ensures that no value is present for IssuerId, not even an explicit nil
### GetUserId

`func (o *IdpUserInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *IdpUserInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *IdpUserInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *IdpUserInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *IdpUserInfo) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *IdpUserInfo) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetVendor

`func (o *IdpUserInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *IdpUserInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *IdpUserInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *IdpUserInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *IdpUserInfo) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *IdpUserInfo) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


