# MongoDBProtectionGroupParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdpInfo** | Pointer to [**MongoDBCdpJobInfo**](MongoDBCdpJobInfo.md) | Specifies the CDP related information for a given protection group. This field will only be populated when protection group is configured with a CDP policy. | [optional] 

## Methods

### NewMongoDBProtectionGroupParamsAllOf

`func NewMongoDBProtectionGroupParamsAllOf() *MongoDBProtectionGroupParamsAllOf`

NewMongoDBProtectionGroupParamsAllOf instantiates a new MongoDBProtectionGroupParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBProtectionGroupParamsAllOfWithDefaults

`func NewMongoDBProtectionGroupParamsAllOfWithDefaults() *MongoDBProtectionGroupParamsAllOf`

NewMongoDBProtectionGroupParamsAllOfWithDefaults instantiates a new MongoDBProtectionGroupParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdpInfo

`func (o *MongoDBProtectionGroupParamsAllOf) GetCdpInfo() MongoDBCdpJobInfo`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *MongoDBProtectionGroupParamsAllOf) GetCdpInfoOk() (*MongoDBCdpJobInfo, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *MongoDBProtectionGroupParamsAllOf) SetCdpInfo(v MongoDBCdpJobInfo)`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *MongoDBProtectionGroupParamsAllOf) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


