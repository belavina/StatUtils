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

On both platforms, `value` for a specific hardware group indicates performance measure in the following format:

- `cpu`: % Processor Time is the percentage of elapsed time that the processor spends to execute a non-Idle thread (CPU Utilization, see [Microsoft Docs](https://social.technet.microsoft.com/wiki/contents/articles/12984.understanding-processor-processor-time-and-process-processor-time.aspx))
- `disk`: Percentage available of total storage space
- `memory`: Memory Available in Bytes

For windows, `instanceName` indicates cpu instance or disk name if applicable;

```javascript
{
   "status": "success",
   "message": "",
   "data": {
      "system": {
			"machine": "win.host",
			"platform": "windows",
			"softwareVersion": "0.0.0"
		},
      "cpu": {
         "stats": [
            {
               "counterType": "Timer100NsInverse",
               "defaultScale": "0",
               "instanceName": "_total",
               "multipleCount": "1",
               "path": "\\\\desktop-fo9p270\\processor information(_total)\\% processor time",
               "rawValue": "689984375000",
               "secondValue": "132054295798539933",
               "status": "0",
               "timeBase": "10000000",
               "timestamp": "6\/19\/2019 10:52:59 AM",
               "timestamp100NSec": "132054151798539933",
               "value": "9.41475786277403"
            }
            { ... }
         ],
         "date": "20190619145258" // UNIX timestamp in UTC
      },
      "disk": {
         "stats": [
            {
               "counterType": "RawFraction",
               "defaultScale": "0",
               "instanceName": "harddiskvolume1",
               "multipleCount": "1",
               "path": "\\\\desktop-fo9p270\\logicaldisk(harddiskvolume1)\\% free space",
               "rawValue": "175",
               "secondValue": "578",
               "status": "0",
               "timeBase": "10000000",
               "timestamp": "6\/19\/2019 10:53:01 AM",
               "timestamp100NSec": "132054151816460000",
               "value": "30.2768166089965"
            },
            { ... }
        ],
         "date": "20190619145256"
      },
      "memory": {
         "stats": [
            {
               "counterType": "NumberOfItems64",
               "defaultScale": "4294967290",
               "instanceName": "",
               "multipleCount": "1",
               "path": "\\\\desktop-fo9p270\\memory\\available bytes",
               "rawValue": "1515257856",
               "secondValue": "0",
               "status": "0",
               "timeBase": "10000000",
               "timestamp": "6\/19\/2019 10:52:58 AM",
               "timestamp100NSec": "132054151780150000",
               "value": "1515257856"
            }
         ],
         "date": "20190619145256"
      }
   }
}
```

Identification of hardware devices on linux differs from the windows output (for example, `cpuName` for cpu stats along with a combination of `filesystem` and `mounted` can be used). Note that disk query reports slightly modified `df` command output and you will need to filter out some filesystems. 


```javascript
{
   "status": "success",
   "message": "",
   "data": {
      "system": { ... },
      "cpu": {
         "stats": [
            {
               "cpuName": "cpu1",
               "value": "7.575758",
               "valueType": "% CPU Utilization"
            },
            { ... }
         ],
         "date": "20190619145258"
      },
      "disk": {
         "stats": [
            {
               "1K-blocks": "224021164",
               "available": "171205740",
               "filesystem": "/dev/mapper/fedora_localhost--live-home",
               "mounted": "/home",
               "use%": "20%",
               "used": "41366080",
               "value": "80.540186",
               "valueType": "Space Available"
            },
            { ... }
        ],
         "date": "20190619145256"
      },
      "memory": {
         "stats": [
            {
               "value": "22358220800",
               "valueType": "Memory Available (bytes)"
            }
         ],
         "date": "20190619145256"
      }
   }
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
      "platform":"linux", // or "windows"
      "softwareVersion": "0.2.0" // perfmonitor version
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
