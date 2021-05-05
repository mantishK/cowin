# CoWin CLI

Parses CoWin website to get available slots. Currently works only for Hubli, Blr urban, BBMP and Pune.

## Installation
`go get github.com/mantishK/cowin`    
`go install gitgub.com/mantishK/cowin`    

#### Optional
If you want export data to emails, you will have to get sendgrid API keys. Note that you need a from email ID that is verified by sendgrid.    
Export API Key    
`export SENDGRID_API_KEY='S5frn-7VMSpS'`    
Export verified "from" email id     
`export SENDGRID_FROM_EMAIL='aaa@gmail.com'`    

## Flags
- a - min age limit    
- d - districts (hbl, blr, pn)    
- o - operations (export, display)    
- usr - user list to export data to. Should be in the format \<email\>:\<district\>:\<min-age\>    
e.g `abc@gmail.com:blr:18`     
This flag is a list, you can add multiple users. Check Usage for details    

## Usage
`cowin --a=18 --d=hbl`    
`cowin --o=export --usr=abc@gmail.com:blr:18 --usr=xyz@gmail.com:hbl:45`    
Can use it in a cron to get alerts
