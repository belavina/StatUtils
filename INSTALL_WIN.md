**Windows Installation**

Steps to install `perfmonitor` release on a win system;


1) Extract the release bundle as `C:\Program Files\perfmonitor_win_amd64`

2) Open `Task Scheduler` as `Administrator`

3) Import pre-defined task by picking `Import Task` on action tab
 
    a) select `Perfmonitor` xml file in `C:\Program Files\perfmonitor_win_amd64`

    b) Click `Ok`, select `Task Scheduler Library` on the left, right click on `Perfmonitor` & select `Enable` and then `Run` in the dropdown menu

4) Alternatively, you can create a new task by choosing `Create task` option

    a) On tab `General`, set name as `Perfmonitor`, tick `Run whether user is logged on or not` and `Run with highest privileges`

    b) On tab `Triggers`, set `At startup` as task trigger

    c) On tab `Action`, point `Program/script` to `cmd`, set `Add arguments` as `/c perfmonitor.exe > log.txt` and `Start in` as `C:\Program Files\perfmonitor_win_amd64`

    d) Click `Ok`, select `Task Scheduler Library` on the left, right click on `Perfmonitor` & select `Run` in the dropdown menu
