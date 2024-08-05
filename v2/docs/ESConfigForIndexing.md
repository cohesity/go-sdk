# ESConfigForIndexing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsDomain** | **NullableString** | Fully qualified ES domain name. | 
**EsIamRoleArn** | **NullableString** | IAM role ARN which has access to ES instance. | 

## Methods

### NewESConfigForIndexing

`func NewESConfigForIndexing(esDomain NullableString, esIamRoleArn NullableString, ) *ESConfigForIndexing`

NewESConfigForIndexing instantiates a new ESConfigForIndexing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESConfigForIndexingWithDefaults

`func NewESConfigForIndexingWithDefaults() *ESConfigForIndexing`

NewESConfigForIndexingWithDefaults instantiates a new ESConfigForIndexing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsDomain

`func (o *ESConfigForIndexing) GetEsDomain() string`

GetEsDomain returns the EsDomain field if non-nil, zero value otherwise.

### GetEsDomainOk

`func (o *ESConfigForIndexing) GetEsDomainOk() (*string, bool)`

GetEsDomainOk returns a tuple with the EsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsDomain

`func (o *ESConfigForIndexing) SetEsDomain(v string)`

SetEsDomain sets EsDomain field to given value.


### SetEsDomainNil

`func (o *ESConfigForIndexing) SetEsDomainNil(b bool)`

 SetEsDomainNil sets the value for EsDomain to be an explicit nil

### UnsetEsDomain
`func (o *ESConfigForIndexing) UnsetEsDomain()`

UnsetEsDomain ensures that no value is present for EsDomain, not even an explicit nil
### GetEsIamRoleArn

`func (o *ESConfigForIndexing) GetEsIamRoleArn() string`

GetEsIamRoleArn returns the EsIamRoleArn field if non-nil, zero value otherwise.

### GetEsIamRoleArnOk

`func (o *ESConfigForIndexing) GetEsIamRoleArnOk() (*string, bool)`

GetEsIamRoleArnOk returns a tuple with the EsIamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsIamRoleArn

`func (o *ESConfigForIndexing) SetEsIamRoleArn(v string)`

SetEsIamRoleArn sets EsIamRoleArn field to given value.


### SetEsIamRoleArnNil

`func (o *ESConfigForIndexing) SetEsIamRoleArnNil(b bool)`

 SetEsIamRoleArnNil sets the value for EsIamRoleArn to be an explicit nil

### UnsetEsIamRoleArn
`func (o *ESConfigForIndexing) UnsetEsIamRoleArn()`

UnsetEsIamRoleArn ensures that no value is present for EsIamRoleArn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


