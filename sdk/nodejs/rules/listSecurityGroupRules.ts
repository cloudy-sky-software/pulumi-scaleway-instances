// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listSecurityGroupRules(args: ListSecurityGroupRulesArgs, opts?: pulumi.InvokeOptions): Promise<ListSecurityGroupRulesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:rules:listSecurityGroupRules", {
        "security_group_id": args.security_group_id,
        "zone": args.zone,
    }, opts);
}

export interface ListSecurityGroupRulesArgs {
    /**
     * UUID of the security group
     */
    security_group_id: string;
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface ListSecurityGroupRulesResult {
    readonly items: outputs.rules.ScalewayInstanceV1ListSecurityGroupRulesResponse;
}

export function listSecurityGroupRulesOutput(args: ListSecurityGroupRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListSecurityGroupRulesResult> {
    return pulumi.output(args).apply(a => listSecurityGroupRules(a, opts))
}

export interface ListSecurityGroupRulesOutputArgs {
    /**
     * UUID of the security group
     */
    security_group_id: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
