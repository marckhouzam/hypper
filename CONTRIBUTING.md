# Contributing Guidelines

## Sign Your Work

The sign-off is a simple line at the end of the explanation for a commit. All commits needs to be
signed. Your signature certifies that you wrote the patch or otherwise have the right to contribute
the material. The rules are pretty simple, if you can certify the below (from
[developercertificate.org](https://developercertificate.org/)):

```txt
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.

Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

Then you just add a line to every git commit message:

```terminal
Signed-off-by: Joe Smith <joe.smith@example.com>
```

Use your real name (sorry, no pseudonyms or anonymous contributions.)

If you set your `user.name` and `user.email` git configs, you can sign your commit automatically
with `git commit -s`.

Note: If your git config information is set properly then viewing the `git log` information for your
 commit will look something like this:

```terminal
Author: Joe Smith <joe.smith@example.com>
Date:   Thu Feb 2 11:41:15 2018 -0800

    Update README

    Signed-off-by: Joe Smith <joe.smith@example.com>
```

Notice the `Author` and `Signed-off-by` lines match.

## Consistent output

We aim for a good user experience so we want to have a consistent cli-output. When creating output aim for the clean style we have been using for the existing cli-outputs and place emoticons at the beginning of the line. We are also using emoticons as *"eye-candy"*, please refer to the following table to style the cli-output of your contribution.

|event|suitable emoticons   |shortcode|example|
|:-----|:------------------|:-----|:-----|
|action failed|`❌`|`:x:`|`❌ Failed to install the chart.`|
|action in progress|`🧰 🛳️`|`:toolbox: :cruise_ship:`|`🛳️ Chart is installing`|
|action successful|`🥳👏✅`|`:partying_face: :clap: :white_check_mark`|`✅ Chart has been installed`|
|informational output|`ℹ️`|`:info:`|`ℹ️ The chart is already installed on the cluster, skipping.`|
|removing stuff|`🔥`|`:fire:`|`🔥 Uninstalling chart...`
|running a query|`🔍`|`:mag:`|`🔍 Searching for charts` |

**Note:** Emoticons must not be entered as UTF-8 glyphs into the code, we are using the kyokomis emoji library which can be found under [https://github.com/kyokomi/emoji](https://github.com/kyokomi/emoji). All glyphs must be encoded as ASCII-strings. All full list of available emojis can be found [here](https://github.com/kyokomi/emoji/blob/master/emoji_codemap.go).
