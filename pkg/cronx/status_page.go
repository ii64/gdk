package cronx

import (
	"html/template"
	"sync"
)

const statusPageTemplate = `
<!-- This page is only being used for development, the real html page is on status_page.go -->
<!DOCTYPE html>
<html lang="en">
<head>
	<!-- Standard Meta -->
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
	<!-- Site Properties -->
	<title>Cron Status</title>
	<link
	   rel="stylesheet"
	   type="text/css"
	   href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.2/dist/semantic.min.css">
	<script
	   src="https://code.jquery.com/jquery-3.1.1.min.js"
	   integrity="sha256-hVVnYaiADRTO2PzUGmuLJr8BLUSjGIZsDYGmIJLv2b8="
	   crossorigin="anonymous"></script>
	<script
	   src="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.2/dist/semantic.min.js"
	   crossorigin="anonymous"></script>
	<style type="text/css">
        body > .ui.container {
            margin-top: 3em;
        }
	</style>
	<title>Cron Status</title>
</head>
<body>
<div class="ui container">
	<div class="ui left fixed vertical stackable inverted main menu">
		<div class="header item">
			<i class="stopwatch icon"></i>
			Cronx
		</div>
		<a class="item active" href="javascript:window.location.reload(true)">
			<i class="tasks icon"></i>
			Status
		</a>
	</div>
	<div class="ui three steps">
		<div class="step">
			<i class="arrow up icon"></i>
			<div class="content">
				<div class="title">Up</div>
				<div class="description">Job has just been created</div>
			</div>
		</div>
		<div class="step">
			<i class="sync icon"></i>
			<div class="content">
				<div class="title">Running</div>
				<div class="description">Job is currently running</div>
			</div>
		</div>
		<div class="step">
			<i class="play icon"></i>
			<div class="content">
				<div class="title">Idle</div>
				<div class="description">Job is waiting for next execution time</div>
			</div>
		</div>
	</div>
	<div id="table_status">
		<table class="ui sortable selectable right aligned celled table">
			<thead>
			<tr>
				<th class="left aligned">ID</th>
				<th class="left aligned">Name</th>
				<th class="center aligned">Status</th>
				<th>Last run</th>
				<th class="sorted ascending">Next run</th>
				<th>Latency</th>
			</tr>
			</thead>
			<tbody>
            {{range .}}
				<tr
                        {{if eq .Job.Status "RUNNING"}} class="warning"
                        {{else if eq .Job.Status "IDLE"}} class="positive"
                        {{end}}
				>
					<td class="left aligned">{{.ID}}</td>
					<td class="left aligned">{{.Job.Name}}</td>
					<td class="center aligned">
                        {{if eq .Job.Status "RUNNING"}}
							<div class="ui yellow label">
								<i class="sync icon"></i>
                                {{.Job.Status}}
							</div>
                        {{else if eq .Job.Status "IDLE"}}
							<div class="ui green label">
								<i class="play icon"></i>
                                {{.Job.Status}}
							</div>

                        {{else}}
							<div class="ui label">
								<i class="arrow up icon"></i>
                                {{.Job.Status}}
							</div>
                        {{end}}
					</td>
					<td>{{if not .Prev.IsZero}}{{.Prev.Format "2006-01-02 15:04:05"}}{{end}}</td>
					<td>{{if not .Next.IsZero}}{{.Next.Format "2006-01-02 15:04:05"}}{{end}}</td>
					<td>{{.Job.Latency}}</td>
				</tr>
            {{end}}
			</tbody>
		</table>
	</div>
</div>
</body>
</html>
`

var (
	once       sync.Once
	statusPage *template.Template
	err        error
)

func GetStatusPageTemplate() (*template.Template, error) {
	once.Do(func() {
		t := template.New("status_page.html")
		statusPage, err = t.Parse(statusPageTemplate)
	})

	return statusPage, err
}