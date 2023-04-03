# Pulumi Native Provider for Scaleway Instances (Preview)

[Scaleway Instances](https://www.scaleway.com/en/virtual-instances/). Build, deploy and scale applications on Europe's most complete cloud ecosystem.

> This provider was generated using [`pulschema`](https://github.com/cloudy-sky-software/pulschema) and [`pulumi-provider-framework`](https://github.com/cloudy-sky-software/pulumi-provider-framework).

## Package SDKs

- Node.js: https://www.npmjs.com/package/@cloudyskysoftware/pulumi-scaleway-instances
- Python: https://pypi.org/project/pulumi_scaleway_instances/
- .NET: https://www.nuget.org/packages/Pulumi.ScalewayInstances
- Go: `import github.com/cloudy-sky-software/pulumi-scaleway-instances/sdk/go/sclwyinst`

## Using The Provider

You'll need an API key. Follow Scaleway's [docs](https://developers.scaleway.com/en/quickstart/#authentication) to create one or head straight to the [credential section](https://console.scaleway.com/project/credentials) in their web console.
Then set the API key as a secret with `pulumi config set --secret scaleway-instances:apiKey`.

## Releasing A New Version

:info: Switch to the `main` branch first and get the latest `git pull origin main && git fetch`. Check what the last release tag was.

1. Regular releases should just increment the patch version unless a minor or a major (breaking changes) version bump is warranted.
1. Update the `CHANGELOG.md` with notes about what will be included in this release.
1. Commit the changelog with `git commit -am "vX.Y.Z"` or something similar and push `git push origin main`.
1. Tag the commit with the release version by running

   ```bash
   git tag vX.Y.Z
   git tag sdk/vX.Y.Z
   ```

1. Push the tags.

   ```bash
   git push --tags
   ```
