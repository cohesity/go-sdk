# CouchbaseSourceRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seeds** | **[]string** | Specifies the IP Addresses or hostnames of the Couchbase cluster seed nodes. | 
**IsSslRequired** | **NullableBool** | Set to true if connection to couchbase has to be using SSL. | 
**HttpPort** | **NullableInt32** | HTTP direct or HTTP SSL port. | 
**CarrierPort** | **NullableInt32** | Carrier direct or Carrier SSL port. | 

## Methods

### NewCouchbaseSourceRegistrationParamsAllOf

`func NewCouchbaseSourceRegistrationParamsAllOf(seeds []string, isSslRequired NullableBool, httpPort NullableInt32, carrierPort NullableInt32, ) *CouchbaseSourceRegistrationParamsAllOf`

NewCouchbaseSourceRegistrationParamsAllOf instantiates a new CouchbaseSourceRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseSourceRegistrationParamsAllOfWithDefaults

`func NewCouchbaseSourceRegistrationParamsAllOfWithDefaults() *CouchbaseSourceRegistrationParamsAllOf`

NewCouchbaseSourceRegistrationParamsAllOfWithDefaults instantiates a new CouchbaseSourceRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeeds

`func (o *CouchbaseSourceRegistrationParamsAllOf) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *CouchbaseSourceRegistrationParamsAllOf) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *CouchbaseSourceRegistrationParamsAllOf) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.


### GetIsSslRequired

`func (o *CouchbaseSourceRegistrationParamsAllOf) GetIsSslRequired() bool`

GetIsSslRequired returns the IsSslRequired field if non-nil, zero value otherwise.

### GetIsSslRequiredOk

`func (o *CouchbaseSourceRegistrationParamsAllOf) GetIsSslRequiredOk() (*bool, bool)`

GetIsSslRequiredOk returns a tuple with the IsSslRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSslRequired

`func (o *CouchbaseSourceRegistrationParamsAllOf) SetIsSslRequired(v bool)`

SetIsSslRequired sets IsSslRequired field to given value.


### SetIsSslRequiredNil

`func (o *CouchbaseSourceRegistrationParamsAllOf) SetIsSslRequiredNil(b bool)`

 SetIsSslRequiredNil sets the value for IsSslRequired to be an explicit nil

### UnsetIsSslRequired
`func (o *CouchbaseSourceRegistrationParamsAllOf) UnsetIsSslRequired()`

UnsetIsSslRequired ensures that no value is present for IsSslRequired, not even an explicit nil
### GetHttpPort

`func (o *CouchbaseSourceRegistrationParamsAllOf) GetHttpPort() int32`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *CouchbaseSourceRegistrationParamsAllOf) GetHttpPortOk() (*int32, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *CouchbaseSourceRegistrationParamsAllOf) SetHttpPort(v int32)`

SetHttpPort sets HttpPort field to given value.


### SetHttpPortNil

`func (o *CouchbaseSourceRegistrationParamsAllOf) SetHttpPortNil(b bool)`

 SetHttpPortNil sets the value for HttpPort to be an explicit nil

### UnsetHttpPort
`func (o *CouchbaseSourceRegistrationParamsAllOf) UnsetHttpPort()`

UnsetHttpPort ensures that no value is present for HttpPort, not even an explicit nil
### GetCarrierPort

`func (o *CouchbaseSourceRegistrationParamsAllOf) GetCarrierPort() int32`

GetCarrierPort returns the CarrierPort field if non-nil, zero value otherwise.

### GetCarrierPortOk

`func (o *CouchbaseSourceRegistrationParamsAllOf) GetCarrierPortOk() (*int32, bool)`

GetCarrierPortOk returns a tuple with the CarrierPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierPort

`func (o *CouchbaseSourceRegistrationParamsAllOf) SetCarrierPort(v int32)`

SetCarrierPort sets CarrierPort field to given value.


### SetCarrierPortNil

`func (o *CouchbaseSourceRegistrationParamsAllOf) SetCarrierPortNil(b bool)`

 SetCarrierPortNil sets the value for CarrierPort to be an explicit nil

### UnsetCarrierPort
`func (o *CouchbaseSourceRegistrationParamsAllOf) UnsetCarrierPort()`

UnsetCarrierPort ensures that no value is present for CarrierPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


