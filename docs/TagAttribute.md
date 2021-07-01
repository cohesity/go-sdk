# TagAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcpTagType** | Pointer to **NullableString** | Specifies the tag attribute type of GCP. Going forward, there will be an additional TagTypes for AWS as well. Specifies the type of the tag GCP entity refers to. &#39;kNetworkTag&#39; indicates a network tag present on instances. &#39;kLabel&#39; indicates a label (key-value pair) present on instances. &#39;kCustomMetadata&#39; indicates custom Metadata (key-value pair) present on instances. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the Coheisty id of the VM tag. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the VMware name of the VM tag. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the VMware Universally Unique Identifier (UUID) of the VM tag. | [optional] 

## Methods

### NewTagAttribute

`func NewTagAttribute() *TagAttribute`

NewTagAttribute instantiates a new TagAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagAttributeWithDefaults

`func NewTagAttributeWithDefaults() *TagAttribute`

NewTagAttributeWithDefaults instantiates a new TagAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcpTagType

`func (o *TagAttribute) GetGcpTagType() string`

GetGcpTagType returns the GcpTagType field if non-nil, zero value otherwise.

### GetGcpTagTypeOk

`func (o *TagAttribute) GetGcpTagTypeOk() (*string, bool)`

GetGcpTagTypeOk returns a tuple with the GcpTagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTagType

`func (o *TagAttribute) SetGcpTagType(v string)`

SetGcpTagType sets GcpTagType field to given value.

### HasGcpTagType

`func (o *TagAttribute) HasGcpTagType() bool`

HasGcpTagType returns a boolean if a field has been set.

### SetGcpTagTypeNil

`func (o *TagAttribute) SetGcpTagTypeNil(b bool)`

 SetGcpTagTypeNil sets the value for GcpTagType to be an explicit nil

### UnsetGcpTagType
`func (o *TagAttribute) UnsetGcpTagType()`

UnsetGcpTagType ensures that no value is present for GcpTagType, not even an explicit nil
### GetId

`func (o *TagAttribute) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagAttribute) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagAttribute) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TagAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TagAttribute) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TagAttribute) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TagAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TagAttribute) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TagAttribute) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUuid

`func (o *TagAttribute) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TagAttribute) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TagAttribute) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TagAttribute) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *TagAttribute) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *TagAttribute) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


