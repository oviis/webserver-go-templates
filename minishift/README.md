## This part is the installation of the app on minishift, to create the possibility to deploy it in a Openshift CI/CD environment


# Requirements
* minishift installed, here an install Link  [minishift installation](https://docs.okd.io/latest/minishift/getting-started/installing.html)
* openshift cli, here an [install guide](https://docs.openshift.com/enterprise/3.1/cli_reference/get_started_cli.html#installing-the-cli);
or over brew 
```bash
brew install openshift-cli
```

# Deployment of following objects in minishift
* Service
* Route
* ImageStream
* BuildConfig
* DeploymentConfig

1. first of all fork a copy of webserver-go-templates

2. you can deploy all of this with following template and command
```bash
oc new-app templates/webserver-go.yaml -p SOURCE_REPOSITORY_URL=https://github.com/<yourusername>/webserver-go-templates -p APPLICATION_DOMAIN=echo-example.<your-private-minishift-ip>.nip.io
```
