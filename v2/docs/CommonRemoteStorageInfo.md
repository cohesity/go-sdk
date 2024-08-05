# CommonRemoteStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies unique id of the registered remote storage. | [optional] [readonly] 
**Product** | **NullableString** | Specifies the product type of the remote storage. | 

## Methods

### NewCommonRemoteStorageInfo

`func NewCommonRemoteStorageInfo(product NullableString, ) *CommonRemoteStorageInfo`

NewCommonRemoteStorageInfo instantiates a new CommonRemoteStorageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonRemoteStorageInfoWithDefaults

`func NewCommonRemoteStorageInfoWithDefaults() *CommonRemoteStorageInfo`

NewCommonRemoteStorageInfoWithDefaults instantiates a new CommonRemoteStorageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonRemoteStorageInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonRemoteStorageInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonRemoteStorageInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CommonRemoteStorageInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommonRemoteStorageInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonRemoteStorageInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProduct

`func (o *CommonRemoteStorageInfo) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CommonRemoteStorageInfo) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CommonRemoteStorageInfo) SetProduct(v string)`

SetProduct sets Product field to given value.


### SetProductNil

`func (o *CommonRemoteStorageInfo) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *CommonRemoteStorageInfo) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


