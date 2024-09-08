// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function listSnapshots(args: ListSnapshotsArgs, opts?: pulumi.InvokeOptions): Promise<outputs.snapshots.ScalewayInstanceV1ListSnapshotsResponse> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:snapshots:listSnapshots", {
        "zone": args.zone,
    }, opts);
}

export interface ListSnapshotsArgs {
    /**
     * The zone you want to target
     */
    zone: string;
}
export function listSnapshotsOutput(args: ListSnapshotsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.snapshots.ScalewayInstanceV1ListSnapshotsResponse> {
    return pulumi.output(args).apply((a: any) => listSnapshots(a, opts))
}

export interface ListSnapshotsOutputArgs {
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
