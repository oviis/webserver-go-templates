## This part is the installation of the app on minishift, to create the possibility to deploy it in a Openshift CI/CD environment


# Requirements
* minishift installed, here an install Link  [minishift installation](https://docs.okd.io/latest/minishift/getting-started/installing.html)
* openshift cli, here an [install guide](https://docs.openshift.com/enterprise/3.1/cli_reference/get_started_cli.html#installing-the-cli);
or over brew 
```bash
brew cask install minishift
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
#this will deploy every object and parameter, change the parameter to fit your environment
$ oc new-app templates/webserver-go.yaml -p SOURCE_REPOSITORY_URL=https://github.com/<yourusername>/webserver-go-templates -p APPLICATION_DOMAIN=echo-example.<your-private-minishift-ip>.nip.io
```

3. once the app is running, you need to wait 3-5min to be build and deployed, you can check the svc and the route with following commands
```bash
#get the IP of the services
$ oc get svc
NAME                     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)     AGE
echo-example             ClusterIP   172.30.61.88     <none>        1323/TCP    9m

#get the dns hostname and access the application
#if your router is configured correctly you can access over the named route
$ oc get routes | awk '{print $1,$2}'
NAME            HOST/PORT
echo-example    echo-example.192.168.99.100.nip.io

#get the logs from the build
oc logs -f bc/echo-example
````

4. Housekeeping, deleting the whole objects
```bash
oc delete all --selector app=echo-example
```
