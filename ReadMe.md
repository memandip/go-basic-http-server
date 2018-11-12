# Simple Http Server created in go #

## There are four routes

* /view/:filename - This route is used to view the content of the file. If file does not exist it creates new file.

* /edit/:filename - This routes is used to edit the content of the file. If file does not exist it creates new file.

* /save/:filename - This route is used to save the edited content.

* / - This is the index route. It just shows the list of available files.
