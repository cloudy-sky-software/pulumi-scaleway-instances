// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Volumes
{
    public static class ListVolumes
    {
        public static Task<Outputs.ScalewayInstanceV1ListVolumesResponse> InvokeAsync(ListVolumesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ScalewayInstanceV1ListVolumesResponse>("scaleway-instances:volumes:listVolumes", args ?? new ListVolumesArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1ListVolumesResponse> Invoke(ListVolumesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1ListVolumesResponse>("scaleway-instances:volumes:listVolumes", args ?? new ListVolumesInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1ListVolumesResponse> Invoke(ListVolumesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1ListVolumesResponse>("scaleway-instances:volumes:listVolumes", args ?? new ListVolumesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListVolumesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public ListVolumesArgs()
        {
        }
        public static new ListVolumesArgs Empty => new ListVolumesArgs();
    }

    public sealed class ListVolumesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ListVolumesInvokeArgs()
        {
        }
        public static new ListVolumesInvokeArgs Empty => new ListVolumesInvokeArgs();
    }
}
