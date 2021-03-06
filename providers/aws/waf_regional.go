// Copyright 2020 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

type WafRegionalGenerator struct {
	AWSService
}

func (g *WafRegionalGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := wafregional.New(config)

	if err := g.loadWebACL(svc); err != nil {
		return err
	}
	// AWS WAF Regional API doesn't provide API to build aws_wafregional_web_acl_association resources
	if err := g.loadByteMatchSet(svc); err != nil {
		return err
	}
	if err := g.loadGeoMatchSet(svc); err != nil {
		return err
	}
	if err := g.loadIPSet(svc); err != nil {
		return err
	}
	if err := g.loadRateBasedRules(svc); err != nil {
		return err
	}
	if err := g.loadRegexMatchSets(svc); err != nil {
		return err
	}
	if err := g.loadRegexPatternSets(svc); err != nil {
		return err
	}
	if err := g.loadWafRules(svc); err != nil {
		return err
	}
	if err := g.loadWafRuleGroups(svc); err != nil {
		return err
	}
	if err := g.loadSizeConstraintSets(svc); err != nil {
		return err
	}
	if err := g.loadSqlInjectionMatchSets(svc); err != nil {
		return err
	}
	if err := g.loadXssMatchSet(svc); err != nil {
		return err
	}

	return nil
}

func (g *WafRegionalGenerator) loadWebACL(svc *wafregional.Client) error {
	output, err := svc.ListWebACLsRequest(&wafregional.ListWebACLsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, acl := range output.WebACLs {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*acl.WebACLId,
			*acl.Name,
			"aws_wafregional_web_acl",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadByteMatchSet(svc *wafregional.Client) error {
	output, err := svc.ListByteMatchSetsRequest(&wafregional.ListByteMatchSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, byteMatchSet := range output.ByteMatchSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*byteMatchSet.ByteMatchSetId,
			*byteMatchSet.Name,
			"aws_wafregional_byte_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadGeoMatchSet(svc *wafregional.Client) error {
	output, err := svc.ListGeoMatchSetsRequest(&wafregional.ListGeoMatchSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, matchSet := range output.GeoMatchSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*matchSet.GeoMatchSetId,
			*matchSet.Name,
			"aws_wafregional_geo_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadIPSet(svc *wafregional.Client) error {
	output, err := svc.ListIPSetsRequest(&wafregional.ListIPSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, IPSet := range output.IPSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*IPSet.IPSetId,
			*IPSet.Name,
			"aws_wafregional_ipset",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadRateBasedRules(svc *wafregional.Client) error {
	output, err := svc.ListRateBasedRulesRequest(&wafregional.ListRateBasedRulesInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, rule := range output.Rules {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*rule.RuleId,
			*rule.Name,
			"aws_wafregional_rate_based_rule",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadRegexMatchSets(svc *wafregional.Client) error {
	output, err := svc.ListRegexMatchSetsRequest(&wafregional.ListRegexMatchSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, regexMatchSet := range output.RegexMatchSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*regexMatchSet.RegexMatchSetId,
			*regexMatchSet.Name,
			"aws_wafregional_regex_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadRegexPatternSets(svc *wafregional.Client) error {
	output, err := svc.ListRegexPatternSetsRequest(&wafregional.ListRegexPatternSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, regexPatternSet := range output.RegexPatternSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*regexPatternSet.RegexPatternSetId,
			*regexPatternSet.Name,
			"aws_wafregional_regex_pattern_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadWafRules(svc *wafregional.Client) error {
	output, err := svc.ListRulesRequest(&wafregional.ListRulesInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, rule := range output.Rules {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*rule.RuleId,
			*rule.Name,
			"aws_wafregional_rule",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadWafRuleGroups(svc *wafregional.Client) error {
	output, err := svc.ListRuleGroupsRequest(&wafregional.ListRuleGroupsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, ruleGroup := range output.RuleGroups {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*ruleGroup.RuleGroupId,
			*ruleGroup.Name,
			"aws_wafregional_rule_group",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadSizeConstraintSets(svc *wafregional.Client) error {
	output, err := svc.ListSizeConstraintSetsRequest(&wafregional.ListSizeConstraintSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, sizeConstraintSet := range output.SizeConstraintSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*sizeConstraintSet.SizeConstraintSetId,
			*sizeConstraintSet.Name,
			"aws_wafregional_size_constraint_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadSqlInjectionMatchSets(svc *wafregional.Client) error {
	output, err := svc.ListSqlInjectionMatchSetsRequest(&wafregional.ListSqlInjectionMatchSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, sqlInjectionMatchSet := range output.SqlInjectionMatchSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*sqlInjectionMatchSet.SqlInjectionMatchSetId,
			*sqlInjectionMatchSet.Name,
			"aws_wafregional_sql_injection_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}

func (g *WafRegionalGenerator) loadXssMatchSet(svc *wafregional.Client) error {
	output, err := svc.ListXssMatchSetsRequest(&wafregional.ListXssMatchSetsInput{}).Send(context.Background())
	if err != nil {
		return err
	}
	for _, xssMatchSet := range output.XssMatchSets {
		g.Resources = append(g.Resources, terraform_utils.NewSimpleResource(
			*xssMatchSet.XssMatchSetId,
			*xssMatchSet.Name,
			"aws_wafregional_xss_match_set",
			"aws",
			wafAllowEmptyValues))
	}
	return nil
}
