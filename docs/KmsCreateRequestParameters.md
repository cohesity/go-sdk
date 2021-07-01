# KmsCreateRequestParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKms** | Pointer to [**AwsKmsConfiguration**](AwsKmsConfiguration.md) |  | [optional] 
**CryptsoftKms** | Pointer to [**CryptsoftKmsConfiguration**](CryptsoftKmsConfiguration.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | The Id of a KMS server. | [optional] 
**ServerName** | Pointer to **NullableString** | Specifies the name given to the KMS Server. | [optional] 
**ServerType** | Pointer to **NullableString** | Specifies the type of key mangement system. &#39;kInternalKms&#39; indicates an internal KMS object. &#39;kAwsKms&#39; indicates an Aws KMS object. &#39;kCryptsoftKms&#39; indicates a Cryptsoft KMS object. | [optional] 

## Methods

### NewKmsCreateRequestParameters

`func NewKmsCreateRequestParameters() *KmsCreateRequestParameters`

NewKmsCreateRequestParameters instantiates a new KmsCreateRequestParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateRequestParametersWithDefaults

`func NewKmsCreateRequestParametersWithDefaults() *KmsCreateRequestParameters`

NewKmsCreateRequestParametersWithDefaults instantiates a new KmsCreateRequestParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKms

`func (o *KmsCreateRequestParameters) GetAwsKms() AwsKmsConfiguration`

GetAwsKms returns the AwsKms field if non-nil, zero value otherwise.

### GetAwsKmsOk

`func (o *KmsCreateRequestParameters) GetAwsKmsOk() (*AwsKmsConfiguration, bool)`

GetAwsKmsOk returns a tuple with the AwsKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKms

`func (o *KmsCreateRequestParameters) SetAwsKms(v AwsKmsConfiguration)`

SetAwsKms sets AwsKms field to given value.

### HasAwsKms

`func (o *KmsCreateRequestParameters) HasAwsKms() bool`

HasAwsKms returns a boolean if a field has been set.

### GetCryptsoftKms

`func (o *KmsCreateRequestParameters) GetCryptsoftKms() CryptsoftKmsConfiguration`

GetCryptsoftKms returns the CryptsoftKms field if non-nil, zero value otherwise.

### GetCryptsoftKmsOk

`func (o *KmsCreateRequestParameters) GetCryptsoftKmsOk() (*CryptsoftKmsConfiguration, bool)`

GetCryptsoftKmsOk returns a tuple with the CryptsoftKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptsoftKms

`func (o *KmsCreateRequestParameters) SetCryptsoftKms(v CryptsoftKmsConfiguration)`

SetCryptsoftKms sets CryptsoftKms field to given value.

### HasCryptsoftKms

`func (o *KmsCreateRequestParameters) HasCryptsoftKms() bool`

HasCryptsoftKms returns a boolean if a field has been set.

### GetId

`func (o *KmsCreateRequestParameters) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsCreateRequestParameters) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsCreateRequestParameters) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KmsCreateRequestParameters) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KmsCreateRequestParameters) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KmsCreateRequestParameters) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetServerName

`func (o *KmsCreateRequestParameters) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *KmsCreateRequestParameters) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *KmsCreateRequestParameters) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *KmsCreateRequestParameters) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *KmsCreateRequestParameters) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *KmsCreateRequestParameters) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetServerType

`func (o *KmsCreateRequestParameters) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *KmsCreateRequestParameters) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *KmsCreateRequestParameters) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *KmsCreateRequestParameters) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### SetServerTypeNil

`func (o *KmsCreateRequestParameters) SetServerTypeNil(b bool)`

 SetServerTypeNil sets the value for ServerType to be an explicit nil

### UnsetServerType
`func (o *KmsCreateRequestParameters) UnsetServerType()`

UnsetServerType ensures that no value is present for ServerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


