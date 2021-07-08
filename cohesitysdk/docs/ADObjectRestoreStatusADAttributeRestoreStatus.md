# ADObjectRestoreStatusADAttributeRestoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttrstatusVec** | Pointer to [**[]ErrorProto**](ErrorProto.md) | Error status. If the &#39;attrstatus_vec&#39; is empty or contains kNoError, treat the attribute restore as success. For multi-valued properties such as &#39;memberOf&#39;, this vector may contain failure to add or remove specific value within the multi-value set. | [optional] 
**LdapName** | Pointer to **NullableString** | LDAP name of the attribute. | [optional] 

## Methods

### NewADObjectRestoreStatusADAttributeRestoreStatus

`func NewADObjectRestoreStatusADAttributeRestoreStatus() *ADObjectRestoreStatusADAttributeRestoreStatus`

NewADObjectRestoreStatusADAttributeRestoreStatus instantiates a new ADObjectRestoreStatusADAttributeRestoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADObjectRestoreStatusADAttributeRestoreStatusWithDefaults

`func NewADObjectRestoreStatusADAttributeRestoreStatusWithDefaults() *ADObjectRestoreStatusADAttributeRestoreStatus`

NewADObjectRestoreStatusADAttributeRestoreStatusWithDefaults instantiates a new ADObjectRestoreStatusADAttributeRestoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttrstatusVec

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) GetAttrstatusVec() []ErrorProto`

GetAttrstatusVec returns the AttrstatusVec field if non-nil, zero value otherwise.

### GetAttrstatusVecOk

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) GetAttrstatusVecOk() (*[]ErrorProto, bool)`

GetAttrstatusVecOk returns a tuple with the AttrstatusVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrstatusVec

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) SetAttrstatusVec(v []ErrorProto)`

SetAttrstatusVec sets AttrstatusVec field to given value.

### HasAttrstatusVec

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) HasAttrstatusVec() bool`

HasAttrstatusVec returns a boolean if a field has been set.

### SetAttrstatusVecNil

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) SetAttrstatusVecNil(b bool)`

 SetAttrstatusVecNil sets the value for AttrstatusVec to be an explicit nil

### UnsetAttrstatusVec
`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) UnsetAttrstatusVec()`

UnsetAttrstatusVec ensures that no value is present for AttrstatusVec, not even an explicit nil
### GetLdapName

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) GetLdapName() string`

GetLdapName returns the LdapName field if non-nil, zero value otherwise.

### GetLdapNameOk

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) GetLdapNameOk() (*string, bool)`

GetLdapNameOk returns a tuple with the LdapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapName

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) SetLdapName(v string)`

SetLdapName sets LdapName field to given value.

### HasLdapName

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) HasLdapName() bool`

HasLdapName returns a boolean if a field has been set.

### SetLdapNameNil

`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) SetLdapNameNil(b bool)`

 SetLdapNameNil sets the value for LdapName to be an explicit nil

### UnsetLdapName
`func (o *ADObjectRestoreStatusADAttributeRestoreStatus) UnsetLdapName()`

UnsetLdapName ensures that no value is present for LdapName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


