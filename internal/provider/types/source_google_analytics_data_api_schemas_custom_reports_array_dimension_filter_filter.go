// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

type SourceGoogleAnalyticsDataAPISchemasCustomReportsArrayDimensionFilterFilter struct {
	BetweenFilter *SourceGoogleAnalyticsDataAPIUpdateSchemasCustomReportsArrayBetweenFilter `tfsdk:"between_filter" tfPlanOnly:"true"`
	InListFilter  *SourceGoogleAnalyticsDataAPIUpdateSchemasCustomReportsArrayInListFilter  `tfsdk:"in_list_filter" tfPlanOnly:"true"`
	NumericFilter *SourceGoogleAnalyticsDataAPIUpdateSchemasCustomReportsArrayNumericFilter `tfsdk:"numeric_filter" tfPlanOnly:"true"`
	StringFilter  *SourceGoogleAnalyticsDataAPIUpdateSchemasCustomReportsArrayStringFilter  `tfsdk:"string_filter" tfPlanOnly:"true"`
}