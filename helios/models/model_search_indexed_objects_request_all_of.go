/*
 * Cohesity REST API
 *
 * Cohesity API provides a RESTful interface to access the various data management operations on Cohesity cluster and Helios.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

// package helios
package models

import (
	"encoding/json"
	. "github.com/cohesity/go-sdk/helios/utils"
)

var _ = NullableBool{}

// SearchIndexedObjectsRequestAllOf struct for SearchIndexedObjectsRequestAllOf
type SearchIndexedObjectsRequestAllOf struct {
	EmailParams *SearchEmailRequestParams `json:"emailParams,omitempty"`
	FileParams *SearchFileRequestParams `json:"fileParams,omitempty"`
	CassandraParams *CassandraOnPremSearchParams `json:"cassandraParams,omitempty"`
	CouchbaseParams *CouchBaseOnPremSearchParams `json:"couchbaseParams,omitempty"`
	HbaseParams *HbaseOnPremSearchParams `json:"hbaseParams,omitempty"`
	HiveParams *HiveOnPremSearchParams `json:"hiveParams,omitempty"`
	MongodbParams *MongoDbOnPremSearchParams `json:"mongodbParams,omitempty"`
	HdfsParams *HDFSOnPremSearchParams `json:"hdfsParams,omitempty"`
	ExchangeParams *SearchExchangeObjectsRequestParams `json:"exchangeParams,omitempty"`
	PublicFolderParams *SearchPublicFolderRequestParams `json:"publicFolderParams,omitempty"`
	MsTeamsParams *SearchMsTeamsRequestParams `json:"msTeamsParams,omitempty"`
	SharepointParams *SearchDocumentLibraryRequestParams `json:"sharepointParams,omitempty"`
}

// NewSearchIndexedObjectsRequestAllOf instantiates a new SearchIndexedObjectsRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIndexedObjectsRequestAllOf() *SearchIndexedObjectsRequestAllOf {
	this := SearchIndexedObjectsRequestAllOf{}
	return &this
}

// NewSearchIndexedObjectsRequestAllOfWithDefaults instantiates a new SearchIndexedObjectsRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIndexedObjectsRequestAllOfWithDefaults() *SearchIndexedObjectsRequestAllOf {
	this := SearchIndexedObjectsRequestAllOf{}
	return &this
}

// GetEmailParams returns the EmailParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetEmailParams() SearchEmailRequestParams {
	if o == nil || o.EmailParams == nil {
		var ret SearchEmailRequestParams
		return ret
	}
	return *o.EmailParams
}

// GetEmailParamsOk returns a tuple with the EmailParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetEmailParamsOk() (*SearchEmailRequestParams, bool) {
	if o == nil || o.EmailParams == nil {
		return nil, false
	}
	return o.EmailParams, true
}

// HasEmailParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasEmailParams() bool {
	if o != nil && o.EmailParams != nil {
		return true
	}

	return false
}

// SetEmailParams gets a reference to the given SearchEmailRequestParams and assigns it to the EmailParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetEmailParams(v SearchEmailRequestParams) {
	o.EmailParams = &v
}

// GetFileParams returns the FileParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetFileParams() SearchFileRequestParams {
	if o == nil || o.FileParams == nil {
		var ret SearchFileRequestParams
		return ret
	}
	return *o.FileParams
}

// GetFileParamsOk returns a tuple with the FileParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetFileParamsOk() (*SearchFileRequestParams, bool) {
	if o == nil || o.FileParams == nil {
		return nil, false
	}
	return o.FileParams, true
}

// HasFileParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasFileParams() bool {
	if o != nil && o.FileParams != nil {
		return true
	}

	return false
}

// SetFileParams gets a reference to the given SearchFileRequestParams and assigns it to the FileParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetFileParams(v SearchFileRequestParams) {
	o.FileParams = &v
}

// GetCassandraParams returns the CassandraParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetCassandraParams() CassandraOnPremSearchParams {
	if o == nil || o.CassandraParams == nil {
		var ret CassandraOnPremSearchParams
		return ret
	}
	return *o.CassandraParams
}

// GetCassandraParamsOk returns a tuple with the CassandraParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetCassandraParamsOk() (*CassandraOnPremSearchParams, bool) {
	if o == nil || o.CassandraParams == nil {
		return nil, false
	}
	return o.CassandraParams, true
}

// HasCassandraParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasCassandraParams() bool {
	if o != nil && o.CassandraParams != nil {
		return true
	}

	return false
}

// SetCassandraParams gets a reference to the given CassandraOnPremSearchParams and assigns it to the CassandraParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetCassandraParams(v CassandraOnPremSearchParams) {
	o.CassandraParams = &v
}

// GetCouchbaseParams returns the CouchbaseParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetCouchbaseParams() CouchBaseOnPremSearchParams {
	if o == nil || o.CouchbaseParams == nil {
		var ret CouchBaseOnPremSearchParams
		return ret
	}
	return *o.CouchbaseParams
}

// GetCouchbaseParamsOk returns a tuple with the CouchbaseParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetCouchbaseParamsOk() (*CouchBaseOnPremSearchParams, bool) {
	if o == nil || o.CouchbaseParams == nil {
		return nil, false
	}
	return o.CouchbaseParams, true
}

// HasCouchbaseParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasCouchbaseParams() bool {
	if o != nil && o.CouchbaseParams != nil {
		return true
	}

	return false
}

// SetCouchbaseParams gets a reference to the given CouchBaseOnPremSearchParams and assigns it to the CouchbaseParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetCouchbaseParams(v CouchBaseOnPremSearchParams) {
	o.CouchbaseParams = &v
}

// GetHbaseParams returns the HbaseParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetHbaseParams() HbaseOnPremSearchParams {
	if o == nil || o.HbaseParams == nil {
		var ret HbaseOnPremSearchParams
		return ret
	}
	return *o.HbaseParams
}

// GetHbaseParamsOk returns a tuple with the HbaseParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetHbaseParamsOk() (*HbaseOnPremSearchParams, bool) {
	if o == nil || o.HbaseParams == nil {
		return nil, false
	}
	return o.HbaseParams, true
}

// HasHbaseParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasHbaseParams() bool {
	if o != nil && o.HbaseParams != nil {
		return true
	}

	return false
}

// SetHbaseParams gets a reference to the given HbaseOnPremSearchParams and assigns it to the HbaseParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetHbaseParams(v HbaseOnPremSearchParams) {
	o.HbaseParams = &v
}

// GetHiveParams returns the HiveParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetHiveParams() HiveOnPremSearchParams {
	if o == nil || o.HiveParams == nil {
		var ret HiveOnPremSearchParams
		return ret
	}
	return *o.HiveParams
}

// GetHiveParamsOk returns a tuple with the HiveParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetHiveParamsOk() (*HiveOnPremSearchParams, bool) {
	if o == nil || o.HiveParams == nil {
		return nil, false
	}
	return o.HiveParams, true
}

// HasHiveParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasHiveParams() bool {
	if o != nil && o.HiveParams != nil {
		return true
	}

	return false
}

// SetHiveParams gets a reference to the given HiveOnPremSearchParams and assigns it to the HiveParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetHiveParams(v HiveOnPremSearchParams) {
	o.HiveParams = &v
}

// GetMongodbParams returns the MongodbParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetMongodbParams() MongoDbOnPremSearchParams {
	if o == nil || o.MongodbParams == nil {
		var ret MongoDbOnPremSearchParams
		return ret
	}
	return *o.MongodbParams
}

// GetMongodbParamsOk returns a tuple with the MongodbParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetMongodbParamsOk() (*MongoDbOnPremSearchParams, bool) {
	if o == nil || o.MongodbParams == nil {
		return nil, false
	}
	return o.MongodbParams, true
}

// HasMongodbParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasMongodbParams() bool {
	if o != nil && o.MongodbParams != nil {
		return true
	}

	return false
}

// SetMongodbParams gets a reference to the given MongoDbOnPremSearchParams and assigns it to the MongodbParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetMongodbParams(v MongoDbOnPremSearchParams) {
	o.MongodbParams = &v
}

// GetHdfsParams returns the HdfsParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetHdfsParams() HDFSOnPremSearchParams {
	if o == nil || o.HdfsParams == nil {
		var ret HDFSOnPremSearchParams
		return ret
	}
	return *o.HdfsParams
}

// GetHdfsParamsOk returns a tuple with the HdfsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetHdfsParamsOk() (*HDFSOnPremSearchParams, bool) {
	if o == nil || o.HdfsParams == nil {
		return nil, false
	}
	return o.HdfsParams, true
}

// HasHdfsParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasHdfsParams() bool {
	if o != nil && o.HdfsParams != nil {
		return true
	}

	return false
}

// SetHdfsParams gets a reference to the given HDFSOnPremSearchParams and assigns it to the HdfsParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetHdfsParams(v HDFSOnPremSearchParams) {
	o.HdfsParams = &v
}

// GetExchangeParams returns the ExchangeParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetExchangeParams() SearchExchangeObjectsRequestParams {
	if o == nil || o.ExchangeParams == nil {
		var ret SearchExchangeObjectsRequestParams
		return ret
	}
	return *o.ExchangeParams
}

// GetExchangeParamsOk returns a tuple with the ExchangeParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetExchangeParamsOk() (*SearchExchangeObjectsRequestParams, bool) {
	if o == nil || o.ExchangeParams == nil {
		return nil, false
	}
	return o.ExchangeParams, true
}

// HasExchangeParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasExchangeParams() bool {
	if o != nil && o.ExchangeParams != nil {
		return true
	}

	return false
}

// SetExchangeParams gets a reference to the given SearchExchangeObjectsRequestParams and assigns it to the ExchangeParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetExchangeParams(v SearchExchangeObjectsRequestParams) {
	o.ExchangeParams = &v
}

// GetPublicFolderParams returns the PublicFolderParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetPublicFolderParams() SearchPublicFolderRequestParams {
	if o == nil || o.PublicFolderParams == nil {
		var ret SearchPublicFolderRequestParams
		return ret
	}
	return *o.PublicFolderParams
}

// GetPublicFolderParamsOk returns a tuple with the PublicFolderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetPublicFolderParamsOk() (*SearchPublicFolderRequestParams, bool) {
	if o == nil || o.PublicFolderParams == nil {
		return nil, false
	}
	return o.PublicFolderParams, true
}

// HasPublicFolderParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasPublicFolderParams() bool {
	if o != nil && o.PublicFolderParams != nil {
		return true
	}

	return false
}

// SetPublicFolderParams gets a reference to the given SearchPublicFolderRequestParams and assigns it to the PublicFolderParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetPublicFolderParams(v SearchPublicFolderRequestParams) {
	o.PublicFolderParams = &v
}

// GetMsTeamsParams returns the MsTeamsParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetMsTeamsParams() SearchMsTeamsRequestParams {
	if o == nil || o.MsTeamsParams == nil {
		var ret SearchMsTeamsRequestParams
		return ret
	}
	return *o.MsTeamsParams
}

// GetMsTeamsParamsOk returns a tuple with the MsTeamsParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetMsTeamsParamsOk() (*SearchMsTeamsRequestParams, bool) {
	if o == nil || o.MsTeamsParams == nil {
		return nil, false
	}
	return o.MsTeamsParams, true
}

// HasMsTeamsParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasMsTeamsParams() bool {
	if o != nil && o.MsTeamsParams != nil {
		return true
	}

	return false
}

// SetMsTeamsParams gets a reference to the given SearchMsTeamsRequestParams and assigns it to the MsTeamsParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetMsTeamsParams(v SearchMsTeamsRequestParams) {
	o.MsTeamsParams = &v
}

// GetSharepointParams returns the SharepointParams field value if set, zero value otherwise.
func (o *SearchIndexedObjectsRequestAllOf) GetSharepointParams() SearchDocumentLibraryRequestParams {
	if o == nil || o.SharepointParams == nil {
		var ret SearchDocumentLibraryRequestParams
		return ret
	}
	return *o.SharepointParams
}

// GetSharepointParamsOk returns a tuple with the SharepointParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexedObjectsRequestAllOf) GetSharepointParamsOk() (*SearchDocumentLibraryRequestParams, bool) {
	if o == nil || o.SharepointParams == nil {
		return nil, false
	}
	return o.SharepointParams, true
}

// HasSharepointParams returns a boolean if a field has been set.
func (o *SearchIndexedObjectsRequestAllOf) HasSharepointParams() bool {
	if o != nil && o.SharepointParams != nil {
		return true
	}

	return false
}

// SetSharepointParams gets a reference to the given SearchDocumentLibraryRequestParams and assigns it to the SharepointParams field.
func (o *SearchIndexedObjectsRequestAllOf) SetSharepointParams(v SearchDocumentLibraryRequestParams) {
	o.SharepointParams = &v
}

func (o SearchIndexedObjectsRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailParams != nil {
		toSerialize["emailParams"] = o.EmailParams
	}
	if o.FileParams != nil {
		toSerialize["fileParams"] = o.FileParams
	}
	if o.CassandraParams != nil {
		toSerialize["cassandraParams"] = o.CassandraParams
	}
	if o.CouchbaseParams != nil {
		toSerialize["couchbaseParams"] = o.CouchbaseParams
	}
	if o.HbaseParams != nil {
		toSerialize["hbaseParams"] = o.HbaseParams
	}
	if o.HiveParams != nil {
		toSerialize["hiveParams"] = o.HiveParams
	}
	if o.MongodbParams != nil {
		toSerialize["mongodbParams"] = o.MongodbParams
	}
	if o.HdfsParams != nil {
		toSerialize["hdfsParams"] = o.HdfsParams
	}
	if o.ExchangeParams != nil {
		toSerialize["exchangeParams"] = o.ExchangeParams
	}
	if o.PublicFolderParams != nil {
		toSerialize["publicFolderParams"] = o.PublicFolderParams
	}
	if o.MsTeamsParams != nil {
		toSerialize["msTeamsParams"] = o.MsTeamsParams
	}
	if o.SharepointParams != nil {
		toSerialize["sharepointParams"] = o.SharepointParams
	}
	return json.Marshal(toSerialize)
}

type NullableSearchIndexedObjectsRequestAllOf struct {
	value *SearchIndexedObjectsRequestAllOf
	isSet bool
}

func (v NullableSearchIndexedObjectsRequestAllOf) Get() *SearchIndexedObjectsRequestAllOf {
	return v.value
}

func (v *NullableSearchIndexedObjectsRequestAllOf) Set(val *SearchIndexedObjectsRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchIndexedObjectsRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchIndexedObjectsRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchIndexedObjectsRequestAllOf(val *SearchIndexedObjectsRequestAllOf) *NullableSearchIndexedObjectsRequestAllOf {
	return &NullableSearchIndexedObjectsRequestAllOf{value: val, isSet: true}
}

func (v NullableSearchIndexedObjectsRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchIndexedObjectsRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


