# SourceRegistrationCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | Specifies the username used to register a source. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password used to register a source. | [optional] 
**SmbCredentials** | Pointer to [**SmbCredentials**](SmbCredentials.md) |  | [optional] 
**Vcenters** | Pointer to [**[]VcenterCredentialInfo**](VcenterCredentialInfo.md) | Specifies the list of child vcenter credentials. This will only be populated in the case of VCD. | [optional] 
**Office365AppCredentialsList** | Pointer to [**[]Office365AppCredentials**](Office365AppCredentials.md) | Specifies a list of office365 azure application credentials needed to authenticate &amp; authorize users. | [optional] 

## Methods

### NewSourceRegistrationCredentials

`func NewSourceRegistrationCredentials() *SourceRegistrationCredentials`

NewSourceRegistrationCredentials instantiates a new SourceRegistrationCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceRegistrationCredentialsWithDefaults

`func NewSourceRegistrationCredentialsWithDefaults() *SourceRegistrationCredentials`

NewSourceRegistrationCredentialsWithDefaults instantiates a new SourceRegistrationCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SourceRegistrationCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SourceRegistrationCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SourceRegistrationCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SourceRegistrationCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *SourceRegistrationCredentials) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SourceRegistrationCredentials) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *SourceRegistrationCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SourceRegistrationCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SourceRegistrationCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SourceRegistrationCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SourceRegistrationCredentials) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SourceRegistrationCredentials) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSmbCredentials

`func (o *SourceRegistrationCredentials) GetSmbCredentials() SmbCredentials`

GetSmbCredentials returns the SmbCredentials field if non-nil, zero value otherwise.

### GetSmbCredentialsOk

`func (o *SourceRegistrationCredentials) GetSmbCredentialsOk() (*SmbCredentials, bool)`

GetSmbCredentialsOk returns a tuple with the SmbCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbCredentials

`func (o *SourceRegistrationCredentials) SetSmbCredentials(v SmbCredentials)`

SetSmbCredentials sets SmbCredentials field to given value.

### HasSmbCredentials

`func (o *SourceRegistrationCredentials) HasSmbCredentials() bool`

HasSmbCredentials returns a boolean if a field has been set.

### GetVcenters

`func (o *SourceRegistrationCredentials) GetVcenters() []VcenterCredentialInfo`

GetVcenters returns the Vcenters field if non-nil, zero value otherwise.

### GetVcentersOk

`func (o *SourceRegistrationCredentials) GetVcentersOk() (*[]VcenterCredentialInfo, bool)`

GetVcentersOk returns a tuple with the Vcenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenters

`func (o *SourceRegistrationCredentials) SetVcenters(v []VcenterCredentialInfo)`

SetVcenters sets Vcenters field to given value.

### HasVcenters

`func (o *SourceRegistrationCredentials) HasVcenters() bool`

HasVcenters returns a boolean if a field has been set.

### GetOffice365AppCredentialsList

`func (o *SourceRegistrationCredentials) GetOffice365AppCredentialsList() []Office365AppCredentials`

GetOffice365AppCredentialsList returns the Office365AppCredentialsList field if non-nil, zero value otherwise.

### GetOffice365AppCredentialsListOk

`func (o *SourceRegistrationCredentials) GetOffice365AppCredentialsListOk() (*[]Office365AppCredentials, bool)`

GetOffice365AppCredentialsListOk returns a tuple with the Office365AppCredentialsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365AppCredentialsList

`func (o *SourceRegistrationCredentials) SetOffice365AppCredentialsList(v []Office365AppCredentials)`

SetOffice365AppCredentialsList sets Office365AppCredentialsList field to given value.

### HasOffice365AppCredentialsList

`func (o *SourceRegistrationCredentials) HasOffice365AppCredentialsList() bool`

HasOffice365AppCredentialsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


