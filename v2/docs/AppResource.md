# AppResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]ResourceEndpoint**](ResourceEndpoint.md) | Specifies the information about endpoint associated with this resorce. | [optional] 
**Id** | Pointer to **NullableString** | Specifies the unique id of the app resource. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the app resource. | [optional] 

## Methods

### NewAppResource

`func NewAppResource() *AppResource`

NewAppResource instantiates a new AppResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppResourceWithDefaults

`func NewAppResourceWithDefaults() *AppResource`

NewAppResourceWithDefaults instantiates a new AppResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *AppResource) GetEndpoints() []ResourceEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *AppResource) GetEndpointsOk() (*[]ResourceEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *AppResource) SetEndpoints(v []ResourceEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *AppResource) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### SetEndpointsNil

`func (o *AppResource) SetEndpointsNil(b bool)`

 SetEndpointsNil sets the value for Endpoints to be an explicit nil

### UnsetEndpoints
`func (o *AppResource) UnsetEndpoints()`

UnsetEndpoints ensures that no value is present for Endpoints, not even an explicit nil
### GetId

`func (o *AppResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppResource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AppResource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AppResource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *AppResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppResource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AppResource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AppResource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


