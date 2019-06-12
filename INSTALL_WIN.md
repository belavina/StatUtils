**Windows Installation**

Steps to install `perfmonitor` release on a win system;

1) Open `PowerShell` as `Administrator` and run the following command:

    `Set-ExecutionPolicy RemoteSigned`

2) Copy the release folder `perfmonitor_win_amd64` to `C:\Program Files`

3) Open "Task Scheduler"

    a) On tab "General", set name as "Perfmonitor" and tick "Run whether user is logged on or not"

    b) On tab "Triggers", set "At startup" as task trigger

    c) On tab "Action", point "Program/script" to perfmonitor location `C:\Program Files\perfmonitor_win_amd64\perfmonitor`

    d) Click "Ok", select "Task Scheduler Library" on the left, right click on "Perfmonitor" & select "Run" in the dropdown menu
