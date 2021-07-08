# OracleDatabaseNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCount** | Pointer to **NullableInt32** | Specifies the number of channels user wants for the backup/recovery of this node. | [optional] 
**Node** | Pointer to **NullableString** | Specifies the ip of the database node. | [optional] 
**Port** | Pointer to **NullableInt64** | Specifies the port on which user wants to run the backup/recovery. | [optional] 
**SbtHostParams** | Pointer to [**OracleSbtHostParams**](OracleSbtHostParams.md) |  | [optional] 

## Methods

### NewOracleDatabaseNode

`func NewOracleDatabaseNode() *OracleDatabaseNode`

NewOracleDatabaseNode instantiates a new OracleDatabaseNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDatabaseNodeWithDefaults

`func NewOracleDatabaseNodeWithDefaults() *OracleDatabaseNode`

NewOracleDatabaseNodeWithDefaults instantiates a new OracleDatabaseNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCount

`func (o *OracleDatabaseNode) GetChannelCount() int32`

GetChannelCount returns the ChannelCount field if non-nil, zero value otherwise.

### GetChannelCountOk

`func (o *OracleDatabaseNode) GetChannelCountOk() (*int32, bool)`

GetChannelCountOk returns a tuple with the ChannelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCount

`func (o *OracleDatabaseNode) SetChannelCount(v int32)`

SetChannelCount sets ChannelCount field to given value.

### HasChannelCount

`func (o *OracleDatabaseNode) HasChannelCount() bool`

HasChannelCount returns a boolean if a field has been set.

### SetChannelCountNil

`func (o *OracleDatabaseNode) SetChannelCountNil(b bool)`

 SetChannelCountNil sets the value for ChannelCount to be an explicit nil

### UnsetChannelCount
`func (o *OracleDatabaseNode) UnsetChannelCount()`

UnsetChannelCount ensures that no value is present for ChannelCount, not even an explicit nil
### GetNode

`func (o *OracleDatabaseNode) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *OracleDatabaseNode) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *OracleDatabaseNode) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *OracleDatabaseNode) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *OracleDatabaseNode) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *OracleDatabaseNode) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetPort

`func (o *OracleDatabaseNode) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OracleDatabaseNode) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OracleDatabaseNode) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *OracleDatabaseNode) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *OracleDatabaseNode) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *OracleDatabaseNode) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSbtHostParams

`func (o *OracleDatabaseNode) GetSbtHostParams() OracleSbtHostParams`

GetSbtHostParams returns the SbtHostParams field if non-nil, zero value otherwise.

### GetSbtHostParamsOk

`func (o *OracleDatabaseNode) GetSbtHostParamsOk() (*OracleSbtHostParams, bool)`

GetSbtHostParamsOk returns a tuple with the SbtHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbtHostParams

`func (o *OracleDatabaseNode) SetSbtHostParams(v OracleSbtHostParams)`

SetSbtHostParams sets SbtHostParams field to given value.

### HasSbtHostParams

`func (o *OracleDatabaseNode) HasSbtHostParams() bool`

HasSbtHostParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


