# StatUtils

Collection of utilities for querying hardware statistics and information for Anvil! systems;

## perfmonitor

Perfomonitor is a server returning system performance data upon a request;

Open `PowerShell` as `Administrator` and run the following command:


    Set-ExecutionPolicy RemoteSigned

Start perfmonitor by running:


    perfmonitor.exe


**Building**
The .exe can be built for windows as:


    GOOS=windows GOARCH=amd64 go build
