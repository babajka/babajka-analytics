## Pushing Analytics from Yandex.Metrics to our DB

### Crontab setup

every half an hour silently. twice a day with slack notifications.

`crontab -e` to edit. `crontab -l` to list

**DEV:**

```
3,33 * * * * /home/wir-dev/deployed/analytics/babajka-analytics --secretPath=/home/wir-dev/secret-staging.json --env=staging

14 5,17 * * * /home/wir-dev/deployed/analytics/babajka-analytics --secretPath=/home/wir-dev/secret-staging.json --env=staging --enableSlack
```

**PROD:**

```
6,36 * * * * /home/wir-prod/deployed/analytics/babajka-analytics --secretPath=/home/wir-prod/secret-production.json --env=production

15 5,17 * * * /home/wir-prod/deployed/analytics/babajka-analytics --secretPath=/home/wir-prod/secret-production.json --env=production --enableSlack

```
