
package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/wafv2"
	"github.com/aws/aws-sdk-go/service/wafv2/wafv2iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type WAFV2Activities struct {
    client wafv2iface.WAFV2API
}

func NewWAFV2Activities(session *session.Session, config ...*aws.Config) *WAFV2Activities {
    client := wafv2.New(session, config...)
    return &WAFV2Activities{client: client}
}

func (a *WAFV2Activities) AssociateWebACL(ctx context.Context, input *wafv2.AssociateWebACLInput) (*wafv2.AssociateWebACLOutput, error) {
    return a.client.AssociateWebACLWithContext(ctx, input)
}

func (a *WAFV2Activities) CheckCapacity(ctx context.Context, input *wafv2.CheckCapacityInput) (*wafv2.CheckCapacityOutput, error) {
    return a.client.CheckCapacityWithContext(ctx, input)
}

func (a *WAFV2Activities) CreateIPSet(ctx context.Context, input *wafv2.CreateIPSetInput) (*wafv2.CreateIPSetOutput, error) {
    return a.client.CreateIPSetWithContext(ctx, input)
}

func (a *WAFV2Activities) CreateRegexPatternSet(ctx context.Context, input *wafv2.CreateRegexPatternSetInput) (*wafv2.CreateRegexPatternSetOutput, error) {
    return a.client.CreateRegexPatternSetWithContext(ctx, input)
}

func (a *WAFV2Activities) CreateRuleGroup(ctx context.Context, input *wafv2.CreateRuleGroupInput) (*wafv2.CreateRuleGroupOutput, error) {
    return a.client.CreateRuleGroupWithContext(ctx, input)
}

func (a *WAFV2Activities) CreateWebACL(ctx context.Context, input *wafv2.CreateWebACLInput) (*wafv2.CreateWebACLOutput, error) {
    return a.client.CreateWebACLWithContext(ctx, input)
}

func (a *WAFV2Activities) DeleteFirewallManagerRuleGroups(ctx context.Context, input *wafv2.DeleteFirewallManagerRuleGroupsInput) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error) {
    return a.client.DeleteFirewallManagerRuleGroupsWithContext(ctx, input)
}

func (a *WAFV2Activities) DeleteIPSet(ctx context.Context, input *wafv2.DeleteIPSetInput) (*wafv2.DeleteIPSetOutput, error) {
    return a.client.DeleteIPSetWithContext(ctx, input)
}

func (a *WAFV2Activities) DeleteLoggingConfiguration(ctx context.Context, input *wafv2.DeleteLoggingConfigurationInput) (*wafv2.DeleteLoggingConfigurationOutput, error) {
    return a.client.DeleteLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFV2Activities) DeletePermissionPolicy(ctx context.Context, input *wafv2.DeletePermissionPolicyInput) (*wafv2.DeletePermissionPolicyOutput, error) {
    return a.client.DeletePermissionPolicyWithContext(ctx, input)
}

func (a *WAFV2Activities) DeleteRegexPatternSet(ctx context.Context, input *wafv2.DeleteRegexPatternSetInput) (*wafv2.DeleteRegexPatternSetOutput, error) {
    return a.client.DeleteRegexPatternSetWithContext(ctx, input)
}

func (a *WAFV2Activities) DeleteRuleGroup(ctx context.Context, input *wafv2.DeleteRuleGroupInput) (*wafv2.DeleteRuleGroupOutput, error) {
    return a.client.DeleteRuleGroupWithContext(ctx, input)
}

func (a *WAFV2Activities) DeleteWebACL(ctx context.Context, input *wafv2.DeleteWebACLInput) (*wafv2.DeleteWebACLOutput, error) {
    return a.client.DeleteWebACLWithContext(ctx, input)
}

func (a *WAFV2Activities) DescribeManagedRuleGroup(ctx context.Context, input *wafv2.DescribeManagedRuleGroupInput) (*wafv2.DescribeManagedRuleGroupOutput, error) {
    return a.client.DescribeManagedRuleGroupWithContext(ctx, input)
}

func (a *WAFV2Activities) DisassociateWebACL(ctx context.Context, input *wafv2.DisassociateWebACLInput) (*wafv2.DisassociateWebACLOutput, error) {
    return a.client.DisassociateWebACLWithContext(ctx, input)
}

func (a *WAFV2Activities) GetIPSet(ctx context.Context, input *wafv2.GetIPSetInput) (*wafv2.GetIPSetOutput, error) {
    return a.client.GetIPSetWithContext(ctx, input)
}

func (a *WAFV2Activities) GetLoggingConfiguration(ctx context.Context, input *wafv2.GetLoggingConfigurationInput) (*wafv2.GetLoggingConfigurationOutput, error) {
    return a.client.GetLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFV2Activities) GetPermissionPolicy(ctx context.Context, input *wafv2.GetPermissionPolicyInput) (*wafv2.GetPermissionPolicyOutput, error) {
    return a.client.GetPermissionPolicyWithContext(ctx, input)
}

func (a *WAFV2Activities) GetRateBasedStatementManagedKeys(ctx context.Context, input *wafv2.GetRateBasedStatementManagedKeysInput) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
    return a.client.GetRateBasedStatementManagedKeysWithContext(ctx, input)
}

func (a *WAFV2Activities) GetRegexPatternSet(ctx context.Context, input *wafv2.GetRegexPatternSetInput) (*wafv2.GetRegexPatternSetOutput, error) {
    return a.client.GetRegexPatternSetWithContext(ctx, input)
}

func (a *WAFV2Activities) GetRuleGroup(ctx context.Context, input *wafv2.GetRuleGroupInput) (*wafv2.GetRuleGroupOutput, error) {
    return a.client.GetRuleGroupWithContext(ctx, input)
}

func (a *WAFV2Activities) GetSampledRequests(ctx context.Context, input *wafv2.GetSampledRequestsInput) (*wafv2.GetSampledRequestsOutput, error) {
    return a.client.GetSampledRequestsWithContext(ctx, input)
}

func (a *WAFV2Activities) GetWebACL(ctx context.Context, input *wafv2.GetWebACLInput) (*wafv2.GetWebACLOutput, error) {
    return a.client.GetWebACLWithContext(ctx, input)
}

func (a *WAFV2Activities) GetWebACLForResource(ctx context.Context, input *wafv2.GetWebACLForResourceInput) (*wafv2.GetWebACLForResourceOutput, error) {
    return a.client.GetWebACLForResourceWithContext(ctx, input)
}

func (a *WAFV2Activities) ListAvailableManagedRuleGroups(ctx context.Context, input *wafv2.ListAvailableManagedRuleGroupsInput) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
    return a.client.ListAvailableManagedRuleGroupsWithContext(ctx, input)
}

func (a *WAFV2Activities) ListIPSets(ctx context.Context, input *wafv2.ListIPSetsInput) (*wafv2.ListIPSetsOutput, error) {
    return a.client.ListIPSetsWithContext(ctx, input)
}

func (a *WAFV2Activities) ListLoggingConfigurations(ctx context.Context, input *wafv2.ListLoggingConfigurationsInput) (*wafv2.ListLoggingConfigurationsOutput, error) {
    return a.client.ListLoggingConfigurationsWithContext(ctx, input)
}

func (a *WAFV2Activities) ListRegexPatternSets(ctx context.Context, input *wafv2.ListRegexPatternSetsInput) (*wafv2.ListRegexPatternSetsOutput, error) {
    return a.client.ListRegexPatternSetsWithContext(ctx, input)
}

func (a *WAFV2Activities) ListResourcesForWebACL(ctx context.Context, input *wafv2.ListResourcesForWebACLInput) (*wafv2.ListResourcesForWebACLOutput, error) {
    return a.client.ListResourcesForWebACLWithContext(ctx, input)
}

func (a *WAFV2Activities) ListRuleGroups(ctx context.Context, input *wafv2.ListRuleGroupsInput) (*wafv2.ListRuleGroupsOutput, error) {
    return a.client.ListRuleGroupsWithContext(ctx, input)
}

func (a *WAFV2Activities) ListTagsForResource(ctx context.Context, input *wafv2.ListTagsForResourceInput) (*wafv2.ListTagsForResourceOutput, error) {
    return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *WAFV2Activities) ListWebACLs(ctx context.Context, input *wafv2.ListWebACLsInput) (*wafv2.ListWebACLsOutput, error) {
    return a.client.ListWebACLsWithContext(ctx, input)
}

func (a *WAFV2Activities) PutLoggingConfiguration(ctx context.Context, input *wafv2.PutLoggingConfigurationInput) (*wafv2.PutLoggingConfigurationOutput, error) {
    return a.client.PutLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFV2Activities) PutPermissionPolicy(ctx context.Context, input *wafv2.PutPermissionPolicyInput) (*wafv2.PutPermissionPolicyOutput, error) {
    return a.client.PutPermissionPolicyWithContext(ctx, input)
}

func (a *WAFV2Activities) TagResource(ctx context.Context, input *wafv2.TagResourceInput) (*wafv2.TagResourceOutput, error) {
    return a.client.TagResourceWithContext(ctx, input)
}

func (a *WAFV2Activities) UntagResource(ctx context.Context, input *wafv2.UntagResourceInput) (*wafv2.UntagResourceOutput, error) {
    return a.client.UntagResourceWithContext(ctx, input)
}

func (a *WAFV2Activities) UpdateIPSet(ctx context.Context, input *wafv2.UpdateIPSetInput) (*wafv2.UpdateIPSetOutput, error) {
    return a.client.UpdateIPSetWithContext(ctx, input)
}

func (a *WAFV2Activities) UpdateRegexPatternSet(ctx context.Context, input *wafv2.UpdateRegexPatternSetInput) (*wafv2.UpdateRegexPatternSetOutput, error) {
    return a.client.UpdateRegexPatternSetWithContext(ctx, input)
}

func (a *WAFV2Activities) UpdateRuleGroup(ctx context.Context, input *wafv2.UpdateRuleGroupInput) (*wafv2.UpdateRuleGroupOutput, error) {
    return a.client.UpdateRuleGroupWithContext(ctx, input)
}

func (a *WAFV2Activities) UpdateWebACL(ctx context.Context, input *wafv2.UpdateWebACLInput) (*wafv2.UpdateWebACLOutput, error) {
    return a.client.UpdateWebACLWithContext(ctx, input)
}
