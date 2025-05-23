// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getPlacementGroupServers(args: GetPlacementGroupServersArgs, opts?: pulumi.InvokeOptions): Promise<outputs.servers.ScalewayInstanceV1GetPlacementGroupServersResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:servers:getPlacementGroupServers", {
        "placementGroupId": args.placementGroupId,
        "zone": args.zone,
    }, opts);
}

export interface GetPlacementGroupServersArgs {
    /**
     * UUID of the placement group
     */
    placementGroupId: string;
    /**
     * The zone you want to target
     */
    zone: string;
}
export function getPlacementGroupServersOutput(args: GetPlacementGroupServersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<outputs.servers.ScalewayInstanceV1GetPlacementGroupServersResponse> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway-instances:servers:getPlacementGroupServers", {
        "placementGroupId": args.placementGroupId,
        "zone": args.zone,
    }, opts);
}

export interface GetPlacementGroupServersOutputArgs {
    /**
     * UUID of the placement group
     */
    placementGroupId: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
