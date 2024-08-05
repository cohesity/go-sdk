# CouchbaseSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouchbaseObjectTypes** | **[]string** | Specifies Couchbase object types be searched. For Couchbase it can only be set to &#39;CouchbaseBuckets&#39;. | 
**SearchString** | **NullableString** | Specifies the search string to search the Couchbase Objects | 

## Methods

### NewCouchbaseSearchParams

`func NewCouchbaseSearchParams(couchbaseObjectTypes []string, searchString NullableString, ) *CouchbaseSearchParams`

NewCouchbaseSearchParams instantiates a new CouchbaseSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchbaseSearchParamsWithDefaults

`func NewCouchbaseSearchParamsWithDefaults() *CouchbaseSearchParams`

NewCouchbaseSearchParamsWithDefaults instantiates a new CouchbaseSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouchbaseObjectTypes

`func (o *CouchbaseSearchParams) GetCouchbaseObjectTypes() []string`

GetCouchbaseObjectTypes returns the CouchbaseObjectTypes field if non-nil, zero value otherwise.

### GetCouchbaseObjectTypesOk

`func (o *CouchbaseSearchParams) GetCouchbaseObjectTypesOk() (*[]string, bool)`

GetCouchbaseObjectTypesOk returns a tuple with the CouchbaseObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseObjectTypes

`func (o *CouchbaseSearchParams) SetCouchbaseObjectTypes(v []string)`

SetCouchbaseObjectTypes sets CouchbaseObjectTypes field to given value.


### GetSearchString

`func (o *CouchbaseSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *CouchbaseSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *CouchbaseSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *CouchbaseSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *CouchbaseSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


