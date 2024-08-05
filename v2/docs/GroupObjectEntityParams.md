# GroupObjectEntityParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMailEnabled** | Pointer to **NullableBool** | Specifies whether the Group is mail enabled. Mail enabled groups are used within Microsoft to distribute messages. | [optional] 
**IsSecurityEnabled** | Pointer to **NullableBool** | Specifies whether the Group is security enabled. Security enabled groups are used to grant access permissions to resources in Exchange and Active Directory. | [optional] 
**MemberCount** | Pointer to **NullableInt64** | Specifies the count of members within the Group. | [optional] 

## Methods

### NewGroupObjectEntityParams

`func NewGroupObjectEntityParams() *GroupObjectEntityParams`

NewGroupObjectEntityParams instantiates a new GroupObjectEntityParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupObjectEntityParamsWithDefaults

`func NewGroupObjectEntityParamsWithDefaults() *GroupObjectEntityParams`

NewGroupObjectEntityParamsWithDefaults instantiates a new GroupObjectEntityParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMailEnabled

`func (o *GroupObjectEntityParams) GetIsMailEnabled() bool`

GetIsMailEnabled returns the IsMailEnabled field if non-nil, zero value otherwise.

### GetIsMailEnabledOk

`func (o *GroupObjectEntityParams) GetIsMailEnabledOk() (*bool, bool)`

GetIsMailEnabledOk returns a tuple with the IsMailEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMailEnabled

`func (o *GroupObjectEntityParams) SetIsMailEnabled(v bool)`

SetIsMailEnabled sets IsMailEnabled field to given value.

### HasIsMailEnabled

`func (o *GroupObjectEntityParams) HasIsMailEnabled() bool`

HasIsMailEnabled returns a boolean if a field has been set.

### SetIsMailEnabledNil

`func (o *GroupObjectEntityParams) SetIsMailEnabledNil(b bool)`

 SetIsMailEnabledNil sets the value for IsMailEnabled to be an explicit nil

### UnsetIsMailEnabled
`func (o *GroupObjectEntityParams) UnsetIsMailEnabled()`

UnsetIsMailEnabled ensures that no value is present for IsMailEnabled, not even an explicit nil
### GetIsSecurityEnabled

`func (o *GroupObjectEntityParams) GetIsSecurityEnabled() bool`

GetIsSecurityEnabled returns the IsSecurityEnabled field if non-nil, zero value otherwise.

### GetIsSecurityEnabledOk

`func (o *GroupObjectEntityParams) GetIsSecurityEnabledOk() (*bool, bool)`

GetIsSecurityEnabledOk returns a tuple with the IsSecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecurityEnabled

`func (o *GroupObjectEntityParams) SetIsSecurityEnabled(v bool)`

SetIsSecurityEnabled sets IsSecurityEnabled field to given value.

### HasIsSecurityEnabled

`func (o *GroupObjectEntityParams) HasIsSecurityEnabled() bool`

HasIsSecurityEnabled returns a boolean if a field has been set.

### SetIsSecurityEnabledNil

`func (o *GroupObjectEntityParams) SetIsSecurityEnabledNil(b bool)`

 SetIsSecurityEnabledNil sets the value for IsSecurityEnabled to be an explicit nil

### UnsetIsSecurityEnabled
`func (o *GroupObjectEntityParams) UnsetIsSecurityEnabled()`

UnsetIsSecurityEnabled ensures that no value is present for IsSecurityEnabled, not even an explicit nil
### GetMemberCount

`func (o *GroupObjectEntityParams) GetMemberCount() int64`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *GroupObjectEntityParams) GetMemberCountOk() (*int64, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *GroupObjectEntityParams) SetMemberCount(v int64)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *GroupObjectEntityParams) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### SetMemberCountNil

`func (o *GroupObjectEntityParams) SetMemberCountNil(b bool)`

 SetMemberCountNil sets the value for MemberCount to be an explicit nil

### UnsetMemberCount
`func (o *GroupObjectEntityParams) UnsetMemberCount()`

UnsetMemberCount ensures that no value is present for MemberCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


