# ZYT

Time tools. Built with Go.

## Taxonomy

ZYT's most basic record type is **events**, which describe an individual, indivisible entry of time. Events can be grouped using two different composite structures: **stories** and **tags**. Stories are intended to have a narrative, to which each event contributes. Tags are intended to cluster related events without expectation of a narrative, optionally providing functionality. Topologically, they are identical in that both tags and stories relate to events in a many-to-many style. As such, events can belong to multiple stories and have multiple tags; likewise, stories and tags should each contain multiple events.

### events
- start time
- end time
- duration
- name
- notes
- relationships
  - users: one
  - stories: many
  - tags: many

### stories
- start time
- end time
- duration
- name
- notes
- relationships
  - users: one
  - events: many

### tags
- name
- function
- relationships
  - users: one
  - events: many
- examples
  - name: Pro Bono
  - name: Company, Inc.
  - name: Standard Rate; function: `rate = 120 * (duration * 60 * 60)`
  - name: Discount; function: `rate *= 0.75`

## Commands
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
