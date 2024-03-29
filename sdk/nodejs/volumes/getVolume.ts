// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getVolume(args: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:volumes:getVolume", {
        "id": args.id,
        "zone": args.zone,
    }, opts);
}

export interface GetVolumeArgs {
    /**
     * UUID of the volume you want to get
     */
    id: string;
    /**
     * The zone you want to target
     */
    zone: string;
}

export interface GetVolumeResult {
    readonly items: outputs.volumes.ScalewayInstanceV1GetVolumeResponse;
}
export function getVolumeOutput(args: GetVolumeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeResult> {
    return pulumi.output(args).apply((a: any) => getVolume(a, opts))
}

export interface GetVolumeOutputArgs {
    /**
     * UUID of the volume you want to get
     */
    id: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
