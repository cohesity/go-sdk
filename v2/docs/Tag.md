# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimeUsecs** | Pointer to **NullableInt32** | Specifies the timestamp in microseconds since the epoch when this Tag was created. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Description of the Tag. | [optional] 
**Id** | Pointer to **NullableString** | Specifies unique id of the Tag. | [optional] [readonly] 
**LastUpdatedTimeUsecs** | Pointer to **NullableInt32** | Specifies the timestamp in microseconds since the epoch when this Tag was last updated. | [optional] [readonly] 
**MarkedForDeletion** | Pointer to **NullableBool** | If true, Tag is marked for deletion. | [optional] [readonly] 
**Name** | **NullableString** | Name of the Tag. Name has to be unique under Namespace. | 
**Namespace** | **NullableString** | Namespace of the Tag. This is used to filter tags based on application or usecase. For example all tags related to vcenter can be put under one namespace or different departments could have their own namespaces e.g. finance/tag1 or operations/tag2 etc. | 
**TenantId** | Pointer to **NullableString** | Tenant Id to whom the Tag belongs. | [optional] [readonly] 
**UiColor** | Pointer to **NullableString** | Color of the tag in UI. | [optional] 
**UiPathElements** | Pointer to **[]string** | Path of the tag for UI nesting purposes. | [optional] 

## Methods

### NewTag

`func NewTag(name NullableString, namespace NullableString, ) *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimeUsecs

`func (o *Tag) GetCreatedTimeUsecs() int32`

GetCreatedTimeUsecs returns the CreatedTimeUsecs field if non-nil, zero value otherwise.

### GetCreatedTimeUsecsOk

`func (o *Tag) GetCreatedTimeUsecsOk() (*int32, bool)`

GetCreatedTimeUsecsOk returns a tuple with the CreatedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeUsecs

`func (o *Tag) SetCreatedTimeUsecs(v int32)`

SetCreatedTimeUsecs sets CreatedTimeUsecs field to given value.

### HasCreatedTimeUsecs

`func (o *Tag) HasCreatedTimeUsecs() bool`

HasCreatedTimeUsecs returns a boolean if a field has been set.

### SetCreatedTimeUsecsNil

`func (o *Tag) SetCreatedTimeUsecsNil(b bool)`

 SetCreatedTimeUsecsNil sets the value for CreatedTimeUsecs to be an explicit nil

### UnsetCreatedTimeUsecs
`func (o *Tag) UnsetCreatedTimeUsecs()`

UnsetCreatedTimeUsecs ensures that no value is present for CreatedTimeUsecs, not even an explicit nil
### GetDescription

`func (o *Tag) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tag) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tag) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tag) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Tag) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Tag) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Tag) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tag) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tag) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Tag) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Tag) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Tag) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastUpdatedTimeUsecs

`func (o *Tag) GetLastUpdatedTimeUsecs() int32`

GetLastUpdatedTimeUsecs returns the LastUpdatedTimeUsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeUsecsOk

`func (o *Tag) GetLastUpdatedTimeUsecsOk() (*int32, bool)`

GetLastUpdatedTimeUsecsOk returns a tuple with the LastUpdatedTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeUsecs

`func (o *Tag) SetLastUpdatedTimeUsecs(v int32)`

SetLastUpdatedTimeUsecs sets LastUpdatedTimeUsecs field to given value.

### HasLastUpdatedTimeUsecs

`func (o *Tag) HasLastUpdatedTimeUsecs() bool`

HasLastUpdatedTimeUsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeUsecsNil

`func (o *Tag) SetLastUpdatedTimeUsecsNil(b bool)`

 SetLastUpdatedTimeUsecsNil sets the value for LastUpdatedTimeUsecs to be an explicit nil

### UnsetLastUpdatedTimeUsecs
`func (o *Tag) UnsetLastUpdatedTimeUsecs()`

UnsetLastUpdatedTimeUsecs ensures that no value is present for LastUpdatedTimeUsecs, not even an explicit nil
### GetMarkedForDeletion

`func (o *Tag) GetMarkedForDeletion() bool`

GetMarkedForDeletion returns the MarkedForDeletion field if non-nil, zero value otherwise.

### GetMarkedForDeletionOk

`func (o *Tag) GetMarkedForDeletionOk() (*bool, bool)`

GetMarkedForDeletionOk returns a tuple with the MarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedForDeletion

`func (o *Tag) SetMarkedForDeletion(v bool)`

SetMarkedForDeletion sets MarkedForDeletion field to given value.

### HasMarkedForDeletion

`func (o *Tag) HasMarkedForDeletion() bool`

HasMarkedForDeletion returns a boolean if a field has been set.

### SetMarkedForDeletionNil

`func (o *Tag) SetMarkedForDeletionNil(b bool)`

 SetMarkedForDeletionNil sets the value for MarkedForDeletion to be an explicit nil

### UnsetMarkedForDeletion
`func (o *Tag) UnsetMarkedForDeletion()`

UnsetMarkedForDeletion ensures that no value is present for MarkedForDeletion, not even an explicit nil
### GetName

`func (o *Tag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tag) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Tag) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Tag) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *Tag) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Tag) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Tag) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *Tag) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *Tag) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetTenantId

`func (o *Tag) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Tag) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Tag) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Tag) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Tag) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Tag) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUiColor

`func (o *Tag) GetUiColor() string`

GetUiColor returns the UiColor field if non-nil, zero value otherwise.

### GetUiColorOk

`func (o *Tag) GetUiColorOk() (*string, bool)`

GetUiColorOk returns a tuple with the UiColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiColor

`func (o *Tag) SetUiColor(v string)`

SetUiColor sets UiColor field to given value.

### HasUiColor

`func (o *Tag) HasUiColor() bool`

HasUiColor returns a boolean if a field has been set.

### SetUiColorNil

`func (o *Tag) SetUiColorNil(b bool)`

 SetUiColorNil sets the value for UiColor to be an explicit nil

### UnsetUiColor
`func (o *Tag) UnsetUiColor()`

UnsetUiColor ensures that no value is present for UiColor, not even an explicit nil
### GetUiPathElements

`func (o *Tag) GetUiPathElements() []string`

GetUiPathElements returns the UiPathElements field if non-nil, zero value otherwise.

### GetUiPathElementsOk

`func (o *Tag) GetUiPathElementsOk() (*[]string, bool)`

GetUiPathElementsOk returns a tuple with the UiPathElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiPathElements

`func (o *Tag) SetUiPathElements(v []string)`

SetUiPathElements sets UiPathElements field to given value.

### HasUiPathElements

`func (o *Tag) HasUiPathElements() bool`

HasUiPathElements returns a boolean if a field has been set.

### SetUiPathElementsNil

`func (o *Tag) SetUiPathElementsNil(b bool)`

 SetUiPathElementsNil sets the value for UiPathElements to be an explicit nil

### UnsetUiPathElements
`func (o *Tag) UnsetUiPathElements()`

UnsetUiPathElements ensures that no value is present for UiPathElements, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


