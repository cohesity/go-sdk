# KmsUpdateRequestParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKms** | Pointer to [**AwsKmsUpdateParams**](AwsKmsUpdateParams.md) |  | [optional] 
**CryptsoftKms** | Pointer to [**CryptsoftKmsUpdateParams**](CryptsoftKmsUpdateParams.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | The Id of a KMS server. | [optional] 
**ServerName** | Pointer to **NullableString** | Specifies the name given to the KMS Server. | [optional] 

## Methods

### NewKmsUpdateRequestParameters

`func NewKmsUpdateRequestParameters() *KmsUpdateRequestParameters`

NewKmsUpdateRequestParameters instantiates a new KmsUpdateRequestParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsUpdateRequestParametersWithDefaults

`func NewKmsUpdateRequestParametersWithDefaults() *KmsUpdateRequestParameters`

NewKmsUpdateRequestParametersWithDefaults instantiates a new KmsUpdateRequestParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKms

`func (o *KmsUpdateRequestParameters) GetAwsKms() AwsKmsUpdateParams`

GetAwsKms returns the AwsKms field if non-nil, zero value otherwise.

### GetAwsKmsOk

`func (o *KmsUpdateRequestParameters) GetAwsKmsOk() (*AwsKmsUpdateParams, bool)`

GetAwsKmsOk returns a tuple with the AwsKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKms

`func (o *KmsUpdateRequestParameters) SetAwsKms(v AwsKmsUpdateParams)`

SetAwsKms sets AwsKms field to given value.

### HasAwsKms

`func (o *KmsUpdateRequestParameters) HasAwsKms() bool`

HasAwsKms returns a boolean if a field has been set.

### GetCryptsoftKms

`func (o *KmsUpdateRequestParameters) GetCryptsoftKms() CryptsoftKmsUpdateParams`

GetCryptsoftKms returns the CryptsoftKms field if non-nil, zero value otherwise.

### GetCryptsoftKmsOk

`func (o *KmsUpdateRequestParameters) GetCryptsoftKmsOk() (*CryptsoftKmsUpdateParams, bool)`

GetCryptsoftKmsOk returns a tuple with the CryptsoftKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptsoftKms

`func (o *KmsUpdateRequestParameters) SetCryptsoftKms(v CryptsoftKmsUpdateParams)`

SetCryptsoftKms sets CryptsoftKms field to given value.

### HasCryptsoftKms

`func (o *KmsUpdateRequestParameters) HasCryptsoftKms() bool`

HasCryptsoftKms returns a boolean if a field has been set.

### GetId

`func (o *KmsUpdateRequestParameters) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsUpdateRequestParameters) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsUpdateRequestParameters) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KmsUpdateRequestParameters) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KmsUpdateRequestParameters) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KmsUpdateRequestParameters) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetServerName

`func (o *KmsUpdateRequestParameters) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *KmsUpdateRequestParameters) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *KmsUpdateRequestParameters) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *KmsUpdateRequestParameters) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *KmsUpdateRequestParameters) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *KmsUpdateRequestParameters) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


