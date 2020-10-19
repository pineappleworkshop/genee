# Genee
Genee is a microservice project generator written in golang.
This powerful CLI tool is based on a project called `Cobra`
(https://github.com/spf13/cobra) which is a library providing
a simple interface to create modern CLI interfaces and helps
to generate a application scaffolding to rapidly develop a
Cobra-based application.

## Getting Started

First make sure you have `git` installed in your computer.
To do that, run the following command in your terminal.
```bash
git --version
```

If you get a message like `command not found`, follow the following
guide to install it in your computer
```
https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
```

Also, you need to have GO installed and in your PATH. You can run:

```bash
go --help
```

If is not installed:

https://golang.org/doc/install

And be sure to add it to your terminal `$PATH` variable in the terminal source file.

Clone the `genee` project into your projects/working directory
```bash
cd ~/Projects
git clone git@github.com:pineappleworkshop/genee.git
```

Go into the `genee` folder and install dependencies
```bash
cd genee
go install
```

Now run `genee` in your terminal to make sure it was installed properly. If the command is not recognized,
be sure that you have your `GOPATH` bin in your `$PATH`. Check your terminal source file, i.e `.bashrc`, `.bash_profile`, `.zshrc`, etc with the proper values:

```bash
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$(go env GOPATH)/bin
```

## Genee commands

For now Genee has only two commands available:
- help
- project

### help
This command provides information any command.

### project
This command is used to generates a microservice project from a template.

Project command also needs 3 flags to be use with it in order to make it
work, the expected flags are:

| Flag                     |                            Description                             |
| ------------------------ | :----------------------------------------------------------------: |
| -c, --config string      |      Sets the configuration file to be used during generation      |
| -d, --destination string | Sets the destination directory where the resulting project will be |
| -t, --template string    |  Sets the template directory in which to generate a project from   |
| -r, --repo string        |  Sets the repository link in which to generate a project from      |


You can also use "genee [command] --help" for more information about a command.


## How to generate a Go service

We can use `genee` to help us generate a new Go service but to
do so we are also going to need a Go project template.
The project called `pw-go-template` (https://github.com/pineappleworkshop/pw-go-template)
provides use a Go project scaffold with integration to MongoDB
and the initial configuration to deploy your code in Kubernetes.

Yoo don't need to clone `pw-go-template` project into your projects/working directory,
but if you want, here's the command to do it:
```bash
cd ~/Projects
git clone git@github.com:pineappleworkshop/pw-go-template.git
```

Then, copy in your projects/working the `genee-test.yml` file that
is into the `pw-go-template` project. If you didn't clone, pull a copy
directly from github.
```bash
cp pw-go-template/genee-test.yml .
mv genee-test.yml genee.yml
```

The `genee-test.yml` file has all the basic initial configuration file with all the
variables and values needed by `genee` to generate the service.
Edit the `genee.yml` and set the values for each of the variables
declared in there.

In the following line replace `SERVICE_NAME` with the value set
on the `service_name` variable in `genee.yml`.

If you cloned the repo locally:
```bash
genee project -t ./pw-go-template -d SERVICE_NAME -c ./genee.yml
```

If you didn't clone:
```bash
genee project -r git@github.com:pineappleworkshop/pw-go-template.git -d SERVICE_NAME -c ./genee.yml
```

Finally, go to the terminal, paste the line above (once you have replace
`SERVICE_NAME`), hit Enter and let `genee` do its magic.
