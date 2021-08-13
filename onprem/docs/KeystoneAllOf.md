# KeystoneAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the Keystone configuration name. | 
**Id** | Pointer to **NullableInt64** | Specifies the Keystone configuration id. | [optional] [readonly] 
**AuthUrl** | **NullableString** | Specifies the url points to the Keystone service. | 

## Methods

### NewKeystoneAllOf

`func NewKeystoneAllOf(name NullableString, authUrl NullableString, ) *KeystoneAllOf`

NewKeystoneAllOf instantiates a new KeystoneAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneAllOfWithDefaults

`func NewKeystoneAllOfWithDefaults() *KeystoneAllOf`

NewKeystoneAllOfWithDefaults instantiates a new KeystoneAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeystoneAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeystoneAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeystoneAllOf) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *KeystoneAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeystoneAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *KeystoneAllOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeystoneAllOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeystoneAllOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeystoneAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KeystoneAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KeystoneAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAuthUrl

`func (o *KeystoneAllOf) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *KeystoneAllOf) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *KeystoneAllOf) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### SetAuthUrlNil

`func (o *KeystoneAllOf) SetAuthUrlNil(b bool)`

 SetAuthUrlNil sets the value for AuthUrl to be an explicit nil

### UnsetAuthUrl
`func (o *KeystoneAllOf) UnsetAuthUrl()`

UnsetAuthUrl ensures that no value is present for AuthUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


