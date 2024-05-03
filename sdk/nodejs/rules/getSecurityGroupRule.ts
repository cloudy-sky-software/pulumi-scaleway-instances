// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getSecurityGroupRule(args: GetSecurityGroupRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupRuleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:rules:getSecurityGroupRule", {
        "id": args.id,
        "securityGroupId": args.securityGroupId,
        "zone": args.zone,
    }, opts);
}

export interface GetSecurityGroupRuleArgs {
    id: string;
    securityGroupId: string;
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface GetSecurityGroupRuleResult {
    readonly items: outputs.rules.ScalewayInstanceV1GetSecurityGroupRuleResponse;
}
export function getSecurityGroupRuleOutput(args: GetSecurityGroupRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupRuleResult> {
    return pulumi.output(args).apply((a: any) => getSecurityGroupRule(a, opts))
}

export interface GetSecurityGroupRuleOutputArgs {
    id: pulumi.Input<string>;
    securityGroupId: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
