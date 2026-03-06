# Time Tracker Go
This is a simple app to manage your time without leaving the console.
You can have multiple activities and track when you start and end doing them.

## Commands
### add 
Add activities

### remove
Remove activities

### start
Start doing activities

### end
End doing activities

### list
Get a list of all your activities

### report
Get a report of your activities. You can either pass a natural language time expression as arguments (examples below), or use the `--activity` flag to filter by activity.

- Examples: `report last week`, `report 2 days ago`, `report monday`, `report "yesterday at 5pm"`

- activity(a): Use this if you want to only get the report for a single activity, default `all`.

## Development
The project is built with go, so you can just clone the repo, run
```bash
go mod tidy
```
and then
```
go run .
```

## Others
- If you want to find the sqlite database where your activities are stored for whatever reason, it is in your [config folder](https://pkg.go.dev/os#UserConfigDir) according to your OS.
