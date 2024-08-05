# RunObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | Pointer to **[]int64** | Specifies a list of ids of applications. | [optional] 
**Id** | **NullableInt64** | Specifies the id of object. | 
**PhysicalParams** | Pointer to [**RunObjectPhysicalParams**](RunObjectPhysicalParams.md) |  | [optional] 

## Methods

### NewRunObject

`func NewRunObject(id NullableInt64, ) *RunObject`

NewRunObject instantiates a new RunObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunObjectWithDefaults

`func NewRunObjectWithDefaults() *RunObject`

NewRunObjectWithDefaults instantiates a new RunObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *RunObject) GetAppIds() []int64`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *RunObject) GetAppIdsOk() (*[]int64, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *RunObject) SetAppIds(v []int64)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *RunObject) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### SetAppIdsNil

`func (o *RunObject) SetAppIdsNil(b bool)`

 SetAppIdsNil sets the value for AppIds to be an explicit nil

### UnsetAppIds
`func (o *RunObject) UnsetAppIds()`

UnsetAppIds ensures that no value is present for AppIds, not even an explicit nil
### GetId

`func (o *RunObject) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunObject) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunObject) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *RunObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RunObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPhysicalParams

`func (o *RunObject) GetPhysicalParams() RunObjectPhysicalParams`

GetPhysicalParams returns the PhysicalParams field if non-nil, zero value otherwise.

### GetPhysicalParamsOk

`func (o *RunObject) GetPhysicalParamsOk() (*RunObjectPhysicalParams, bool)`

GetPhysicalParamsOk returns a tuple with the PhysicalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalParams

`func (o *RunObject) SetPhysicalParams(v RunObjectPhysicalParams)`

SetPhysicalParams sets PhysicalParams field to given value.

### HasPhysicalParams

`func (o *RunObject) HasPhysicalParams() bool`

HasPhysicalParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


