# rappelz-golang-launcher

This is a very simple program to launch Rappelz client without using the official launcher.

---

You can download a precompiled executable from the [Releases]([https://github.com/Aodai/rappelz-golang-launcher/releases](https://github.com/Aodai/rappelz-golang-launcher/releases) page or build the executable yourself by cloning the repository and running the appropriate command to build the executable.

## Launcher Configuration

The launcher uses a JSON file for its configuration because it's pretty simple and fits for this purpose, the launcher requires the file **config.json** to be present in the same directory as the executable. An example configuration file is provided in the repository and in the release archive as well.

## Building

To build the executable:

```powershell
go build
```

To build the binary without debug symbols  which would result in a smaller executable size you'd add a couple of flags to the build command:

```powershell
go build -ldflags "-s -w"
```
