// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export interface ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs {
    /**
     * Action to apply when the rule matches a packet
     */
    action?: pulumi.Input<enums.rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleAction>;
    /**
     * Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY
     */
    dest_port_from?: pulumi.Input<number>;
    /**
     * End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from
     */
    dest_port_to?: pulumi.Input<number>;
    /**
     * Direction the rule applies to
     */
    direction?: pulumi.Input<enums.rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleDirection>;
    /**
     * Indicates if this rule is editable. Rules with the value false will be ignored
     */
    editable?: pulumi.Input<boolean>;
    /**
     * UUID of the security rule to update. If no value is provided, a new rule will be created
     */
    id?: pulumi.Input<string>;
    /**
     * The range of IP address this rules applies to (IP network)
     */
    ip_range?: pulumi.Input<string>;
    /**
     * Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined
     */
    position?: pulumi.Input<number>;
    /**
     * Protocol family this rule applies to
     */
    protocol?: pulumi.Input<enums.rules.ScalewayInstanceV1SetSecurityGroupRulesRequestRuleProtocol>;
    /**
     * Zone of the rule. This field is ignored
     */
    zone?: pulumi.Input<string>;
}
/**
 * scalewayInstanceV1SetSecurityGroupRulesRequestRuleArgsProvideDefaults sets the appropriate defaults for ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs
 */
export function scalewayInstanceV1SetSecurityGroupRulesRequestRuleArgsProvideDefaults(val: ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs): ScalewayInstanceV1SetSecurityGroupRulesRequestRuleArgs {
    return {
        ...val,
        action: (val.action) ?? "accept",
        direction: (val.direction) ?? "inbound",
        protocol: (val.protocol) ?? "TCP",
    };
}