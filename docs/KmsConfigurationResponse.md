# KmsConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKms** | Pointer to [**AwsKmsConfiguration**](AwsKmsConfiguration.md) |  | [optional] 
**ConnectionStatus** | Pointer to **NullableBool** | Specifies if connection to this KMS exists. | [optional] 
**CryptsoftKms** | Pointer to [**CryptsoftKmsConfigResponse**](CryptsoftKmsConfigResponse.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | The Id of a KMS server. | [optional] 
**ServerName** | Pointer to **NullableString** | Specifies the name given to the KMS Server. | [optional] 
**ServerType** | Pointer to **NullableString** | Specifies the type of key mangement system. &#39;kInternalKms&#39; indicates an internal KMS object. &#39;kAwsKms&#39; indicates an Aws KMS object. &#39;kCryptsoftKms&#39; indicates a Cryptsoft KMS object. | [optional] 

## Methods

### NewKmsConfigurationResponse

`func NewKmsConfigurationResponse() *KmsConfigurationResponse`

NewKmsConfigurationResponse instantiates a new KmsConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigurationResponseWithDefaults

`func NewKmsConfigurationResponseWithDefaults() *KmsConfigurationResponse`

NewKmsConfigurationResponseWithDefaults instantiates a new KmsConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKms

`func (o *KmsConfigurationResponse) GetAwsKms() AwsKmsConfiguration`

GetAwsKms returns the AwsKms field if non-nil, zero value otherwise.

### GetAwsKmsOk

`func (o *KmsConfigurationResponse) GetAwsKmsOk() (*AwsKmsConfiguration, bool)`

GetAwsKmsOk returns a tuple with the AwsKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKms

`func (o *KmsConfigurationResponse) SetAwsKms(v AwsKmsConfiguration)`

SetAwsKms sets AwsKms field to given value.

### HasAwsKms

`func (o *KmsConfigurationResponse) HasAwsKms() bool`

HasAwsKms returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *KmsConfigurationResponse) GetConnectionStatus() bool`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *KmsConfigurationResponse) GetConnectionStatusOk() (*bool, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *KmsConfigurationResponse) SetConnectionStatus(v bool)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *KmsConfigurationResponse) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### SetConnectionStatusNil

`func (o *KmsConfigurationResponse) SetConnectionStatusNil(b bool)`

 SetConnectionStatusNil sets the value for ConnectionStatus to be an explicit nil

### UnsetConnectionStatus
`func (o *KmsConfigurationResponse) UnsetConnectionStatus()`

UnsetConnectionStatus ensures that no value is present for ConnectionStatus, not even an explicit nil
### GetCryptsoftKms

`func (o *KmsConfigurationResponse) GetCryptsoftKms() CryptsoftKmsConfigResponse`

GetCryptsoftKms returns the CryptsoftKms field if non-nil, zero value otherwise.

### GetCryptsoftKmsOk

`func (o *KmsConfigurationResponse) GetCryptsoftKmsOk() (*CryptsoftKmsConfigResponse, bool)`

GetCryptsoftKmsOk returns a tuple with the CryptsoftKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptsoftKms

`func (o *KmsConfigurationResponse) SetCryptsoftKms(v CryptsoftKmsConfigResponse)`

SetCryptsoftKms sets CryptsoftKms field to given value.

### HasCryptsoftKms

`func (o *KmsConfigurationResponse) HasCryptsoftKms() bool`

HasCryptsoftKms returns a boolean if a field has been set.

### GetId

`func (o *KmsConfigurationResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsConfigurationResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsConfigurationResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KmsConfigurationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KmsConfigurationResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KmsConfigurationResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetServerName

`func (o *KmsConfigurationResponse) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *KmsConfigurationResponse) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *KmsConfigurationResponse) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *KmsConfigurationResponse) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *KmsConfigurationResponse) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *KmsConfigurationResponse) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetServerType

`func (o *KmsConfigurationResponse) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *KmsConfigurationResponse) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *KmsConfigurationResponse) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *KmsConfigurationResponse) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### SetServerTypeNil

`func (o *KmsConfigurationResponse) SetServerTypeNil(b bool)`

 SetServerTypeNil sets the value for ServerType to be an explicit nil

### UnsetServerType
`func (o *KmsConfigurationResponse) UnsetServerType()`

UnsetServerType ensures that no value is present for ServerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


