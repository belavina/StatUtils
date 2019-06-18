# StatUtils [![CircleCI](https://circleci.com/gh/Seneca-CDOT/StatUtils/tree/master.svg?style=svg)](https://circleci.com/gh/Seneca-CDOT/StatUtils/tree/master)

Collection of utilities for querying hardware statistics and information for Anvil! systems;

## perfmonitor

Perfomonitor is a server returning system performance data upon a request;


### Release

Navigate to [releases page](https://github.com/Seneca-CDOT/StatUtils/releases) to grab the latest version of perfmonitor. Platform-specific bundles contain ready-to-use executables as well as instructions on how to install it on the target system.


**HTTP Interface**

Current performance stats can be accessed at HTTP port `:9159` & route `/sysstats`:

```bash
$ curl -i "http://{{vm_host}}:9159/sysstats"
```

data is returned in the following JSON format (this is windows example, Keys won't match the linux output):

```javascript
{
   "status": "success",
   "message": "",
   "data": [
      {  
         "date":"6/9/2019 10:16:39 PM",
         "key":"\\\\pc-name\\logicaldisk(q:)\\% free space",
         "value":"60.1844908902267"
      },
      {  
         "date":"6/9/2019 10:16:39 PM",
         "key":"\\\\pc-name\\processor information(0,0)\\% processor time",
         "value":"17.484500998004"
      },
      {  
         "date":"6/9/2019 10:16:39 PM",
         "key":"\\\\pc-name\\memory\\available bytes",
         "value":"3154997248"
      }
      ...
   ]
}
```

You can find os of the vm `perfmonitor` is running on by querying:

```bash
$ curl -i "http://{{vm_host}}:9159/platform"
```

```javascript
{
   "status": "success",
   "message": "",
   "data": {
      "machine":"my.hostname",
      "platform":"linux" // or "windows"
   }
}
```

### Development

Building & development requires golang binary release present on your system, navigate to the [official "Getting Started" page](https://golang.org/doc/install) to grab the latest version.

Note that it can be installed on fedora systems as:

```bash
$ dnf install golang -y
```

**Package Installation**

Grab the latest `perfmonitor` by running:

```bash
# installs in ~/go/src/github.com/Seneca-CDOT/StatUtils
go get github.com/Seneca-CDOT/StatUtils 
```

**Compilation**

The .exe can be built for windows (on linux) as:

`GOOS=windows GOARCH=amd64 go build -o perfmonitor.exe`

Or it can be compiled on windows as:

`$env:GOOS="windows"; $env:GOARCH="amd64";C:\Go\bin\go build -o perfmonitor.exe`

Linux executable can be created with the following command:

`GOOS=linux GOARCH=amd64 go build -o perfmonitor`

**Running Windows**

Open `PowerShell` as `Administrator` and run the following command:

`Set-ExecutionPolicy RemoteSigned`

Start perfmonitor by running:

`./perfmonitor.exe`

**Running Linux**

Start perfmonitor by running:

`./perfmonitor`
