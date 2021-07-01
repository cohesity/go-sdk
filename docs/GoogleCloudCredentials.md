# GoogleCloudCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientEmailAddress** | Pointer to **NullableString** | Specifies the client email address used to access Google Cloud Storage. | [optional] 
**ClientPrivateKey** | Pointer to **NullableString** | Specifies the private key used to access Google Cloud Storage that is generated when the service account is created. | [optional] 
**ProjectId** | Pointer to **NullableString** | Specifies the project id of an existing Google Cloud project to store objects. | [optional] 
**TierType** | Pointer to **NullableString** | Specifies the storage class of GCP. GoogleTierType specifies the storage class for Google. &#39;kGoogleStandard&#39; indicates a tier type of Google properties. &#39;kGoogleNearline&#39; indicates a tier type of Google properties that is not accessed frequently. &#39;kGoogleColdline&#39; indicates a tier type of Google properties that is rarely accessed. &#39;kGoogleRegional&#39; indicates a tier type of Google properties that stores frequently accessed data in the same region. &#39;kGoogleMultiRegional&#39; indicates a tier type of Google properties that is frequently accessed (\&quot;hot\&quot; objects) around the world. | [optional] 

## Methods

### NewGoogleCloudCredentials

`func NewGoogleCloudCredentials() *GoogleCloudCredentials`

NewGoogleCloudCredentials instantiates a new GoogleCloudCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudCredentialsWithDefaults

`func NewGoogleCloudCredentialsWithDefaults() *GoogleCloudCredentials`

NewGoogleCloudCredentialsWithDefaults instantiates a new GoogleCloudCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientEmailAddress

`func (o *GoogleCloudCredentials) GetClientEmailAddress() string`

GetClientEmailAddress returns the ClientEmailAddress field if non-nil, zero value otherwise.

### GetClientEmailAddressOk

`func (o *GoogleCloudCredentials) GetClientEmailAddressOk() (*string, bool)`

GetClientEmailAddressOk returns a tuple with the ClientEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmailAddress

`func (o *GoogleCloudCredentials) SetClientEmailAddress(v string)`

SetClientEmailAddress sets ClientEmailAddress field to given value.

### HasClientEmailAddress

`func (o *GoogleCloudCredentials) HasClientEmailAddress() bool`

HasClientEmailAddress returns a boolean if a field has been set.

### SetClientEmailAddressNil

`func (o *GoogleCloudCredentials) SetClientEmailAddressNil(b bool)`

 SetClientEmailAddressNil sets the value for ClientEmailAddress to be an explicit nil

### UnsetClientEmailAddress
`func (o *GoogleCloudCredentials) UnsetClientEmailAddress()`

UnsetClientEmailAddress ensures that no value is present for ClientEmailAddress, not even an explicit nil
### GetClientPrivateKey

`func (o *GoogleCloudCredentials) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *GoogleCloudCredentials) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *GoogleCloudCredentials) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *GoogleCloudCredentials) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### SetClientPrivateKeyNil

`func (o *GoogleCloudCredentials) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *GoogleCloudCredentials) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil
### GetProjectId

`func (o *GoogleCloudCredentials) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GoogleCloudCredentials) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GoogleCloudCredentials) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GoogleCloudCredentials) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *GoogleCloudCredentials) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *GoogleCloudCredentials) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetTierType

`func (o *GoogleCloudCredentials) GetTierType() string`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *GoogleCloudCredentials) GetTierTypeOk() (*string, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *GoogleCloudCredentials) SetTierType(v string)`

SetTierType sets TierType field to given value.

### HasTierType

`func (o *GoogleCloudCredentials) HasTierType() bool`

HasTierType returns a boolean if a field has been set.

### SetTierTypeNil

`func (o *GoogleCloudCredentials) SetTierTypeNil(b bool)`

 SetTierTypeNil sets the value for TierType to be an explicit nil

### UnsetTierType
`func (o *GoogleCloudCredentials) UnsetTierType()`

UnsetTierType ensures that no value is present for TierType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


