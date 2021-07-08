# ADObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies the &#39;description&#39; of an AD object. | [optional] 
**DestinationGuid** | Pointer to **NullableString** | Specifies the guid of object in the Production AD which is equivalent to the object in the Snapshot AD. | [optional] 
**DisplayName** | Pointer to **NullableString** | Specifies the display name of the AD object. | [optional] 
**DistinguishedName** | Pointer to **NullableString** | Specifies the distinguished name of the AD object. Eg: CN&#x3D;Jone Doe,OU&#x3D;Users,DC&#x3D;corp,DC&#x3D;cohesity,DC&#x3D;com | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error message while fetching the AD object. | [optional] 
**ObjectClass** | Pointer to **NullableString** | Specifies the class name of an AD Object such as &#39;user&#39;,&#39;computer&#39;, &#39;organizationalUnit&#39;. | [optional] 
**SearchResultFlags** | Pointer to **[]string** | Specifies the SearchResultFlags of the AD object. &#39;kEqual&#39; indicates the AD Object from Snapshot and Production AD are equal. &#39;kNotEqual&#39; indicates the AD Object from snapshot and production AD are not equal. &#39;kRestorePasswordRequired&#39; indicates when restoring this AD Object from Snapshot AD to Production AD, a password is required. &#39;kMovedOnDestination&#39; indicates the object has moved to another container or OU in Production AD compared to  Snapshot AD. &#39;kDisableSupported&#39; indicates the enable and disable is supported on the AD Object. AD Objects of type &#39;User&#39; and &#39;Computers&#39; support this operation. | [optional] 
**SourceGuid** | Pointer to **NullableString** | Specifies the guid of the AD object in Snapshot AD. | [optional] 

## Methods

### NewADObject

`func NewADObject() *ADObject`

NewADObject instantiates a new ADObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADObjectWithDefaults

`func NewADObjectWithDefaults() *ADObject`

NewADObjectWithDefaults instantiates a new ADObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ADObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ADObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ADObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ADObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ADObject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ADObject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDestinationGuid

`func (o *ADObject) GetDestinationGuid() string`

GetDestinationGuid returns the DestinationGuid field if non-nil, zero value otherwise.

### GetDestinationGuidOk

`func (o *ADObject) GetDestinationGuidOk() (*string, bool)`

GetDestinationGuidOk returns a tuple with the DestinationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGuid

`func (o *ADObject) SetDestinationGuid(v string)`

SetDestinationGuid sets DestinationGuid field to given value.

### HasDestinationGuid

`func (o *ADObject) HasDestinationGuid() bool`

HasDestinationGuid returns a boolean if a field has been set.

### SetDestinationGuidNil

`func (o *ADObject) SetDestinationGuidNil(b bool)`

 SetDestinationGuidNil sets the value for DestinationGuid to be an explicit nil

### UnsetDestinationGuid
`func (o *ADObject) UnsetDestinationGuid()`

UnsetDestinationGuid ensures that no value is present for DestinationGuid, not even an explicit nil
### GetDisplayName

`func (o *ADObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ADObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ADObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ADObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ADObject) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ADObject) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDistinguishedName

`func (o *ADObject) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *ADObject) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *ADObject) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *ADObject) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *ADObject) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *ADObject) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetErrorMessage

`func (o *ADObject) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ADObject) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ADObject) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ADObject) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ADObject) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ADObject) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetObjectClass

`func (o *ADObject) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ADObject) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ADObject) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ADObject) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### SetObjectClassNil

`func (o *ADObject) SetObjectClassNil(b bool)`

 SetObjectClassNil sets the value for ObjectClass to be an explicit nil

### UnsetObjectClass
`func (o *ADObject) UnsetObjectClass()`

UnsetObjectClass ensures that no value is present for ObjectClass, not even an explicit nil
### GetSearchResultFlags

`func (o *ADObject) GetSearchResultFlags() []string`

GetSearchResultFlags returns the SearchResultFlags field if non-nil, zero value otherwise.

### GetSearchResultFlagsOk

`func (o *ADObject) GetSearchResultFlagsOk() (*[]string, bool)`

GetSearchResultFlagsOk returns a tuple with the SearchResultFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResultFlags

`func (o *ADObject) SetSearchResultFlags(v []string)`

SetSearchResultFlags sets SearchResultFlags field to given value.

### HasSearchResultFlags

`func (o *ADObject) HasSearchResultFlags() bool`

HasSearchResultFlags returns a boolean if a field has been set.

### SetSearchResultFlagsNil

`func (o *ADObject) SetSearchResultFlagsNil(b bool)`

 SetSearchResultFlagsNil sets the value for SearchResultFlags to be an explicit nil

### UnsetSearchResultFlags
`func (o *ADObject) UnsetSearchResultFlags()`

UnsetSearchResultFlags ensures that no value is present for SearchResultFlags, not even an explicit nil
### GetSourceGuid

`func (o *ADObject) GetSourceGuid() string`

GetSourceGuid returns the SourceGuid field if non-nil, zero value otherwise.

### GetSourceGuidOk

`func (o *ADObject) GetSourceGuidOk() (*string, bool)`

GetSourceGuidOk returns a tuple with the SourceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuid

`func (o *ADObject) SetSourceGuid(v string)`

SetSourceGuid sets SourceGuid field to given value.

### HasSourceGuid

`func (o *ADObject) HasSourceGuid() bool`

HasSourceGuid returns a boolean if a field has been set.

### SetSourceGuidNil

`func (o *ADObject) SetSourceGuidNil(b bool)`

 SetSourceGuidNil sets the value for SourceGuid to be an explicit nil

### UnsetSourceGuid
`func (o *ADObject) UnsetSourceGuid()`

UnsetSourceGuid ensures that no value is present for SourceGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


