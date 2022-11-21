// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

export interface ScalewayInstanceV1BootscriptArgs {
    /**
     * The bootscript arch
     */
    arch?: pulumi.Input<enums.images.ScalewayInstanceV1BootscriptArch>;
    /**
     * The bootscript arguments
     */
    bootcmdargs?: pulumi.Input<string>;
    /**
     * Dispmay if the bootscript is the default bootscript if no other boot option is configured
     */
    default?: pulumi.Input<boolean>;
    /**
     * Provide information regarding a Device Tree Binary (dtb) for use with C1 servers
     */
    dtb?: pulumi.Input<string>;
    /**
     * The bootscript ID
     */
    id?: pulumi.Input<string>;
    /**
     * The initrd (initial ramdisk) configuration
     */
    initrd?: pulumi.Input<string>;
    /**
     * The server kernel version
     */
    kernel?: pulumi.Input<string>;
    /**
     * The bootscript organization ID
     */
    organization?: pulumi.Input<string>;
    /**
     * The bootscript project ID
     */
    project?: pulumi.Input<string>;
    /**
     * Provide information if the bootscript is public
     */
    public?: pulumi.Input<boolean>;
    /**
     * The bootscript title
     */
    title?: pulumi.Input<string>;
    /**
     * The zone in which is the bootscript
     */
    zone?: pulumi.Input<string>;
}
/**
 * scalewayInstanceV1BootscriptArgsProvideDefaults sets the appropriate defaults for ScalewayInstanceV1BootscriptArgs
 */
export function scalewayInstanceV1BootscriptArgsProvideDefaults(val: ScalewayInstanceV1BootscriptArgs): ScalewayInstanceV1BootscriptArgs {
    return {
        ...val,
        arch: (val.arch) ?? "x86_64",
    };
}

export interface ScalewayInstanceV1VolumeArgs {
    /**
     * The volume creation date (RFC 3339 format)
     */
    creation_date?: pulumi.Input<string>;
    /**
     * Show the volume NBD export URI
     */
    export_uri?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    /**
     * The volume modification date (RFC 3339 format)
     */
    modification_date?: pulumi.Input<string>;
    /**
     * The volume name
     */
    name: pulumi.Input<string>;
    /**
     * The volume organization ID
     */
    organization?: pulumi.Input<string>;
    /**
     * The volume project ID
     */
    project: pulumi.Input<string>;
    /**
     * The server attached to the volume
     */
    server?: pulumi.Input<inputs.images.ScalewayInstanceV1VolumeServerPropertiesArgs>;
    /**
     * The volume disk size (in bytes)
     */
    size?: pulumi.Input<number>;
    state?: pulumi.Input<enums.images.ScalewayInstanceV1VolumeState>;
    /**
     * The volume tags
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    volume_type?: pulumi.Input<enums.images.ScalewayInstanceV1VolumeVolumeType>;
    /**
     * The zone in which is the volume
     */
    zone?: pulumi.Input<string>;
}
/**
 * scalewayInstanceV1VolumeArgsProvideDefaults sets the appropriate defaults for ScalewayInstanceV1VolumeArgs
 */
export function scalewayInstanceV1VolumeArgsProvideDefaults(val: ScalewayInstanceV1VolumeArgs): ScalewayInstanceV1VolumeArgs {
    return {
        ...val,
        state: (val.state) ?? "available",
        volume_type: (val.volume_type) ?? "l_ssd",
    };
}

/**
 * The server attached to the volume
 */
export interface ScalewayInstanceV1VolumeServerPropertiesArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
}

export interface ScalewayInstanceV1VolumeSummaryArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * (in bytes)
     */
    size?: pulumi.Input<number>;
    volume_type?: pulumi.Input<enums.images.ScalewayInstanceV1VolumeSummaryVolumeType>;
}
/**
 * scalewayInstanceV1VolumeSummaryArgsProvideDefaults sets the appropriate defaults for ScalewayInstanceV1VolumeSummaryArgs
 */
export function scalewayInstanceV1VolumeSummaryArgsProvideDefaults(val: ScalewayInstanceV1VolumeSummaryArgs): ScalewayInstanceV1VolumeSummaryArgs {
    return {
        ...val,
        volume_type: (val.volume_type) ?? "l_ssd",
    };
}