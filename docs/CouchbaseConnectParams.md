# CouchbaseConnectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierDirectPort** | Pointer to **NullableInt32** | Specifies the Carrier direct/sll port. | [optional] 
**HttpDirectPort** | Pointer to **NullableInt32** | Specifies the HTTP direct/sll port. | [optional] 
**RequiresSsl** | Pointer to **NullableBool** | Specifies whether this cluster allows connection through SSL only. | [optional] 
**Seeds** | Pointer to **[]string** | Specifies the Seeds of this Couchbase Cluster. | [optional] 

## Methods

### NewCouchbaseConnectParams

`func NewCouchbaseConnectParams() *CouchbaseConnectParams`

NewCouchbaseConnectParams instantiates a new CouchbaseConnectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseConnectParamsWithDefaults

`func NewCouchbaseConnectParamsWithDefaults() *CouchbaseConnectParams`

NewCouchbaseConnectParamsWithDefaults instantiates a new CouchbaseConnectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierDirectPort

`func (o *CouchbaseConnectParams) GetCarrierDirectPort() int32`

GetCarrierDirectPort returns the CarrierDirectPort field if non-nil, zero value otherwise.

### GetCarrierDirectPortOk

`func (o *CouchbaseConnectParams) GetCarrierDirectPortOk() (*int32, bool)`

GetCarrierDirectPortOk returns a tuple with the CarrierDirectPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierDirectPort

`func (o *CouchbaseConnectParams) SetCarrierDirectPort(v int32)`

SetCarrierDirectPort sets CarrierDirectPort field to given value.

### HasCarrierDirectPort

`func (o *CouchbaseConnectParams) HasCarrierDirectPort() bool`

HasCarrierDirectPort returns a boolean if a field has been set.

### SetCarrierDirectPortNil

`func (o *CouchbaseConnectParams) SetCarrierDirectPortNil(b bool)`

 SetCarrierDirectPortNil sets the value for CarrierDirectPort to be an explicit nil

### UnsetCarrierDirectPort
`func (o *CouchbaseConnectParams) UnsetCarrierDirectPort()`

UnsetCarrierDirectPort ensures that no value is present for CarrierDirectPort, not even an explicit nil
### GetHttpDirectPort

`func (o *CouchbaseConnectParams) GetHttpDirectPort() int32`

GetHttpDirectPort returns the HttpDirectPort field if non-nil, zero value otherwise.

### GetHttpDirectPortOk

`func (o *CouchbaseConnectParams) GetHttpDirectPortOk() (*int32, bool)`

GetHttpDirectPortOk returns a tuple with the HttpDirectPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpDirectPort

`func (o *CouchbaseConnectParams) SetHttpDirectPort(v int32)`

SetHttpDirectPort sets HttpDirectPort field to given value.

### HasHttpDirectPort

`func (o *CouchbaseConnectParams) HasHttpDirectPort() bool`

HasHttpDirectPort returns a boolean if a field has been set.

### SetHttpDirectPortNil

`func (o *CouchbaseConnectParams) SetHttpDirectPortNil(b bool)`

 SetHttpDirectPortNil sets the value for HttpDirectPort to be an explicit nil

### UnsetHttpDirectPort
`func (o *CouchbaseConnectParams) UnsetHttpDirectPort()`

UnsetHttpDirectPort ensures that no value is present for HttpDirectPort, not even an explicit nil
### GetRequiresSsl

`func (o *CouchbaseConnectParams) GetRequiresSsl() bool`

GetRequiresSsl returns the RequiresSsl field if non-nil, zero value otherwise.

### GetRequiresSslOk

`func (o *CouchbaseConnectParams) GetRequiresSslOk() (*bool, bool)`

GetRequiresSslOk returns a tuple with the RequiresSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSsl

`func (o *CouchbaseConnectParams) SetRequiresSsl(v bool)`

SetRequiresSsl sets RequiresSsl field to given value.

### HasRequiresSsl

`func (o *CouchbaseConnectParams) HasRequiresSsl() bool`

HasRequiresSsl returns a boolean if a field has been set.

### SetRequiresSslNil

`func (o *CouchbaseConnectParams) SetRequiresSslNil(b bool)`

 SetRequiresSslNil sets the value for RequiresSsl to be an explicit nil

### UnsetRequiresSsl
`func (o *CouchbaseConnectParams) UnsetRequiresSsl()`

UnsetRequiresSsl ensures that no value is present for RequiresSsl, not even an explicit nil
### GetSeeds

`func (o *CouchbaseConnectParams) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *CouchbaseConnectParams) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *CouchbaseConnectParams) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *CouchbaseConnectParams) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### SetSeedsNil

`func (o *CouchbaseConnectParams) SetSeedsNil(b bool)`

 SetSeedsNil sets the value for Seeds to be an explicit nil

### UnsetSeeds
`func (o *CouchbaseConnectParams) UnsetSeeds()`

UnsetSeeds ensures that no value is present for Seeds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


