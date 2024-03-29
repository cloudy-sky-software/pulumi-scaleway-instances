// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.ScalewayInstances.Availability
{
    [EnumType]
    public readonly struct ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability : IEquatable<ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability>
    {
        private readonly string _value;

        private ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability Available { get; } = new ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability("available");
        public static ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability Scarce { get; } = new ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability("scarce");
        public static ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability Shortage { get; } = new ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability("shortage");

        public static bool operator ==(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability left, ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability left, ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability other && Equals(other);
        public bool Equals(ScalewayInstanceV1GetServerTypesAvailabilityResponseAvailabilityAvailability other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
