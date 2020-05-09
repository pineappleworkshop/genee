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
```
git --version
```

If you get a message like `command not found`, follow the following
guide to install it in your computer
```
https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
```

Clone the `genee` project into your projects/working directory
```
cd ~/Projects
git clone git@github.com:pineappleworkshop/genee.git
```

Go into the `genee` folder and install dependencies
```
cd genee
go install
```

Now run `genee` in your terminal to make sure it was installed properly.

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

| Flag                     | Description                                                        | 
| ------------------------ |:------------------------------------------------------------------:|
| -c, --config string      | Sets the configuration file to be used during generation           |
| -d, --destination string | Sets the destination directory where the resulting project will be |
| -t, --template string    | Sets the template directory in which to generate a project from    |


You can also use "genee [command] --help" for more information about a command.


## How to generate a Go service

We can use `genee` to help us generate a new Go service but to
do so we are also going to need a Go project template.
The project called `pw-go-template` (https://github.com/pineappleworkshop/pw-go-template)
provides use a Go project scaffold with integration to MongoDB
and the initial configuration to deploy your code in Kubernetes.

First thing that needs to be done is cloning `pw-go-template`
project into your projects/working directory
```
cd ~/Projects
git clone git@github.com:pineappleworkshop/pw-go-template.git
```

Then, copy in your projects/working the `genee.yml` file that
is into the `pw-go-template` project.
The `genee.yml` file is the configuration file with all the
variables and values needed by `genee` to generate the service.
```
cp pw-go-template/genee.yml .
```

Edit the `genee.yml` and set the values for each of the variables
declared in there.

In the following line replace `SERVICE_NAME` with the value set
on the `service_name` variable in `genee.yml`.
```
genee project -t ./pw-go-template -d SERVICE_NAME -c ./genee-test.yml
```

Finally, go to the terminal, paste the line above (once you have replace
`SERVICE_NAME`), hit Enter and let `genee` do its magic.
