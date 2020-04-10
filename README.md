[![Go Report Card](https://goreportcard.com/badge/github.com/bmv3cg/tf-crud)](https://goreportcard.com/report/github.com/bmv3cg/tf-crud)

Tfc-Workspace-manager
=====================

Terrafrom cloud workspace manager is a command line utility to create, delete, list and update workspaces in [Terraform cloud](http://app.terraform.io). Tfc workspace manager is written in Go lang and uses Hashicrop's go-tfe SDK for managing Terraform cloud workspaces.

Terraform cloud workspace manager provides capablity to list unused workspaces in an oragnisation and sort accordong to creation
time which can be used for lifecycle managemnt of workspace in an Terraform coud organisation.

Currenlty supported commands. 

- Create workspace
- Delete workspace
- List all workspaces
- List unused workspaces


Installation
------------

You can intall tfc workspace manager using homebrew by adding this [tap](bmv3cg/homebrew-tap). Once tap is configured you can use the brew install command to install the binary.

```bash
brew tap bmv3cg/homebrew-tap
brew install bmv3cg/tap/tfc 
```

Pre-requistes
-------------

Terraform cloud API token is requried for authentciating with Terraform cloud account. You can follow the steps in this [link](https://www.terraform.io/docs/cloud/users-teams-organizations/api-tokens.html#team-api-tokens) to create a token. You can export Terraform token as environment variable to manage workspaces using the following command.

```bash 
export TFE_TOKEN="Replace with Terraform cloud token"
```

Tfc-Workspace-manager Configuration
-----------------------------------

TfC workspace manager supports configuration either using configuration file or using command line arguments passed as flags.

1. TFC-ws-manager configuration using configuration file
--------------------------------------------------------

Terraform workspace manager can be configured using a config file. You can see a sample of config file [here](examples/configuration/tfe-ws-manager-config.yaml)


2. TFC-ws-manager configuration using command arguments
-------------------------------------------------------

Terraform cloud can be configured by passing following command line arguments.

| Argument      | Description                                                 |  
|---------------|-------------------------------------------------------------|
| organisation  | Organisation in which Terraform cloud workspaces are hosted |
| wsname        | Workspace name which is to be managed                       |
| verbosity     | Vebosity level for debugging                                |

Create workspace 
----------------

You can create a Terraform cloud workspace using the following commands. 

1. Using config file and create command

```bash
tfe-ws-manager create --config ~/.tfe-ws-manager-config.yaml
```

2. Using config file and command alias

```bash
tfe-ws-manager mk --config ~/.tfe-ws-manager-config.yaml
```

3. Using command line arguments 

```bash
tfe-ws-manager mk --organisation test --wsname ws 
```

Delete workspace
----------------

You can delete a Terraform cloud workspace using the following commands.  

1. Using config file and delete command 

```bash
tfe-ws-manager delete --config ~/.tfe-ws-manager-config.yaml
```

2. Using config file and command alias

```bash
tfe-ws-manager rm --config ~/.tfe-ws-manager-config.yaml
```

3. Using command line arguments 

```bash 
tfe-ws-manager rm --organisation test --wsname ws 
```

List workspace
---------------

You can list all Terraform cloud workspace in an organisation using the following command. This will list all workspaces along with workspace name and workspace ID in an organisation in a tabular format.

1. Using config file and list command 

```bash
tfe-ws-manager list --config ~/.tfe-ws-manager-config.yaml
```

2. Using config file and command alias

```bash
tfe-ws-manager ls --config ~/.tfe-ws-manager-config.yaml
```

3. Using command line arguments 

```bash
tfe-ws-manager ls  --organisation test
```

List unused workspace
---------------------

List unused workspace sorts all Terraform workspaces in an organisation, filters according to date of creation and show only unused workspaces in organsition. By default all unused workspaces will be listed by using command emptyws. Sorting can be enabled using the delta flag to sort the workspaces accoding to creation time. 

1. Using config file and sort command 

```bash
tfe-ws-manager emptyws --config ~/.tfe-ws-manager-config.yaml
```

2. Using config file and command alias

```bash
tfe-ws-manager ews --organisation test --delta 20
```

3. Using command line arguments 

```bash
tfe-ws-manager emptyws --organisation test --delta 20
```

Log level
---------

You can pass the log level variable for enabling detailed logs.

```bash
tfe-ws-manager ls --config ~/.tfe-ws-manager-config.yaml -v 2
```