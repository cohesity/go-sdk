# ViewClientsSummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfsClientCount** | Pointer to **NullableInt64** | Specifies the number of NFS clients. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies the node ip the clients are connected to. | [optional] 
**ServerIp** | Pointer to **NullableString** | Specifies the server ip the clients are connected to. | [optional] 
**SmbClientCount** | Pointer to **NullableInt64** | Specifies the number of SMB clients. | [optional] 

## Methods

### NewViewClientsSummaryInfo

`func NewViewClientsSummaryInfo() *ViewClientsSummaryInfo`

NewViewClientsSummaryInfo instantiates a new ViewClientsSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewClientsSummaryInfoWithDefaults

`func NewViewClientsSummaryInfoWithDefaults() *ViewClientsSummaryInfo`

NewViewClientsSummaryInfoWithDefaults instantiates a new ViewClientsSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfsClientCount

`func (o *ViewClientsSummaryInfo) GetNfsClientCount() int64`

GetNfsClientCount returns the NfsClientCount field if non-nil, zero value otherwise.

### GetNfsClientCountOk

`func (o *ViewClientsSummaryInfo) GetNfsClientCountOk() (*int64, bool)`

GetNfsClientCountOk returns a tuple with the NfsClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsClientCount

`func (o *ViewClientsSummaryInfo) SetNfsClientCount(v int64)`

SetNfsClientCount sets NfsClientCount field to given value.

### HasNfsClientCount

`func (o *ViewClientsSummaryInfo) HasNfsClientCount() bool`

HasNfsClientCount returns a boolean if a field has been set.

### SetNfsClientCountNil

`func (o *ViewClientsSummaryInfo) SetNfsClientCountNil(b bool)`

 SetNfsClientCountNil sets the value for NfsClientCount to be an explicit nil

### UnsetNfsClientCount
`func (o *ViewClientsSummaryInfo) UnsetNfsClientCount()`

UnsetNfsClientCount ensures that no value is present for NfsClientCount, not even an explicit nil
### GetNodeIp

`func (o *ViewClientsSummaryInfo) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *ViewClientsSummaryInfo) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *ViewClientsSummaryInfo) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *ViewClientsSummaryInfo) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *ViewClientsSummaryInfo) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *ViewClientsSummaryInfo) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetServerIp

`func (o *ViewClientsSummaryInfo) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *ViewClientsSummaryInfo) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *ViewClientsSummaryInfo) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *ViewClientsSummaryInfo) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### SetServerIpNil

`func (o *ViewClientsSummaryInfo) SetServerIpNil(b bool)`

 SetServerIpNil sets the value for ServerIp to be an explicit nil

### UnsetServerIp
`func (o *ViewClientsSummaryInfo) UnsetServerIp()`

UnsetServerIp ensures that no value is present for ServerIp, not even an explicit nil
### GetSmbClientCount

`func (o *ViewClientsSummaryInfo) GetSmbClientCount() int64`

GetSmbClientCount returns the SmbClientCount field if non-nil, zero value otherwise.

### GetSmbClientCountOk

`func (o *ViewClientsSummaryInfo) GetSmbClientCountOk() (*int64, bool)`

GetSmbClientCountOk returns a tuple with the SmbClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbClientCount

`func (o *ViewClientsSummaryInfo) SetSmbClientCount(v int64)`

SetSmbClientCount sets SmbClientCount field to given value.

### HasSmbClientCount

`func (o *ViewClientsSummaryInfo) HasSmbClientCount() bool`

HasSmbClientCount returns a boolean if a field has been set.

### SetSmbClientCountNil

`func (o *ViewClientsSummaryInfo) SetSmbClientCountNil(b bool)`

 SetSmbClientCountNil sets the value for SmbClientCount to be an explicit nil

### UnsetSmbClientCount
`func (o *ViewClientsSummaryInfo) UnsetSmbClientCount()`

UnsetSmbClientCount ensures that no value is present for SmbClientCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


