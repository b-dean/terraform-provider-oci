// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAssetsRequest wrapper for the ListAssets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAssets.go.html to see an example of how to use ListAssetsRequest.
type ListAssetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only assets whose lifecycleState matches the given lifecycleState.
	LifecycleState AssetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Source key from where the assets originate.
	SourceKey *string `mandatory:"false" contributesTo:"query" name:"sourceKey"`

	// External asset key.
	ExternalAssetKey *string `mandatory:"false" contributesTo:"query" name:"externalAssetKey"`

	// The type of asset.
	AssetType ListAssetsAssetTypeEnum `mandatory:"false" contributesTo:"query" name:"assetType" omitEmpty:"true"`

	// Unique asset identifier.
	AssetId *string `mandatory:"false" contributesTo:"query" name:"assetId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListAssetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique Inventory identifier.
	InventoryId *string `mandatory:"false" contributesTo:"query" name:"inventoryId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAssetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssetLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAssetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssetsAssetTypeEnum(string(request.AssetType)); !ok && request.AssetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssetType: %s. Supported values are: %s.", request.AssetType, strings.Join(GetListAssetsAssetTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAssetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAssetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAssetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAssetsResponse wrapper for the ListAssets operation
type ListAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssetCollection instances
	AssetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssetsAssetTypeEnum Enum with underlying type: string
type ListAssetsAssetTypeEnum string

// Set of constants representing the allowable values for ListAssetsAssetTypeEnum
const (
	ListAssetsAssetTypeVmwareVm ListAssetsAssetTypeEnum = "VMWARE_VM"
	ListAssetsAssetTypeVm       ListAssetsAssetTypeEnum = "VM"
)

var mappingListAssetsAssetTypeEnum = map[string]ListAssetsAssetTypeEnum{
	"VMWARE_VM": ListAssetsAssetTypeVmwareVm,
	"VM":        ListAssetsAssetTypeVm,
}

var mappingListAssetsAssetTypeEnumLowerCase = map[string]ListAssetsAssetTypeEnum{
	"vmware_vm": ListAssetsAssetTypeVmwareVm,
	"vm":        ListAssetsAssetTypeVm,
}

// GetListAssetsAssetTypeEnumValues Enumerates the set of values for ListAssetsAssetTypeEnum
func GetListAssetsAssetTypeEnumValues() []ListAssetsAssetTypeEnum {
	values := make([]ListAssetsAssetTypeEnum, 0)
	for _, v := range mappingListAssetsAssetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssetsAssetTypeEnumStringValues Enumerates the set of values in String for ListAssetsAssetTypeEnum
func GetListAssetsAssetTypeEnumStringValues() []string {
	return []string{
		"VMWARE_VM",
		"VM",
	}
}

// GetMappingListAssetsAssetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssetsAssetTypeEnum(val string) (ListAssetsAssetTypeEnum, bool) {
	enum, ok := mappingListAssetsAssetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssetsSortOrderEnum Enum with underlying type: string
type ListAssetsSortOrderEnum string

// Set of constants representing the allowable values for ListAssetsSortOrderEnum
const (
	ListAssetsSortOrderAsc  ListAssetsSortOrderEnum = "ASC"
	ListAssetsSortOrderDesc ListAssetsSortOrderEnum = "DESC"
)

var mappingListAssetsSortOrderEnum = map[string]ListAssetsSortOrderEnum{
	"ASC":  ListAssetsSortOrderAsc,
	"DESC": ListAssetsSortOrderDesc,
}

var mappingListAssetsSortOrderEnumLowerCase = map[string]ListAssetsSortOrderEnum{
	"asc":  ListAssetsSortOrderAsc,
	"desc": ListAssetsSortOrderDesc,
}

// GetListAssetsSortOrderEnumValues Enumerates the set of values for ListAssetsSortOrderEnum
func GetListAssetsSortOrderEnumValues() []ListAssetsSortOrderEnum {
	values := make([]ListAssetsSortOrderEnum, 0)
	for _, v := range mappingListAssetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssetsSortOrderEnumStringValues Enumerates the set of values in String for ListAssetsSortOrderEnum
func GetListAssetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAssetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssetsSortOrderEnum(val string) (ListAssetsSortOrderEnum, bool) {
	enum, ok := mappingListAssetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAssetsSortByEnum Enum with underlying type: string
type ListAssetsSortByEnum string

// Set of constants representing the allowable values for ListAssetsSortByEnum
const (
	ListAssetsSortByTimecreated ListAssetsSortByEnum = "timeCreated"
	ListAssetsSortByTimeupdated ListAssetsSortByEnum = "timeUpdated"
	ListAssetsSortByDisplayname ListAssetsSortByEnum = "displayName"
)

var mappingListAssetsSortByEnum = map[string]ListAssetsSortByEnum{
	"timeCreated": ListAssetsSortByTimecreated,
	"timeUpdated": ListAssetsSortByTimeupdated,
	"displayName": ListAssetsSortByDisplayname,
}

var mappingListAssetsSortByEnumLowerCase = map[string]ListAssetsSortByEnum{
	"timecreated": ListAssetsSortByTimecreated,
	"timeupdated": ListAssetsSortByTimeupdated,
	"displayname": ListAssetsSortByDisplayname,
}

// GetListAssetsSortByEnumValues Enumerates the set of values for ListAssetsSortByEnum
func GetListAssetsSortByEnumValues() []ListAssetsSortByEnum {
	values := make([]ListAssetsSortByEnum, 0)
	for _, v := range mappingListAssetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAssetsSortByEnumStringValues Enumerates the set of values in String for ListAssetsSortByEnum
func GetListAssetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListAssetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAssetsSortByEnum(val string) (ListAssetsSortByEnum, bool) {
	enum, ok := mappingListAssetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
