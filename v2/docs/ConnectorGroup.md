# ConnectorGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to **[]int64** | Specifies the ids of the connectors in the group. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the group. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the group. | [optional] 

## Methods

### NewConnectorGroup

`func NewConnectorGroup() *ConnectorGroup`

NewConnectorGroup instantiates a new ConnectorGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorGroupWithDefaults

`func NewConnectorGroupWithDefaults() *ConnectorGroup`

NewConnectorGroupWithDefaults instantiates a new ConnectorGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *ConnectorGroup) GetConnectors() []int64`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *ConnectorGroup) GetConnectorsOk() (*[]int64, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *ConnectorGroup) SetConnectors(v []int64)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *ConnectorGroup) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetId

`func (o *ConnectorGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConnectorGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConnectorGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ConnectorGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConnectorGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConnectorGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


