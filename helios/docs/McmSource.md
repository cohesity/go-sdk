# McmSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the unique identifier of the Source. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Protection Source. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment of the Protection Source. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the object type of the Protection Source. | [optional] 
**SourceInfoList** | Pointer to [**[]McmSourceInfo**](McmSourceInfo.md) | Specifies the Source information list across all platforms. | [optional] 

## Methods

### NewMcmSource

`func NewMcmSource() *McmSource`

NewMcmSource instantiates a new McmSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSourceWithDefaults

`func NewMcmSourceWithDefaults() *McmSource`

NewMcmSourceWithDefaults instantiates a new McmSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McmSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McmSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McmSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *McmSource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *McmSource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *McmSource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *McmSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McmSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McmSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *McmSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *McmSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *McmSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnvironment

`func (o *McmSource) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *McmSource) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *McmSource) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *McmSource) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *McmSource) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *McmSource) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetType

`func (o *McmSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *McmSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *McmSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *McmSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *McmSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *McmSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSourceInfoList

`func (o *McmSource) GetSourceInfoList() []McmSourceInfo`

GetSourceInfoList returns the SourceInfoList field if non-nil, zero value otherwise.

### GetSourceInfoListOk

`func (o *McmSource) GetSourceInfoListOk() (*[]McmSourceInfo, bool)`

GetSourceInfoListOk returns a tuple with the SourceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfoList

`func (o *McmSource) SetSourceInfoList(v []McmSourceInfo)`

SetSourceInfoList sets SourceInfoList field to given value.

### HasSourceInfoList

`func (o *McmSource) HasSourceInfoList() bool`

HasSourceInfoList returns a boolean if a field has been set.

### SetSourceInfoListNil

`func (o *McmSource) SetSourceInfoListNil(b bool)`

 SetSourceInfoListNil sets the value for SourceInfoList to be an explicit nil

### UnsetSourceInfoList
`func (o *McmSource) UnsetSourceInfoList()`

UnsetSourceInfoList ensures that no value is present for SourceInfoList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


