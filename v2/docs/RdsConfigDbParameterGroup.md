# RdsConfigDbParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 

## Methods

### NewRdsConfigDbParameterGroup

`func NewRdsConfigDbParameterGroup(id NullableInt64, ) *RdsConfigDbParameterGroup`

NewRdsConfigDbParameterGroup instantiates a new RdsConfigDbParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsConfigDbParameterGroupWithDefaults

`func NewRdsConfigDbParameterGroupWithDefaults() *RdsConfigDbParameterGroup`

NewRdsConfigDbParameterGroupWithDefaults instantiates a new RdsConfigDbParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RdsConfigDbParameterGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RdsConfigDbParameterGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RdsConfigDbParameterGroup) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *RdsConfigDbParameterGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RdsConfigDbParameterGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *RdsConfigDbParameterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RdsConfigDbParameterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RdsConfigDbParameterGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RdsConfigDbParameterGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RdsConfigDbParameterGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RdsConfigDbParameterGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


