# Saarthi

Saarthi is your friendly neighbourhood task scheduler. It is responsible for scheduling external tasks provided as a HTTP endpoint. The scheduler guarantees that the task will run at the time requested.

Most applications require scheduling tasks from time to time, be it some kind of report generation or scheduling mails and notifications for your customers. But since scheduling is not a core business requirement for most applications, it makes little sense to invest the all important resources to ensure uptime and availability of the scheduler.


Saarthi allows you to schedule tasks without caring about what task you want to run. The tasks are accepted as HTTP enpoints that Saarthi fires at the mentioned schedule. The scheduling is possible as a one off firing or as a cron-like schedule.
