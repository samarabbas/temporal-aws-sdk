package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"go.temporal.io/sdk/workflow"
)

type CloudSearchClient interface {
       BuildSuggesters(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error)
       BuildSuggestersAsync(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) *CloudsearchBuildSuggestersResult

       CreateDomain(ctx workflow.Context, input *cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error)
       CreateDomainAsync(ctx workflow.Context, input *cloudsearch.CreateDomainInput) *CloudsearchCreateDomainResult

       DefineAnalysisScheme(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error)
       DefineAnalysisSchemeAsync(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) *CloudsearchDefineAnalysisSchemeResult

       DefineExpression(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error)
       DefineExpressionAsync(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) *CloudsearchDefineExpressionResult

       DefineIndexField(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error)
       DefineIndexFieldAsync(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) *CloudsearchDefineIndexFieldResult

       DefineSuggester(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error)
       DefineSuggesterAsync(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) *CloudsearchDefineSuggesterResult

       DeleteAnalysisScheme(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error)
       DeleteAnalysisSchemeAsync(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) *CloudsearchDeleteAnalysisSchemeResult

       DeleteDomain(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error)
       DeleteDomainAsync(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) *CloudsearchDeleteDomainResult

       DeleteExpression(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error)
       DeleteExpressionAsync(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) *CloudsearchDeleteExpressionResult

       DeleteIndexField(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error)
       DeleteIndexFieldAsync(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) *CloudsearchDeleteIndexFieldResult

       DeleteSuggester(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error)
       DeleteSuggesterAsync(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) *CloudsearchDeleteSuggesterResult

       DescribeAnalysisSchemes(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error)
       DescribeAnalysisSchemesAsync(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) *CloudsearchDescribeAnalysisSchemesResult

       DescribeAvailabilityOptions(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error)
       DescribeAvailabilityOptionsAsync(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) *CloudsearchDescribeAvailabilityOptionsResult

       DescribeDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error)
       DescribeDomainEndpointOptionsAsync(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) *CloudsearchDescribeDomainEndpointOptionsResult

       DescribeDomains(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error)
       DescribeDomainsAsync(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) *CloudsearchDescribeDomainsResult

       DescribeExpressions(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error)
       DescribeExpressionsAsync(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) *CloudsearchDescribeExpressionsResult

       DescribeIndexFields(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error)
       DescribeIndexFieldsAsync(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) *CloudsearchDescribeIndexFieldsResult

       DescribeScalingParameters(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error)
       DescribeScalingParametersAsync(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) *CloudsearchDescribeScalingParametersResult

       DescribeServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error)
       DescribeServiceAccessPoliciesAsync(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) *CloudsearchDescribeServiceAccessPoliciesResult

       DescribeSuggesters(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error)
       DescribeSuggestersAsync(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) *CloudsearchDescribeSuggestersResult

       IndexDocuments(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error)
       IndexDocumentsAsync(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) *CloudsearchIndexDocumentsResult

       ListDomainNames(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error)
       ListDomainNamesAsync(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) *CloudsearchListDomainNamesResult

       UpdateAvailabilityOptions(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error)
       UpdateAvailabilityOptionsAsync(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) *CloudsearchUpdateAvailabilityOptionsResult

       UpdateDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error)
       UpdateDomainEndpointOptionsAsync(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) *CloudsearchUpdateDomainEndpointOptionsResult

       UpdateScalingParameters(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error)
       UpdateScalingParametersAsync(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) *CloudsearchUpdateScalingParametersResult

       UpdateServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error)
       UpdateServiceAccessPoliciesAsync(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) *CloudsearchUpdateServiceAccessPoliciesResult
}

type CloudSearchStub struct {
}

func NewCloudSearchStub() CloudSearchClient {
    return &CloudSearchStub{}
}

type CloudsearchBuildSuggestersResult struct {
	Result workflow.Future
}

func (r *CloudsearchBuildSuggestersResult) Get(ctx workflow.Context) (*cloudsearch.BuildSuggestersOutput, error) {
    var output cloudsearch.BuildSuggestersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchCreateDomainResult struct {
	Result workflow.Future
}

func (r *CloudsearchCreateDomainResult) Get(ctx workflow.Context) (*cloudsearch.CreateDomainOutput, error) {
    var output cloudsearch.CreateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDefineAnalysisSchemeResult struct {
	Result workflow.Future
}

func (r *CloudsearchDefineAnalysisSchemeResult) Get(ctx workflow.Context) (*cloudsearch.DefineAnalysisSchemeOutput, error) {
    var output cloudsearch.DefineAnalysisSchemeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDefineExpressionResult struct {
	Result workflow.Future
}

func (r *CloudsearchDefineExpressionResult) Get(ctx workflow.Context) (*cloudsearch.DefineExpressionOutput, error) {
    var output cloudsearch.DefineExpressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDefineIndexFieldResult struct {
	Result workflow.Future
}

func (r *CloudsearchDefineIndexFieldResult) Get(ctx workflow.Context) (*cloudsearch.DefineIndexFieldOutput, error) {
    var output cloudsearch.DefineIndexFieldOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDefineSuggesterResult struct {
	Result workflow.Future
}

func (r *CloudsearchDefineSuggesterResult) Get(ctx workflow.Context) (*cloudsearch.DefineSuggesterOutput, error) {
    var output cloudsearch.DefineSuggesterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDeleteAnalysisSchemeResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteAnalysisSchemeResult) Get(ctx workflow.Context) (*cloudsearch.DeleteAnalysisSchemeOutput, error) {
    var output cloudsearch.DeleteAnalysisSchemeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDeleteDomainResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteDomainResult) Get(ctx workflow.Context) (*cloudsearch.DeleteDomainOutput, error) {
    var output cloudsearch.DeleteDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDeleteExpressionResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteExpressionResult) Get(ctx workflow.Context) (*cloudsearch.DeleteExpressionOutput, error) {
    var output cloudsearch.DeleteExpressionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDeleteIndexFieldResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteIndexFieldResult) Get(ctx workflow.Context) (*cloudsearch.DeleteIndexFieldOutput, error) {
    var output cloudsearch.DeleteIndexFieldOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDeleteSuggesterResult struct {
	Result workflow.Future
}

func (r *CloudsearchDeleteSuggesterResult) Get(ctx workflow.Context) (*cloudsearch.DeleteSuggesterOutput, error) {
    var output cloudsearch.DeleteSuggesterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeAnalysisSchemesResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeAnalysisSchemesResult) Get(ctx workflow.Context) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
    var output cloudsearch.DescribeAnalysisSchemesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeAvailabilityOptionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeAvailabilityOptionsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
    var output cloudsearch.DescribeAvailabilityOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeDomainEndpointOptionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeDomainEndpointOptionsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
    var output cloudsearch.DescribeDomainEndpointOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeDomainsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeDomainsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeDomainsOutput, error) {
    var output cloudsearch.DescribeDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeExpressionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeExpressionsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeExpressionsOutput, error) {
    var output cloudsearch.DescribeExpressionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeIndexFieldsResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeIndexFieldsResult) Get(ctx workflow.Context) (*cloudsearch.DescribeIndexFieldsOutput, error) {
    var output cloudsearch.DescribeIndexFieldsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeScalingParametersResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeScalingParametersResult) Get(ctx workflow.Context) (*cloudsearch.DescribeScalingParametersOutput, error) {
    var output cloudsearch.DescribeScalingParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeServiceAccessPoliciesResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeServiceAccessPoliciesResult) Get(ctx workflow.Context) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
    var output cloudsearch.DescribeServiceAccessPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchDescribeSuggestersResult struct {
	Result workflow.Future
}

func (r *CloudsearchDescribeSuggestersResult) Get(ctx workflow.Context) (*cloudsearch.DescribeSuggestersOutput, error) {
    var output cloudsearch.DescribeSuggestersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchIndexDocumentsResult struct {
	Result workflow.Future
}

func (r *CloudsearchIndexDocumentsResult) Get(ctx workflow.Context) (*cloudsearch.IndexDocumentsOutput, error) {
    var output cloudsearch.IndexDocumentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchListDomainNamesResult struct {
	Result workflow.Future
}

func (r *CloudsearchListDomainNamesResult) Get(ctx workflow.Context) (*cloudsearch.ListDomainNamesOutput, error) {
    var output cloudsearch.ListDomainNamesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchUpdateAvailabilityOptionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchUpdateAvailabilityOptionsResult) Get(ctx workflow.Context) (*cloudsearch.UpdateAvailabilityOptionsOutput, error) {
    var output cloudsearch.UpdateAvailabilityOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchUpdateDomainEndpointOptionsResult struct {
	Result workflow.Future
}

func (r *CloudsearchUpdateDomainEndpointOptionsResult) Get(ctx workflow.Context) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error) {
    var output cloudsearch.UpdateDomainEndpointOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchUpdateScalingParametersResult struct {
	Result workflow.Future
}

func (r *CloudsearchUpdateScalingParametersResult) Get(ctx workflow.Context) (*cloudsearch.UpdateScalingParametersOutput, error) {
    var output cloudsearch.UpdateScalingParametersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type CloudsearchUpdateServiceAccessPoliciesResult struct {
	Result workflow.Future
}

func (r *CloudsearchUpdateServiceAccessPoliciesResult) Get(ctx workflow.Context) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error) {
    var output cloudsearch.UpdateServiceAccessPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) BuildSuggesters(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error) {
    var output cloudsearch.BuildSuggestersOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.BuildSuggesters", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) BuildSuggestersAsync(ctx workflow.Context, input *cloudsearch.BuildSuggestersInput) *CloudsearchBuildSuggestersResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.BuildSuggesters", input)
    return &CloudsearchBuildSuggestersResult{Result: future}
}

func (a *CloudSearchStub) CreateDomain(ctx workflow.Context, input *cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error) {
    var output cloudsearch.CreateDomainOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.CreateDomain", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) CreateDomainAsync(ctx workflow.Context, input *cloudsearch.CreateDomainInput) *CloudsearchCreateDomainResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.CreateDomain", input)
    return &CloudsearchCreateDomainResult{Result: future}
}

func (a *CloudSearchStub) DefineAnalysisScheme(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error) {
    var output cloudsearch.DefineAnalysisSchemeOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DefineAnalysisScheme", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DefineAnalysisSchemeAsync(ctx workflow.Context, input *cloudsearch.DefineAnalysisSchemeInput) *CloudsearchDefineAnalysisSchemeResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DefineAnalysisScheme", input)
    return &CloudsearchDefineAnalysisSchemeResult{Result: future}
}

func (a *CloudSearchStub) DefineExpression(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error) {
    var output cloudsearch.DefineExpressionOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DefineExpression", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DefineExpressionAsync(ctx workflow.Context, input *cloudsearch.DefineExpressionInput) *CloudsearchDefineExpressionResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DefineExpression", input)
    return &CloudsearchDefineExpressionResult{Result: future}
}

func (a *CloudSearchStub) DefineIndexField(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error) {
    var output cloudsearch.DefineIndexFieldOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DefineIndexField", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DefineIndexFieldAsync(ctx workflow.Context, input *cloudsearch.DefineIndexFieldInput) *CloudsearchDefineIndexFieldResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DefineIndexField", input)
    return &CloudsearchDefineIndexFieldResult{Result: future}
}

func (a *CloudSearchStub) DefineSuggester(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error) {
    var output cloudsearch.DefineSuggesterOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DefineSuggester", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DefineSuggesterAsync(ctx workflow.Context, input *cloudsearch.DefineSuggesterInput) *CloudsearchDefineSuggesterResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DefineSuggester", input)
    return &CloudsearchDefineSuggesterResult{Result: future}
}

func (a *CloudSearchStub) DeleteAnalysisScheme(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error) {
    var output cloudsearch.DeleteAnalysisSchemeOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteAnalysisScheme", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteAnalysisSchemeAsync(ctx workflow.Context, input *cloudsearch.DeleteAnalysisSchemeInput) *CloudsearchDeleteAnalysisSchemeResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteAnalysisScheme", input)
    return &CloudsearchDeleteAnalysisSchemeResult{Result: future}
}

func (a *CloudSearchStub) DeleteDomain(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error) {
    var output cloudsearch.DeleteDomainOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteDomain", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteDomainAsync(ctx workflow.Context, input *cloudsearch.DeleteDomainInput) *CloudsearchDeleteDomainResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteDomain", input)
    return &CloudsearchDeleteDomainResult{Result: future}
}

func (a *CloudSearchStub) DeleteExpression(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error) {
    var output cloudsearch.DeleteExpressionOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteExpression", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteExpressionAsync(ctx workflow.Context, input *cloudsearch.DeleteExpressionInput) *CloudsearchDeleteExpressionResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteExpression", input)
    return &CloudsearchDeleteExpressionResult{Result: future}
}

func (a *CloudSearchStub) DeleteIndexField(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error) {
    var output cloudsearch.DeleteIndexFieldOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteIndexField", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteIndexFieldAsync(ctx workflow.Context, input *cloudsearch.DeleteIndexFieldInput) *CloudsearchDeleteIndexFieldResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteIndexField", input)
    return &CloudsearchDeleteIndexFieldResult{Result: future}
}

func (a *CloudSearchStub) DeleteSuggester(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error) {
    var output cloudsearch.DeleteSuggesterOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteSuggester", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DeleteSuggesterAsync(ctx workflow.Context, input *cloudsearch.DeleteSuggesterInput) *CloudsearchDeleteSuggesterResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DeleteSuggester", input)
    return &CloudsearchDeleteSuggesterResult{Result: future}
}

func (a *CloudSearchStub) DescribeAnalysisSchemes(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error) {
    var output cloudsearch.DescribeAnalysisSchemesOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeAnalysisSchemes", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeAnalysisSchemesAsync(ctx workflow.Context, input *cloudsearch.DescribeAnalysisSchemesInput) *CloudsearchDescribeAnalysisSchemesResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeAnalysisSchemes", input)
    return &CloudsearchDescribeAnalysisSchemesResult{Result: future}
}

func (a *CloudSearchStub) DescribeAvailabilityOptions(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error) {
    var output cloudsearch.DescribeAvailabilityOptionsOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeAvailabilityOptions", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeAvailabilityOptionsAsync(ctx workflow.Context, input *cloudsearch.DescribeAvailabilityOptionsInput) *CloudsearchDescribeAvailabilityOptionsResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeAvailabilityOptions", input)
    return &CloudsearchDescribeAvailabilityOptionsResult{Result: future}
}

func (a *CloudSearchStub) DescribeDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) (*cloudsearch.DescribeDomainEndpointOptionsOutput, error) {
    var output cloudsearch.DescribeDomainEndpointOptionsOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeDomainEndpointOptions", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeDomainEndpointOptionsAsync(ctx workflow.Context, input *cloudsearch.DescribeDomainEndpointOptionsInput) *CloudsearchDescribeDomainEndpointOptionsResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeDomainEndpointOptions", input)
    return &CloudsearchDescribeDomainEndpointOptionsResult{Result: future}
}

func (a *CloudSearchStub) DescribeDomains(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error) {
    var output cloudsearch.DescribeDomainsOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeDomains", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeDomainsAsync(ctx workflow.Context, input *cloudsearch.DescribeDomainsInput) *CloudsearchDescribeDomainsResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeDomains", input)
    return &CloudsearchDescribeDomainsResult{Result: future}
}

func (a *CloudSearchStub) DescribeExpressions(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error) {
    var output cloudsearch.DescribeExpressionsOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeExpressions", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeExpressionsAsync(ctx workflow.Context, input *cloudsearch.DescribeExpressionsInput) *CloudsearchDescribeExpressionsResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeExpressions", input)
    return &CloudsearchDescribeExpressionsResult{Result: future}
}

func (a *CloudSearchStub) DescribeIndexFields(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error) {
    var output cloudsearch.DescribeIndexFieldsOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeIndexFields", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeIndexFieldsAsync(ctx workflow.Context, input *cloudsearch.DescribeIndexFieldsInput) *CloudsearchDescribeIndexFieldsResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeIndexFields", input)
    return &CloudsearchDescribeIndexFieldsResult{Result: future}
}

func (a *CloudSearchStub) DescribeScalingParameters(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error) {
    var output cloudsearch.DescribeScalingParametersOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeScalingParameters", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeScalingParametersAsync(ctx workflow.Context, input *cloudsearch.DescribeScalingParametersInput) *CloudsearchDescribeScalingParametersResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeScalingParameters", input)
    return &CloudsearchDescribeScalingParametersResult{Result: future}
}

func (a *CloudSearchStub) DescribeServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error) {
    var output cloudsearch.DescribeServiceAccessPoliciesOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeServiceAccessPolicies", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeServiceAccessPoliciesAsync(ctx workflow.Context, input *cloudsearch.DescribeServiceAccessPoliciesInput) *CloudsearchDescribeServiceAccessPoliciesResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeServiceAccessPolicies", input)
    return &CloudsearchDescribeServiceAccessPoliciesResult{Result: future}
}

func (a *CloudSearchStub) DescribeSuggesters(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error) {
    var output cloudsearch.DescribeSuggestersOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeSuggesters", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) DescribeSuggestersAsync(ctx workflow.Context, input *cloudsearch.DescribeSuggestersInput) *CloudsearchDescribeSuggestersResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.DescribeSuggesters", input)
    return &CloudsearchDescribeSuggestersResult{Result: future}
}

func (a *CloudSearchStub) IndexDocuments(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error) {
    var output cloudsearch.IndexDocumentsOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.IndexDocuments", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) IndexDocumentsAsync(ctx workflow.Context, input *cloudsearch.IndexDocumentsInput) *CloudsearchIndexDocumentsResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.IndexDocuments", input)
    return &CloudsearchIndexDocumentsResult{Result: future}
}

func (a *CloudSearchStub) ListDomainNames(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error) {
    var output cloudsearch.ListDomainNamesOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.ListDomainNames", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) ListDomainNamesAsync(ctx workflow.Context, input *cloudsearch.ListDomainNamesInput) *CloudsearchListDomainNamesResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.ListDomainNames", input)
    return &CloudsearchListDomainNamesResult{Result: future}
}

func (a *CloudSearchStub) UpdateAvailabilityOptions(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error) {
    var output cloudsearch.UpdateAvailabilityOptionsOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.UpdateAvailabilityOptions", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) UpdateAvailabilityOptionsAsync(ctx workflow.Context, input *cloudsearch.UpdateAvailabilityOptionsInput) *CloudsearchUpdateAvailabilityOptionsResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.UpdateAvailabilityOptions", input)
    return &CloudsearchUpdateAvailabilityOptionsResult{Result: future}
}

func (a *CloudSearchStub) UpdateDomainEndpointOptions(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) (*cloudsearch.UpdateDomainEndpointOptionsOutput, error) {
    var output cloudsearch.UpdateDomainEndpointOptionsOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.UpdateDomainEndpointOptions", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) UpdateDomainEndpointOptionsAsync(ctx workflow.Context, input *cloudsearch.UpdateDomainEndpointOptionsInput) *CloudsearchUpdateDomainEndpointOptionsResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.UpdateDomainEndpointOptions", input)
    return &CloudsearchUpdateDomainEndpointOptionsResult{Result: future}
}

func (a *CloudSearchStub) UpdateScalingParameters(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error) {
    var output cloudsearch.UpdateScalingParametersOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.UpdateScalingParameters", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) UpdateScalingParametersAsync(ctx workflow.Context, input *cloudsearch.UpdateScalingParametersInput) *CloudsearchUpdateScalingParametersResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.UpdateScalingParameters", input)
    return &CloudsearchUpdateScalingParametersResult{Result: future}
}

func (a *CloudSearchStub) UpdateServiceAccessPolicies(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error) {
    var output cloudsearch.UpdateServiceAccessPoliciesOutput
    err := workflow.ExecuteActivity(ctx, "CloudSearch.UpdateServiceAccessPolicies", input).Get(ctx, &output)
    return &output, err
}

func (a *CloudSearchStub) UpdateServiceAccessPoliciesAsync(ctx workflow.Context, input *cloudsearch.UpdateServiceAccessPoliciesInput) *CloudsearchUpdateServiceAccessPoliciesResult {
    future := workflow.ExecuteActivity(ctx, "CloudSearch.UpdateServiceAccessPolicies", input)
    return &CloudsearchUpdateServiceAccessPoliciesResult{Result: future}
}
