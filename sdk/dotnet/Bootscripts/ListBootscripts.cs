// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Bootscripts
{
    public static class ListBootscripts
    {
        public static Task<ListBootscriptsResult> InvokeAsync(ListBootscriptsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListBootscriptsResult>("scaleway-instances:bootscripts:listBootscripts", args ?? new ListBootscriptsArgs(), options.WithDefaults());

        public static Output<ListBootscriptsResult> Invoke(ListBootscriptsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListBootscriptsResult>("scaleway-instances:bootscripts:listBootscripts", args ?? new ListBootscriptsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListBootscriptsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public ListBootscriptsArgs()
        {
        }
        public static new ListBootscriptsArgs Empty => new ListBootscriptsArgs();
    }

    public sealed class ListBootscriptsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ListBootscriptsInvokeArgs()
        {
        }
        public static new ListBootscriptsInvokeArgs Empty => new ListBootscriptsInvokeArgs();
    }


    [OutputType]
    public sealed class ListBootscriptsResult
    {
        public readonly Outputs.ScalewayInstanceV1ListBootscriptsResponse Items;

        [OutputConstructor]
        private ListBootscriptsResult(Outputs.ScalewayInstanceV1ListBootscriptsResponse items)
        {
            Items = items;
        }
    }
}
