# Office365SharepointSiteObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]Office365ObjectProtectionObjectParams**](Office365ObjectProtectionObjectParams.md) | Specifies the objects to be included in the Object Protection. | 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 

## Methods

### NewOffice365SharepointSiteObjectProtectionParams

`func NewOffice365SharepointSiteObjectProtectionParams(objects []Office365ObjectProtectionObjectParams, ) *Office365SharepointSiteObjectProtectionParams`

NewOffice365SharepointSiteObjectProtectionParams instantiates a new Office365SharepointSiteObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365SharepointSiteObjectProtectionParamsWithDefaults

`func NewOffice365SharepointSiteObjectProtectionParamsWithDefaults() *Office365SharepointSiteObjectProtectionParams`

NewOffice365SharepointSiteObjectProtectionParamsWithDefaults instantiates a new Office365SharepointSiteObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *Office365SharepointSiteObjectProtectionParams) GetObjects() []Office365ObjectProtectionObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Office365SharepointSiteObjectProtectionParams) GetObjectsOk() (*[]Office365ObjectProtectionObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Office365SharepointSiteObjectProtectionParams) SetObjects(v []Office365ObjectProtectionObjectParams)`

SetObjects sets Objects field to given value.


### GetIndexingPolicy

`func (o *Office365SharepointSiteObjectProtectionParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *Office365SharepointSiteObjectProtectionParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *Office365SharepointSiteObjectProtectionParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *Office365SharepointSiteObjectProtectionParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetSourceId

`func (o *Office365SharepointSiteObjectProtectionParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Office365SharepointSiteObjectProtectionParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Office365SharepointSiteObjectProtectionParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Office365SharepointSiteObjectProtectionParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *Office365SharepointSiteObjectProtectionParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *Office365SharepointSiteObjectProtectionParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *Office365SharepointSiteObjectProtectionParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Office365SharepointSiteObjectProtectionParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Office365SharepointSiteObjectProtectionParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *Office365SharepointSiteObjectProtectionParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *Office365SharepointSiteObjectProtectionParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *Office365SharepointSiteObjectProtectionParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


