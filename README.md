## Pushing Analytics from Yandex.Metrics to our DB

### Crontab setup

every half an hour. `crontab -e` to edit

`3,33 * * * * /home/wir-dev/deployed/analytics/babajka-analytics --secretPath=/home/wir-dev/secret-staging.json`
