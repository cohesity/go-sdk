# CouchBaseOnPremSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouchbaseObjectTypes** | **[]string** | Specifies Couchbase object types be searched. For Couchbase it can only be set to &#39;CouchbaseBuckets&#39;. | 
**SearchString** | **NullableString** | Specifies the search string to search the Couchbase Objects | 
**SourceIds** | Pointer to **[]int64** | Specifies a list of source ids. Only files found in these sources will be returned. | [optional] 

## Methods

### NewCouchBaseOnPremSearchParams

`func NewCouchBaseOnPremSearchParams(couchbaseObjectTypes []string, searchString NullableString, ) *CouchBaseOnPremSearchParams`

NewCouchBaseOnPremSearchParams instantiates a new CouchBaseOnPremSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouchBaseOnPremSearchParamsWithDefaults

`func NewCouchBaseOnPremSearchParamsWithDefaults() *CouchBaseOnPremSearchParams`

NewCouchBaseOnPremSearchParamsWithDefaults instantiates a new CouchBaseOnPremSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouchbaseObjectTypes

`func (o *CouchBaseOnPremSearchParams) GetCouchbaseObjectTypes() []string`

GetCouchbaseObjectTypes returns the CouchbaseObjectTypes field if non-nil, zero value otherwise.

### GetCouchbaseObjectTypesOk

`func (o *CouchBaseOnPremSearchParams) GetCouchbaseObjectTypesOk() (*[]string, bool)`

GetCouchbaseObjectTypesOk returns a tuple with the CouchbaseObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseObjectTypes

`func (o *CouchBaseOnPremSearchParams) SetCouchbaseObjectTypes(v []string)`

SetCouchbaseObjectTypes sets CouchbaseObjectTypes field to given value.


### GetSearchString

`func (o *CouchBaseOnPremSearchParams) GetSearchString() string`

GetSearchString returns the SearchString field if non-nil, zero value otherwise.

### GetSearchStringOk

`func (o *CouchBaseOnPremSearchParams) GetSearchStringOk() (*string, bool)`

GetSearchStringOk returns a tuple with the SearchString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchString

`func (o *CouchBaseOnPremSearchParams) SetSearchString(v string)`

SetSearchString sets SearchString field to given value.


### SetSearchStringNil

`func (o *CouchBaseOnPremSearchParams) SetSearchStringNil(b bool)`

 SetSearchStringNil sets the value for SearchString to be an explicit nil

### UnsetSearchString
`func (o *CouchBaseOnPremSearchParams) UnsetSearchString()`

UnsetSearchString ensures that no value is present for SearchString, not even an explicit nil
### GetSourceIds

`func (o *CouchBaseOnPremSearchParams) GetSourceIds() []int64`

GetSourceIds returns the SourceIds field if non-nil, zero value otherwise.

### GetSourceIdsOk

`func (o *CouchBaseOnPremSearchParams) GetSourceIdsOk() (*[]int64, bool)`

GetSourceIdsOk returns a tuple with the SourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIds

`func (o *CouchBaseOnPremSearchParams) SetSourceIds(v []int64)`

SetSourceIds sets SourceIds field to given value.

### HasSourceIds

`func (o *CouchBaseOnPremSearchParams) HasSourceIds() bool`

HasSourceIds returns a boolean if a field has been set.

### SetSourceIdsNil

`func (o *CouchBaseOnPremSearchParams) SetSourceIdsNil(b bool)`

 SetSourceIdsNil sets the value for SourceIds to be an explicit nil

### UnsetSourceIds
`func (o *CouchBaseOnPremSearchParams) UnsetSourceIds()`

UnsetSourceIds ensures that no value is present for SourceIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


