# DSESolrInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SolrNodeVec** | Pointer to **[]string** | Solr node IP Addresses. | [optional] 
**SolrPort** | Pointer to **NullableInt32** | Solr node Port. | [optional] 

## Methods

### NewDSESolrInfo

`func NewDSESolrInfo() *DSESolrInfo`

NewDSESolrInfo instantiates a new DSESolrInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSESolrInfoWithDefaults

`func NewDSESolrInfoWithDefaults() *DSESolrInfo`

NewDSESolrInfoWithDefaults instantiates a new DSESolrInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSolrNodeVec

`func (o *DSESolrInfo) GetSolrNodeVec() []string`

GetSolrNodeVec returns the SolrNodeVec field if non-nil, zero value otherwise.

### GetSolrNodeVecOk

`func (o *DSESolrInfo) GetSolrNodeVecOk() (*[]string, bool)`

GetSolrNodeVecOk returns a tuple with the SolrNodeVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolrNodeVec

`func (o *DSESolrInfo) SetSolrNodeVec(v []string)`

SetSolrNodeVec sets SolrNodeVec field to given value.

### HasSolrNodeVec

`func (o *DSESolrInfo) HasSolrNodeVec() bool`

HasSolrNodeVec returns a boolean if a field has been set.

### SetSolrNodeVecNil

`func (o *DSESolrInfo) SetSolrNodeVecNil(b bool)`

 SetSolrNodeVecNil sets the value for SolrNodeVec to be an explicit nil

### UnsetSolrNodeVec
`func (o *DSESolrInfo) UnsetSolrNodeVec()`

UnsetSolrNodeVec ensures that no value is present for SolrNodeVec, not even an explicit nil
### GetSolrPort

`func (o *DSESolrInfo) GetSolrPort() int32`

GetSolrPort returns the SolrPort field if non-nil, zero value otherwise.

### GetSolrPortOk

`func (o *DSESolrInfo) GetSolrPortOk() (*int32, bool)`

GetSolrPortOk returns a tuple with the SolrPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolrPort

`func (o *DSESolrInfo) SetSolrPort(v int32)`

SetSolrPort sets SolrPort field to given value.

### HasSolrPort

`func (o *DSESolrInfo) HasSolrPort() bool`

HasSolrPort returns a boolean if a field has been set.

### SetSolrPortNil

`func (o *DSESolrInfo) SetSolrPortNil(b bool)`

 SetSolrPortNil sets the value for SolrPort to be an explicit nil

### UnsetSolrPort
`func (o *DSESolrInfo) UnsetSolrPort()`

UnsetSolrPort ensures that no value is present for SolrPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


