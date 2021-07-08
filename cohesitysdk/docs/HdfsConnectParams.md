# HdfsConnectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HadoopDistribution** | Pointer to **NullableString** | Specifies the Hadoop Distribution. Hadoop distribution.  &#39;CDH&#39; indicates Hadoop distribution type Cloudera. &#39;HDP&#39; indicates Hadoop distribution type Hortonworks. | [optional] 
**HadoopVersion** | Pointer to **NullableString** | Specifies the Hadoop version | [optional] 
**KerberosPrincipal** | Pointer to **NullableString** | Specifies the kerberos principal. | [optional] 
**Namenode** | Pointer to **NullableString** | Specifies the Namenode host or Nameservice. | [optional] 
**Port** | Pointer to **NullableInt32** | Specifies the Webhdfs Port | [optional] 

## Methods

### NewHdfsConnectParams

`func NewHdfsConnectParams() *HdfsConnectParams`

NewHdfsConnectParams instantiates a new HdfsConnectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdfsConnectParamsWithDefaults

`func NewHdfsConnectParamsWithDefaults() *HdfsConnectParams`

NewHdfsConnectParamsWithDefaults instantiates a new HdfsConnectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHadoopDistribution

`func (o *HdfsConnectParams) GetHadoopDistribution() string`

GetHadoopDistribution returns the HadoopDistribution field if non-nil, zero value otherwise.

### GetHadoopDistributionOk

`func (o *HdfsConnectParams) GetHadoopDistributionOk() (*string, bool)`

GetHadoopDistributionOk returns a tuple with the HadoopDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadoopDistribution

`func (o *HdfsConnectParams) SetHadoopDistribution(v string)`

SetHadoopDistribution sets HadoopDistribution field to given value.

### HasHadoopDistribution

`func (o *HdfsConnectParams) HasHadoopDistribution() bool`

HasHadoopDistribution returns a boolean if a field has been set.

### SetHadoopDistributionNil

`func (o *HdfsConnectParams) SetHadoopDistributionNil(b bool)`

 SetHadoopDistributionNil sets the value for HadoopDistribution to be an explicit nil

### UnsetHadoopDistribution
`func (o *HdfsConnectParams) UnsetHadoopDistribution()`

UnsetHadoopDistribution ensures that no value is present for HadoopDistribution, not even an explicit nil
### GetHadoopVersion

`func (o *HdfsConnectParams) GetHadoopVersion() string`

GetHadoopVersion returns the HadoopVersion field if non-nil, zero value otherwise.

### GetHadoopVersionOk

`func (o *HdfsConnectParams) GetHadoopVersionOk() (*string, bool)`

GetHadoopVersionOk returns a tuple with the HadoopVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHadoopVersion

`func (o *HdfsConnectParams) SetHadoopVersion(v string)`

SetHadoopVersion sets HadoopVersion field to given value.

### HasHadoopVersion

`func (o *HdfsConnectParams) HasHadoopVersion() bool`

HasHadoopVersion returns a boolean if a field has been set.

### SetHadoopVersionNil

`func (o *HdfsConnectParams) SetHadoopVersionNil(b bool)`

 SetHadoopVersionNil sets the value for HadoopVersion to be an explicit nil

### UnsetHadoopVersion
`func (o *HdfsConnectParams) UnsetHadoopVersion()`

UnsetHadoopVersion ensures that no value is present for HadoopVersion, not even an explicit nil
### GetKerberosPrincipal

`func (o *HdfsConnectParams) GetKerberosPrincipal() string`

GetKerberosPrincipal returns the KerberosPrincipal field if non-nil, zero value otherwise.

### GetKerberosPrincipalOk

`func (o *HdfsConnectParams) GetKerberosPrincipalOk() (*string, bool)`

GetKerberosPrincipalOk returns a tuple with the KerberosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrincipal

`func (o *HdfsConnectParams) SetKerberosPrincipal(v string)`

SetKerberosPrincipal sets KerberosPrincipal field to given value.

### HasKerberosPrincipal

`func (o *HdfsConnectParams) HasKerberosPrincipal() bool`

HasKerberosPrincipal returns a boolean if a field has been set.

### SetKerberosPrincipalNil

`func (o *HdfsConnectParams) SetKerberosPrincipalNil(b bool)`

 SetKerberosPrincipalNil sets the value for KerberosPrincipal to be an explicit nil

### UnsetKerberosPrincipal
`func (o *HdfsConnectParams) UnsetKerberosPrincipal()`

UnsetKerberosPrincipal ensures that no value is present for KerberosPrincipal, not even an explicit nil
### GetNamenode

`func (o *HdfsConnectParams) GetNamenode() string`

GetNamenode returns the Namenode field if non-nil, zero value otherwise.

### GetNamenodeOk

`func (o *HdfsConnectParams) GetNamenodeOk() (*string, bool)`

GetNamenodeOk returns a tuple with the Namenode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamenode

`func (o *HdfsConnectParams) SetNamenode(v string)`

SetNamenode sets Namenode field to given value.

### HasNamenode

`func (o *HdfsConnectParams) HasNamenode() bool`

HasNamenode returns a boolean if a field has been set.

### SetNamenodeNil

`func (o *HdfsConnectParams) SetNamenodeNil(b bool)`

 SetNamenodeNil sets the value for Namenode to be an explicit nil

### UnsetNamenode
`func (o *HdfsConnectParams) UnsetNamenode()`

UnsetNamenode ensures that no value is present for Namenode, not even an explicit nil
### GetPort

`func (o *HdfsConnectParams) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HdfsConnectParams) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HdfsConnectParams) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HdfsConnectParams) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HdfsConnectParams) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HdfsConnectParams) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


