# CouchbaseBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketType** | Pointer to **NullableString** | Type of this bucket. | [optional] 
**DocumentCount** | Pointer to **NullableInt64** | Number of documents in this bucket. | [optional] 

## Methods

### NewCouchbaseBucket

`func NewCouchbaseBucket() *CouchbaseBucket`

NewCouchbaseBucket instantiates a new CouchbaseBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseBucketWithDefaults

`func NewCouchbaseBucketWithDefaults() *CouchbaseBucket`

NewCouchbaseBucketWithDefaults instantiates a new CouchbaseBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketType

`func (o *CouchbaseBucket) GetBucketType() string`

GetBucketType returns the BucketType field if non-nil, zero value otherwise.

### GetBucketTypeOk

`func (o *CouchbaseBucket) GetBucketTypeOk() (*string, bool)`

GetBucketTypeOk returns a tuple with the BucketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketType

`func (o *CouchbaseBucket) SetBucketType(v string)`

SetBucketType sets BucketType field to given value.

### HasBucketType

`func (o *CouchbaseBucket) HasBucketType() bool`

HasBucketType returns a boolean if a field has been set.

### SetBucketTypeNil

`func (o *CouchbaseBucket) SetBucketTypeNil(b bool)`

 SetBucketTypeNil sets the value for BucketType to be an explicit nil

### UnsetBucketType
`func (o *CouchbaseBucket) UnsetBucketType()`

UnsetBucketType ensures that no value is present for BucketType, not even an explicit nil
### GetDocumentCount

`func (o *CouchbaseBucket) GetDocumentCount() int64`

GetDocumentCount returns the DocumentCount field if non-nil, zero value otherwise.

### GetDocumentCountOk

`func (o *CouchbaseBucket) GetDocumentCountOk() (*int64, bool)`

GetDocumentCountOk returns a tuple with the DocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCount

`func (o *CouchbaseBucket) SetDocumentCount(v int64)`

SetDocumentCount sets DocumentCount field to given value.

### HasDocumentCount

`func (o *CouchbaseBucket) HasDocumentCount() bool`

HasDocumentCount returns a boolean if a field has been set.

### SetDocumentCountNil

`func (o *CouchbaseBucket) SetDocumentCountNil(b bool)`

 SetDocumentCountNil sets the value for DocumentCount to be an explicit nil

### UnsetDocumentCount
`func (o *CouchbaseBucket) UnsetDocumentCount()`

UnsetDocumentCount ensures that no value is present for DocumentCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


