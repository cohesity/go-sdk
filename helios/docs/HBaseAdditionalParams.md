# HBaseAdditionalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZookeeperQuorum** | Pointer to **[]string** | The &#39;Zookeeper Quorum&#39; for this HBase. | [optional] [readonly] 
**DataRootDirectory** | Pointer to **NullableString** | The &#39;Data root directory&#39; for this HBase. | [optional] [readonly] 
**AuthType** | Pointer to **NullableString** | Authentication type. | [optional] [readonly] 

## Methods

### NewHBaseAdditionalParams

`func NewHBaseAdditionalParams() *HBaseAdditionalParams`

NewHBaseAdditionalParams instantiates a new HBaseAdditionalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHBaseAdditionalParamsWithDefaults

`func NewHBaseAdditionalParamsWithDefaults() *HBaseAdditionalParams`

NewHBaseAdditionalParamsWithDefaults instantiates a new HBaseAdditionalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZookeeperQuorum

`func (o *HBaseAdditionalParams) GetZookeeperQuorum() []string`

GetZookeeperQuorum returns the ZookeeperQuorum field if non-nil, zero value otherwise.

### GetZookeeperQuorumOk

`func (o *HBaseAdditionalParams) GetZookeeperQuorumOk() (*[]string, bool)`

GetZookeeperQuorumOk returns a tuple with the ZookeeperQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZookeeperQuorum

`func (o *HBaseAdditionalParams) SetZookeeperQuorum(v []string)`

SetZookeeperQuorum sets ZookeeperQuorum field to given value.

### HasZookeeperQuorum

`func (o *HBaseAdditionalParams) HasZookeeperQuorum() bool`

HasZookeeperQuorum returns a boolean if a field has been set.

### SetZookeeperQuorumNil

`func (o *HBaseAdditionalParams) SetZookeeperQuorumNil(b bool)`

 SetZookeeperQuorumNil sets the value for ZookeeperQuorum to be an explicit nil

### UnsetZookeeperQuorum
`func (o *HBaseAdditionalParams) UnsetZookeeperQuorum()`

UnsetZookeeperQuorum ensures that no value is present for ZookeeperQuorum, not even an explicit nil
### GetDataRootDirectory

`func (o *HBaseAdditionalParams) GetDataRootDirectory() string`

GetDataRootDirectory returns the DataRootDirectory field if non-nil, zero value otherwise.

### GetDataRootDirectoryOk

`func (o *HBaseAdditionalParams) GetDataRootDirectoryOk() (*string, bool)`

GetDataRootDirectoryOk returns a tuple with the DataRootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRootDirectory

`func (o *HBaseAdditionalParams) SetDataRootDirectory(v string)`

SetDataRootDirectory sets DataRootDirectory field to given value.

### HasDataRootDirectory

`func (o *HBaseAdditionalParams) HasDataRootDirectory() bool`

HasDataRootDirectory returns a boolean if a field has been set.

### SetDataRootDirectoryNil

`func (o *HBaseAdditionalParams) SetDataRootDirectoryNil(b bool)`

 SetDataRootDirectoryNil sets the value for DataRootDirectory to be an explicit nil

### UnsetDataRootDirectory
`func (o *HBaseAdditionalParams) UnsetDataRootDirectory()`

UnsetDataRootDirectory ensures that no value is present for DataRootDirectory, not even an explicit nil
### GetAuthType

`func (o *HBaseAdditionalParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *HBaseAdditionalParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *HBaseAdditionalParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *HBaseAdditionalParams) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *HBaseAdditionalParams) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *HBaseAdditionalParams) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


