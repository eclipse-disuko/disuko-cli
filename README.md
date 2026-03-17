# disuko-cli

## Introduction
The disuko-cli provides an easy way to access the public api of the eclipse-disuko Disclosure Portal.

With the disuko-cli users can:
- Create new project versions
- Access policy rules
- Upload SBOM files
- Set a reference to the corresponding source code
- Get general information about the project

## Table of Contents
* [Download and build](#download-and-build)
* [How-To](#how-to)
* [Guided example](#guided-example)
* [Contributing](#contributing)
* [Code of Conduct](#code-of-conduct)
* [License](#license)
* [Provider Information](#provider-information)

## Download and build
You can download the source code and build it yourself. 
```
git clone https://github.com/eclipse-disuko/disuko-cli.git
cd disuko-cli
go build -o disuko-cli
```

## How-To

The recommended way to use the disuko-cli as executable is with a config file, but you can also set environment variables or use flags with your commands instead. 
The config file needs to be in the same folder as the disuko-cli, needs to be named `config.yml` and must have the following structure:

``` yml
projecttoken: "project-token"
projectuuid: "project-uuid"
projectversion: "1.0"
host: "PUBLIC API"
```

The environment variables mus be named as follows:
``` 
INPUT_TOKEN
INPUT_PROJECT_UUID
INPUT_PROJECT_VERSION
INPUT_HOST
```

### How to use the disuko-cli

```
./disuko-cli

A client for the disclosure public api, to manage your projects.

Usage:
  ./disuko-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  project     Execute a project command
  sbom        Execute a project version sbom command
  sha256      Generates a sha256 hash
  version     Execute a project version command

Flags:
  -c, --configFile string   location of the config file (default "./config.yml")
  -h, --help                help for disuko-cli
  -t, --token string        disuko-cli project token
  -u, --uuid string         uuid of the project
  -v, --version string      version of the project (default "1.0")

Use "./disuko-cli [command] --help" for more information about a command.
```

### Sample commands
```
// Retrieving project detail
./disuko-cli project details -H HOST -u PROJECT_UUID -t TOKEN

// Upload sbom to a project version
./disuko-cli version sbomUpload FILENAME -H HOST -u PROJECT_UUID -t TOKEN -v VERSION
```

### Available Commands

```
project details       Returning the project details
project policyrules   Returning the project policy rules
project sbomCheck     On demand check for SBOM files
project schema        Returning the project schema
project status        Returning the project status

version ccs           CCS status
version ccsAdd        Add reference (url) to complete corresponding source code
version create        Create version
version details       Returning the project version details
version list          Returning the project version list
version sbomDetails   Details of SBOM
version sbomNotice    Get third party notice information for a SBOM as html / json / text 
version sbomStatus    Status information of SBOM
version sbomUpload    Uploads SBOM file to a project version
version sboms         List of all uploaded SBOMS

sbom tag              Add tag to a sbom

```




### Help on commands
```
version sbomUpload -h
```


## Guided example
The next few steps will guide you through a eclipse-disuko Disclosure Portal project with the disuko-cli.
In this guided example we will use a config file to access our project, but you can also use flags or environment variables. 

### Step 0 - Preparation
For the next few steps you need a project in the eclipse-disuko Disclosure Portal, and the unique identifier and token of your project.  
You can skip to the disuko-cli subsection, if you already have this data.

**eclipse-disuko Disclosure Portal**
*Only project owners in the eclipse-disuko Disclosure Portal have the permissions to do these steps.*
a) Create a new Project in the eclipse-disuko Disclosure Portal. Take note of the Unique identifier of the project, you will need it in a moment.  
b) Create a token. You will need this token to access your project with the disuko-cli.

**disuko-cli**  
a) Create a new folder wherever you want. I will call this folder `disuko-cli`, but you can name it differently.  
b) Download the source code for the disuko-cli client and build it.
```
go build -o disuko-cli
```
c) Move `disuko-cli` into your created folder (a)  
d) Create a new config file named `config.yml` with the following attributes and values in the same folder
``` yml
projecttoken: "project-token"
projectuuid: "project-uuid"
projectversion: "1.0"
host: "" // host needs to be changed to public api
```
If you have just created a new project, your project on the eclipse-disuko Disclosure Portal won't have a project version yet.
We will create a version `1.0` in the following steps.

### Step 1 - Your first command: Get the project details

```
./disuko-cli project details
```
```
{
    "name": "disuko-cli-example",
    "uuid": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeee",
    "created": "2022-03-16T14:54:13.306719888Z",
    "updated": "2022-03-16T15:13:45.461701011Z",
    "schema": "DefaultSchema",
    "description": "This is the example project for the cli client."
}
```

### Step 2 - Create a project version
Our project still does not have a project version. Let's create one.
```
./disuko-cli version create "1.0" "First iteration of the disuko-cli example"
```
```
{
    "success": true,
    "message": "version created",
    "key": "ab02a5b5-8923-41pb-9a9b-cd5836d66dbe"
}
```
Let's have a look at the project versions of our project. There should only be one ;)
```
./disuko-cli version list
```
```
[
    "1.0"
]
```

### Step 3 - Upload the SBOM of your project
Next we will upload the SBOM and look at the metadata of the SBOM.
The SBOM is always related to a specific project version. In our case it is `1.0` as described in our `config.yml`.

```
./disuko-cli version sbomUpload ./disuko-cli/sbom.json
```
```
{
    "docIsValid": true,
    "validationFailedMessage": "",
    "hash": "dummye85477db7580d3052403a8d1ebe4ad9a3b912ad18e3a3b4f0ccecf36c65",
    "fileUploaded": true,
    "id": "SPDXRef-DOCUMENT"
}
```
```
./disuko-cli version sbomDetails
```
```
{
    "name": "disuko-cli,
    "id": "SPDXRef-DOCUMENT",
    "version": "SPDX-2.2",
    "creators": "Tool: xxx",
    "created": "2022-03-16T14:54:13.306719888Z",
    "updated": "2022-03-16T15:13:45.461701011Z",
    "status": true
}
```

## Contributing

Contributions are welcome and appreciated.

This project follows the Eclipse Foundation development and contribution processes.
Before contributing, please make sure you are familiar with the following resources:

- Eclipse Foundation Contributing Guide  
  https://www.eclipse.org/contribute/

- Eclipse Contributor Agreement (ECA)  
  https://www.eclipse.org/legal/ECA.php

By submitting a pull request, you confirm that you have the right to contribute the code
and that you agree to the terms of the Eclipse Contributor Agreement.

No additional project specific contribution guidelines are required at this time.

### Security

This project provides a Gitleaks configuration file to help contributors
detect accidental secret commits. Usage is optional and can be integrated
locally or in CI environments.

## Code of Conduct

This project follows the Eclipse Foundation Code of Conduct to ensure a respectful,
inclusive, and harassment free environment for everyone involved.

All participants are expected to adhere to the rules defined here:  
https://www.eclipse.org/org/documents/Community_Code_of_Conduct.php

By participating in this project, you agree to uphold this Code of Conduct in all project related spaces.

## License

This project is licensed under the [Apache-2.0](LICENSE).

## Provider Information

Please visit <https://github.com/mercedes-benz/foss/blob/master/PROVIDER_INFORMATION.md> for information on the provider.

Notice: Before you use the program in productive use, please take all necessary precautions,
e.g. testing and verifying the program with regard to your specific use.
The program was tested solely for our own use cases, which might differ from yours.
