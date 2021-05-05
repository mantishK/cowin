# CoWin CLI

Parses CoWin website to get available slots. Currently works only for Hubli, Blr urban, BBMP and Pune.

## Installation
`go get github.com/mantishK/cowin`    
`go install gitgub.com/mantishK/cowin`    

## Flags
a - min age limit    
d - districts (hbl, blr, pn)    
o - operations (export, display)    
usr - user list to export data to. Should be in the format <email>:<district>:<min-age> e.g abc@gmail.com:blr:18. Note that this is a list. You can add multiple users    

## Usage
`cowin --a=18 --d=hbl`    
`cowin --o=export --usr=abc@gmail.com:blr:18 --usr=xyz@gmail.com:hbl:45`
Can use it in a cron to get alerts
