// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Snapshots
{
    public static class GetSnapshot
    {
        public static Task<Outputs.ScalewayInstanceV1GetSnapshotResponse> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Outputs.ScalewayInstanceV1GetSnapshotResponse>("scaleway-instances:snapshots:getSnapshot", args ?? new GetSnapshotArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1GetSnapshotResponse> Invoke(GetSnapshotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1GetSnapshotResponse>("scaleway-instances:snapshots:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());

        public static Output<Outputs.ScalewayInstanceV1GetSnapshotResponse> Invoke(GetSnapshotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<Outputs.ScalewayInstanceV1GetSnapshotResponse>("scaleway-instances:snapshots:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the snapshot you want to get
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
        public static new GetSnapshotArgs Empty => new GetSnapshotArgs();
    }

    public sealed class GetSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of the snapshot you want to get
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetSnapshotInvokeArgs()
        {
        }
        public static new GetSnapshotInvokeArgs Empty => new GetSnapshotInvokeArgs();
    }
}
