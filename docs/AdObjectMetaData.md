# AdObjectMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinguishedName** | Pointer to **NullableString** | Specifies the Distinguished name of the AD object. | [optional] 
**Domain** | Pointer to **NullableString** | Domain of the AD object. | [optional] 
**Email** | Pointer to **NullableString** | Specifies the email of the AD object of type user or group. | [optional] 
**Guid** | Pointer to **NullableString** | Specifies the Guid of the AD object. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the AD object. | [optional] 
**ObjectType** | Pointer to **NullableString** | Specifies the type of the AD Object. The type may be user, computer, group or ou. | [optional] 
**SamAccountName** | Pointer to **NullableString** | Specifies the sam account name of the AD object. | [optional] 

## Methods

### NewAdObjectMetaData

`func NewAdObjectMetaData() *AdObjectMetaData`

NewAdObjectMetaData instantiates a new AdObjectMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdObjectMetaDataWithDefaults

`func NewAdObjectMetaDataWithDefaults() *AdObjectMetaData`

NewAdObjectMetaDataWithDefaults instantiates a new AdObjectMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinguishedName

`func (o *AdObjectMetaData) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *AdObjectMetaData) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *AdObjectMetaData) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *AdObjectMetaData) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *AdObjectMetaData) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *AdObjectMetaData) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetDomain

`func (o *AdObjectMetaData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AdObjectMetaData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AdObjectMetaData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AdObjectMetaData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *AdObjectMetaData) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *AdObjectMetaData) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEmail

`func (o *AdObjectMetaData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdObjectMetaData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdObjectMetaData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AdObjectMetaData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AdObjectMetaData) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdObjectMetaData) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGuid

`func (o *AdObjectMetaData) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *AdObjectMetaData) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *AdObjectMetaData) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *AdObjectMetaData) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### SetGuidNil

`func (o *AdObjectMetaData) SetGuidNil(b bool)`

 SetGuidNil sets the value for Guid to be an explicit nil

### UnsetGuid
`func (o *AdObjectMetaData) UnsetGuid()`

UnsetGuid ensures that no value is present for Guid, not even an explicit nil
### GetName

`func (o *AdObjectMetaData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdObjectMetaData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdObjectMetaData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdObjectMetaData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdObjectMetaData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdObjectMetaData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetObjectType

`func (o *AdObjectMetaData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdObjectMetaData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdObjectMetaData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AdObjectMetaData) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *AdObjectMetaData) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *AdObjectMetaData) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetSamAccountName

`func (o *AdObjectMetaData) GetSamAccountName() string`

GetSamAccountName returns the SamAccountName field if non-nil, zero value otherwise.

### GetSamAccountNameOk

`func (o *AdObjectMetaData) GetSamAccountNameOk() (*string, bool)`

GetSamAccountNameOk returns a tuple with the SamAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamAccountName

`func (o *AdObjectMetaData) SetSamAccountName(v string)`

SetSamAccountName sets SamAccountName field to given value.

### HasSamAccountName

`func (o *AdObjectMetaData) HasSamAccountName() bool`

HasSamAccountName returns a boolean if a field has been set.

### SetSamAccountNameNil

`func (o *AdObjectMetaData) SetSamAccountNameNil(b bool)`

 SetSamAccountNameNil sets the value for SamAccountName to be an explicit nil

### UnsetSamAccountName
`func (o *AdObjectMetaData) UnsetSamAccountName()`

UnsetSamAccountName ensures that no value is present for SamAccountName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


