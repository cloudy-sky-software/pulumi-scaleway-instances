// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.ScalewayInstances.Images
{
    [EnumType]
    public readonly struct Arch : IEquatable<Arch>
    {
        private readonly string _value;

        private Arch(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Arch X8664 { get; } = new Arch("x86_64");
        public static Arch Arm { get; } = new Arch("arm");

        public static bool operator ==(Arch left, Arch right) => left.Equals(right);
        public static bool operator !=(Arch left, Arch right) => !left.Equals(right);

        public static explicit operator string(Arch value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Arch other && Equals(other);
        public bool Equals(Arch other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The bootscript arch
    /// </summary>
    [EnumType]
    public readonly struct ScalewayInstanceV1BootscriptArch : IEquatable<ScalewayInstanceV1BootscriptArch>
    {
        private readonly string _value;

        private ScalewayInstanceV1BootscriptArch(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1BootscriptArch X8664 { get; } = new ScalewayInstanceV1BootscriptArch("x86_64");
        public static ScalewayInstanceV1BootscriptArch Arm { get; } = new ScalewayInstanceV1BootscriptArch("arm");

        public static bool operator ==(ScalewayInstanceV1BootscriptArch left, ScalewayInstanceV1BootscriptArch right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1BootscriptArch left, ScalewayInstanceV1BootscriptArch right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1BootscriptArch value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1BootscriptArch other && Equals(other);
        public bool Equals(ScalewayInstanceV1BootscriptArch other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScalewayInstanceV1ImageArch : IEquatable<ScalewayInstanceV1ImageArch>
    {
        private readonly string _value;

        private ScalewayInstanceV1ImageArch(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1ImageArch X8664 { get; } = new ScalewayInstanceV1ImageArch("x86_64");
        public static ScalewayInstanceV1ImageArch Arm { get; } = new ScalewayInstanceV1ImageArch("arm");

        public static bool operator ==(ScalewayInstanceV1ImageArch left, ScalewayInstanceV1ImageArch right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1ImageArch left, ScalewayInstanceV1ImageArch right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1ImageArch value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1ImageArch other && Equals(other);
        public bool Equals(ScalewayInstanceV1ImageArch other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScalewayInstanceV1ImageState : IEquatable<ScalewayInstanceV1ImageState>
    {
        private readonly string _value;

        private ScalewayInstanceV1ImageState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1ImageState Available { get; } = new ScalewayInstanceV1ImageState("available");
        public static ScalewayInstanceV1ImageState Creating { get; } = new ScalewayInstanceV1ImageState("creating");
        public static ScalewayInstanceV1ImageState Error { get; } = new ScalewayInstanceV1ImageState("error");

        public static bool operator ==(ScalewayInstanceV1ImageState left, ScalewayInstanceV1ImageState right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1ImageState left, ScalewayInstanceV1ImageState right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1ImageState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1ImageState other && Equals(other);
        public bool Equals(ScalewayInstanceV1ImageState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScalewayInstanceV1VolumeState : IEquatable<ScalewayInstanceV1VolumeState>
    {
        private readonly string _value;

        private ScalewayInstanceV1VolumeState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1VolumeState Available { get; } = new ScalewayInstanceV1VolumeState("available");
        public static ScalewayInstanceV1VolumeState Snapshotting { get; } = new ScalewayInstanceV1VolumeState("snapshotting");
        public static ScalewayInstanceV1VolumeState Error { get; } = new ScalewayInstanceV1VolumeState("error");
        public static ScalewayInstanceV1VolumeState Fetching { get; } = new ScalewayInstanceV1VolumeState("fetching");
        public static ScalewayInstanceV1VolumeState Resizing { get; } = new ScalewayInstanceV1VolumeState("resizing");
        public static ScalewayInstanceV1VolumeState Saving { get; } = new ScalewayInstanceV1VolumeState("saving");
        public static ScalewayInstanceV1VolumeState Hotsyncing { get; } = new ScalewayInstanceV1VolumeState("hotsyncing");

        public static bool operator ==(ScalewayInstanceV1VolumeState left, ScalewayInstanceV1VolumeState right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1VolumeState left, ScalewayInstanceV1VolumeState right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1VolumeState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1VolumeState other && Equals(other);
        public bool Equals(ScalewayInstanceV1VolumeState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScalewayInstanceV1VolumeSummaryVolumeType : IEquatable<ScalewayInstanceV1VolumeSummaryVolumeType>
    {
        private readonly string _value;

        private ScalewayInstanceV1VolumeSummaryVolumeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1VolumeSummaryVolumeType LSsd { get; } = new ScalewayInstanceV1VolumeSummaryVolumeType("l_ssd");
        public static ScalewayInstanceV1VolumeSummaryVolumeType BSsd { get; } = new ScalewayInstanceV1VolumeSummaryVolumeType("b_ssd");
        public static ScalewayInstanceV1VolumeSummaryVolumeType Unified { get; } = new ScalewayInstanceV1VolumeSummaryVolumeType("unified");

        public static bool operator ==(ScalewayInstanceV1VolumeSummaryVolumeType left, ScalewayInstanceV1VolumeSummaryVolumeType right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1VolumeSummaryVolumeType left, ScalewayInstanceV1VolumeSummaryVolumeType right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1VolumeSummaryVolumeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1VolumeSummaryVolumeType other && Equals(other);
        public bool Equals(ScalewayInstanceV1VolumeSummaryVolumeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct ScalewayInstanceV1VolumeVolumeType : IEquatable<ScalewayInstanceV1VolumeVolumeType>
    {
        private readonly string _value;

        private ScalewayInstanceV1VolumeVolumeType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ScalewayInstanceV1VolumeVolumeType LSsd { get; } = new ScalewayInstanceV1VolumeVolumeType("l_ssd");
        public static ScalewayInstanceV1VolumeVolumeType BSsd { get; } = new ScalewayInstanceV1VolumeVolumeType("b_ssd");
        public static ScalewayInstanceV1VolumeVolumeType Unified { get; } = new ScalewayInstanceV1VolumeVolumeType("unified");

        public static bool operator ==(ScalewayInstanceV1VolumeVolumeType left, ScalewayInstanceV1VolumeVolumeType right) => left.Equals(right);
        public static bool operator !=(ScalewayInstanceV1VolumeVolumeType left, ScalewayInstanceV1VolumeVolumeType right) => !left.Equals(right);

        public static explicit operator string(ScalewayInstanceV1VolumeVolumeType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ScalewayInstanceV1VolumeVolumeType other && Equals(other);
        public bool Equals(ScalewayInstanceV1VolumeVolumeType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct State : IEquatable<State>
    {
        private readonly string _value;

        private State(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static State Available { get; } = new State("available");
        public static State Creating { get; } = new State("creating");
        public static State Error { get; } = new State("error");

        public static bool operator ==(State left, State right) => left.Equals(right);
        public static bool operator !=(State left, State right) => !left.Equals(right);

        public static explicit operator string(State value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is State other && Equals(other);
        public bool Equals(State other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
