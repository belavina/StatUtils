**Windows Installation**

Steps to install `perfmonitor` release on a win system;


1) Copy the release folder `perfmonitor_win_amd64` to `C:\Program Files`

2) Open `Task Scheduler` as `Administrator`

3) Import pre-defined task by picking `Import Task` on action tab

    a) select `Perfmonitor` xml file in `C:\Program Files\perfmonitor_win_amd64`

    b) Click `Ok`, select `Task Scheduler Library` on the left, right click on `Perfmonitor` & select `Run` in the dropdown menu

4) Alternatively, you can create a new task by choosing `Create task` option

    a) On tab `General`, set name as `Perfmonitor` and tick `Run whether user is logged on or not`

    b) On tab `Triggers`, set `At startup` as task trigger

    c) On tab `Action`, point `Program/script` to perfmonitor location `C:\Program Files\perfmonitor_win_amd64\perfmonitor.exe`

    d) Click `Ok`, select `Task Scheduler Library` on the left, right click on `Perfmonitor` & select `Run` in the dropdown menu
