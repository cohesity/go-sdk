# RigelConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimToken** | Pointer to **NullableString** | Specifies the claim token in the connection. | [optional] 
**ConnectorGroups** | Pointer to [**[]ConnectorGroup**](ConnectorGroup.md) | Specifies the connector groups in the connection. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the connection. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the connection. | [optional] 
**Scalable** | Pointer to **NullableBool** | Flag to specify if the connection is scalable. | [optional] 
**UngroupedConnectors** | Pointer to **[]int64** | Specifies the ids of the connectors which are not grouped in this connection | [optional] 

## Methods

### NewRigelConnection

`func NewRigelConnection() *RigelConnection`

NewRigelConnection instantiates a new RigelConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRigelConnectionWithDefaults

`func NewRigelConnectionWithDefaults() *RigelConnection`

NewRigelConnectionWithDefaults instantiates a new RigelConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimToken

`func (o *RigelConnection) GetClaimToken() string`

GetClaimToken returns the ClaimToken field if non-nil, zero value otherwise.

### GetClaimTokenOk

`func (o *RigelConnection) GetClaimTokenOk() (*string, bool)`

GetClaimTokenOk returns a tuple with the ClaimToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimToken

`func (o *RigelConnection) SetClaimToken(v string)`

SetClaimToken sets ClaimToken field to given value.

### HasClaimToken

`func (o *RigelConnection) HasClaimToken() bool`

HasClaimToken returns a boolean if a field has been set.

### SetClaimTokenNil

`func (o *RigelConnection) SetClaimTokenNil(b bool)`

 SetClaimTokenNil sets the value for ClaimToken to be an explicit nil

### UnsetClaimToken
`func (o *RigelConnection) UnsetClaimToken()`

UnsetClaimToken ensures that no value is present for ClaimToken, not even an explicit nil
### GetConnectorGroups

`func (o *RigelConnection) GetConnectorGroups() []ConnectorGroup`

GetConnectorGroups returns the ConnectorGroups field if non-nil, zero value otherwise.

### GetConnectorGroupsOk

`func (o *RigelConnection) GetConnectorGroupsOk() (*[]ConnectorGroup, bool)`

GetConnectorGroupsOk returns a tuple with the ConnectorGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorGroups

`func (o *RigelConnection) SetConnectorGroups(v []ConnectorGroup)`

SetConnectorGroups sets ConnectorGroups field to given value.

### HasConnectorGroups

`func (o *RigelConnection) HasConnectorGroups() bool`

HasConnectorGroups returns a boolean if a field has been set.

### GetId

`func (o *RigelConnection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RigelConnection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RigelConnection) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RigelConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RigelConnection) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RigelConnection) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *RigelConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RigelConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RigelConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RigelConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RigelConnection) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RigelConnection) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScalable

`func (o *RigelConnection) GetScalable() bool`

GetScalable returns the Scalable field if non-nil, zero value otherwise.

### GetScalableOk

`func (o *RigelConnection) GetScalableOk() (*bool, bool)`

GetScalableOk returns a tuple with the Scalable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalable

`func (o *RigelConnection) SetScalable(v bool)`

SetScalable sets Scalable field to given value.

### HasScalable

`func (o *RigelConnection) HasScalable() bool`

HasScalable returns a boolean if a field has been set.

### SetScalableNil

`func (o *RigelConnection) SetScalableNil(b bool)`

 SetScalableNil sets the value for Scalable to be an explicit nil

### UnsetScalable
`func (o *RigelConnection) UnsetScalable()`

UnsetScalable ensures that no value is present for Scalable, not even an explicit nil
### GetUngroupedConnectors

`func (o *RigelConnection) GetUngroupedConnectors() []int64`

GetUngroupedConnectors returns the UngroupedConnectors field if non-nil, zero value otherwise.

### GetUngroupedConnectorsOk

`func (o *RigelConnection) GetUngroupedConnectorsOk() (*[]int64, bool)`

GetUngroupedConnectorsOk returns a tuple with the UngroupedConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUngroupedConnectors

`func (o *RigelConnection) SetUngroupedConnectors(v []int64)`

SetUngroupedConnectors sets UngroupedConnectors field to given value.

### HasUngroupedConnectors

`func (o *RigelConnection) HasUngroupedConnectors() bool`

HasUngroupedConnectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


