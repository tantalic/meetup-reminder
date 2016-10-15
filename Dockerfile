FROM scratch
MAINTAINER Kevin Stock <kevinstock@tantalic.com>

ADD certs/ca-certificates.crt /etc/ssl/certs/
ADD build/meetup-reminder-linux-amd64 meetup-reminder

# ENV SLACK_WEBHOOK https://hooks.slack.com/services/xxxxxxxxx/xxxxxxxxx/xxxxxxxxxxxxxxxxxxxx
# ENV MEETUP_NAME meetup_url_name
# ENV DAYS_BEFORE_REMINDER 7
# ENV MESSAGE "is in a week"

CMD ["/meetup-reminder"]
