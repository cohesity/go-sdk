# CommonACLProtoGrantees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllUsers** | Pointer to **NullableBool** | This field indicates if all users are granted ACL permission. | [optional] 
**DeniedReferrerVec** | Pointer to **[]string** | This field holds a list of referers who are denied ACL permission. | [optional] 
**GrantedReferrerVec** | Pointer to **[]string** | This field holds a list of referers who are granted ACL permission. | [optional] 
**Rlistings** | Pointer to **NullableBool** | This fields indicates if container GET and HEAD operations are permitted provided that read access is granted (using referer ACL) on objects. | [optional] 

## Methods

### NewCommonACLProtoGrantees

`func NewCommonACLProtoGrantees() *CommonACLProtoGrantees`

NewCommonACLProtoGrantees instantiates a new CommonACLProtoGrantees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonACLProtoGranteesWithDefaults

`func NewCommonACLProtoGranteesWithDefaults() *CommonACLProtoGrantees`

NewCommonACLProtoGranteesWithDefaults instantiates a new CommonACLProtoGrantees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllUsers

`func (o *CommonACLProtoGrantees) GetAllUsers() bool`

GetAllUsers returns the AllUsers field if non-nil, zero value otherwise.

### GetAllUsersOk

`func (o *CommonACLProtoGrantees) GetAllUsersOk() (*bool, bool)`

GetAllUsersOk returns a tuple with the AllUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUsers

`func (o *CommonACLProtoGrantees) SetAllUsers(v bool)`

SetAllUsers sets AllUsers field to given value.

### HasAllUsers

`func (o *CommonACLProtoGrantees) HasAllUsers() bool`

HasAllUsers returns a boolean if a field has been set.

### SetAllUsersNil

`func (o *CommonACLProtoGrantees) SetAllUsersNil(b bool)`

 SetAllUsersNil sets the value for AllUsers to be an explicit nil

### UnsetAllUsers
`func (o *CommonACLProtoGrantees) UnsetAllUsers()`

UnsetAllUsers ensures that no value is present for AllUsers, not even an explicit nil
### GetDeniedReferrerVec

`func (o *CommonACLProtoGrantees) GetDeniedReferrerVec() []string`

GetDeniedReferrerVec returns the DeniedReferrerVec field if non-nil, zero value otherwise.

### GetDeniedReferrerVecOk

`func (o *CommonACLProtoGrantees) GetDeniedReferrerVecOk() (*[]string, bool)`

GetDeniedReferrerVecOk returns a tuple with the DeniedReferrerVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedReferrerVec

`func (o *CommonACLProtoGrantees) SetDeniedReferrerVec(v []string)`

SetDeniedReferrerVec sets DeniedReferrerVec field to given value.

### HasDeniedReferrerVec

`func (o *CommonACLProtoGrantees) HasDeniedReferrerVec() bool`

HasDeniedReferrerVec returns a boolean if a field has been set.

### SetDeniedReferrerVecNil

`func (o *CommonACLProtoGrantees) SetDeniedReferrerVecNil(b bool)`

 SetDeniedReferrerVecNil sets the value for DeniedReferrerVec to be an explicit nil

### UnsetDeniedReferrerVec
`func (o *CommonACLProtoGrantees) UnsetDeniedReferrerVec()`

UnsetDeniedReferrerVec ensures that no value is present for DeniedReferrerVec, not even an explicit nil
### GetGrantedReferrerVec

`func (o *CommonACLProtoGrantees) GetGrantedReferrerVec() []string`

GetGrantedReferrerVec returns the GrantedReferrerVec field if non-nil, zero value otherwise.

### GetGrantedReferrerVecOk

`func (o *CommonACLProtoGrantees) GetGrantedReferrerVecOk() (*[]string, bool)`

GetGrantedReferrerVecOk returns a tuple with the GrantedReferrerVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedReferrerVec

`func (o *CommonACLProtoGrantees) SetGrantedReferrerVec(v []string)`

SetGrantedReferrerVec sets GrantedReferrerVec field to given value.

### HasGrantedReferrerVec

`func (o *CommonACLProtoGrantees) HasGrantedReferrerVec() bool`

HasGrantedReferrerVec returns a boolean if a field has been set.

### SetGrantedReferrerVecNil

`func (o *CommonACLProtoGrantees) SetGrantedReferrerVecNil(b bool)`

 SetGrantedReferrerVecNil sets the value for GrantedReferrerVec to be an explicit nil

### UnsetGrantedReferrerVec
`func (o *CommonACLProtoGrantees) UnsetGrantedReferrerVec()`

UnsetGrantedReferrerVec ensures that no value is present for GrantedReferrerVec, not even an explicit nil
### GetRlistings

`func (o *CommonACLProtoGrantees) GetRlistings() bool`

GetRlistings returns the Rlistings field if non-nil, zero value otherwise.

### GetRlistingsOk

`func (o *CommonACLProtoGrantees) GetRlistingsOk() (*bool, bool)`

GetRlistingsOk returns a tuple with the Rlistings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlistings

`func (o *CommonACLProtoGrantees) SetRlistings(v bool)`

SetRlistings sets Rlistings field to given value.

### HasRlistings

`func (o *CommonACLProtoGrantees) HasRlistings() bool`

HasRlistings returns a boolean if a field has been set.

### SetRlistingsNil

`func (o *CommonACLProtoGrantees) SetRlistingsNil(b bool)`

 SetRlistingsNil sets the value for Rlistings to be an explicit nil

### UnsetRlistings
`func (o *CommonACLProtoGrantees) UnsetRlistings()`

UnsetRlistings ensures that no value is present for Rlistings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


