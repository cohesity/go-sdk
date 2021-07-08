# AdRootTopologyObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildObjects** | Pointer to **[]map[string]interface{}** | Specifies the array of children of this object. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the &#39;description&#39; of an object. | [optional] 
**DestGuid** | Pointer to **NullableString** | Specifies the guid of matching &#39;source_guid&#39; from production AD. This is looked up  based on source_guid or distinguishedName attribute value. | [optional] 
**DisplayName** | Pointer to **NullableString** | Specifies the display name of the object in AD Topology tree. | [optional] 
**DistinguishedName** | Pointer to **NullableString** | Specifies the distinguished name of the object in AD Topology tree. Eg: CN&#x3D;Jone Doe,OU&#x3D;Users,DC&#x3D;corp,DC&#x3D;cohesity,DC&#x3D;com | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the AD error while fetching the ADRootTopologyObject. | [optional] 
**ObjectClass** | Pointer to **NullableString** | Specifies the LDAP class name such as &#39;user&#39;,&#39;computer&#39;, &#39;organizationalUnit&#39;. | [optional] 
**SourceGuid** | Pointer to **NullableString** | Specifies the guid string of the object in AD snapshot database. | [optional] 

## Methods

### NewAdRootTopologyObject

`func NewAdRootTopologyObject() *AdRootTopologyObject`

NewAdRootTopologyObject instantiates a new AdRootTopologyObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdRootTopologyObjectWithDefaults

`func NewAdRootTopologyObjectWithDefaults() *AdRootTopologyObject`

NewAdRootTopologyObjectWithDefaults instantiates a new AdRootTopologyObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildObjects

`func (o *AdRootTopologyObject) GetChildObjects() []map[string]interface{}`

GetChildObjects returns the ChildObjects field if non-nil, zero value otherwise.

### GetChildObjectsOk

`func (o *AdRootTopologyObject) GetChildObjectsOk() (*[]map[string]interface{}, bool)`

GetChildObjectsOk returns a tuple with the ChildObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildObjects

`func (o *AdRootTopologyObject) SetChildObjects(v []map[string]interface{})`

SetChildObjects sets ChildObjects field to given value.

### HasChildObjects

`func (o *AdRootTopologyObject) HasChildObjects() bool`

HasChildObjects returns a boolean if a field has been set.

### SetChildObjectsNil

`func (o *AdRootTopologyObject) SetChildObjectsNil(b bool)`

 SetChildObjectsNil sets the value for ChildObjects to be an explicit nil

### UnsetChildObjects
`func (o *AdRootTopologyObject) UnsetChildObjects()`

UnsetChildObjects ensures that no value is present for ChildObjects, not even an explicit nil
### GetDescription

`func (o *AdRootTopologyObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdRootTopologyObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdRootTopologyObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdRootTopologyObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AdRootTopologyObject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdRootTopologyObject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDestGuid

`func (o *AdRootTopologyObject) GetDestGuid() string`

GetDestGuid returns the DestGuid field if non-nil, zero value otherwise.

### GetDestGuidOk

`func (o *AdRootTopologyObject) GetDestGuidOk() (*string, bool)`

GetDestGuidOk returns a tuple with the DestGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestGuid

`func (o *AdRootTopologyObject) SetDestGuid(v string)`

SetDestGuid sets DestGuid field to given value.

### HasDestGuid

`func (o *AdRootTopologyObject) HasDestGuid() bool`

HasDestGuid returns a boolean if a field has been set.

### SetDestGuidNil

`func (o *AdRootTopologyObject) SetDestGuidNil(b bool)`

 SetDestGuidNil sets the value for DestGuid to be an explicit nil

### UnsetDestGuid
`func (o *AdRootTopologyObject) UnsetDestGuid()`

UnsetDestGuid ensures that no value is present for DestGuid, not even an explicit nil
### GetDisplayName

`func (o *AdRootTopologyObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdRootTopologyObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdRootTopologyObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdRootTopologyObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AdRootTopologyObject) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AdRootTopologyObject) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDistinguishedName

`func (o *AdRootTopologyObject) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *AdRootTopologyObject) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *AdRootTopologyObject) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *AdRootTopologyObject) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *AdRootTopologyObject) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *AdRootTopologyObject) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetErrorMessage

`func (o *AdRootTopologyObject) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AdRootTopologyObject) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AdRootTopologyObject) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AdRootTopologyObject) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AdRootTopologyObject) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AdRootTopologyObject) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetObjectClass

`func (o *AdRootTopologyObject) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *AdRootTopologyObject) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *AdRootTopologyObject) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *AdRootTopologyObject) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### SetObjectClassNil

`func (o *AdRootTopologyObject) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *AdRootTopologyObject) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetSourceGuid

`func (o *AdRootTopologyObject) GetSourceGuid() string`

GetSourceGuid returns the SourceGuid field if non-nil, zero value otherwise.

### GetSourceGuidOk

`func (o *AdRootTopologyObject) GetSourceGuidOk() (*string, bool)`

GetSourceGuidOk returns a tuple with the SourceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuid

`func (o *AdRootTopologyObject) SetSourceGuid(v string)`

SetSourceGuid sets SourceGuid field to given value.

### HasSourceGuid

`func (o *AdRootTopologyObject) HasSourceGuid() bool`

HasSourceGuid returns a boolean if a field has been set.

### SetSourceGuidNil

`func (o *AdRootTopologyObject) SetSourceGuidNil(b bool)`

 SetSourceGuidNil sets the value for SourceGuid to be an explicit nil

### UnsetSourceGuid
`func (o *AdRootTopologyObject) UnsetSourceGuid()`

UnsetSourceGuid ensures that no value is present for SourceGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


