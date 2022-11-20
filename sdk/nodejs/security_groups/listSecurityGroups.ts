// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listSecurityGroups(args: ListSecurityGroupsArgs, opts?: pulumi.InvokeOptions): Promise<ListSecurityGroupsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:security_groups:listSecurityGroups", {
        "zone": args.zone,
    }, opts);
}

export interface ListSecurityGroupsArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface ListSecurityGroupsResult {
    readonly items: outputs.security_groups.ScalewayInstanceV1ListSecurityGroupsResponse;
}

export function listSecurityGroupsOutput(args: ListSecurityGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ListSecurityGroupsResult> {
    return pulumi.output(args).apply(a => listSecurityGroups(a, opts))
}

export interface ListSecurityGroupsOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
