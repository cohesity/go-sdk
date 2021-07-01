# NasMountCredentialParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | Specifies the domain in which this credential is valid. | [optional] 
**DomainController** | Pointer to **NullableString** | Specifies the domain controller for the selected domain | [optional] 
**ManagePasswordByCohesity** | Pointer to **NullableBool** | Specifies if Cohesity can manage the password after registration | [optional] 
**NasProtocol** | Pointer to **NullableString** | Specifies the protocol used by the NAS server. Specifies the protocol used by a NAS server. &#39;kNfs3&#39; indicates NFS v3 protocol. &#39;kCifs1&#39; indicates CIFS v1.0 protocol. | [optional] 
**NasType** | Pointer to **NullableString** | Specifies the type of a NAS Object such as &#39;kGroup&#39;, or &#39;kHost&#39;. Specifies the kind of NAS mount. &#39;kGroup&#39; indicates top level node that holds individual NAS hosts. &#39;kHost&#39; indicates a single NAS path that can be mounted. &#39;kDfsGroup&#39; indicates a DFS group containing top level directories mapped to different servers. &#39;kDfsTopDir&#39; indicates a top level directory inside a DFS group, discovered when registering a DFS group. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password for the username to use for mounting the NAS. | [optional] 
**SkipValidation** | Pointer to **NullableBool** | Specifies the flag to disable mount point validation during registration process. | [optional] 
**Username** | Pointer to **NullableString** | Specifies a username to use for mounting the NAS. | [optional] 

## Methods

### NewNasMountCredentialParams

`func NewNasMountCredentialParams() *NasMountCredentialParams`

NewNasMountCredentialParams instantiates a new NasMountCredentialParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNasMountCredentialParamsWithDefaults

`func NewNasMountCredentialParamsWithDefaults() *NasMountCredentialParams`

NewNasMountCredentialParamsWithDefaults instantiates a new NasMountCredentialParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *NasMountCredentialParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *NasMountCredentialParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *NasMountCredentialParams) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *NasMountCredentialParams) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *NasMountCredentialParams) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *NasMountCredentialParams) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDomainController

`func (o *NasMountCredentialParams) GetDomainController() string`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *NasMountCredentialParams) GetDomainControllerOk() (*string, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *NasMountCredentialParams) SetDomainController(v string)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *NasMountCredentialParams) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### SetDomainControllerNil

`func (o *NasMountCredentialParams) SetDomainControllerNil(b bool)`

 SetDomainControllerNil sets the value for DomainController to be an explicit nil

### UnsetDomainController
`func (o *NasMountCredentialParams) UnsetDomainController()`

UnsetDomainController ensures that no value is present for DomainController, not even an explicit nil
### GetManagePasswordByCohesity

`func (o *NasMountCredentialParams) GetManagePasswordByCohesity() bool`

GetManagePasswordByCohesity returns the ManagePasswordByCohesity field if non-nil, zero value otherwise.

### GetManagePasswordByCohesityOk

`func (o *NasMountCredentialParams) GetManagePasswordByCohesityOk() (*bool, bool)`

GetManagePasswordByCohesityOk returns a tuple with the ManagePasswordByCohesity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagePasswordByCohesity

`func (o *NasMountCredentialParams) SetManagePasswordByCohesity(v bool)`

SetManagePasswordByCohesity sets ManagePasswordByCohesity field to given value.

### HasManagePasswordByCohesity

`func (o *NasMountCredentialParams) HasManagePasswordByCohesity() bool`

HasManagePasswordByCohesity returns a boolean if a field has been set.

### SetManagePasswordByCohesityNil

`func (o *NasMountCredentialParams) SetManagePasswordByCohesityNil(b bool)`

 SetManagePasswordByCohesityNil sets the value for ManagePasswordByCohesity to be an explicit nil

### UnsetManagePasswordByCohesity
`func (o *NasMountCredentialParams) UnsetManagePasswordByCohesity()`

UnsetManagePasswordByCohesity ensures that no value is present for ManagePasswordByCohesity, not even an explicit nil
### GetNasProtocol

`func (o *NasMountCredentialParams) GetNasProtocol() string`

GetNasProtocol returns the NasProtocol field if non-nil, zero value otherwise.

### GetNasProtocolOk

`func (o *NasMountCredentialParams) GetNasProtocolOk() (*string, bool)`

GetNasProtocolOk returns a tuple with the NasProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasProtocol

`func (o *NasMountCredentialParams) SetNasProtocol(v string)`

SetNasProtocol sets NasProtocol field to given value.

### HasNasProtocol

`func (o *NasMountCredentialParams) HasNasProtocol() bool`

HasNasProtocol returns a boolean if a field has been set.

### SetNasProtocolNil

`func (o *NasMountCredentialParams) SetNasProtocolNil(b bool)`

 SetNasProtocolNil sets the value for NasProtocol to be an explicit nil

### UnsetNasProtocol
`func (o *NasMountCredentialParams) UnsetNasProtocol()`

UnsetNasProtocol ensures that no value is present for NasProtocol, not even an explicit nil
### GetNasType

`func (o *NasMountCredentialParams) GetNasType() string`

GetNasType returns the NasType field if non-nil, zero value otherwise.

### GetNasTypeOk

`func (o *NasMountCredentialParams) GetNasTypeOk() (*string, bool)`

GetNasTypeOk returns a tuple with the NasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasType

`func (o *NasMountCredentialParams) SetNasType(v string)`

SetNasType sets NasType field to given value.

### HasNasType

`func (o *NasMountCredentialParams) HasNasType() bool`

HasNasType returns a boolean if a field has been set.

### SetNasTypeNil

`func (o *NasMountCredentialParams) SetNasTypeNil(b bool)`

 SetNasTypeNil sets the value for NasType to be an explicit nil

### UnsetNasType
`func (o *NasMountCredentialParams) UnsetNasType()`

UnsetNasType ensures that no value is present for NasType, not even an explicit nil
### GetPassword

`func (o *NasMountCredentialParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NasMountCredentialParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NasMountCredentialParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NasMountCredentialParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *NasMountCredentialParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *NasMountCredentialParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSkipValidation

`func (o *NasMountCredentialParams) GetSkipValidation() bool`

GetSkipValidation returns the SkipValidation field if non-nil, zero value otherwise.

### GetSkipValidationOk

`func (o *NasMountCredentialParams) GetSkipValidationOk() (*bool, bool)`

GetSkipValidationOk returns a tuple with the SkipValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipValidation

`func (o *NasMountCredentialParams) SetSkipValidation(v bool)`

SetSkipValidation sets SkipValidation field to given value.

### HasSkipValidation

`func (o *NasMountCredentialParams) HasSkipValidation() bool`

HasSkipValidation returns a boolean if a field has been set.

### SetSkipValidationNil

`func (o *NasMountCredentialParams) SetSkipValidationNil(b bool)`

 SetSkipValidationNil sets the value for SkipValidation to be an explicit nil

### UnsetSkipValidation
`func (o *NasMountCredentialParams) UnsetSkipValidation()`

UnsetSkipValidation ensures that no value is present for SkipValidation, not even an explicit nil
### GetUsername

`func (o *NasMountCredentialParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NasMountCredentialParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NasMountCredentialParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NasMountCredentialParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *NasMountCredentialParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *NasMountCredentialParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


