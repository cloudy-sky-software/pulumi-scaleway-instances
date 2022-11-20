// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getSecurityGroup(args: GetSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:security_groups:getSecurityGroup", {
        "security_group_id": args.security_group_id,
        "zone": args.zone,
    }, opts);
}

export interface GetSecurityGroupArgs {
    /**
     * UUID of the security group you want to get
     */
    security_group_id: string;
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface GetSecurityGroupResult {
    readonly items: outputs.security_groups.ScalewayInstanceV1GetSecurityGroupResponse;
}

export function getSecurityGroupOutput(args: GetSecurityGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecurityGroupResult> {
    return pulumi.output(args).apply(a => getSecurityGroup(a, opts))
}

export interface GetSecurityGroupOutputArgs {
    /**
     * UUID of the security group you want to get
     */
    security_group_id: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
