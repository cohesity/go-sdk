# NodeDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthStatus** | Pointer to **string** | Node Health Status | [optional] 
**Id** | Pointer to **string** | Node ID | [optional] 
**Role** | Pointer to **string** | Node Role | [optional] 

## Methods

### NewNodeDef

`func NewNodeDef() *NodeDef`

NewNodeDef instantiates a new NodeDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDefWithDefaults

`func NewNodeDefWithDefaults() *NodeDef`

NewNodeDefWithDefaults instantiates a new NodeDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthStatus

`func (o *NodeDef) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *NodeDef) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *NodeDef) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *NodeDef) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetId

`func (o *NodeDef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeDef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeDef) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeDef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *NodeDef) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NodeDef) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NodeDef) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NodeDef) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


