# ViewProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies a unique id of a Protection Source for a View. The id is unique across Cohesity Clusters. | [optional] 
**Name** | Pointer to **NullableString** | Specifies a human readable name of the Protection Source of a View. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of managed Object in a View Protection Source environment. Examples of View Objects include &#39;kViewBox&#39; or &#39;kView&#39;. &#39;kViewBox&#39; indicates Storage Domain as a Protection Source type. &#39;kView&#39; indicates View as a Protection Source type. | [optional] 

## Methods

### NewViewProtectionSource

`func NewViewProtectionSource() *ViewProtectionSource`

NewViewProtectionSource instantiates a new ViewProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProtectionSourceWithDefaults

`func NewViewProtectionSourceWithDefaults() *ViewProtectionSource`

NewViewProtectionSourceWithDefaults instantiates a new ViewProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewProtectionSource) GetId() UniversalId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewProtectionSource) GetIdOk() (*UniversalId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewProtectionSource) SetId(v UniversalId)`

SetId sets Id field to given value.

### HasId

`func (o *ViewProtectionSource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ViewProtectionSource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ViewProtectionSource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ViewProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ViewProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ViewProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *ViewProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ViewProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ViewProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


