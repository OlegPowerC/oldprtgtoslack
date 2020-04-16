In PRTG add "Execute HTTP Action" notify and configure it like:

URL: http://Your_Docker_Host:8788/slack

Postdata:
{
    "text": "%device %status %down (%message)... Alerting @channel",
    "link_names": 1,
	"attachments": [
        {
            "title": "Sensor: %device %name",
            "title_link": "%homesensor.htm?id=%sensorid",
            "text": "*Status:* %status %down \n*Date/Time:* %datetime (UTC) \n*Last Result:* %lastvalue \n*Last Message:* %message \n*Probe:* %probe \n*Group:* %group \n*Device:* %device () \n*Last Scan:* %lastcheck \n*Last Up:* %lastup \n*Last Down:* %lastdown \n*Uptime:* %uptime \n*Downtime:* %downtime \n*Cumulated since:* %cumsince \n*Location:* %location",
			"color": "%colorofstate",
			"mrkdwn_in": ["text", "pretext"]
        }
    ]
}

In docker-compose.yml:
set SLACK_URL variable - it is Your Slack webhook (https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX)