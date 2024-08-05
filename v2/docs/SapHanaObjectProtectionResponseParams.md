# SapHanaObjectProtectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Concurrency** | Pointer to **NullableInt32** | Specifies the maximum number of concurrent IO Streams thatwill be created to exchange data with the cluster. If not specified, the default value is taken as 1. | [optional] [default to 8]
**Delta** | Pointer to **NullableString** | Specifies the incremental backup delta (incremental/differential) | [optional] [default to "incremental"]
**Objects** | [**[]UdaObjectProtectionObjectParams**](UdaObjectProtectionObjectParams.md) | Specifies the objects to be included in the Object Protection. | 

## Methods

### NewSapHanaObjectProtectionResponseParams

`func NewSapHanaObjectProtectionResponseParams(objects []UdaObjectProtectionObjectParams, ) *SapHanaObjectProtectionResponseParams`

NewSapHanaObjectProtectionResponseParams instantiates a new SapHanaObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSapHanaObjectProtectionResponseParamsWithDefaults

`func NewSapHanaObjectProtectionResponseParamsWithDefaults() *SapHanaObjectProtectionResponseParams`

NewSapHanaObjectProtectionResponseParamsWithDefaults instantiates a new SapHanaObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrency

`func (o *SapHanaObjectProtectionResponseParams) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *SapHanaObjectProtectionResponseParams) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *SapHanaObjectProtectionResponseParams) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *SapHanaObjectProtectionResponseParams) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### SetConcurrencyNil

`func (o *SapHanaObjectProtectionResponseParams) SetConcurrencyNil(b bool)`

 SetConcurrencyNil sets the value for Concurrency to be an explicit nil

### UnsetConcurrency
`func (o *SapHanaObjectProtectionResponseParams) UnsetConcurrency()`

UnsetConcurrency ensures that no value is present for Concurrency, not even an explicit nil
### GetDelta

`func (o *SapHanaObjectProtectionResponseParams) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *SapHanaObjectProtectionResponseParams) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *SapHanaObjectProtectionResponseParams) SetDelta(v string)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *SapHanaObjectProtectionResponseParams) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDeltaNil

`func (o *SapHanaObjectProtectionResponseParams) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *SapHanaObjectProtectionResponseParams) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetObjects

`func (o *SapHanaObjectProtectionResponseParams) GetObjects() []UdaObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *SapHanaObjectProtectionResponseParams) GetObjectsOk() (*[]UdaObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *SapHanaObjectProtectionResponseParams) SetObjects(v []UdaObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


