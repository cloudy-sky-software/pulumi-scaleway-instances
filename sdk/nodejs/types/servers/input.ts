// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export interface ScalewayInstanceV1VolumeServerTemplateArgs {
    /**
     * The ID of the snapshot on which this volume will be based
     */
    base_snapshot?: pulumi.Input<string>;
    /**
     * Force the server to boot on this volume
     */
    boot?: pulumi.Input<boolean>;
    /**
     * UUID of the volume
     */
    id?: pulumi.Input<string>;
    /**
     * Name of the volume
     */
    name?: pulumi.Input<string>;
    /**
     * Organization ID of the volume
     */
    organization?: pulumi.Input<string>;
    /**
     * Project ID of the volume
     */
    project?: pulumi.Input<string>;
    /**
     * Disk size of the volume, must be a multiple of 512 (in bytes)
     */
    size?: pulumi.Input<number>;
    volume_type?: pulumi.Input<enums.servers.ScalewayInstanceV1VolumeServerTemplateVolumeType>;
}
/**
 * scalewayInstanceV1VolumeServerTemplateArgsProvideDefaults sets the appropriate defaults for ScalewayInstanceV1VolumeServerTemplateArgs
 */
export function scalewayInstanceV1VolumeServerTemplateArgsProvideDefaults(val: ScalewayInstanceV1VolumeServerTemplateArgs): ScalewayInstanceV1VolumeServerTemplateArgs {
    return {
        ...val,
        boot: (val.boot) ?? false,
        volume_type: (val.volume_type) ?? "l_ssd",
    };
}
