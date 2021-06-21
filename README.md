# The App
The app is a simple two route Go app that adds in some random delays to response times. 
# Install Tilt
Install Tilt @ https://docs.tilt.dev/install.html

# Mock some traffic
This repo uses Locust to simulate calls to the two endpoints
## Install Locust
Install Locust @ https://docs.locust.io/en/stable/installation.html

## Run the Load Test
Generate some Load

```
locust --headless -f ./auto/locust_file.py --only-summary -u 5 -r 10 -t 15m --stop-timeout 10 -H http://localhost:8080
```

# Visualise in Grafana
- Open up `http://localhost:3000`
- Explore your data
```
avg by (quantile,method,route) (
  rate(http_request_summary[5m])
)
```