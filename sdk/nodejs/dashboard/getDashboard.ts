// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getDashboard(args: GetDashboardArgs, opts?: pulumi.InvokeOptions): Promise<outputs.dashboard.ScalewayInstanceV1GetDashboardResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:dashboard:getDashboard", {
        "zone": args.zone,
    }, opts);
}

export interface GetDashboardArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}
export function getDashboardOutput(args: GetDashboardOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.dashboard.ScalewayInstanceV1GetDashboardResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway-instances:dashboard:getDashboard", {
        "zone": args.zone,
    }, opts);
}

export interface GetDashboardOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
