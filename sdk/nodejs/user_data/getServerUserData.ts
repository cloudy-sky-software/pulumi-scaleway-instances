// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getServerUserData(args: GetServerUserDataArgs, opts?: pulumi.InvokeOptions): Promise<outputs.user_data.ScalewayStdFile> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway-instances:user_data:getServerUserData", {
        "key": args.key,
        "serverId": args.serverId,
        "zone": args.zone,
    }, opts);
}

export interface GetServerUserDataArgs {
    /**
     * Key of the user data to get
     */
    key: string;
    /**
     * UUID of the server
     */
    serverId: string;
    /**
     * The zone you want to target
     */
    zone: string;
}
export function getServerUserDataOutput(args: GetServerUserDataOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<outputs.user_data.ScalewayStdFile> {
    return pulumi.output(args).apply((a: any) => getServerUserData(a, opts))
}

export interface GetServerUserDataOutputArgs {
    /**
     * Key of the user data to get
     */
    key: pulumi.Input<string>;
    /**
     * UUID of the server
     */
    serverId: pulumi.Input<string>;
    /**
     * The zone you want to target
     */
    zone: pulumi.Input<string>;
}
