# ComparedADObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdAttributes** | Pointer to [**[]AdAttribute**](AdAttribute.md) | Specifies the list of AD attributes for the AD object. | [optional] 
**AdObjectFlags** | Pointer to **[]string** | Specifies the flags related to this AD Object. &#39;kEqual&#39; indicates all the attributes of the AD Object on the Snapshot and Production are equal. &#39;kNotEqual&#39; indicates atleast one of the attribute of the AD Object on the Snapshot and Production AD are not equal. &#39;kRestorePasswordRequired&#39; indicates the AD Object is of &#39;User&#39; object class type. when restoring this object from Snapshot AD to Priduction AD, a password is required. &#39;kMovedOnDestination&#39; indicates the object has moved to another container or OU in production AD compared to AD snapshot. In this case, the distinguishedName will be different for these objects &#39;kDestinationNotFound&#39; indicates the object corresponding to dest_guid specified is missing from Production AD. Caller should check this flag and empty &#39;dest_guid&#39; first to find out destination is missing. &#39;kDisableSupported&#39; indicates the enable and disable is supported on the AD Object. AD Objects of type &#39;User&#39; and &#39;Computers&#39; support this operation. | [optional] 
**DestinationGuid** | Pointer to **NullableString** | Specifies the guid of the object in the Production AD which is equivalent to the one in the Snapshot AD. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Specifies the error message while fetching the AD object. | [optional] 
**MismatchAttrCount** | Pointer to **NullableInt32** | Specifies the number of attributes of AD Object mismatched on the Snapshot and Production AD. | [optional] 
**SourceGuid** | Pointer to **NullableString** | Specifies the guid of the AD object in the Snapshot AD. | [optional] 

## Methods

### NewComparedADObject

`func NewComparedADObject() *ComparedADObject`

NewComparedADObject instantiates a new ComparedADObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparedADObjectWithDefaults

`func NewComparedADObjectWithDefaults() *ComparedADObject`

NewComparedADObjectWithDefaults instantiates a new ComparedADObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdAttributes

`func (o *ComparedADObject) GetAdAttributes() []AdAttribute`

GetAdAttributes returns the AdAttributes field if non-nil, zero value otherwise.

### GetAdAttributesOk

`func (o *ComparedADObject) GetAdAttributesOk() (*[]AdAttribute, bool)`

GetAdAttributesOk returns a tuple with the AdAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAttributes

`func (o *ComparedADObject) SetAdAttributes(v []AdAttribute)`

SetAdAttributes sets AdAttributes field to given value.

### HasAdAttributes

`func (o *ComparedADObject) HasAdAttributes() bool`

HasAdAttributes returns a boolean if a field has been set.

### SetAdAttributesNil

`func (o *ComparedADObject) SetAdAttributesNil(b bool)`

 SetAdAttributesNil sets the value for AdAttributes to be an explicit nil

### UnsetAdAttributes
`func (o *ComparedADObject) UnsetAdAttributes()`

UnsetAdAttributes ensures that no value is present for AdAttributes, not even an explicit nil
### GetAdObjectFlags

`func (o *ComparedADObject) GetAdObjectFlags() []string`

GetAdObjectFlags returns the AdObjectFlags field if non-nil, zero value otherwise.

### GetAdObjectFlagsOk

`func (o *ComparedADObject) GetAdObjectFlagsOk() (*[]string, bool)`

GetAdObjectFlagsOk returns a tuple with the AdObjectFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdObjectFlags

`func (o *ComparedADObject) SetAdObjectFlags(v []string)`

SetAdObjectFlags sets AdObjectFlags field to given value.

### HasAdObjectFlags

`func (o *ComparedADObject) HasAdObjectFlags() bool`

HasAdObjectFlags returns a boolean if a field has been set.

### SetAdObjectFlagsNil

`func (o *ComparedADObject) SetAdObjectFlagsNil(b bool)`

 SetAdObjectFlagsNil sets the value for AdObjectFlags to be an explicit nil

### UnsetAdObjectFlags
`func (o *ComparedADObject) UnsetAdObjectFlags()`

UnsetAdObjectFlags ensures that no value is present for AdObjectFlags, not even an explicit nil
### GetDestinationGuid

`func (o *ComparedADObject) GetDestinationGuid() string`

GetDestinationGuid returns the DestinationGuid field if non-nil, zero value otherwise.

### GetDestinationGuidOk

`func (o *ComparedADObject) GetDestinationGuidOk() (*string, bool)`

GetDestinationGuidOk returns a tuple with the DestinationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGuid

`func (o *ComparedADObject) SetDestinationGuid(v string)`

SetDestinationGuid sets DestinationGuid field to given value.

### HasDestinationGuid

`func (o *ComparedADObject) HasDestinationGuid() bool`

HasDestinationGuid returns a boolean if a field has been set.

### SetDestinationGuidNil

`func (o *ComparedADObject) SetDestinationGuidNil(b bool)`

 SetDestinationGuidNil sets the value for DestinationGuid to be an explicit nil

### UnsetDestinationGuid
`func (o *ComparedADObject) UnsetDestinationGuid()`

UnsetDestinationGuid ensures that no value is present for DestinationGuid, not even an explicit nil
### GetErrorMessage

`func (o *ComparedADObject) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ComparedADObject) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ComparedADObject) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ComparedADObject) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ComparedADObject) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ComparedADObject) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetMismatchAttrCount

`func (o *ComparedADObject) GetMismatchAttrCount() int32`

GetMismatchAttrCount returns the MismatchAttrCount field if non-nil, zero value otherwise.

### GetMismatchAttrCountOk

`func (o *ComparedADObject) GetMismatchAttrCountOk() (*int32, bool)`

GetMismatchAttrCountOk returns a tuple with the MismatchAttrCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchAttrCount

`func (o *ComparedADObject) SetMismatchAttrCount(v int32)`

SetMismatchAttrCount sets MismatchAttrCount field to given value.

### HasMismatchAttrCount

`func (o *ComparedADObject) HasMismatchAttrCount() bool`

HasMismatchAttrCount returns a boolean if a field has been set.

### SetMismatchAttrCountNil

`func (o *ComparedADObject) SetMismatchAttrCountNil(b bool)`

 SetMismatchAttrCountNil sets the value for MismatchAttrCount to be an explicit nil

### UnsetMismatchAttrCount
`func (o *ComparedADObject) UnsetMismatchAttrCount()`

UnsetMismatchAttrCount ensures that no value is present for MismatchAttrCount, not even an explicit nil
### GetSourceGuid

`func (o *ComparedADObject) GetSourceGuid() string`

GetSourceGuid returns the SourceGuid field if non-nil, zero value otherwise.

### GetSourceGuidOk

`func (o *ComparedADObject) GetSourceGuidOk() (*string, bool)`

GetSourceGuidOk returns a tuple with the SourceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGuid

`func (o *ComparedADObject) SetSourceGuid(v string)`

SetSourceGuid sets SourceGuid field to given value.

### HasSourceGuid

`func (o *ComparedADObject) HasSourceGuid() bool`

HasSourceGuid returns a boolean if a field has been set.

### SetSourceGuidNil

`func (o *ComparedADObject) SetSourceGuidNil(b bool)`

 SetSourceGuidNil sets the value for SourceGuid to be an explicit nil

### UnsetSourceGuid
`func (o *ComparedADObject) UnsetSourceGuid()`

UnsetSourceGuid ensures that no value is present for SourceGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


