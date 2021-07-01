# IsilonNfsMountPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessZoneName** | Pointer to **NullableString** | Specifies the Access Zone name. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the NFS mount point. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the Id of the NFS export. | [optional] 

## Methods

### NewIsilonNfsMountPoint

`func NewIsilonNfsMountPoint() *IsilonNfsMountPoint`

NewIsilonNfsMountPoint instantiates a new IsilonNfsMountPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsilonNfsMountPointWithDefaults

`func NewIsilonNfsMountPointWithDefaults() *IsilonNfsMountPoint`

NewIsilonNfsMountPointWithDefaults instantiates a new IsilonNfsMountPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessZoneName

`func (o *IsilonNfsMountPoint) GetAccessZoneName() string`

GetAccessZoneName returns the AccessZoneName field if non-nil, zero value otherwise.

### GetAccessZoneNameOk

`func (o *IsilonNfsMountPoint) GetAccessZoneNameOk() (*string, bool)`

GetAccessZoneNameOk returns a tuple with the AccessZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessZoneName

`func (o *IsilonNfsMountPoint) SetAccessZoneName(v string)`

SetAccessZoneName sets AccessZoneName field to given value.

### HasAccessZoneName

`func (o *IsilonNfsMountPoint) HasAccessZoneName() bool`

HasAccessZoneName returns a boolean if a field has been set.

### SetAccessZoneNameNil

`func (o *IsilonNfsMountPoint) SetAccessZoneNameNil(b bool)`

 SetAccessZoneNameNil sets the value for AccessZoneName to be an explicit nil

### UnsetAccessZoneName
`func (o *IsilonNfsMountPoint) UnsetAccessZoneName()`

UnsetAccessZoneName ensures that no value is present for AccessZoneName, not even an explicit nil
### GetDescription

`func (o *IsilonNfsMountPoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IsilonNfsMountPoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IsilonNfsMountPoint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IsilonNfsMountPoint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IsilonNfsMountPoint) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IsilonNfsMountPoint) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *IsilonNfsMountPoint) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IsilonNfsMountPoint) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IsilonNfsMountPoint) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IsilonNfsMountPoint) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IsilonNfsMountPoint) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IsilonNfsMountPoint) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


