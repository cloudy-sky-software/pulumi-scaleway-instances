// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Ips.Inputs
{

    public sealed class ScalewayInstanceV1ServerSummaryArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public ScalewayInstanceV1ServerSummaryArgs()
        {
        }
        public static new ScalewayInstanceV1ServerSummaryArgs Empty => new ScalewayInstanceV1ServerSummaryArgs();
    }
}
