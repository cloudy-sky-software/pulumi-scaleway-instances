// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Servers
{
    public static class ListServersTypes
    {
        public static Task<ListServersTypesResult> InvokeAsync(ListServersTypesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListServersTypesResult>("scaleway-instances:servers:listServersTypes", args ?? new ListServersTypesArgs(), options.WithDefaults());

        public static Output<ListServersTypesResult> Invoke(ListServersTypesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListServersTypesResult>("scaleway-instances:servers:listServersTypes", args ?? new ListServersTypesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListServersTypesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public ListServersTypesArgs()
        {
        }
        public static new ListServersTypesArgs Empty => new ListServersTypesArgs();
    }

    public sealed class ListServersTypesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The zone you want to target
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public ListServersTypesInvokeArgs()
        {
        }
        public static new ListServersTypesInvokeArgs Empty => new ListServersTypesInvokeArgs();
    }


    [OutputType]
    public sealed class ListServersTypesResult
    {
        public readonly Outputs.ScalewayInstanceV1ListServersTypesResponse Items;

        [OutputConstructor]
        private ListServersTypesResult(Outputs.ScalewayInstanceV1ListServersTypesResponse items)
        {
            Items = items;
        }
    }
}