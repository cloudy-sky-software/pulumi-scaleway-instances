// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.PrivateNics.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1ListPrivateNICsResponse
    {
        public readonly ImmutableArray<Outputs.ScalewayInstanceV1PrivateNIC> Private_nics;

        [OutputConstructor]
        private ScalewayInstanceV1ListPrivateNICsResponse(ImmutableArray<Outputs.ScalewayInstanceV1PrivateNIC> private_nics)
        {
            Private_nics = private_nics;
        }
    }
}
