# CouchbaseSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Specifies the password to access target entity. | 
**Username** | **string** | Specifies the username to access target entity. | 
**CarrierPort** | **NullableInt32** | Carrier direct or Carrier SSL port. | 
**HttpPort** | **NullableInt32** | HTTP direct or HTTP SSL port. | 
**IsSslRequired** | **NullableBool** | Set to true if connection to couchbase has to be using SSL. | 
**Seeds** | **[]string** | Specifies the IP Addresses or hostnames of the Couchbase cluster seed nodes. | 

## Methods

### NewCouchbaseSourceRegistrationParams

`func NewCouchbaseSourceRegistrationParams(password string, username string, carrierPort NullableInt32, httpPort NullableInt32, isSslRequired NullableBool, seeds []string, ) *CouchbaseSourceRegistrationParams`

NewCouchbaseSourceRegistrationParams instantiates a new CouchbaseSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseSourceRegistrationParamsWithDefaults

`func NewCouchbaseSourceRegistrationParamsWithDefaults() *CouchbaseSourceRegistrationParams`

NewCouchbaseSourceRegistrationParamsWithDefaults instantiates a new CouchbaseSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *CouchbaseSourceRegistrationParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CouchbaseSourceRegistrationParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CouchbaseSourceRegistrationParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *CouchbaseSourceRegistrationParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CouchbaseSourceRegistrationParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CouchbaseSourceRegistrationParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCarrierPort

`func (o *CouchbaseSourceRegistrationParams) GetCarrierPort() int32`

GetCarrierPort returns the CarrierPort field if non-nil, zero value otherwise.

### GetCarrierPortOk

`func (o *CouchbaseSourceRegistrationParams) GetCarrierPortOk() (*int32, bool)`

GetCarrierPortOk returns a tuple with the CarrierPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierPort

`func (o *CouchbaseSourceRegistrationParams) SetCarrierPort(v int32)`

SetCarrierPort sets CarrierPort field to given value.


### SetCarrierPortNil

`func (o *CouchbaseSourceRegistrationParams) SetCarrierPortNil(b bool)`

 SetCarrierPortNil sets the value for CarrierPort to be an explicit nil

### UnsetCarrierPort
`func (o *CouchbaseSourceRegistrationParams) UnsetCarrierPort()`

UnsetCarrierPort ensures that no value is present for CarrierPort, not even an explicit nil
### GetHttpPort

`func (o *CouchbaseSourceRegistrationParams) GetHttpPort() int32`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *CouchbaseSourceRegistrationParams) GetHttpPortOk() (*int32, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *CouchbaseSourceRegistrationParams) SetHttpPort(v int32)`

SetHttpPort sets HttpPort field to given value.


### SetHttpPortNil

`func (o *CouchbaseSourceRegistrationParams) SetHttpPortNil(b bool)`

 SetHttpPortNil sets the value for HttpPort to be an explicit nil

### UnsetHttpPort
`func (o *CouchbaseSourceRegistrationParams) UnsetHttpPort()`

UnsetHttpPort ensures that no value is present for HttpPort, not even an explicit nil
### GetIsSslRequired

`func (o *CouchbaseSourceRegistrationParams) GetIsSslRequired() bool`

GetIsSslRequired returns the IsSslRequired field if non-nil, zero value otherwise.

### GetIsSslRequiredOk

`func (o *CouchbaseSourceRegistrationParams) GetIsSslRequiredOk() (*bool, bool)`

GetIsSslRequiredOk returns a tuple with the IsSslRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSslRequired

`func (o *CouchbaseSourceRegistrationParams) SetIsSslRequired(v bool)`

SetIsSslRequired sets IsSslRequired field to given value.


### SetIsSslRequiredNil

`func (o *CouchbaseSourceRegistrationParams) SetIsSslRequiredNil(b bool)`

 SetIsSslRequiredNil sets the value for IsSslRequired to be an explicit nil

### UnsetIsSslRequired
`func (o *CouchbaseSourceRegistrationParams) UnsetIsSslRequired()`

UnsetIsSslRequired ensures that no value is present for IsSslRequired, not even an explicit nil
### GetSeeds

`func (o *CouchbaseSourceRegistrationParams) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *CouchbaseSourceRegistrationParams) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *CouchbaseSourceRegistrationParams) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


