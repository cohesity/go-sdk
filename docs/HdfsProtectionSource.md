# HdfsProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the instance name of the Hdfs entity. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the managed Object in Hdfs Protection Source. Specifies the type of an Hdfs source entity. &#39;kCluster&#39; indicates a Hdfs cluster distributed over several physical nodes. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID for the Hdfs entity. | [optional] 

## Methods

### NewHdfsProtectionSource

`func NewHdfsProtectionSource() *HdfsProtectionSource`

NewHdfsProtectionSource instantiates a new HdfsProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsProtectionSourceWithDefaults

`func NewHdfsProtectionSourceWithDefaults() *HdfsProtectionSource`

NewHdfsProtectionSourceWithDefaults instantiates a new HdfsProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HdfsProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HdfsProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HdfsProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HdfsProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HdfsProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HdfsProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *HdfsProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HdfsProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HdfsProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HdfsProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HdfsProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HdfsProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *HdfsProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HdfsProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HdfsProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HdfsProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *HdfsProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *HdfsProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


