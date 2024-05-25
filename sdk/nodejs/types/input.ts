// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";

import * as utilities from "../utilities";

export namespace action {
    export interface ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs {
        volumeType?: pulumi.Input<enums.action.ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateVolumeType>;
    }
    /**
     * scalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgsProvideDefaults sets the appropriate defaults for ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs
     */
    export function scalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgsProvideDefaults(val: ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs): ScalewayInstanceV1ServerActionRequestVolumeBackupTemplateArgs {
        return {
            ...val,
            volumeType: (val.volumeType) ?? "l_ssd",
        };
    }

}

export namespace availability {
}

export namespace bootscripts {
}

export namespace dashboard {
}

export namespace images {
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
        creationDate?: pulumi.Input<string>;
        /**
         * Show the volume NBD export URI
         */
        exportUri?: pulumi.Input<string>;
        id?: pulumi.Input<string>;
        /**
         * The volume modification date (RFC 3339 format)
         */
        modificationDate?: pulumi.Input<string>;
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
        volumeType?: pulumi.Input<enums.images.ScalewayInstanceV1VolumeVolumeType>;
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
            volumeType: (val.volumeType) ?? "l_ssd",
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
        volumeType?: pulumi.Input<enums.images.ScalewayInstanceV1VolumeSummaryVolumeType>;
    }
    /**
     * scalewayInstanceV1VolumeSummaryArgsProvideDefaults sets the appropriate defaults for ScalewayInstanceV1VolumeSummaryArgs
     */
    export function scalewayInstanceV1VolumeSummaryArgsProvideDefaults(val: ScalewayInstanceV1VolumeSummaryArgs): ScalewayInstanceV1VolumeSummaryArgs {
        return {
            ...val,
            volumeType: (val.volumeType) ?? "l_ssd",
        };
    }
}

export namespace ips {
    export interface ScalewayInstanceV1ServerSummaryArgs {
        id?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
    }
}

export namespace placement_groups {
}

export namespace private_nics {
}

export namespace rules {
}

export namespace security_groups {
}

export namespace servers {
    export interface ScalewayInstanceV1VolumeServerTemplateArgs {
        /**
         * The ID of the snapshot on which this volume will be based
         */
        baseSnapshot?: pulumi.Input<string>;
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
        volumeType?: pulumi.Input<enums.servers.ScalewayInstanceV1VolumeServerTemplateVolumeType>;
    }
    /**
     * scalewayInstanceV1VolumeServerTemplateArgsProvideDefaults sets the appropriate defaults for ScalewayInstanceV1VolumeServerTemplateArgs
     */
    export function scalewayInstanceV1VolumeServerTemplateArgsProvideDefaults(val: ScalewayInstanceV1VolumeServerTemplateArgs): ScalewayInstanceV1VolumeServerTemplateArgs {
        return {
            ...val,
            boot: (val.boot) ?? false,
            volumeType: (val.volumeType) ?? "l_ssd",
        };
    }

}

export namespace snapshot_export {
}

export namespace snapshots {
}

export namespace user_data {
}

export namespace volumes {
}
