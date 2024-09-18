// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listDefaultSecurityGroupRules(args: ListDefaultSecurityGroupRulesArgs, opts?: pulumi.InvokeOptions): Promise<outputs.rules.ScalewayInstanceV1ListSecurityGroupRulesResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:rules:listDefaultSecurityGroupRules", {
        "zone": args.zone,
    }, opts);
}

export interface ListDefaultSecurityGroupRulesArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}
export function listDefaultSecurityGroupRulesOutput(args: ListDefaultSecurityGroupRulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.rules.ScalewayInstanceV1ListSecurityGroupRulesResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway-instances:rules:listDefaultSecurityGroupRules", {
        "zone": args.zone,
    }, opts);
}

export interface ListDefaultSecurityGroupRulesOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
