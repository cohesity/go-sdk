# KubernetesCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientPrivateKey** | Pointer to **NullableString** | Specifies Client private associated with the service account. | [optional] 

## Methods

### NewKubernetesCredentials

`func NewKubernetesCredentials() *KubernetesCredentials`

NewKubernetesCredentials instantiates a new KubernetesCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesCredentialsWithDefaults

`func NewKubernetesCredentialsWithDefaults() *KubernetesCredentials`

NewKubernetesCredentialsWithDefaults instantiates a new KubernetesCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientPrivateKey

`func (o *KubernetesCredentials) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *KubernetesCredentials) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *KubernetesCredentials) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *KubernetesCredentials) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### SetClientPrivateKeyNil

`func (o *KubernetesCredentials) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *KubernetesCredentials) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


