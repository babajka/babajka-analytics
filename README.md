## Pushing Analytics from Yandex.Metrics to our DB

### Crontab setup

every half an hour. `crontab -e` to edit

DEV:

`3,33 * * * * /home/wir-dev/deployed/analytics/babajka-analytics --secretPath=/home/wir-dev/secret-staging.json --env=staging`

PROD:

`6,36 * * * * /home/wir-prod/deployed/analytics/babajka-analytics --secretPath=/home/wir-prod/secret-production.json --env=production`
