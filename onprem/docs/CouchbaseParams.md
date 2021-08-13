# CouchbaseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**RecoverCouchbaseParams** | [**RecoverCouchbaseParams**](RecoverCouchbaseParams.md) |  | 

## Methods

### NewCouchbaseParams

`func NewCouchbaseParams(recoveryAction string, recoverCouchbaseParams RecoverCouchbaseParams, ) *CouchbaseParams`

NewCouchbaseParams instantiates a new CouchbaseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseParamsWithDefaults

`func NewCouchbaseParamsWithDefaults() *CouchbaseParams`

NewCouchbaseParamsWithDefaults instantiates a new CouchbaseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryAction

`func (o *CouchbaseParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *CouchbaseParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *CouchbaseParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.


### GetRecoverCouchbaseParams

`func (o *CouchbaseParams) GetRecoverCouchbaseParams() RecoverCouchbaseParams`

GetRecoverCouchbaseParams returns the RecoverCouchbaseParams field if non-nil, zero value otherwise.

### GetRecoverCouchbaseParamsOk

`func (o *CouchbaseParams) GetRecoverCouchbaseParamsOk() (*RecoverCouchbaseParams, bool)`

GetRecoverCouchbaseParamsOk returns a tuple with the RecoverCouchbaseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverCouchbaseParams

`func (o *CouchbaseParams) SetRecoverCouchbaseParams(v RecoverCouchbaseParams)`

SetRecoverCouchbaseParams sets RecoverCouchbaseParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


