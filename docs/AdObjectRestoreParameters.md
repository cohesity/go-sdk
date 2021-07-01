# AdObjectRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangePasswordOnNextLogon** | Pointer to **NullableBool** | Specifies the option for AD &#39;user&#39; type of objects to change password when they next logon. | [optional] 
**LeaveStateDisabled** | Pointer to **NullableBool** | Specifies the option to leave the restored object in disabled state for AD &#39;account&#39; type of objects that support disabled state so that admin can do manual inspection before enabling the account. This property has no effect when restoring object from recycle bin. &#39;User&#39; and &#39;Computer&#39; are AD account types having enable/disable option. | [optional] 
**ObjectGuids** | Pointer to **[]string** | Specifies the array of AD object guids to restore either from recycle bin or from AD snapshot. The guid should not exist in production AD. If it exits, then kDuplicate error is output in status. | [optional] 
**OrganizationUnitPath** | Pointer to **NullableString** | Specifies the Distinguished name(DN) of the target Organization Unit (OU) to restore non-OU object. This can be empty, in which case objects are restored to their original OU. The &#39;credential&#39; specified must have privileges to add objects to this OU. Example: &#39;OU&#x3D;SJC,OU&#x3D;EngOU,DC&#x3D;contoso,DC&#x3D;com&#39;. This parameter is ignored for OU objects. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password for restoring user type objects (user, inetOrgPerson or organizationalPerson LDAP classes) that is returned in search result with &#39;kRestorePasswordRequired&#39; flag, an initial password is required. The password is UTF8 string encoded in base64 format. Password cannot be empty and must satisfy production AD password complexity. If &#39;kFromDestRecycleBinIfFound&#39; is true and the user is restored from production AD recycle bin, password will not be changed and the restored user retains their original password. For non-user type objects, this password value is ignored. | [optional] 

## Methods

### NewAdObjectRestoreParameters

`func NewAdObjectRestoreParameters() *AdObjectRestoreParameters`

NewAdObjectRestoreParameters instantiates a new AdObjectRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdObjectRestoreParametersWithDefaults

`func NewAdObjectRestoreParametersWithDefaults() *AdObjectRestoreParameters`

NewAdObjectRestoreParametersWithDefaults instantiates a new AdObjectRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangePasswordOnNextLogon

`func (o *AdObjectRestoreParameters) GetChangePasswordOnNextLogon() bool`

GetChangePasswordOnNextLogon returns the ChangePasswordOnNextLogon field if non-nil, zero value otherwise.

### GetChangePasswordOnNextLogonOk

`func (o *AdObjectRestoreParameters) GetChangePasswordOnNextLogonOk() (*bool, bool)`

GetChangePasswordOnNextLogonOk returns a tuple with the ChangePasswordOnNextLogon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePasswordOnNextLogon

`func (o *AdObjectRestoreParameters) SetChangePasswordOnNextLogon(v bool)`

SetChangePasswordOnNextLogon sets ChangePasswordOnNextLogon field to given value.

### HasChangePasswordOnNextLogon

`func (o *AdObjectRestoreParameters) HasChangePasswordOnNextLogon() bool`

HasChangePasswordOnNextLogon returns a boolean if a field has been set.

### SetChangePasswordOnNextLogonNil

`func (o *AdObjectRestoreParameters) SetChangePasswordOnNextLogonNil(b bool)`

 SetChangePasswordOnNextLogonNil sets the value for ChangePasswordOnNextLogon to be an explicit nil

### UnsetChangePasswordOnNextLogon
`func (o *AdObjectRestoreParameters) UnsetChangePasswordOnNextLogon()`

UnsetChangePasswordOnNextLogon ensures that no value is present for ChangePasswordOnNextLogon, not even an explicit nil
### GetLeaveStateDisabled

`func (o *AdObjectRestoreParameters) GetLeaveStateDisabled() bool`

GetLeaveStateDisabled returns the LeaveStateDisabled field if non-nil, zero value otherwise.

### GetLeaveStateDisabledOk

`func (o *AdObjectRestoreParameters) GetLeaveStateDisabledOk() (*bool, bool)`

GetLeaveStateDisabledOk returns a tuple with the LeaveStateDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaveStateDisabled

`func (o *AdObjectRestoreParameters) SetLeaveStateDisabled(v bool)`

SetLeaveStateDisabled sets LeaveStateDisabled field to given value.

### HasLeaveStateDisabled

`func (o *AdObjectRestoreParameters) HasLeaveStateDisabled() bool`

HasLeaveStateDisabled returns a boolean if a field has been set.

### SetLeaveStateDisabledNil

`func (o *AdObjectRestoreParameters) SetLeaveStateDisabledNil(b bool)`

 SetLeaveStateDisabledNil sets the value for LeaveStateDisabled to be an explicit nil

### UnsetLeaveStateDisabled
`func (o *AdObjectRestoreParameters) UnsetLeaveStateDisabled()`

UnsetLeaveStateDisabled ensures that no value is present for LeaveStateDisabled, not even an explicit nil
### GetObjectGuids

`func (o *AdObjectRestoreParameters) GetObjectGuids() []string`

GetObjectGuids returns the ObjectGuids field if non-nil, zero value otherwise.

### GetObjectGuidsOk

`func (o *AdObjectRestoreParameters) GetObjectGuidsOk() (*[]string, bool)`

GetObjectGuidsOk returns a tuple with the ObjectGuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGuids

`func (o *AdObjectRestoreParameters) SetObjectGuids(v []string)`

SetObjectGuids sets ObjectGuids field to given value.

### HasObjectGuids

`func (o *AdObjectRestoreParameters) HasObjectGuids() bool`

HasObjectGuids returns a boolean if a field has been set.

### SetObjectGuidsNil

`func (o *AdObjectRestoreParameters) SetObjectGuidsNil(b bool)`

 SetObjectGuidsNil sets the value for ObjectGuids to be an explicit nil

### UnsetObjectGuids
`func (o *AdObjectRestoreParameters) UnsetObjectGuids()`

UnsetObjectGuids ensures that no value is present for ObjectGuids, not even an explicit nil
### GetOrganizationUnitPath

`func (o *AdObjectRestoreParameters) GetOrganizationUnitPath() string`

GetOrganizationUnitPath returns the OrganizationUnitPath field if non-nil, zero value otherwise.

### GetOrganizationUnitPathOk

`func (o *AdObjectRestoreParameters) GetOrganizationUnitPathOk() (*string, bool)`

GetOrganizationUnitPathOk returns a tuple with the OrganizationUnitPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnitPath

`func (o *AdObjectRestoreParameters) SetOrganizationUnitPath(v string)`

SetOrganizationUnitPath sets OrganizationUnitPath field to given value.

### HasOrganizationUnitPath

`func (o *AdObjectRestoreParameters) HasOrganizationUnitPath() bool`

HasOrganizationUnitPath returns a boolean if a field has been set.

### SetOrganizationUnitPathNil

`func (o *AdObjectRestoreParameters) SetOrganizationUnitPathNil(b bool)`

 SetOrganizationUnitPathNil sets the value for OrganizationUnitPath to be an explicit nil

### UnsetOrganizationUnitPath
`func (o *AdObjectRestoreParameters) UnsetOrganizationUnitPath()`

UnsetOrganizationUnitPath ensures that no value is present for OrganizationUnitPath, not even an explicit nil
### GetPassword

`func (o *AdObjectRestoreParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdObjectRestoreParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdObjectRestoreParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AdObjectRestoreParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AdObjectRestoreParameters) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AdObjectRestoreParameters) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


