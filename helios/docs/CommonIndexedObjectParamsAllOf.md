# CommonIndexedObjectParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path of the object. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | \&quot;Specifies the protection group id which contains this object.\&quot; | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | \&quot;Specifies the protection group name which contains this object.\&quot; | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | \&quot;Specifies the Storage Domain id where the backup data of Object is present.\&quot; | [optional] 
**SourceInfo** | Pointer to [**ObjectSummary**](ObjectSummary.md) | Specifies the Source Object information. | [optional] 

## Methods

### NewCommonIndexedObjectParamsAllOf

`func NewCommonIndexedObjectParamsAllOf() *CommonIndexedObjectParamsAllOf`

NewCommonIndexedObjectParamsAllOf instantiates a new CommonIndexedObjectParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonIndexedObjectParamsAllOfWithDefaults

`func NewCommonIndexedObjectParamsAllOfWithDefaults() *CommonIndexedObjectParamsAllOf`

NewCommonIndexedObjectParamsAllOfWithDefaults instantiates a new CommonIndexedObjectParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommonIndexedObjectParamsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommonIndexedObjectParamsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommonIndexedObjectParamsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommonIndexedObjectParamsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CommonIndexedObjectParamsAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CommonIndexedObjectParamsAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *CommonIndexedObjectParamsAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommonIndexedObjectParamsAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommonIndexedObjectParamsAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CommonIndexedObjectParamsAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *CommonIndexedObjectParamsAllOf) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CommonIndexedObjectParamsAllOf) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProtectionGroupId

`func (o *CommonIndexedObjectParamsAllOf) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *CommonIndexedObjectParamsAllOf) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *CommonIndexedObjectParamsAllOf) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *CommonIndexedObjectParamsAllOf) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *CommonIndexedObjectParamsAllOf) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *CommonIndexedObjectParamsAllOf) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *CommonIndexedObjectParamsAllOf) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *CommonIndexedObjectParamsAllOf) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *CommonIndexedObjectParamsAllOf) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *CommonIndexedObjectParamsAllOf) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *CommonIndexedObjectParamsAllOf) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *CommonIndexedObjectParamsAllOf) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetStorageDomainId

`func (o *CommonIndexedObjectParamsAllOf) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *CommonIndexedObjectParamsAllOf) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *CommonIndexedObjectParamsAllOf) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *CommonIndexedObjectParamsAllOf) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *CommonIndexedObjectParamsAllOf) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *CommonIndexedObjectParamsAllOf) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetSourceInfo

`func (o *CommonIndexedObjectParamsAllOf) GetSourceInfo() ObjectSummary`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *CommonIndexedObjectParamsAllOf) GetSourceInfoOk() (*ObjectSummary, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *CommonIndexedObjectParamsAllOf) SetSourceInfo(v ObjectSummary)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *CommonIndexedObjectParamsAllOf) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


