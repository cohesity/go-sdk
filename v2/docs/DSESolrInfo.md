# DSESolrInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SolrNodes** | **[]string** | Solr node IP Addresses | 
**SolrPort** | **NullableInt32** | Solr node Port. | 

## Methods

### NewDSESolrInfo

`func NewDSESolrInfo(solrNodes []string, solrPort NullableInt32, ) *DSESolrInfo`

NewDSESolrInfo instantiates a new DSESolrInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSESolrInfoWithDefaults

`func NewDSESolrInfoWithDefaults() *DSESolrInfo`

NewDSESolrInfoWithDefaults instantiates a new DSESolrInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSolrNodes

`func (o *DSESolrInfo) GetSolrNodes() []string`

GetSolrNodes returns the SolrNodes field if non-nil, zero value otherwise.

### GetSolrNodesOk

`func (o *DSESolrInfo) GetSolrNodesOk() (*[]string, bool)`

GetSolrNodesOk returns a tuple with the SolrNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolrNodes

`func (o *DSESolrInfo) SetSolrNodes(v []string)`

SetSolrNodes sets SolrNodes field to given value.


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


### SetSolrPortNil

`func (o *DSESolrInfo) SetSolrPortNil(b bool)`

 SetSolrPortNil sets the value for SolrPort to be an explicit nil

### UnsetSolrPort
`func (o *DSESolrInfo) UnsetSolrPort()`

UnsetSolrPort ensures that no value is present for SolrPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


