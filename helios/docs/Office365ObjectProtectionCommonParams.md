# Office365ObjectProtectionCommonParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]Office365ObjectProtectionObjectParams**](Office365ObjectProtectionObjectParams.md) | Specifies the objects to be included in the Object Protection. | 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 

## Methods

### NewOffice365ObjectProtectionCommonParams

`func NewOffice365ObjectProtectionCommonParams(objects []Office365ObjectProtectionObjectParams, ) *Office365ObjectProtectionCommonParams`

NewOffice365ObjectProtectionCommonParams instantiates a new Office365ObjectProtectionCommonParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365ObjectProtectionCommonParamsWithDefaults

`func NewOffice365ObjectProtectionCommonParamsWithDefaults() *Office365ObjectProtectionCommonParams`

NewOffice365ObjectProtectionCommonParamsWithDefaults instantiates a new Office365ObjectProtectionCommonParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *Office365ObjectProtectionCommonParams) GetObjects() []Office365ObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Office365ObjectProtectionCommonParams) GetObjectsOk() (*[]Office365ObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Office365ObjectProtectionCommonParams) SetObjects(v []Office365ObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.


### GetIndexingPolicy

`func (o *Office365ObjectProtectionCommonParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *Office365ObjectProtectionCommonParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *Office365ObjectProtectionCommonParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *Office365ObjectProtectionCommonParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetSourceId

`func (o *Office365ObjectProtectionCommonParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Office365ObjectProtectionCommonParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Office365ObjectProtectionCommonParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Office365ObjectProtectionCommonParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *Office365ObjectProtectionCommonParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *Office365ObjectProtectionCommonParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *Office365ObjectProtectionCommonParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Office365ObjectProtectionCommonParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Office365ObjectProtectionCommonParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *Office365ObjectProtectionCommonParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *Office365ObjectProtectionCommonParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *Office365ObjectProtectionCommonParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


