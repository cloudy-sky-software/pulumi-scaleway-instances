// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.ScalewayInstances.Bootscripts.Outputs
{

    [OutputType]
    public sealed class ScalewayInstanceV1GetBootscriptResponse
    {
        public readonly Outputs.ScalewayInstanceV1Bootscript? Bootscript;

        [OutputConstructor]
        private ScalewayInstanceV1GetBootscriptResponse(Outputs.ScalewayInstanceV1Bootscript? bootscript)
        {
            Bootscript = bootscript;
        }
    }
}
