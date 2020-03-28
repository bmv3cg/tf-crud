Tfc-Workspace-manager
=====================

TFC-Workspace-manager is a command line utility to create, delete, list and update workspaces in [Terraform cloud](http://app.terraform.io). TFC-Workspace-manager is written in Go lang. This uses cobra for CLI commands and viper for configuration  file management.

Pre-requistes
-------------

Terraform cloud token is used for authentciating with Terraform cloud account. You can export Terraform token as environment variable to manage workspaces using the following command.

```bash 
export TFE_TOKEN="Replace with Terraform cloud token"
```

Tfc-Workspace-manager Configuration
-----------------------------------

TfC workspace manger supports configuration either using configuration file or using command line arguments passed through.

1. TFC-ws-manager configuration uing configuration file
-------------------------------------------------------

Terraform workspace manager can be configured using and config file. You can see a sample of config file [here](examples/configuration/tfe-ws-manager-config.yaml)


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
tfe-ws-manager create   --config ~/.tf-crud.yaml
```

2. Using config file and command alias

```bash
tfe-ws-manager mk --config ~/.tf-crud.yaml
```

3. Using command line arguments 

```bash
tfe-ws-manager mk   --organisation aexp --wsname ws 
```

Delete workspace
----------------

You can delete a Terraform cloud workspace using the following commands. 


1. Using config file and delete command 

```bash
tfe-ws-manager delete   --config ~/.tf-crud.yaml
```

2. Using config file and command alias

```bash
tfe-ws-manager rm   --config ~/.tf-crud.yaml
```

3. Using command line arguments 

```bash 
tfe-ws-manager rm   --organisation aexp --wsname ws 
```

List workspace
---------------

You can list all Terraform cloud workspace in an organisation using the following command. This will list all workspaces along with workspace name and workspace ID in an organisation in a tabular format.

1. Using config file and list command 

```bash
tfe-ws-manager list  --organisation aexp
```

2. Using config file and command alias

```bash
tfe-ws-manager ls --config ~/.tf-crud.yaml
```

3. Using command line arguments 

```bash
tfe-ws-manager ls  --organisation aexp
```

Log level
---------

You can pass the log level variable for enabling detailed logs.

```bash
tfe-ws-manager ls --config ~/.tf-crud.yaml -v 2
```