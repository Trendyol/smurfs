# Command Manifest (WIP)

You know host is orchestrator of CLI and It has to know all plugins as subcommands so you have to pass plugins as manifest object to Host Initializer. If you want to store manifest as JSON or YAML. You can design like below example

```yaml
kind: Manifest
metadata:
    creationTimestamp: ""
spec:
    plugins:
     - name: onboarding
       supportedOS:
        - darwin_amd64
        - windows
        - linux_x86
      provider:
        name: gitlab
        target: https://gitlab.trendyol.com/devx/producitivty/onboarding

```

When we design plugin/manifest structure. We should design extensible file. It must not get Trendyol specific things. Provider should be extensible. We'll support Gitlab releases but if it will be open source project. We have to support other providers like github, bitbucket or direct binary

We also can add specific data as metadata or other things. For example we would like to apply restrictions team/role based or out of Trendyol users can add specific data into yaml file

SupportedOS field has to be in plugin object. If any different OS wants to run unsupported command. We can throw exception earlier and also we can filter plugins for better experience. We can hide commands If It's not supported