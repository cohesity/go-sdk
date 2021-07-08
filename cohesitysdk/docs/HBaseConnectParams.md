# HBaseConnectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HdfsEntityId** | Pointer to **NullableInt64** | The entity id of the HDFS source for this HBase | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | Specifies the kerberos principal. | [optional] 
**RootDataDirectory** | Pointer to **NullableString** | Specifies the HBase data root directory. | [optional] 
**ZookeeperQuorum** | Pointer to **[]string** | Specifies the HBase zookeeper quorum. | [optional] 

## Methods

### NewHBaseConnectParams

`func NewHBaseConnectParams() *HBaseConnectParams`

NewHBaseConnectParams instantiates a new HBaseConnectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHBaseConnectParamsWithDefaults

`func NewHBaseConnectParamsWithDefaults() *HBaseConnectParams`

NewHBaseConnectParamsWithDefaults instantiates a new HBaseConnectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHdfsEntityId

`func (o *HBaseConnectParams) GetHdfsEntityId() int64`

GetHdfsEntityId returns the HdfsEntityId field if non-nil, zero value otherwise.

### GetHdfsEntityIdOk

`func (o *HBaseConnectParams) GetHdfsEntityIdOk() (*int64, bool)`

GetHdfsEntityIdOk returns a tuple with the HdfsEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsEntityId

`func (o *HBaseConnectParams) SetHdfsEntityId(v int64)`

SetHdfsEntityId sets HdfsEntityId field to given value.

### HasHdfsEntityId

`func (o *HBaseConnectParams) HasHdfsEntityId() bool`

HasHdfsEntityId returns a boolean if a field has been set.

### SetHdfsEntityIdNil

`func (o *HBaseConnectParams) SetHdfsEntityIdNil(b bool)`

 SetHdfsEntityIdNil sets the value for HdfsEntityId to be an explicit nil

### UnsetHdfsEntityId
`func (o *HBaseConnectParams) UnsetHdfsEntityId()`

UnsetHdfsEntityId ensures that no value is present for HdfsEntityId, not even an explicit nil
### GetKerberosPrincipal

`func (o *HBaseConnectParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HBaseConnectParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HBaseConnectParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HBaseConnectParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HBaseConnectParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HBaseConnectParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetRootDataDirectory

`func (o *HBaseConnectParams) GetRootDataDirectory() string`

GetRootDataDirectory returns the RootDataDirectory field if non-nil, zero value otherwise.

### GetRootDataDirectoryOk

`func (o *HBaseConnectParams) GetRootDataDirectoryOk() (*string, bool)`

GetRootDataDirectoryOk returns a tuple with the RootDataDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDataDirectory

`func (o *HBaseConnectParams) SetRootDataDirectory(v string)`

SetRootDataDirectory sets RootDataDirectory field to given value.

### HasRootDataDirectory

`func (o *HBaseConnectParams) HasRootDataDirectory() bool`

HasRootDataDirectory returns a boolean if a field has been set.

### SetRootDataDirectoryNil

`func (o *HBaseConnectParams) SetRootDataDirectoryNil(b bool)`

 SetRootDataDirectoryNil sets the value for RootDataDirectory to be an explicit nil

### UnsetRootDataDirectory
`func (o *HBaseConnectParams) UnsetRootDataDirectory()`

UnsetRootDataDirectory ensures that no value is present for RootDataDirectory, not even an explicit nil
### GetZookeeperQuorum

`func (o *HBaseConnectParams) GetZookeeperQuorum() []string`

GetZookeeperQuorum returns the ZookeeperQuorum field if non-nil, zero value otherwise.

### GetZookeeperQuorumOk

`func (o *HBaseConnectParams) GetZookeeperQuorumOk() (*[]string, bool)`

GetZookeeperQuorumOk returns a tuple with the ZookeeperQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZookeeperQuorum

`func (o *HBaseConnectParams) SetZookeeperQuorum(v []string)`

SetZookeeperQuorum sets ZookeeperQuorum field to given value.

### HasZookeeperQuorum

`func (o *HBaseConnectParams) HasZookeeperQuorum() bool`

HasZookeeperQuorum returns a boolean if a field has been set.

### SetZookeeperQuorumNil

`func (o *HBaseConnectParams) SetZookeeperQuorumNil(b bool)`

 SetZookeeperQuorumNil sets the value for ZookeeperQuorum to be an explicit nil

### UnsetZookeeperQuorum
`func (o *HBaseConnectParams) UnsetZookeeperQuorum()`

UnsetZookeeperQuorum ensures that no value is present for ZookeeperQuorum, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


