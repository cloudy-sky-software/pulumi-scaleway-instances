// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getPlacementGroupServer(args: GetPlacementGroupServerArgs, opts?: pulumi.InvokeOptions): Promise<GetPlacementGroupServerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:servers:getPlacementGroupServer", {
        "placementGroupId": args.placementGroupId,
        "zone": args.zone,
    }, opts);
}

export interface GetPlacementGroupServerArgs {
    /**
     * UUID of the placement group
     */
    placementGroupId: string;
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface GetPlacementGroupServerResult {
    readonly items: outputs.servers.ScalewayInstanceV1GetPlacementGroupServersResponse;
}
export function getPlacementGroupServerOutput(args: GetPlacementGroupServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlacementGroupServerResult> {
    return pulumi.output(args).apply((a: any) => getPlacementGroupServer(a, opts))
}

export interface GetPlacementGroupServerOutputArgs {
    /**
     * UUID of the placement group
     */
    placementGroupId: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}