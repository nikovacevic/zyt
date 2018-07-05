# zyt

Time tools. Built with Go.

## Data model
```
events
.id           uuid
.user_id      uuid
.activity_id  uuid      null
.start        datetime
.end          datetime
.duration     int

tags
.id           uuid
.user_id      uuid
.code         string
.name         string

event_tag
.id           uuid
.event_id     uuid
.tag_id       uuid

activity
.id           uuid
.name         string

users
.id           uuid
.email        
.password     
```

## Command line
```
zyt
  
  // Analyze
  
  status
  • list of running timers
    [id] [duration] [start / -:-] [name] [tags|.|.]
  
  day
    [date? "YYYY-MM-DD"]
  yesterday
  • list of zyts recorded on given day (default today)
    [id] [duration] [start / end] [name] [tags|.|.]
  
  week
  last-week
  • list of zyts recorded during the given week
    [id] [duration] [start / end] [name] [tags|.|.]
  
  month
    [month? "MM (YYYY)?"]
  last-month
  • list of zyts recorded today
    [id] [duration] [start / end] [name] [tags|.|.]
    
  // Record
  
  start
    [name? "name"]
    [id? "-i 248"]
    [start? "-t (YYYY-MM-DD)? HH:II(:SS)?"]
  • start a zyt with a name or from a previous zyt id
  
  stop
    [name? "name"]
    [id? "-i 248"]
    [end? "-t (YYYY-MM-DD)? HH:II(:SS)?"]
  • stop a zyt; stops all zyts unless a name or id is given
  
  make
    [start: "-t (YYYY-MM-DD)? HH:II(:SS)?"]
    [end: "-t (YYYY-MM-DD)? HH:II(:SS)?"]
    [name? "name"]
    [id? "-i 248"]
  • make a zyt; start and end are required, but name and id are optional
  • giving a future date schedules zyt
```
