# StatUtils

Collection of utilities for querying hardware statistics and information for Anvil! systems;

## perfmonitor

Perfomonitor is a server returning system performance data upon a request;

### Running

Open `PowerShell` as `Administrator` and run the following command:


    Set-ExecutionPolicy RemoteSigned

Start perfmonitor by running:


    perfmonitor.exe

### HTTP Interface

Current performance stats can be accessed at HTTP port `:9159`, data is returned in the following JSON format:

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

### Compilation

Building step requires golang binary release present on your system, navigate to the [official downlad page](https://golang.org/dl/) to grab the latest 
version.

The .exe can be built for windows as:


    GOOS=windows GOARCH=amd64 go build -o perfmonitor.exe
