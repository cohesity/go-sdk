# IsilonSmbMountPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessZoneId** | Pointer to **NullableInt64** | Specifies the Access Zone Id. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the NFS mount point. | [optional] 
**ShareName** | Pointer to **NullableString** | Specifies the name of the SMB/CIFS share. | [optional] 

## Methods

### NewIsilonSmbMountPoint

`func NewIsilonSmbMountPoint() *IsilonSmbMountPoint`

NewIsilonSmbMountPoint instantiates a new IsilonSmbMountPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonSmbMountPointWithDefaults

`func NewIsilonSmbMountPointWithDefaults() *IsilonSmbMountPoint`

NewIsilonSmbMountPointWithDefaults instantiates a new IsilonSmbMountPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessZoneId

`func (o *IsilonSmbMountPoint) GetAccessZoneId() int64`

GetAccessZoneId returns the AccessZoneId field if non-nil, zero value otherwise.

### GetAccessZoneIdOk

`func (o *IsilonSmbMountPoint) GetAccessZoneIdOk() (*int64, bool)`

GetAccessZoneIdOk returns a tuple with the AccessZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessZoneId

`func (o *IsilonSmbMountPoint) SetAccessZoneId(v int64)`

SetAccessZoneId sets AccessZoneId field to given value.

### HasAccessZoneId

`func (o *IsilonSmbMountPoint) HasAccessZoneId() bool`

HasAccessZoneId returns a boolean if a field has been set.

### SetAccessZoneIdNil

`func (o *IsilonSmbMountPoint) SetAccessZoneIdNil(b bool)`

 SetAccessZoneIdNil sets the value for AccessZoneId to be an explicit nil

### UnsetAccessZoneId
`func (o *IsilonSmbMountPoint) UnsetAccessZoneId()`

UnsetAccessZoneId ensures that no value is present for AccessZoneId, not even an explicit nil
### GetDescription

`func (o *IsilonSmbMountPoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IsilonSmbMountPoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IsilonSmbMountPoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IsilonSmbMountPoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IsilonSmbMountPoint) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IsilonSmbMountPoint) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShareName

`func (o *IsilonSmbMountPoint) GetShareName() string`

GetShareName returns the ShareName field if non-nil, zero value otherwise.

### GetShareNameOk

`func (o *IsilonSmbMountPoint) GetShareNameOk() (*string, bool)`

GetShareNameOk returns a tuple with the ShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareName

`func (o *IsilonSmbMountPoint) SetShareName(v string)`

SetShareName sets ShareName field to given value.

### HasShareName

`func (o *IsilonSmbMountPoint) HasShareName() bool`

HasShareName returns a boolean if a field has been set.

### SetShareNameNil

`func (o *IsilonSmbMountPoint) SetShareNameNil(b bool)`

 SetShareNameNil sets the value for ShareName to be an explicit nil

### UnsetShareName
`func (o *IsilonSmbMountPoint) UnsetShareName()`

UnsetShareName ensures that no value is present for ShareName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


