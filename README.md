# StatUtils [![CircleCI](https://circleci.com/gh/Seneca-CDOT/StatUtils/tree/master.svg?style=svg)](https://circleci.com/gh/Seneca-CDOT/StatUtils/tree/master)

Collection of utilities for querying hardware statistics and information for Anvil! systems;

## perfmonitor

Perfomonitor is a server returning system performance data upon a request;

### Running Windows

Open `PowerShell` as `Administrator` and run the following command:

`Set-ExecutionPolicy RemoteSigned`

Start perfmonitor by running:

`./perfmonitor.exe`

### Running Linux

Start perfmonitor by running:

`./perfmonitor`


### HTTP Interface

Current performance stats can be accessed at HTTP port `:9159` & route `/sysstats`:

`$ curl -i "http://{{vm_host}}:9159/sysstats"`

data is returned in the following JSON format (this is windows example, Keys won't match the linux output):

```
[  
   {  
      "Date":"6/9/2019 10:16:39 PM",
      "Key":"\\\\pc-name\\logicaldisk(q:)\\% free space",
      "Value":"60.1844908902267"
   },
   {  
      "Date":"6/9/2019 10:16:39 PM",
      "Key":"\\\\pc-name\\processor information(0,0)\\% processor time",
      "Value":"17.484500998004"
   },
   {  
      "Date":"6/9/2019 10:16:39 PM",
      "Key":"\\\\pc-name\\memory\\available bytes",
      "Value":"3154997248"
   }
   ...
]
```

You can find os of the vm `perfmonitor` is running on by querying:

`$ curl -i "http://{{vm_host}}:9159/platform"`

```
{
   "machine":"my.hostname",
   "platform":"linux" // or "windows"
}
```

### Compilation

Building step requires golang binary release present on your system, navigate to the [official downlad page](https://golang.org/dl/) to grab the latest 
version.

The .exe can be built for windows (on linux) as:

`GOOS=windows GOARCH=amd64 go build -o perfmonitor.exe`

Or it can be compiled on windows as:

`$env:GOOS="windows; $env:GOARCH="amd64";C:\Go\bin\go build -o perfmonitor.exe`

Linux executable can be created with the following command:

`GOOS=linux GOARCH=amd64 go build -o perfmonitor`