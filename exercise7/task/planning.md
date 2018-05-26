# task


One bucket, tasks
key - task name
value - task status (complete/incomplete)

Downside of this is that removing is done by task name. This is cumbersome. How to improve?
* Have an int
* part matching?


* add task to not done bucket
* do, add task to complete bucket. remove from notdone bucket
* list, iterate through not done bucket


## TODO

* convert to int based id
    * taskname, id bucket
    * id, taskStatus


* Add logic for event when task list is empty