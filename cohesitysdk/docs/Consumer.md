# Consumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the id of the consumer. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the consumer. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of the consumer. Type of the consumer can be one of the following three,  &#39;kViews&#39;, indicates the stats info of Views used per organization (tenant) per view box (storage domain). &#39;kProtectionRuns&#39;, indicates the stats info of Protection Runs used per organization (tenant) per view box (storage domain). &#39;kReplicationRuns&#39;, indicates the stats info of Replication In used per organization (tenant) per view box (storage domain). &#39;kViewProtectionRuns&#39;, indicates the stats info of View Protection Runs used per organization (tenant) per view box (storage domain). | [optional] 

## Methods

### NewConsumer

`func NewConsumer() *Consumer`

NewConsumer instantiates a new Consumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerWithDefaults

`func NewConsumerWithDefaults() *Consumer`

NewConsumerWithDefaults instantiates a new Consumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Consumer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Consumer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Consumer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Consumer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Consumer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Consumer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Consumer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Consumer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Consumer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Consumer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Consumer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Consumer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *Consumer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Consumer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Consumer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Consumer) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Consumer) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Consumer) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


