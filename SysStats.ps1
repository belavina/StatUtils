<#
This script collects some performance counters
Usage:
    ./SysStats.ps1 # to stdout
    ./SysStats.ps1 -export -outpath ./perf_log.csv
#>

[CmdLetBinding(DefaultParameterSetName = "NormalRun")]
Param(
    [Parameter(Mandatory = $False, Position = 1, ParameterSetName = "NormalRun")] $Server = "",
    [switch]$export = $false,
    [string]$outpath = ""
)

if ($Server -ne ""){$Server = ("\\") + $Server} else {$Server = ""}

$Counters = @(
"$Server\LogicalDisk(*)\% Free Space",
"$Server\Processor Information(*)\% Processor Time",
"$Server\Memory\Available Bytes"
)

$results = @()

Get-Counter -Counter $Counters | ForEach {
    $_.CounterSamples | ForEach {
        $details = @{
                Date  = get-date
                Path  = $_.Path
                Value = $_.CookedValue
        }
        $results += New-Object PSObject -Property $details 
    }
}

if ($export -And $outpath -ne "") {
    $results | export-csv -Path $outpath -NoTypeInformation
}

Write-Host ($results | convertto-csv -NoTypeInformation)