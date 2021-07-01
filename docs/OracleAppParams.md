# OracleAppParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseAppId** | Pointer to **NullableInt64** | Specifies the source entity id of the selected app entity. | [optional] 
**NodeChannelList** | Pointer to [**[]OracleDatabaseNodeChannel**](OracleDatabaseNodeChannel.md) | Array of database node channel info.  Specifies the node channel info for all the databases of app entity. Length of this array will be 1 for RAC and Standalone setups. | [optional] 

## Methods

### NewOracleAppParams

`func NewOracleAppParams() *OracleAppParams`

NewOracleAppParams instantiates a new OracleAppParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleAppParamsWithDefaults

`func NewOracleAppParamsWithDefaults() *OracleAppParams`

NewOracleAppParamsWithDefaults instantiates a new OracleAppParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseAppId

`func (o *OracleAppParams) GetDatabaseAppId() int64`

GetDatabaseAppId returns the DatabaseAppId field if non-nil, zero value otherwise.

### GetDatabaseAppIdOk

`func (o *OracleAppParams) GetDatabaseAppIdOk() (*int64, bool)`

GetDatabaseAppIdOk returns a tuple with the DatabaseAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseAppId

`func (o *OracleAppParams) SetDatabaseAppId(v int64)`

SetDatabaseAppId sets DatabaseAppId field to given value.

### HasDatabaseAppId

`func (o *OracleAppParams) HasDatabaseAppId() bool`

HasDatabaseAppId returns a boolean if a field has been set.

### SetDatabaseAppIdNil

`func (o *OracleAppParams) SetDatabaseAppIdNil(b bool)`

 SetDatabaseAppIdNil sets the value for DatabaseAppId to be an explicit nil

### UnsetDatabaseAppId
`func (o *OracleAppParams) UnsetDatabaseAppId()`

UnsetDatabaseAppId ensures that no value is present for DatabaseAppId, not even an explicit nil
### GetNodeChannelList

`func (o *OracleAppParams) GetNodeChannelList() []OracleDatabaseNodeChannel`

GetNodeChannelList returns the NodeChannelList field if non-nil, zero value otherwise.

### GetNodeChannelListOk

`func (o *OracleAppParams) GetNodeChannelListOk() (*[]OracleDatabaseNodeChannel, bool)`

GetNodeChannelListOk returns a tuple with the NodeChannelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeChannelList

`func (o *OracleAppParams) SetNodeChannelList(v []OracleDatabaseNodeChannel)`

SetNodeChannelList sets NodeChannelList field to given value.

### HasNodeChannelList

`func (o *OracleAppParams) HasNodeChannelList() bool`

HasNodeChannelList returns a boolean if a field has been set.

### SetNodeChannelListNil

`func (o *OracleAppParams) SetNodeChannelListNil(b bool)`

 SetNodeChannelListNil sets the value for NodeChannelList to be an explicit nil

### UnsetNodeChannelList
`func (o *OracleAppParams) UnsetNodeChannelList()`

UnsetNodeChannelList ensures that no value is present for NodeChannelList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


